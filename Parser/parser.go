package Parser

import (
	"fmt"
	"strconv"

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
	INDEX
)

var precedences = map[Token.TokenType]int{
	Token.EQ:       EQUALS,
	Token.NOT_EQ:   EQUALS,
	Token.LT:       LESSGREATER,
	Token.GT:       LESSGREATER,
	Token.PLUS:     SUM,
	Token.MINUS:    SUM,
	Token.SLASH:    PRODUCT,
	Token.ASTERISK: PRODUCT,
	Token.LBRACKET: INDEX,
	Token.LPAREN:   CALL,
}

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

	// Registro de cada função que deve ser chamada ao encontrar determinado "prefix"
	p.prefixParseFns = make(map[Token.TokenType]prefixParseFn)
	p.registerPrefix(Token.IDENT, p.parseIdentifier)
	p.registerPrefix(Token.INT, p.parseIntegerLiteral)
	p.registerPrefix(Token.FLOAT, p.parseFloatLiteral)
	p.registerPrefix(Token.BANG, p.parsePrefixExpression)
	p.registerPrefix(Token.MINUS, p.parsePrefixExpression)
	p.registerPrefix(Token.TRUE, p.parseBoolean)
	p.registerPrefix(Token.FALSE, p.parseBoolean)
	p.registerPrefix(Token.LPAREN, p.parseGroupedExpression)
	p.registerPrefix(Token.IF, p.parseIfExpression)
	p.registerPrefix(Token.FUNCTION, p.parseFunctionLiteral)
	p.registerPrefix(Token.STRING, p.parseStringLiteral)
	p.registerPrefix(Token.LBRACKET, p.parseArrayLiteral)

	// Registro de cada função que deve ser chamada ao encontrar determinado "infix"
	p.infixParseFns = make(map[Token.TokenType]infixParseFn)
	p.registerInfix(Token.PLUS, p.parseInfixExpression)
	p.registerInfix(Token.MINUS, p.parseInfixExpression)
	p.registerInfix(Token.SLASH, p.parseInfixExpression)
	p.registerInfix(Token.ASTERISK, p.parseInfixExpression)
	p.registerInfix(Token.EQ, p.parseInfixExpression)
	p.registerInfix(Token.NOT_EQ, p.parseInfixExpression)
	p.registerInfix(Token.LT, p.parseInfixExpression)
	p.registerInfix(Token.GT, p.parseInfixExpression)
	p.registerInfix(Token.LPAREN, p.parseCallExpression)
	p.registerInfix(Token.LBRACKET, p.parseIndexExpression)

	// Duas vezes para termos valores tanto no curToken como no peekToken
	p.nextToken()
	p.nextToken()

	return p
}

// Retorna um wrape de AST.Identifier contendo o Token e o valor do Identifier
func (p *Parser) parseIdentifier() AST.Expression {
	return &AST.Identifier{Token: p.currentToken, Value: p.currentToken.Literal}
}

func (p *Parser) nextToken() {
	p.currentToken = p.peekToken
	p.peekToken = p.l.NextToken()
}

func (p *Parser) noPrefixParseFnError(t Token.TokenType) {
	msg := fmt.Sprintf("no prefix parse function for %s found", t)
	p.errors = append(p.errors, msg)
}

func (p *Parser) ParseProgram() *AST.Program {
	// Inicializa a AST do nosso programa
	program := &AST.Program{}
	// Inicia a cadeia de statements da AST
	program.Statements = []AST.Statement{}

	for !p.currentTokenIs(Token.EOF) {
		// Olha o proximo token e "parseia" ele de acordo com o que ele representa
		stmt := p.ParseStatement()
		if stmt != nil {
			program.Statements = append(program.Statements, stmt)
		}
		p.nextToken()
	}

	return program
}

