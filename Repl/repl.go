package Repl

import (
	"bufio"
	"fmt"
	"io"

	"github.com/Mellotonio/Andrei_lang/Lexer"
	"github.com/Mellotonio/Andrei_lang/Parser"
)

// READ EVAL PRINT LOOP

const PROMPT = ">> "

func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)

	for {
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

		io.WriteString(out, program.String())
		io.WriteString(out, "\n")

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
