package main

import (
	"fmt"
	"os"
	"bufio"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	text := ""
	for text != "q\n" {
		fmt.Print("Enter text: ")
		text, _ = reader.ReadString('\n')
		fmt.Println(text)

		lexer := NewLexer(text)
		interpreter := NewInterpreter(lexer)
		fmt.Println(interpreter.Expression())
	}
}