func (p *Parser) ParseStatement() AST.Statement {
	switch p.currentToken.Type {
	case Token.MOONVAR:
		// Tokens como "moonvar" ou "return" são faceis de diferenciar
		// Ambos tem um shape pre-definido de como eles devem ser
		return p.ParseMoonvarStatement()
	case Token.RETURN:
		return p.ParseReturnStatement()
	default:
		// Expressões representam qualquer expressão depois do "="
		// O principal cuidado que se deve ter é no momento de realizar operações que possuem precedencia
		return p.ParseExpressionStatement()
	}
}

// Constroi um statement a partir do MoonvarStatement e faz asserções sobre
// como um MoonvarStatement deveria se comportar ->  ex: "moonvar x = 5"
func (p *Parser) ParseMoonvarStatement() *AST.MoonvarStatement {
	statement := &AST.MoonvarStatement{Token: p.currentToken}

	// Espera que o proximo token seja um identifier, ou seja, o nome da variavel
	if !p.expectPeek(Token.IDENT) {
		return nil
	}

	// Guarda o nome do statement como o "nome" da variavel ex: "moonvar x = 1" -> guardamos o x
	statement.Name = &AST.Identifier{Token: p.currentToken, Value: p.currentToken.Literal}

	// Espera que o proximo token seja um "="
	if !p.expectPeek(Token.ASSIGN) {
		return nil
	}

	p.nextToken()

	// Avalia a proxima expressão, após o sinal de =
	statement.Value = p.parseExpression(LOWEST)

	for p.peekTokenIs(Token.SEMICOLON) {
		p.nextToken()
	}

	return statement
}

// Constroi um statement a partir do returnStatement e faz asserções sobre
func (p *Parser) ParseReturnStatement() *AST.ReturnStatement {
	statement := &AST.ReturnStatement{Token: p.currentToken}

	p.nextToken()

	// Return (expression) -> Essa expression que estamos parseando aqui
	statement.ReturnValue = p.parseExpression(LOWEST)

	for p.peekTokenIs(Token.SEMICOLON) {
		p.nextToken()
	}

	return statement
}

func (p *Parser) ParseExpressionStatement() *AST.ExpressionStatement {
	statement := &AST.ExpressionStatement{Token: p.currentToken}

	// Vemos o inicio do algoritmo de Vaughan Pratt aqui
	statement.Expression = p.parseExpression(LOWEST)

	if p.peekTokenIs(Token.SEMICOLON) {
		p.nextToken()
	}

	return statement
}

func (p *Parser) parseExpression(precedence int) AST.Expression {
	// Acha uma função (prefix) que se encaixa com o token atual
	prefix := p.prefixParseFns[p.currentToken.Type]

	if prefix == nil {
		p.noPrefixParseFnError(p.currentToken.Type)
		return nil
	}

	// Executa a função antes capturada, a funçao mais à esquerda
	leftexp := prefix()

	// Função recursiva -> aonde pegamos o infix que é sempre a função do proximo token
	// Avançamos um token
	// Jogamos na função do proximo token, a funçao mais à esquerda (salvando ela na "AST")
	// Para então no final formarmos uma expressão completa
	for !p.peekTokenIs(Token.SEMICOLON) && precedence < p.peekPrecedence() {
		infix := p.infixParseFns[p.peekToken.Type]

		if infix == nil {
			return leftexp
		}
		p.nextToken()

		leftexp = infix(leftexp)
	}

	return leftexp
}

func (p *Parser) parsePrefixExpression() AST.Expression {
	expression := &AST.PrefixExpression{
		Token:    p.currentToken,
		Operator: p.currentToken.Literal,
	}

	p.nextToken()

	// O valor "PREFIX" indica o nivel de precedencia, que é quase o maior possivel
	// Fazendo com que dessa forma ele retorne sem entrar no loop recursivo
	expression.Right = p.parseExpression(PREFIX)

	return expression
}

