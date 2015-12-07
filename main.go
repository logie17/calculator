package main

import (
	"fmt"
	"os"
	"bufio"
	"strconv"
	"log"
)

var Null byte = '\x00'

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


type Token struct {
	ttype string
	value interface{}
}

func NewToken(ttype string, value interface{})(*Token) {
	return &Token{ttype, value}
}

type Lexer struct {
	text string
	pos int
	currentChar byte
}

// NewLexer returns a lexer
func NewLexer(s string)(*Lexer) {
	return &Lexer{s, 0, s[0]}
}

func (l *Lexer) Integer() int {
	result := ""
	for l.currentChar != Null && l.IsNumber() {
		result = result + string(l.currentChar)
		l.Advance()
	}
	i, _ := strconv.Atoi(result)
	return i
}

func (l *Lexer) Error() {
	log.Fatal("Invalid Character!")
}

func (l *Lexer) GetNextToken() (*Token) {
	for l.currentChar != Null {
		if l.currentChar == ' ' {
			l.SkipWhiteSpace()
			continue
		}

		if l.IsNumber() {
			return NewToken("INTEGER", l.Integer())
		}

		if l.currentChar == '+' {
			l.Advance()
			return NewToken("PLUS", "+")
		}

		if l.currentChar == '-' {
			l.Advance()
			return NewToken("MINUS", "-")
		}

		if l.currentChar == '*' {
			l.Advance()
			return NewToken("MUL", "*")
		}
		
		if l.currentChar == '/' {
			l.Advance()
			return NewToken("DIV", "/")
		}

		if l.currentChar == '(' {
			l.Advance()
			return NewToken("LPAREN", "(")
		}

		if l.currentChar == ')' {
			l.Advance()
			return NewToken("RPAREN", ")")
		}

		l.Error()
	}
	return NewToken("EOF","Nil")
}

func (l *Lexer) IsNumber() bool {
	return l.currentChar == '0' || l.currentChar == '1' || l.currentChar == '2' || l.currentChar == '3' || l.currentChar == '4' || l.currentChar == '5' || l.currentChar == '6' || l.currentChar == '7' || l.currentChar == '8' || l.currentChar == '9'
}

func(l *Lexer) SkipWhiteSpace() {
	for l.currentChar != Null && l.currentChar == ' ' {
		l.Advance()
	}
}

func (l *Lexer) Advance() {
	l.pos = l.pos + 1

	if l.pos == (len(l.text) - 1) {
		l.currentChar = Null
	} else {
		l.currentChar = l.text[l.pos]
	}
}

type Interpreter struct {
	lexer *Lexer
	currentToken *Token
}

func NewInterpreter(l *Lexer) (*Interpreter) {
	return &Interpreter{l, l.GetNextToken()}
}

func (i *Interpreter) Error() {
	log.Fatal("Invalid Syntax")
}
func (i *Interpreter) Eat(ttype string) {
	if i.currentToken.ttype == ttype {
		i.currentToken = i.lexer.GetNextToken()
	} else {
		i.Error()
	}
}

func (i *Interpreter) Factor() int {
	token := i.currentToken

	if token.ttype == "INTEGER" {
		i.Eat("INTEGER")
		return token.value.(int)
	} else if token.ttype == "LPAREN" {
		i.Eat("LPAREN")
		result := i.Expression()
		i.Eat("RPAREN")
		return result
	}
	i.Error()
	return 0
	
}

func (i *Interpreter) Term() int {
	result := i.Factor()

	for i.currentToken.ttype == "MUL" || i.currentToken.ttype == "DIV" {
		token := i.currentToken

		if token.ttype == "MUL" {
			i.Eat("MUL")
			result = result * i.Factor()
		} else if token.ttype == "DIV" {
			i.Eat("DIV")
			result = result / i.Factor()
		}
		
	}

	return result

}

func (i *Interpreter) Expression() (int) {

	result := i.Term()
	for i.currentToken.ttype == "PLUS" || i.currentToken.ttype == "MINUS" {
		token := i.currentToken

		if token.ttype == "PLUS" {
			i.Eat("PLUS")
			result = result + i.Term()
		} else if token.ttype == "MINUS" {
			i.Eat("MINUS")
			result = result - i.Term()
		}
		
	}
	
	return result

}
