package AST

import "github.com/Mellotonio/Andrei_lang/Token"

type Node interface {
	TokenLiteral() string
}

type Statement interface {
	Node
	statementNode()
}
type Expression interface {
	Node
	expressionNode()
}

// Guarda a serie de statements
type Program struct {
	Statements []Statement
}

type MoonvarStatement struct {
	Token Token.Token
	Name  *Identifier // Nome da variavel
	Value Expression  // Expressão que a variavel está recebendo
}

func (p *Program) TokenLiteral() string {
	if len(p.Statements) > 0 {
		return p.Statements[0].TokenLiteral()
	} else {
		return ""
	}
}

func (mv *MoonvarStatement) statementNode()       {}
func (mv *MoonvarStatement) TokenLiteral() string { return mv.Token.Literal }

type Identifier struct {
	Token Token.Token
	Value string
}

func (i *Identifier) expressionNode()      {}
func (i *Identifier) TokenLiteral() string { return i.Token.Literal }

// AST TREE
/*
STATEMENTS -> MoonvarStatement (name, value)
								 |		|_____Expression
						     Identifier
*/
