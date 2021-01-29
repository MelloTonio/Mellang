package main

import (
	"fmt"
	"os"
	"os/user"

	"github.com/Mellotonio/Andrei_lang/FileMatch"
	"github.com/Mellotonio/Andrei_lang/Repl"
)

func main() {
	var filename string
	//var matches []string

	user, err := user.Current()

	if err != nil {
		panic(err)
	}

	filename = FileMatch.FindFilename()

	fmt.Printf("File '%s' has been loaded\n\n", filename)

	fmt.Printf("Hello %s! This is the Mellang programming language!\n", user.Username)

	fmt.Printf("Feel free to type in commands\n")

	Repl.Start(os.Stdin, os.Stdout, filename)
}
