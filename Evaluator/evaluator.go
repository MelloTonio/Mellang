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
	case *AST.InfixExpression:
		left := Eval(node.Left)
		right := Eval(node.Right)

		return evalInfixExpression(node.Operator, left, right)
	case *AST.BlockStatement:
		return evalStatements(node.Statements)
	case *AST.IfExpression:
		return evalIfExpression(node)
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

func evalInfixExpression(operator string, left, right Object.Object) Object.Object {
	switch {
	case left.Type() == Object.INTEGER_OBJ && right.Type() == Object.INTEGER_OBJ:
		return evalIntegerInfixExpression(operator, left, right)
	case operator == "==":
		return nativeBoolToBooleanObject(left == right)
	case operator == "!=":
		return nativeBoolToBooleanObject(left != right)
	default:
		return NULL
	}
}

func evalIntegerInfixExpression(operator string, left Object.Object, right Object.Object) Object.Object {
	leftVal := left.(*Object.Integer).Value
	rightVal := right.(*Object.Integer).Value

	switch operator {
	case "+":
		return &Object.Integer{Value: leftVal + rightVal}
	case "-":
		return &Object.Integer{Value: leftVal - rightVal}
	case "*":
		return &Object.Integer{Value: leftVal * rightVal}
	case "/":
		return &Object.Integer{Value: leftVal / rightVal}
	case "<":
		return nativeBoolToBooleanObject(leftVal < rightVal)
	case ">":
		return nativeBoolToBooleanObject(leftVal > rightVal)
	case "==":
		return nativeBoolToBooleanObject(leftVal == rightVal)
	case "!=":
		return nativeBoolToBooleanObject(leftVal != rightVal)
	default:
		return NULL
	}
}

func evalIfExpression(ie *AST.IfExpression) Object.Object {
	condition := Eval(ie.Condition)

	if isTruthy(condition) {
		return Eval(ie.Consequence)
	} else if ie.Alternative != nil {
		return Eval(ie.Alternative)
	} else {
		return NULL
	}
}

func isTruthy(obj Object.Object) bool {
	switch obj {
	case NULL:
		return false
	case TRUE:
		return true
	case FALSE:
		return false
	default:
		return true
	}
}
