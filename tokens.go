package main

type tokenType int

const (
	_ tokenType = iota
	CHAR_TOKEN
	LPAREN_TOKEN
	SPACE_TOKEN
	RPAREN_TOKEN
	UNKNOWN_TOKEN
)

type Token struct {
	kind   tokenType
	lexeme rune
}
