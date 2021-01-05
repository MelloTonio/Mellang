package Evaluator

import "github.com/Mellotonio/Andrei_lang/Object"

var builtins = map[string]*Object.Builtin{
	"len": &Object.Builtin{
		Fn: func(args ...Object.Object) Object.Object {
			if len(args) != 1 {
				return newError("wrong number of arguments. got=%d, want=1", len(args))
			}

			switch arg := args[0].(type) {
			case *Object.String:
				return &Object.Integer{Value: int64(len(arg.Value))}
			default:
				return newError("argument to `len` not supported, got %s", args[0].Type())
			}
		},
	},
	"Benicio": &Object.Builtin{
		Fn: func(args ...Object.Object) Object.Object {
			return &Object.String{Value: "Benicio está assistindo disney +"}
		},
	},
}
