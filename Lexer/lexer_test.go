package Lexer

import (
	"testing"
)

func TestNextToken(t *testing.T) {
	/*input := `=+(){},;`

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
		tok := l.NextToken()

		if tok.Type != tt.expectedType {
			t.Fatalf("tests[%d] - tokentype wrong. expected=%q, got=%q", i, tt.expectedType, tok.Type)
		}
		if tok.Literal != tt.expectedLiteral {
			t.Fatalf("tests[%d] - literal wrong. expected=%q, got=%q", i, tt.expectedLiteral, tok.Literal)
		}
	}*/

	/*
		input := `moonvar five = 5;moonvar ten = 10;moonvar add = fn(x, y) {x + y;};moonvar result = add(five, ten);`

		tests := []struct {
			expectedType    Token.TokenType
			expectedLiteral string
		}{
			{Token.MOONVAR, "moonvar"},
			{Token.IDENT, "five"},
			{Token.ASSIGN, "="},
			{Token.INT, "5"},
			{Token.SEMICOLON, ";"},
			{Token.MOONVAR, "moonvar"},
			{Token.IDENT, "ten"},
			{Token.ASSIGN, "="},
			{Token.INT, "10"},
			{Token.SEMICOLON, ";"},
			{Token.MOONVAR, "moonvar"},
			{Token.IDENT, "add"},
			{Token.ASSIGN, "="},
			{Token.FUNCTION, "fn"},
			{Token.LPAREN, "("},
			{Token.IDENT, "x"},
			{Token.COMMA, ","},
			{Token.IDENT, "y"},
			{Token.RPAREN, ")"},
			{Token.LBRACE, "{"},
			{Token.IDENT, "x"},
			{Token.PLUS, "+"},
			{Token.IDENT, "y"},
			{Token.SEMICOLON, ";"},
			{Token.RBRACE, "}"},
			{Token.SEMICOLON, ";"},
			{Token.MOONVAR, "moonvar"},
			{Token.IDENT, "result"},
			{Token.ASSIGN, "="},
			{Token.IDENT, "add"},
			{Token.LPAREN, "("},
			{Token.IDENT, "five"},
			{Token.COMMA, ","},
			{Token.IDENT, "ten"},
			{Token.RPAREN, ")"},
			{Token.SEMICOLON, ";"},
			{Token.EOF, ""},
		}
	*/

	input := `let five = 5;
	let ten = 10;
	
	let add = fn(x, y) {
	  x + y;
	};
	
	let result = add(five, ten);
	!-/*5;
	5 < 10 > 5;
	
	if (5 < 10) {
		return true;
	} else {
		return false;
	}
	
	10 == 10;
	10 != 9;
	`

	l := New(input)

	for i, tt := range tests {
		tok := l.NextToken()

		if tok.Type != tt.expectedType {
			t.Fatalf("tests[%d] - tokentype wrong. expected=%q, got=%q", i, tt.expectedType, tok.Type)
		}
		if tok.Literal != tt.expectedLiteral {
			t.Fatalf("tests[%d] - literal wrong. expected=%q, got=%q", i, tt.expectedLiteral, tok.Literal)
		}
	}
}
