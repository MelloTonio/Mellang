package Evaluator

import (
	"github.com/Mellotonio/Andrei_lang/AST"
	"github.com/Mellotonio/Andrei_lang/Object"
)

var (
	TRUE  = &Object.Boolean{Value: true}
	FALSE = &Object.Boolean{Value: false}
	NULL  = &Object.Null{}
)

func Eval(node AST.Node) Object.Object {
	switch node := node.(type) {
	case *AST.Program:
		return evalStatements(node.Statements)
	case *AST.ExpressionStatement:
		return Eval(node.Expression)
	case *AST.LiteralInteger:
		return &Object.Integer{Value: node.Value}
	case *AST.Boolean:
		return nativeBoolToBooleanObject(node.Value)
	case *AST.PrefixExpression:
		// Pega o numero ou o booleano no lado direto
		right := Eval(node.Right)
		return evalPrefixExpression(node.Operator, right)
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

func nativeBoolToBooleanObject(input bool) *Object.Boolean {
	if input {
		return TRUE
	}
	return FALSE
}

func evalPrefixExpression(operator string, right Object.Object) Object.Object {
	switch operator {
	case "!":
		return evalBangOperatorExpression(right)
	case "-":
		return evalMinusOperatorExpression(right)
	default:
		return NULL
	}
}

func evalBangOperatorExpression(right Object.Object) Object.Object {
	switch right {
	case TRUE:
		return FALSE
	case FALSE:
		return TRUE
	case NULL:
		return FALSE
	default:
		return FALSE
	}
}

func evalMinusOperatorExpression(right Object.Object) Object.Object {
	if right.Type() != Object.INTEGER_OBJ {
		return NULL
	}

	value := right.(*Object.Integer).Value
	return &Object.Integer{Value: -value}
}
