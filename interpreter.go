package main

import (
	"log"
)
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
