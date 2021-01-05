package Evaluator

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/Mellotonio/Andrei_lang/Object"
)

var builtins = map[string]*Object.Builtin{
	"len": &Object.Builtin{
		Fn: func(args ...Object.Object) Object.Object {
			if len(args) != 1 {
				return newError("wrong number of arguments. got=%d, want=1", len(args))
			}
			switch arg := args[0].(type) {
			case *Object.String:
				return &Object.Integer{Value: int64(len(arg.Value))}
			case *Object.Array:
				return &Object.Integer{Value: int64(len(arg.Elements))}
			default:
				return newError("argument to `len` not supported, got %s", args[0].Type())
			}
		},
	},
	"Benicio": &Object.Builtin{
		Fn: func(args ...Object.Object) Object.Object {
			return &Object.String{Value: "Não foi possivel aniquilar lisbete, tente novamente mais tarde..."}
		},
	},
	"Stefano": &Object.Builtin{
		Fn: func(args ...Object.Object) Object.Object {
			if len(args) != 1 {
				return newError("wrong number of arguments. got=%d, want=3", len(args))
			}

			s1 := rand.NewSource(time.Now().UnixNano())
			r1 := rand.New(s1)

			array_assert := []string{"Verdadeira", "Falsa"}

			assert_1 := r1.Intn(1)

			s := fmt.Sprintf("A pessoa %s é com toda certeza -> %s\nE possui um total de %d de QI", args[0].Inspect(), array_assert[assert_1], r1.Intn(120))

			return &Object.String{Value: s}
		},
	},
	"first": &Object.Builtin{
		Fn: func(args ...Object.Object) Object.Object {
			if len(args) != 1 {
				return newError("wrong number of arguments. got=%d, want=1", len(args))
			}
			if args[0].Type() != Object.ARRAY_OBJ {
				return newError("Argument to first must be array. got=%s, want=Array", args[0].Type())
			}

			arr := args[0].(*Object.Array)

			if len(arr.Elements) > 0 {
				return arr.Elements[0]
			}

			return NULL
		},
	},
	"last": &Object.Builtin{
		Fn: func(args ...Object.Object) Object.Object {
			if len(args) != 1 {
				return newError("wrong number of arguments. got=%d, want=1", len(args))
			}
			if args[0].Type() != Object.ARRAY_OBJ {
				return newError("Argument to first must be array. got=%s, want=Array", args[0].Type())
			}

			arr := args[0].(*Object.Array)
			length := len(arr.Elements)

			if len(arr.Elements) > 0 {
				return arr.Elements[length-1]
			}

			return NULL
		},
	},
	"rest": &Object.Builtin{
		Fn: func(args ...Object.Object) Object.Object {
			if len(args) != 1 {
				return newError("wrong number of arguments. got=%d, want=1", len(args))
			}
			if args[0].Type() != Object.ARRAY_OBJ {
				return newError("Argument to first must be array. got=%s, want=Array", args[0].Type())
			}

			arr := args[0].(*Object.Array)
			length := len(arr.Elements)

			if len(arr.Elements) > 0 {
				newElements := make([]Object.Object, length-1, length-1)
				copy(newElements, arr.Elements[1:length])
				return &Object.Array{Elements: newElements}
			}

			return NULL
		},
	},
	"push": &Object.Builtin{
		Fn: func(args ...Object.Object) Object.Object {
			if len(args) != 2 {
				return newError("wrong number of arguments. got=%d, want=1", len(args))
			}
			if args[0].Type() != Object.ARRAY_OBJ {
				return newError("Argument to first must be array. got=%s, want=Array", args[0].Type())
			}

			arr := args[0].(*Object.Array)
			length := len(arr.Elements)

			newElements := make([]Object.Object, length+1, length+1)
			copy(newElements, arr.Elements)
			newElements[length] = args[1]

			return &Object.Array{Elements: newElements}

		},
	},
}
