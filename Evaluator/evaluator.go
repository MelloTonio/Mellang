package Evaluator

import (
	"fmt"

	"github.com/Mellotonio/Andrei_lang/AST"
	"github.com/Mellotonio/Andrei_lang/Object"
)

var (
	TRUE  = &Object.Boolean{Value: true}
	FALSE = &Object.Boolean{Value: false}
	NULL  = &Object.Null{}
)

func Eval(node AST.Node, env *Object.Environment) Object.Object {
	switch node := node.(type) {
	case *AST.Program:
		return evalProgram(node, env)
	case *AST.ExpressionStatement:
		return Eval(node.Expression, env)
	case *AST.LiteralInteger:
		return &Object.Integer{Value: node.Value}
	case *AST.Boolean:
		return nativeBoolToBooleanObject(node.Value)
	case *AST.PrefixExpression:
		// Pega o numero ou o booleano no lado direto
		right := Eval(node.Right, env)
		if isError(right) {
			return right
		}
		return evalPrefixExpression(node.Operator, right)
	case *AST.InfixExpression:
		left := Eval(node.Left, env)
		if isError(left) {
			return left
		}

		right := Eval(node.Right, env)
		if isError(right) {
			return right
		}

		return evalInfixExpression(node.Operator, left, right)
	case *AST.BlockStatement:
		return evalBlockStatement(node, env)
	case *AST.IfExpression:
		return evalIfExpression(node, env)
	case *AST.ReturnStatement:
		val := Eval(node.ReturnValue, env)
		if isError(val) {
			return val
		}
		return &Object.ReturnValue{Value: val}
	case *AST.MoonvarStatement:
		val := Eval(node.Value, env)
		if isError(val) {
			return val
		}
	}
	return nil
}

func evalProgram(program *AST.Program, env *Object.Environment) Object.Object {
	var result Object.Object

	for _, statement := range program.Statements {
		result = Eval(statement, env)

		switch result := result.(type) {
		case *Object.ReturnValue:
			return result.Value
		case *Object.Error:
			return result
		}
	}

	// Encontrou um valor de retorno? retorna automaticamente.
	if returnValue, ok := result.(*Object.ReturnValue); ok {
		return returnValue.Value
	}

	return result
}

/*func evalStatements(statements []AST.Statement) Object.Object {
	var result Object.Object

	for _, statement := range statements {
		result = Eval(statement)
	}

	// Encontrou um valor de retorno? retorna automaticamente.
	if returnValue, ok := result.(*Object.ReturnValue); ok {
		return returnValue.Value
	}

	return result
}*/

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
		return newError("unknown operator: %s%s", operator, right.Type())
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
		return newError("unknown operator: -%s", right.Type())
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
	case left.Type() != right.Type():
		return newError("type mismatch: %s %s %s", left.Type(), operator, right.Type())
	default:
		return newError("unknown operator: %s %s %s", left.Type(), operator, right.Type())
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
		return newError("unknown operator: %s %s %s", left.Type(), operator, right.Type())
	}
}

func evalIfExpression(ie *AST.IfExpression, env *Object.Environment) Object.Object {
	condition := Eval(ie.Condition, env)
	if isError(condition) {
		return condition
	}

	if isTruthy(condition) {
		return Eval(ie.Consequence, env)
	} else if ie.Alternative != nil {
		return Eval(ie.Alternative, env)
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

func evalBlockStatement(block *AST.BlockStatement, env *Object.Environment) Object.Object {
	var result Object.Object

	for _, statement := range block.Statements {
		result = Eval(statement, env)

		if result != nil {
			rt := result.Type()
			if rt == Object.RETURN_VALUE_OBJ || rt == Object.ERROR_OBJ {
				return result
			}
		}
	}
	return result
}

func newError(format string, a ...interface{}) *Object.Error {
	return &Object.Error{Message: fmt.Sprintf(format, a...)}
}

func isError(obj Object.Object) bool {
	if obj != nil {
		return obj.Type() == Object.ERROR_OBJ
	}
	return false
}
