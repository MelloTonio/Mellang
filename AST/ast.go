package AST

import (
	"bytes"
	"strings"

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

// DONT STOP TILL GET ENOUGH.... TATANANANANANANA
type Identifier struct {
	Token Token.Token
	Value string
}

type ExpressionStatement struct {
	Token      Token.Token
	Expression Expression
}

type PrefixExpression struct {
	Token    Token.Token
	Operator string
	Right    Expression
}

type InfixExpression struct {
	Token    Token.Token
	Left     Expression
	Operator string
	Right    Expression
}

type Boolean struct {
	Token Token.Token
	Value bool
}

type IfExpression struct {
	Token       Token.Token
	Condition   Expression
	Consequence *BlockStatement
	Alternative *BlockStatement
}

type BlockStatement struct {
	Token      Token.Token // { token
	Statements []Statement
}

type FunctionLiteral struct {
	Token      Token.Token
	Parameters []*Identifier
	Body       *BlockStatement
}

type CallExpression struct {
	Token     Token.Token
	Function  Expression
	Arguments []Expression
}

type StringLiteral struct {
	Token Token.Token
	Value string
}

type IndexExpression struct {
	Token Token.Token
	Left  Expression
	Index Expression
}

type BindExpression struct {
	Token Token.Token // The := token
	Left  string
	Value Expression
}

func (be *BindExpression) expressionNode() {}

// TokenLiteral prints the literal value of the token associated with this node
func (be *BindExpression) TokenLiteral() string { return be.Token.Literal }

// String returns a stringified version of the AST for debugging
func (be *BindExpression) String() string {
	var out bytes.Buffer

	out.WriteString(be.Left)
	out.WriteString(be.TokenLiteral())
	out.WriteString(be.Value.String())

	return out.String()
}

func (ie *IndexExpression) expressionNode()      {}
func (ie *IndexExpression) TokenLiteral() string { return ie.Token.Literal }
func (ie *IndexExpression) String() string {
	var out bytes.Buffer
	out.WriteString("(")
	out.WriteString(ie.Left.String())
	out.WriteString("[")
	out.WriteString(ie.Index.String())
	out.WriteString("])")
	return out.String()
}

func (sl *StringLiteral) expressionNode()      {}
func (sl *StringLiteral) TokenLiteral() string { return sl.Token.Literal }
func (sl *StringLiteral) String() string       { return sl.Token.Literal }

func (ce *CallExpression) expressionNode()      {}
func (ce *CallExpression) TokenLiteral() string { return ce.Token.Literal }
func (ce *CallExpression) String() string {
	var out bytes.Buffer
	args := []string{}
	for _, a := range ce.Arguments {
		args = append(args, a.String())
	}
	out.WriteString(ce.Function.String())
	out.WriteString("(")
	out.WriteString(strings.Join(args, ", "))
	out.WriteString(")")
	return out.String()
}

func (p *Program) TokenLiteral() string {
	if len(p.Statements) > 0 {
		return p.Statements[0].TokenLiteral()
	} else {
		return ""
	}
}

type LiteralInteger struct {
	Token Token.Token
	Value int64
}

type LiteralFloat struct {
	Token Token.Token
	Value float64
}

func (lf *LiteralFloat) expressionNode()      {}
func (lf *LiteralFloat) TokenLiteral() string { return lf.Token.Literal }
func (lf *LiteralFloat) String() string       { return lf.Token.Literal }

func (mv *MoonvarStatement) statementNode()       {}
func (mv *MoonvarStatement) TokenLiteral() string { return mv.Token.Literal }

func (mv *ReturnStatement) statementNode()       {}
func (mv *ReturnStatement) TokenLiteral() string { return mv.Token.Literal }

func (mv *ExpressionStatement) statementNode()       {}
func (mv *ExpressionStatement) TokenLiteral() string { return mv.Token.Literal }

func (il *LiteralInteger) expressionNode()      {}
func (il *LiteralInteger) TokenLiteral() string { return il.Token.Literal }
func (il *LiteralInteger) String() string       { return il.Token.Literal }

func (oe *InfixExpression) expressionNode()      {}
func (oe *InfixExpression) TokenLiteral() string { return oe.Token.Literal }
func (oe *InfixExpression) String() string {
	var out bytes.Buffer

	out.WriteString("(")
	out.WriteString(oe.Left.String())
	out.WriteString(" " + oe.Operator + " ")
	out.WriteString(oe.Right.String())
	out.WriteString(")")

	return out.String()
}

func (b *Boolean) expressionNode()      {}
func (b *Boolean) TokenLiteral() string { return b.Token.Literal }
func (b *Boolean) String() string       { return b.Token.Literal }

func (i *Identifier) expressionNode()      {}
func (i *Identifier) TokenLiteral() string { return i.Token.Literal }

func (pe *PrefixExpression) expressionNode()      {}
func (pe *PrefixExpression) TokenLiteral() string { return pe.Token.Literal }

func (bs *BlockStatement) statementNode()       {}
func (bs *BlockStatement) TokenLiteral() string { return bs.Token.Literal }
func (bs *BlockStatement) String() string {
	var out bytes.Buffer

	for _, s := range bs.Statements {
		out.WriteString(s.String())
	}

	return out.String()
}

func (fl *FunctionLiteral) expressionNode()      {}
func (fl *FunctionLiteral) TokenLiteral() string { return fl.Token.Literal }
func (fl *FunctionLiteral) String() string {
	var out bytes.Buffer

	params := []string{}

	for _, p := range fl.Parameters {
		params = append(params, p.String())
	}

	out.WriteString(fl.TokenLiteral())
	out.WriteString("(")
	out.WriteString(strings.Join(params, ", "))
	out.WriteString(") ")
	out.WriteString(fl.Body.String())
	return out.String()
}

func (pe *PrefixExpression) String() string {
	var out bytes.Buffer

	out.WriteString("(")
	out.WriteString(pe.Operator)
	out.WriteString(pe.Right.String())
	out.WriteString(")")

	return out.String()
}

func (ie *IfExpression) expressionNode()      {}
func (ie *IfExpression) TokenLiteral() string { return ie.Token.Literal }
func (ie *IfExpression) String() string {
	var out bytes.Buffer

	out.WriteString("if")
	out.WriteString(ie.Condition.String())

	out.WriteString(" ")
	out.WriteString(ie.Consequence.String())

	if ie.Alternative != nil {
		out.WriteString("else ")
		out.WriteString(ie.Alternative.String())
	}

	return out.String()
}

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

type ArrayLiteral struct {
	Token    Token.Token
	Elements []Expression
}

func (al *ArrayLiteral) expressionNode()      {}
func (al *ArrayLiteral) TokenLiteral() string { return al.Token.Literal }
func (al *ArrayLiteral) String() string {
	var out bytes.Buffer

	elements := []string{}

	for _, el := range al.Elements {
		elements = append(elements, el.String())
	}

	out.WriteString("[")
	out.WriteString(strings.Join(elements, ", "))
	out.WriteString("]")

	return out.String()
}

type HashLiteral struct {
	Token Token.Token // the '{' token
	Pairs map[Expression]Expression
}

func (hl *HashLiteral) expressionNode()      {}
func (hl *HashLiteral) TokenLiteral() string { return hl.Token.Literal }
func (hl *HashLiteral) String() string {
	var out bytes.Buffer
	pairs := []string{}
	for key, value := range hl.Pairs {
		pairs = append(pairs, key.String()+":"+value.String())
	}

	out.WriteString("{")
	out.WriteString(strings.Join(pairs, ", "))
	out.WriteString("}")
	return out.String()
}

// WhileExpression represents an `while` expression and holds the condition,
// and consequence expression
type WhileExpression struct {
	Token       Token.Token // The 'while' token
	Condition   Expression
	Consequence *BlockStatement
}

func (we *WhileExpression) expressionNode() {}

// TokenLiteral prints the literal value of the token associated with this node
func (we *WhileExpression) TokenLiteral() string { return we.Token.Literal }

// String returns a stringified version of the AST for debugging
func (we *WhileExpression) String() string {
	var out bytes.Buffer

	out.WriteString("while")
	out.WriteString(we.Condition.String())
	out.WriteString(" ")
	out.WriteString(we.Consequence.String())

	return out.String()
}
