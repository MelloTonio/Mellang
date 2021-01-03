package main

import (
	"fmt"
	"os"
	"os/user"

	"github.com/Mellotonio/Andrei_lang/Repl"
)

func main() {
	user, err := user.Current()

	if err != nil {
		panic(err)
	}

	fmt.Printf("Hello %s! This is the Andrei programming language!\n", user.Username)

	fmt.Printf("Feel free to type in commands\n")

	Repl.Start(os.Stdin, os.Stdout)
}
