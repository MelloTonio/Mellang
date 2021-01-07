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
		env.Set(node.Name.Value, val)
	case *AST.BindExpression:
		val := Eval(node.Value, env)
		evalBindExpressions(node.Left, val, env)
	case *AST.Identifier:
		return evalIdentifier(node, env)
	case *AST.FunctionLiteral:
		body := node.Body
		parameters := node.Parameters
		return &Object.Function{Parameters: parameters, Env: env, Body: body}
	case *AST.CallExpression:
		function := Eval(node.Function, env)
		if isError(function) {
			return function
		}
		args := evalExpressions(node.Arguments, env)

		if len(args) == 1 && isError(args[0]) {
			return args[0]
		}
		return applyFunction(function, args)
	case *AST.StringLiteral:
		return &Object.String{Value: node.Value}
	case *AST.ArrayLiteral:
		elements := evalExpressions(node.Elements, env)
		if len(elements) == 1 && isError(elements[0]) {
			return elements[0]
		}
		return &Object.Array{Elements: elements}
	// Infix [
	case *AST.IndexExpression:
		left := Eval(node.Left, env)
		if isError(left) {
			return left
		}
		index := Eval(node.Index, env)
		if isError(index) {
			return index
		}
		return evalIndexExpression(left, index)
	case *AST.LiteralFloat:
		return &Object.Float{Value: node.Value}
	case *AST.HashLiteral:
		return evalHashLiteral(node, env)
	case *AST.WhileExpression:
		return evalWhileExpression(node, env)
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
	case left.Type() == Object.FLOAT && right.Type() == Object.FLOAT:
		return evalFloatInfixExpression(operator, left, right)
	case left.Type() == Object.INTEGER_OBJ && right.Type() == Object.FLOAT:
		return evalIntLeftFloatRight(operator, left, right)
	case left.Type() == Object.FLOAT && right.Type() == Object.INTEGER_OBJ:
		return evalIntRightFloatLeft(operator, left, right)
	case operator == "==":
		return nativeBoolToBooleanObject(left == right)
	case operator == "!=":
		return nativeBoolToBooleanObject(left != right)
	case left.Type() != right.Type():
		return newError("type mismatch: %s %s %s", left.Type(), operator, right.Type())
	case left.Type() == Object.STRING && right.Type() == Object.STRING:
		return evalStringInfixExpression(operator, left, right)
	default:
		return newError("unknown operator: %s %s %s", left.Type(), operator, right.Type())
	}
}

