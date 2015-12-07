package main

type Token struct {
	ttype string
	value interface{}
}

func NewToken(ttype string, value interface{})(*Token) {
	return &Token{ttype, value}
}

