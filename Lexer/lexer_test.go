package Lexer

import (
	"testing"

	"github.com/Mellotonio/Andrei_lang/Token"
)

func TestNextToken(t *testing.T) {
	input := `=+(){},;`

	tests := []struct {
		expectedType    Token.TokenType
		expectedLiteral string
	}{
		{Token.ASSIGN, "="},
		{Token.PLUS, "+"},
		{Token.LPAREN, "("},
		{Token.RPAREN, ")"},
		{Token.LBRACE, "{"},
		{Token.RBRACE, "}"},
		{Token.COMMA, ","},
		{Token.SEMICOLON, ";"},
		{Token.EOF, ""},
	}
	l := New(input)

	for i, tt := range tests {
		tok := l.nextToken()

		if tok.Type != tt.expectedType {
			t.Fatalf("tests[%d] - tokentype wrong. expected=%q, got=%q", i, tt.expectedType, tok.Type)
		}
		if tok.Literal != tt.expectedLiteral {
			t.Fatalf("tests[%d] - literal wrong. expected=%q, got=%q", i, tt.expectedLiteral, tok.Literal)
		}
	}
}
