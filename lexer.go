package main

import (
	"strconv"
	"log"
)

var Null byte = '\x00'


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

