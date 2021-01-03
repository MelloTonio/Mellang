package Parser

import (
	"github.com/Mellotonio/Andrei_lang/AST"
	"github.com/Mellotonio/Andrei_lang/Lexer"
	"github.com/Mellotonio/Andrei_lang/Token"
)

type Parser struct {
	l *Lexer.Lexer

	currentToken Token.Token
	peekToken    Token.Token
}

func New(lexer *Lexer.Lexer) *Parser {
	p := &Parser{l: lexer}

	p.nextToken()
	p.nextToken()

	return p
}

func (p *Parser) nextToken() {
	p.currentToken = p.peekToken
	p.peekToken = p.l.NextToken()
}

func (p *Parser) ParseProgram() *AST.Program { return nil }