func evalIntegerInfixExpression(operator string, left Object.Object, right Object.Object) Object.Object {
	leftVal := left.(*Object.Integer).Value
	rightVal := right.(*Object.Integer).Value

	switch operator {
	// Non-bool
	case "+":
		return &Object.Integer{Value: leftVal + rightVal}
	case "-":
		return &Object.Integer{Value: leftVal - rightVal}
	case "*":
		return &Object.Integer{Value: leftVal * rightVal}
	case "/":
		return &Object.Integer{Value: leftVal / rightVal}
	// Bool
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
func evalFloatInfixExpression(operator string, left Object.Object, right Object.Object) Object.Object {
	leftVal := left.(*Object.Float).Value
	rightVal := right.(*Object.Float).Value

	switch operator {
	// Non-bool
	case "+":
		return &Object.Float{Value: leftVal + rightVal}
	case "-":
		return &Object.Float{Value: leftVal - rightVal}
	case "*":
		return &Object.Float{Value: leftVal * rightVal}
	case "/":
		return &Object.Float{Value: leftVal / rightVal}
	// Bool
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

func evalIntLeftFloatRight(operator string, left Object.Object, right Object.Object) Object.Object {
	leftVal := left.(*Object.Integer).Value
	rightVal := right.(*Object.Float).Value

	switch operator {
	// Non-bool
	case "+":
		return &Object.Float{Value: float64(leftVal) + rightVal}
	case "-":
		return &Object.Float{Value: float64(leftVal) - rightVal}
	case "*":
		return &Object.Float{Value: float64(leftVal) * rightVal}
	case "/":
		return &Object.Float{Value: float64(leftVal) / rightVal}
	// Bool
	case "<":
		return nativeBoolToBooleanObject(float64(leftVal) < rightVal)
	case ">":
		return nativeBoolToBooleanObject(float64(leftVal) > rightVal)
	case "==":
		return nativeBoolToBooleanObject(float64(leftVal) == rightVal)
	case "!=":
		return nativeBoolToBooleanObject(float64(leftVal) != rightVal)
	default:
		return newError("unknown operator: %s %s %s", left.Type(), operator, right.Type())
	}
}

func evalIntRightFloatLeft(operator string, left Object.Object, right Object.Object) Object.Object {
	leftVal := left.(*Object.Float).Value
	rightVal := right.(*Object.Integer).Value

	switch operator {
	// Non-bool
	case "+":
		return &Object.Float{Value: leftVal + float64(rightVal)}
	case "-":
		return &Object.Float{Value: leftVal - float64(rightVal)}
	case "*":
		return &Object.Float{Value: leftVal * float64(rightVal)}
	case "/":
		return &Object.Float{Value: leftVal / float64(rightVal)}
	// Bool
	case "<":
		return nativeBoolToBooleanObject(leftVal < float64(rightVal))
	case ">":
		return nativeBoolToBooleanObject(leftVal > float64(rightVal))
	case "==":
		return nativeBoolToBooleanObject(leftVal == float64(rightVal))
	case "!=":
		return nativeBoolToBooleanObject(leftVal != float64(rightVal))
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

func evalHashLiteral(node *AST.HashLiteral, env *Object.Environment) Object.Object {
	pairs := make(map[Object.HashKey]Object.HashPair)

	for keyNode, valueNode := range node.Pairs {
		key := Eval(keyNode, env)
		if isError(key) {
			return key
		}
		hashKey, ok := key.(Object.Hashable)
		if !ok {
			return newError("Unusable as hash key: %s", key.Type())
		}
		value := Eval(valueNode, env)
		if isError(value) {
			return value
		}
		hashed := hashKey.HashKey()
		pairs[hashed] = Object.HashPair{Key: key, Value: value}
	}
	return &Object.Hash{Pairs: pairs}
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

func evalBindExpressions(name string, val Object.Object, env *Object.Environment) {
	env.Set(name, val)
}

func evalIdentifier(node *AST.Identifier, env *Object.Environment) Object.Object {

	if val, ok := env.Get(node.Value); ok {
		return val
	}

	if builtin, ok := builtins[node.Value]; ok {
		return builtin
	}
	return newError("identifier not found: " + node.Value)
}

func evalExpressions(exps []AST.Expression, env *Object.Environment) []Object.Object {
	var result []Object.Object

	for _, e := range exps {
		evaluated := Eval(e, env)
		if isError(evaluated) {
			return []Object.Object{evaluated}
		}
		result = append(result, evaluated)
	}

	return result
}

// Cria um novo environment para aquela função, uma especie de escopo aonde as variaveis se mantém
func applyFunction(fn Object.Object, args []Object.Object) Object.Object {

	switch fn := fn.(type) {
	case *Object.Function:
		extendedEnv := extendedFunctionEnv(fn, args)
		evaluated := Eval(fn.Body, extendedEnv)
		return unwrapReturnValue(evaluated)
	case *Object.Builtin:
		return fn.Fn(args...)
	default:
		return newError("Not a function: %s", fn.Type())
	}

}

func extendedFunctionEnv(fn *Object.Function, args []Object.Object) *Object.Environment {
	env := Object.NewEnclosedEnvironment(fn.Env)

	for paramID, param := range fn.Parameters {
		// Para cada parametro da função, guardamos dentro desse novo escopo o valor correspondente a ele
		env.Set(param.Value, args[paramID])
	}
	return env
}

// Se encontrar um retorno, não retorne "return" e sim o valor dele, para não parar outras execuções
func unwrapReturnValue(obj Object.Object) Object.Object {
	if returnValue, ok := obj.(*Object.ReturnValue); ok {
		return returnValue.Value
	}

	return obj
}

func evalStringInfixExpression(operator string, left, right Object.Object) Object.Object {
	if operator != "+" {
		return newError("Unknown operator: %s %s %s", left.Type(), operator, right.Type())
	}

	leftVal := left.(*Object.String).Value
	rightVal := right.(*Object.String).Value

	return &Object.String{Value: leftVal + rightVal}
}

func evalIndexExpression(left, index Object.Object) Object.Object {
	switch {
	// array[x]
	case left.Type() == Object.ARRAY_OBJ && index.Type() == Object.INTEGER_OBJ:
		return evalArrayIndexExpression(left, index)
	case left.Type() == Object.HASH_OBJ:
		return evalHashIndexExpression(left, index)
	default:
		return newError("index operator not supported: %s", left.Type())
	}
}

func evalArrayIndexExpression(array, index Object.Object) Object.Object {
	arrayObject := array.(*Object.Array)

	idx := index.(*Object.Integer).Value
	max := int64(len(arrayObject.Elements) - 1)

	if idx < 0 || idx > max {
		return NULL
	}

	return arrayObject.Elements[idx]
}

func evalHashIndexExpression(hash, index Object.Object) Object.Object {
	hashObject := hash.(*Object.Hash)
	key, ok := index.(Object.Hashable)
	if !ok {
		return newError("unusable as hash key: %s", index.Type())
	}
	pair, ok := hashObject.Pairs[key.HashKey()]
	if !ok {
		return NULL
	}
	return pair.Value
}

func evalWhileExpression(we *AST.WhileExpression, env *Object.Environment) Object.Object {
	var result Object.Object

	for {
		condition := Eval(we.Condition, env)
		if isError(condition) {
			return condition
		}

		if isTruthy(condition) {
			result = Eval(we.Consequence, env)
		} else {
			break
		}
	}

	return result

}
