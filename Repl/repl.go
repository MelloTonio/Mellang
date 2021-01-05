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

const PROMPT = ">> "

const REDUCE = `moonvar reduce=fn(arr, initial, f){ moonvar iter = fn(arr,result){
	 if(len(arr)==0){ 
		result 
		} else {
			iter(rest(arr),f(result,first(arr)));}
		};
	iter(arr,initial);
	};`

const MAP = `moonvar map = fn(arr,f){
		moonvar iter = fn(arr,accumulated){
			if(len(arr) == 0){
				accumulated
				} else {
				iter(rest(arr), push(accumulated, f(first(arr))));
			}
		};
	iter(arr,[]);
};`

const SUM = `moonvar sum = fn(arr){reduce(arr, 0, fn(initial,el){ initial + el });};`

func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)
	env := Object.NewEnvironment()
	cont := 0
	injectBuiltin := []string{0: MAP, 1: REDUCE, 2: SUM}

	for {
		if cont == 0 {
			// Gambiarra que injeta funções builtin
			for i := 0; i <= len(injectBuiltin)-1; i++ {

				l := Lexer.New(injectBuiltin[i])
				p := Parser.New(l)

				program := p.ParseProgram()

				Evaluator.Eval(program, env)
			}
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

const Mellus = `
	▓█████  ██▀███   ██▀███   ▒█████   ██▀███  
	▓█   ▀ ▓██ ▒ ██▒▓██ ▒ ██▒▒██▒  ██▒▓██ ▒ ██▒
	▒███   ▓██ ░▄█ ▒▓██ ░▄█ ▒▒██░  ██▒▓██ ░▄█ ▒
	▒▓█  ▄ ▒██▀▀█▄  ▒██▀▀█▄  ▒██   ██░▒██▀▀█▄  
	░▒████▒░██▓ ▒██▒░██▓ ▒██▒░ ████▓▒░░██▓ ▒██▒
	░░ ▒░ ░░ ▒▓ ░▒▓░░ ▒▓ ░▒▓░░ ▒░▒░▒░ ░ ▒▓ ░▒▓░
	░ ░  ░  ░▒ ░ ▒░  ░▒ ░ ▒░  ░ ▒ ▒░   ░▒ ░ ▒░
	░     ░░   ░   ░░   ░ ░ ░ ░ ▒    ░░   ░ 
	░  ░   ░        ░         ░ ░     ░   

`

func printParserErrors(out io.Writer, errors []string) {
	io.WriteString(out, Mellus)
	io.WriteString(out, "Woops! We ran into some monkey business here!\n")
	io.WriteString(out, " parser errors:\n")
	for _, msg := range errors {
		io.WriteString(out, "\t"+msg+"\n")
	}
}
