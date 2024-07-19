package main

import "slices"

var (
	TerminalChars = []rune{'(', ' ', ')'}
)

type Scanner struct {
	symbols_steam []rune
	head_index    int
}

func NewScanner(stream []rune, offset int) *Scanner {
	return &Scanner{
		symbols_steam: stream,
		head_index:    offset,
	}
}

func (s *Scanner) HasNext() bool {
	return s.head_index < len(s.symbols_steam)-1
}

func (s *Scanner) Peek() tokenType {
	n := s.symbols_steam[s.head_index+1]

	if (n < 'a' || n > 'z') && !slices.Contains(TerminalChars, n) {
		return UNKNOWN_TOKEN
	}
	switch n {
	case '(':
		return LPAREN_TOKEN
	case ' ':
		return SPACE_TOKEN

	case ')':
		return RPAREN_TOKEN

	default:
		return CHAR_TOKEN
	}
}

func (s *Scanner) Next() Token {
	tt := s.Peek()

	s.head_index++
	lexeme := s.symbols_steam[s.head_index]

	return Token{
		kind:   tt,
		lexeme: lexeme,
	}
}