func (p *Parser) parseInfixExpression(left AST.Expression) AST.Expression {
	expression := &AST.InfixExpression{
		Token:    p.currentToken,
		Operator: p.currentToken.Literal,
		Left:     left,
	}

	precedence := p.currPrecedence()
	p.nextToken()
	// Recursividade -> chama a parseExpression sempre com um token a frente da precedencia
	// e desse forma gera um loop que guarda as expressões de 2 em 2 até encontrar Token.SEMICOLON
	expression.Right = p.parseExpression(precedence)

	return expression
}

// Grouped expressions baseiam se nos parenteses, cada parenteses podem apenas ter 2 elementos
// Porém isso pode ocorrer (2 / (5 + 5)), dois no parenteses mais a dentro, e no de fora... também dois, infix(infix(2), infix(5,5))
func (p *Parser) parseGroupedExpression() AST.Expression {
	p.nextToken()

	exp := p.parseExpression(LOWEST)

	if !p.expectPeek(Token.RPAREN) {
		return nil
	}
	return exp
}

func (p *Parser) parseIfExpression() AST.Expression {
	expression := &AST.IfExpression{Token: p.currentToken}

	// Espera que o inicio da expressão seja (
	if !p.expectPeek(Token.LPAREN) {
		return nil
	}

	p.nextToken()

	// Parseia o que está entre parenteses
	expression.Condition = p.parseExpression(LOWEST)

	// Espera o fechamento do parenteses )
	if !p.expectPeek(Token.RPAREN) {
		return nil
	}

	// Espera o inicio do {
	if !p.expectPeek(Token.LBRACE) {
		return nil
	}

	expression.Consequence = p.parseBlockStatement()

	// Apos o parse do if (...){} verificamos se o proximo token é um ELSE, para então começarmos a evaluar novamente else{...}
	if p.peekTokenIs(Token.ELSE) {
		p.nextToken()

		if !p.expectPeek(Token.LBRACE) {
			return nil
		}

		expression.Alternative = p.parseBlockStatement()
	}

	return expression
}

func (p *Parser) parseBlockStatement() *AST.BlockStatement {
	block := &AST.BlockStatement{Token: p.currentToken}

	block.Statements = []AST.Statement{}

	p.nextToken()

	// Enquanto um } não é encontrado, continuamos "parseando" a expressão
	for !p.currentTokenIs(Token.RBRACE) && !p.currentTokenIs(Token.EOF) {
		statement := p.ParseStatement()
		if statement != nil {
			block.Statements = append(block.Statements, statement)
		}
		p.nextToken()
	}
	return block
}

func (p *Parser) parseFunctionLiteral() AST.Expression {
	lit := &AST.FunctionLiteral{Token: p.currentToken}

	if !p.expectPeek(Token.LPAREN) {
		return nil
	}

	lit.Parameters = p.parseFunctionParameters()

	if !p.expectPeek(Token.LBRACE) {
		return nil
	}

	lit.Body = p.parseBlockStatement()

	return lit

}

func (p *Parser) parseFunctionParameters() []*AST.Identifier {
	identifiers := []*AST.Identifier{}

	// Checa se é uma função vazia, fn()
	if p.peekTokenIs(Token.RPAREN) {
		p.nextToken()
		return identifiers
	}

	// Captura o primeiro valor
	p.nextToken()

	// Guarda o primeiro valor num identifier
	ident := &AST.Identifier{Token: p.currentToken, Value: p.currentToken.Literal}
	identifiers = append(identifiers, ident)

	// Se o proximo token for uma vigula
	for p.peekTokenIs(Token.COMMA) {
		// Pula valor anterior
		p.nextToken()
		// Pula virgula
		p.nextToken()
		ident := &AST.Identifier{Token: p.currentToken, Value: p.currentToken.Literal}
		identifiers = append(identifiers, ident)
	}

	// Verifica se a função termina com ")"
	if !p.expectPeek(Token.RPAREN) {
		return nil
	}

	return identifiers
}

