package AST

import (
	"testing"

	"github.com/Mellotonio/Andrei_lang/Token"
)

func TestString(t *testing.T) {
	program := &Program{
		Statements: []Statement{
			&MoonvarStatement{
				Token: Token.Token{Type: Token.MOONVAR, Literal: "let"},
				Name: &Identifier{
					Token: Token.Token{Type: Token.IDENT, Literal: "myVar"},
					Value: "myVar",
				},
				Value: &Identifier{
					Token: Token.Token{Type: Token.IDENT, Literal: "anotherVar"},
					Value: "anotherVar",
				},
			},
		},
	}

	if program.String() != "let myVar = anotherVar;" {
		t.Errorf("program.String() wrong. got=%q", program.String())
	}
}
