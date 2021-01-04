package Evaluator

import (
	"github.com/Mellotonio/Andrei_lang/AST"
	"github.com/Mellotonio/Andrei_lang/Object"
)

func Eval(node AST.Node) Object.Object {
	switch node := node.(type) {
	case *AST.Program:
		return evalStatements(node.Statements)
	case *AST.ExpressionStatement:
		return Eval(node.Expression)
	case *AST.LiteralInteger:
		return &Object.Integer{Value: node.Value}
	}
	return nil
}

func evalStatements(statements []AST.Statement) Object.Object {
	var result Object.Object

	for _, statement := range statements {
		result = Eval(statement)
	}

	return result
}
