package AST

import (
	"bytes"

	"github.com/Mellotonio/Andrei_lang/Token"
)

type Node interface {
	TokenLiteral() string
	String() string
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

type ReturnStatement struct {
	Token       Token.Token
	ReturnValue Expression
}

type ExpressionStatement struct {
	Token      Token.Token
	Expression Expression
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

func (mv *ReturnStatement) statementNode()       {}
func (mv *ReturnStatement) TokenLiteral() string { return mv.Token.Literal }

func (mv *ExpressionStatement) statementNode()       {}
func (mv *ExpressionStatement) TokenLiteral() string { return mv.Token.Literal }

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
func (p *Program) String() string {
	var out bytes.Buffer

	for _, s := range p.Statements {
		out.WriteString(s.String())
	}
	return out.String()
}

func (ms *MoonvarStatement) String() string {
	var out bytes.Buffer

	out.WriteString(ms.TokenLiteral() + " ")
	out.WriteString(ms.Name.String())
	out.WriteString(" = ")

	if ms.Value != nil {
		out.WriteString(ms.Value.String())
	}

	out.WriteString(";")

	return out.String()
}

func (rs *ReturnStatement) String() string {
	var out bytes.Buffer

	out.WriteString(rs.TokenLiteral() + " ")

	if rs.ReturnValue != nil {
		out.WriteString(rs.ReturnValue.String())
	}

	out.WriteString(";")

	return out.String()
}

func (es *ExpressionStatement) String() string {
	if es.Expression != nil {
		return es.Expression.String()
	}

	return ""
}

func (i *Identifier) String() string { return i.Value }
