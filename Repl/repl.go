package Repl

import (
	"bufio"
	"fmt"
	"io"

	"github.com/Mellotonio/Andrei_lang/Evaluator"
	"github.com/Mellotonio/Andrei_lang/Lexer"
	"github.com/Mellotonio/Andrei_lang/Object"
	"github.com/Mellotonio/Andrei_lang/Parser"
)

// READ EVAL PRINT LOOP

var Misterious *string

const PROMPT = ">> "

func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)
	env := Object.NewEnvironment()
	var cont = 0

	for {
		// Gambiarra que injeta um map como builtin
		if cont == 0 {
			MAP := `moonvar map = fn(arr,f){
						moonvar iter = fn(arr,accumulated){
							if(len(arr) == 0){
								accumulated
								} else {
								iter(rest(arr), push(accumulated, f(first(arr))));
							}
						};
					iter(arr,[]);
				};`

			l := Lexer.New(MAP)
			p := Parser.New(l)

			program := p.ParseProgram()

			Evaluator.Eval(program, env)
			cont += 1
		}

		fmt.Printf(PROMPT)
		scanned := scanner.Scan()
		if !scanned {
			return
		}

		line := scanner.Text()

		l := Lexer.New(line)
		p := Parser.New(l)

		program := p.ParseProgram()

		if len(p.Errors()) != 0 {
			printParserErrors(out, p.Errors())
			continue
		}

		evaluated := Evaluator.Eval(program, env)
		if evaluated != nil {
			io.WriteString(out, evaluated.Inspect())
			io.WriteString(out, "\n")
		}

		// Read user input, until encounter a new line
	}
}

const Mellus = `░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░
░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░
░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░
░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░
░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░
░░░░░░░░░░░░░░░█░░░░░░█░░░░░░░░░░░░░░░
░██████░░░░░░░░░░░░░░░░░░░░░░░███████░
░░░░░░██░░░░░░░░░░░░░░░░░░░░░██░░░░░░░
░░░░░░░██░░░░░░░██████░░░░░░█░░░░░░░░░
░░░░░░░░░██░░░░░░░░░░░░░░░░██░░░░░░░░░
░░░░░░░░░░███░░░░░░░░░░░░░██░░░░░░░░░░
░░░░░░░░░░░██░░░░░░░░░░░░██░░░░░░░░░░░
░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░
░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░
░░░░░░░░█░░░███░░░███░░███░░███░░░░░░░
░░░░░░░█░░░░█░█░░░█░█░░█░█░░█░█░░░░░░░
░░░░░░░██░░░███░░░███░░█░█░░███░░░░░░░
░░░░░░░█░░░░█░░█░░█░░█░█░█░░█░░█░░░░░░
░░░░░░░██░░░█░░█░░█░░█░███░░█░░█░░░░░░
░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░
░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░
░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░
░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░


`

func printParserErrors(out io.Writer, errors []string) {
	io.WriteString(out, Mellus)
	io.WriteString(out, "Woops! We ran into some monkey business here!\n")
	io.WriteString(out, " parser errors:\n")
	for _, msg := range errors {
		io.WriteString(out, "\t"+msg+"\n")
	}
}
