package Lexer

import (
	"testing"

	"github.com/Mellotonio/Andrei_lang/Token"
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

	input := `moonvar five = 5;
	moonvar ten = 10;
	
	moonvar add = fn(x, y) {
	  x + y;
	};
	
	moonvar result = add(five, ten);
	!-/*5;
	5 < 10 > 5;
	
	if (5 < 10) {
		return true;
	} else {
		return false;
	}
	
	10 == 10;
	10 != 9;
	1.5
	6.7
	five <- 5
	<=
	>=
	OwO
	~>
	`

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
		{Token.BANG, "!"},
		{Token.MINUS, "-"},
		{Token.SLASH, "/"},
		{Token.ASTERISK, "*"},
		{Token.INT, "5"},
		{Token.SEMICOLON, ";"},
		{Token.INT, "5"},
		{Token.LT, "<"},
		{Token.INT, "10"},
		{Token.GT, ">"},
		{Token.INT, "5"},
		{Token.SEMICOLON, ";"},
		{Token.IF, "if"},
		{Token.LPAREN, "("},
		{Token.INT, "5"},
		{Token.LT, "<"},
		{Token.INT, "10"},
		{Token.RPAREN, ")"},
		{Token.LBRACE, "{"},
		{Token.RETURN, "return"},
		{Token.TRUE, "true"},
		{Token.SEMICOLON, ";"},
		{Token.RBRACE, "}"},
		{Token.ELSE, "else"},
		{Token.LBRACE, "{"},
		{Token.RETURN, "return"},
		{Token.FALSE, "false"},
		{Token.SEMICOLON, ";"},
		{Token.RBRACE, "}"},
		{Token.INT, "10"},
		{Token.EQ, "=="},
		{Token.INT, "10"},
		{Token.SEMICOLON, ";"},
		{Token.INT, "10"},
		{Token.NOT_EQ, "!="},
		{Token.INT, "9"},
		{Token.SEMICOLON, ";"},
		{Token.FLOAT, "1.5"},
		{Token.FLOAT, "6.7"},
		{Token.IDENT, "five"},
		{Token.BIND, "<-"},
		{Token.INT, "5"},
		{Token.LTE, "<="},
		{Token.GTE, ">="},
		{Token.OwO, "OwO"},
		{Token.NEXT, "~>"},
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
	}
}
