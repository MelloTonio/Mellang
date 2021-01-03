package Lexer

import (
	"github.com/Mellotonio/Andrei_lang/Token"
)

type Lexer struct {
	input        string // Texto a ser recebido e interpretado
	position     int    // Posição atual
	readPosition int    // Posição atual + 1
	ch           byte   // Caractere atual
}

func New(input string) *Lexer {
	l := &Lexer{input: input} // Começa um novo lexer com o "texto" passado
	l.readChar()              // Inializa o lexer com a primeira posição do texto
	return l
}

func (l *Lexer) readChar() {
	// Verificamos se a proxima posição é um EOF, se for indicamos isso com um 0
	if l.readPosition >= len(l.input) {
		l.ch = 0
	} else {
		// Se não, lemos o caractere que a posição representa
		l.ch = l.input[l.readPosition]
	}

	l.position = l.readPosition // Sempre guardamos a ultima posição no position

	l.readPosition += 1
}

// Para cada elemento do texto, analisamos ele e damos um token correspondente a representação dele
func (l *Lexer) NextToken() Token.Token {
	var tok Token.Token

	l.skipWhitespace()

	switch l.ch {
	case '=':
		if l.peekChar() == '=' {
			ch := l.ch                                                            // ch = "="
			l.readChar()                                                          // l.ch = "="
			tok = Token.Token{Type: Token.EQ, Literal: string(ch) + string(l.ch)} // "="+"=" == "=="
		} else {
			tok = newToken(Token.ASSIGN, l.ch)
		}
	case '+':
		tok = newToken(Token.PLUS, l.ch)
	case '-':
		tok = newToken(Token.MINUS, l.ch)
	case '!':
		if l.peekChar() == '!' {
			ch := l.ch
			l.readChar()
			tok = Token.Token{Type: Token.NOT_EQ, Literal: string(ch) + string(l.ch)}
		}
		tok = newToken(Token.BANG, l.ch)
	case '/':
		tok = newToken(Token.SLASH, l.ch)
	case '*':
		tok = newToken(Token.ASTERISK, l.ch)
	case '<':
		tok = newToken(Token.LT, l.ch)
	case '>':
		tok = newToken(Token.GT, l.ch)
	case ';':
		tok = newToken(Token.SEMICOLON, l.ch)
	case ',':
		tok = newToken(Token.COMMA, l.ch)
	case '{':
		tok = newToken(Token.LBRACE, l.ch)
	case '}':
		tok = newToken(Token.RBRACE, l.ch)
	case '(':
		tok = newToken(Token.LPAREN, l.ch)
	case ')':
		tok = newToken(Token.RPAREN, l.ch)
	case 0:
		tok.Literal = ""
		tok.Type = Token.EOF
	default:
		if isLetter(l.ch) { // É letra?
			tok.Literal = l.readIdentifier() // a partir desta função retornamos a frase
			tok.Type = Token.LookupIdent(tok.Literal)
			return tok
		} else if isDigit(l.ch) {
			tok.Type = Token.INT
			tok.Literal = l.readNumber()
			return tok
		} else {
			tok = newToken(Token.ILLEGAL, l.ch)
		}
	}

	l.readChar()
	return tok
}

func newToken(tokenType Token.TokenType, ch byte) Token.Token {
	return Token.Token{Type: tokenType, Literal: string(ch)}
}

func (l *Lexer) readIdentifier() string {
	first_position := l.position

	for isLetter(l.ch) {
		l.readChar()
	}

	return l.input[first_position:l.position]
}

func isLetter(ch byte) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_'
}

func (l *Lexer) skipWhitespace() {
	for l.ch == ' ' || l.ch == '\t' || l.ch == '\n' || l.ch == '\r' {
		l.readChar()
	}
}

func isDigit(ch byte) bool {
	return '0' <= ch && ch <= '9'
}

func (l *Lexer) readNumber() string {
	first_position := l.position

	for isDigit(l.ch) {
		l.readChar()
	}

	return l.input[first_position:l.position]
}

func (l *Lexer) peekChar() byte {
	if l.readPosition > len(l.input) {
		return 0
	} else {
		return l.input[l.readPosition]
	}
}
