package Parser

import (
	"fmt"

	"github.com/Mellotonio/Andrei_lang/AST"
	"github.com/Mellotonio/Andrei_lang/Lexer"
	"github.com/Mellotonio/Andrei_lang/Token"
)

const (
	_int = iota

	LOWEST
	EQUALS
	LESSGREATER
	SUM
	PRODUCT
	PREFIX
	CALL
)

type Parser struct {
	l *Lexer.Lexer

	currentToken Token.Token
	peekToken    Token.Token

	errors []string

	prefixParseFns map[Token.TokenType]prefixParseFn
	infixParseFns  map[Token.TokenType]infixParseFn
}

type (
	prefixParseFn func() AST.Expression
	infixParseFn  func(AST.Expression) AST.Expression
)

func New(lexer *Lexer.Lexer) *Parser {
	p := &Parser{l: lexer, errors: []string{}}

	// Duas vezes para termos valores tanto no curToken como no peekToken
	p.nextToken()
	p.nextToken()

	return p
}

func (p *Parser) nextToken() {
	p.currentToken = p.peekToken
	p.peekToken = p.l.NextToken()
}

func (p *Parser) ParseProgram() *AST.Program {
	program := &AST.Program{}
	program.Statements = []AST.Statement{}

	// Iteramos sobre os tokens até encontrarmos o fim do texto
	for p.currentToken.Type != Token.EOF {
		statement := p.ParseStatement()
		if statement != nil {
			program.Statements = append(program.Statements, statement)
		}
		p.nextToken()
	}

	return program
}

func (p *Parser) ParseStatement() AST.Statement {
	switch p.currentToken.Type {
	case Token.MOONVAR:
		return p.ParseMoonvarStatement()
	case Token.RETURN:
		return p.ParseReturnStatement()
	default:
		return p.ParseExpressionStatement()
	}
}

// Constroi um statement a partir do MoonvarStatement e faz asserções sobre
// como um MoonvarStatement deveria se comportar ->  ex: "moonvar x = 5"
func (p *Parser) ParseMoonvarStatement() *AST.MoonvarStatement {
	statement := &AST.MoonvarStatement{Token: p.currentToken}

	// Espera que o proximo token seja um identifier
	if !p.expectPeek(Token.IDENT) {
		return nil
	}

	// The x in 'moonvar x = 1'
	statement.Name = &AST.Identifier{Token: p.currentToken, Value: p.currentToken.Literal}

	// Espera que o proximo identifier seja um "="
	if !p.expectPeek(Token.ASSIGN) {
		return nil
	}

	for !p.currentTokenIs(Token.SEMICOLON) {
		p.nextToken()
	}

	return statement
}

func (p *Parser) ParseReturnStatement() *AST.ReturnStatement {
	statement := &AST.ReturnStatement{Token: p.currentToken}

	p.nextToken()

	for !p.currentTokenIs(Token.SEMICOLON) {
		p.nextToken()
	}

	return statement
}

func (p *Parser) ParseExpressionStatement() *AST.ExpressionStatement {
	statement := &AST.ExpressionStatement{Token: p.currentToken}

	statement.Expression = p.parseExpression(LOWEST)

	if p.peekTokenIs(Token.SEMICOLON) {
		p.nextToken()
	}

	return statement
}

func (p *Parser) parseExpression(precedence int) AST.Expression {
	prefix := p.prefixParseFns[p.currentToken.Type]

	if prefix == nil {
		return nil
	}

	leftexp := prefix()

	return leftexp
}

func (p *Parser) currentTokenIs(t Token.TokenType) bool {
	return p.currentToken.Type == t
}

func (p *Parser) peekTokenIs(t Token.TokenType) bool {
	return p.peekToken.Type == t
}

// Checa se o proximo token é realmente um token, esperado (valido)
func (p *Parser) expectPeek(t Token.TokenType) bool {
	if p.peekTokenIs(t) {
		p.nextToken()
		return true
	} else {
		p.peekError(t)
		return false
	}
}

func (p *Parser) Errors() []string {
	return p.errors
}

func (p *Parser) peekError(t Token.TokenType) {
	message := fmt.Sprintf("Expected next token to bem %s, got %s instead", t, p.peekToken.Type)
	p.errors = append(p.errors, message)
}

func (p *Parser) registerPrefix(tokenType Token.TokenType, fn prefixParseFn) {
	p.prefixParseFns[tokenType] = fn
}

func (p *Parser) registerInfix(tokenType Token.TokenType, fn infixParseFn) {
	p.infixParseFns[tokenType] = fn
}