// Captura a call, preenchendo seus argumentos
// callsFunction(2,3,fn(x,y){x+y;});
func (p *Parser) parseCallExpression(function AST.Expression) AST.Expression {
	exp := &AST.CallExpression{Token: p.currentToken, Function: function}
	exp.Arguments = p.parseExpressionList(Token.RPAREN)
	return exp
}

/*
func (p *Parser) parseCallArguments() []AST.Expression {
	args := []AST.Expression{}

	if p.peekTokenIs(Token.RPAREN) {
		p.nextToken()
		return args
	}

	p.nextToken()

	// Inicia o "parseamento" dos argumentos chamados
	args = append(args, p.parseExpression(LOWEST))

	for p.peekTokenIs(Token.COMMA) {
		p.nextToken()
		p.nextToken()
		args = append(args, p.parseExpression(LOWEST))
	}

	if !p.expectPeek(Token.RPAREN) {
		p.nextToken()
		return args
	}

	return args
}*/

func (p *Parser) parseIndexExpression(left AST.Expression) AST.Expression {
	exp := &AST.IndexExpression{Token: p.currentToken, Left: left}
	// Avançamos para o numero
	p.nextToken()

	// parse the number in array->[1]
	exp.Index = p.parseExpression(LOWEST)

	if !p.expectPeek(Token.RBRACKET) {
		return nil
	}

	return exp
}

// Transforma um string (que tem valor inteiro) em inteiro.
func (p *Parser) parseIntegerLiteral() AST.Expression {
	lit := &AST.LiteralInteger{Token: p.currentToken}

	value, err := strconv.ParseInt(p.currentToken.Literal, 0, 64)

	if err != nil {
		msg := fmt.Sprintf("Could not parse %q as integer", p.currentToken.Literal)
		p.errors = append(p.errors, msg)
		return nil
	}

	lit.Value = value

	return lit

}

func (p *Parser) parseFloatLiteral() AST.Expression {
	lit := &AST.LiteralFloat{Token: p.currentToken}

	value, err := strconv.ParseFloat(p.currentToken.Literal, 64)

	if err != nil {
		msg := fmt.Sprintf("Could not parse %q as integer", p.currentToken.Literal)
		p.errors = append(p.errors, msg)
		return nil
	}

	lit.Value = value

	return lit

}
func (p *Parser) parseStringLiteral() AST.Expression {
	return &AST.StringLiteral{Token: p.currentToken, Value: p.currentToken.Literal}
}

func (p *Parser) parseBoolean() AST.Expression {
	return &AST.Boolean{Token: p.currentToken, Value: p.currentTokenIs(Token.TRUE)}
}

// Token Atual
func (p *Parser) currentTokenIs(t Token.TokenType) bool {
	return p.currentToken.Type == t
}

// Proximo Token
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

func (p *Parser) peekPrecedence() int {
	if p, ok := precedences[p.peekToken.Type]; ok {
		return p
	}

	return LOWEST
}

func (p *Parser) currPrecedence() int {
	if p, ok := precedences[p.currentToken.Type]; ok {
		return p
	}

	return LOWEST
}

func (p *Parser) parseArrayLiteral() AST.Expression {
	array := &AST.ArrayLiteral{Token: p.currentToken}

	array.Elements = p.parseExpressionList(Token.RBRACKET)

	return array
}

func (p *Parser) parseExpressionList(end Token.TokenType) []AST.Expression {
	list := []AST.Expression{}

	if p.peekTokenIs(end) {
		p.nextToken()
		return list
	}

	p.nextToken()

	// Inicia o "parseamento" dos argumentos chamados
	list = append(list, p.parseExpression(LOWEST))

	for p.peekTokenIs(Token.COMMA) {
		p.nextToken()
		p.nextToken()
		list = append(list, p.parseExpression(LOWEST))
	}

	if !p.expectPeek(end) {
		return nil
	}

	return list
}
