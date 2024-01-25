package main

import (
	"fmt"
	"io"
	"os"
	"os/user"

	"github.com/3DRX/monkey/evaluator"
	"github.com/3DRX/monkey/lexer"
	"github.com/3DRX/monkey/object"
	"github.com/3DRX/monkey/parser"
	"github.com/3DRX/monkey/repl"
)

func main() {
	args := os.Args
	if len(args) == 1 { // REPL mode
		user, err := user.Current()
		if err != nil {
			panic(err)
		}
		fmt.Printf("Hello %s! This is the Monkey programming language!\n", user.Username)
		fmt.Printf("Feel free to type in commands\n")
		repl.Start(os.Stdin, os.Stdout)
	} else if len(args) == 2 { // File mode
		file, err := os.OpenFile(args[1], os.O_RDONLY, 0644)
		if err != nil {
			fmt.Printf("%s\n", err.Error())
			os.Exit(1)
		}
		defer file.Close()
		buf := make([]byte, 1024)
		currentString := ""
		for {
			n, err := file.Read(buf)
			if err == io.EOF {
				break
			}
			if err != nil {
				fmt.Printf("%s\n", err.Error())
				os.Exit(1)
			}
			if n > 0 {
				currentString += string(buf[:n])
			}
		}
		l := lexer.New(currentString)
		p := parser.New(l)
		program := p.ParseProgram()
		if len(p.Errors()) != 0 {
			repl.PrintParserErrors(os.Stdout, p.Errors())
			os.Exit(1)
		}
		env := object.NewEnvironment()
		evaluated := evaluator.Eval(program, env)
		if evaluated != nil {
			io.WriteString(os.Stdout, evaluated.Inspect())
			io.WriteString(os.Stdout, "\n")
		}
	} else {
		fmt.Println("Usage: monkey [filename]")
	}
}
