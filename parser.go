package main

import "fmt"

type Parser struct {
	sc *Scanner
}

func NewParser(sc *Scanner) *Parser {
	return &Parser{sc}
}

// /* Grammar: */
// Node -> Char | Pair
// Char -> [a-z]
// Pair -> LPAREN . Node . Space . Node . RPAREN
//
// Legend: '.' (dot) is the concatination operator.

func (p *Parser) Parse() *Node {
	tt := p.sc.Peek()

	switch tt {

	case CHAR_TOKEN:
		return p.parseChar()

	case LPAREN_TOKEN:
		return p.parsePair()

	default:
		return NewErrorNode(fmt.Sprintf("Expected a ( or Char[a-z]. Error at column %d\n", p.sc.head_index+1))
	}
}

// Char -> [a-z]
func (p *Parser) parseChar() *Node {
	t := p.sc.Next()

	if t.kind != CHAR_TOKEN {
		return NewErrorNode(fmt.Sprintf("Expected a Char[a-z]. Error at column %d\n", p.sc.head_index))
	}
	return NewCharNode(t)
}

// Node -> Char | Pair
// Char -> [a-z]
// Pair -> LPAREN . Node . Space . Node . RPAREN
func (p *Parser) parsePair() *Node {

	t := p.sc.Next() // lparen
	if t.kind != LPAREN_TOKEN {
		return NewErrorNode(fmt.Sprintf("Expected a ( . Error at column %d\n", p.sc.head_index))
	}

	lnode := p.Parse() // left side of the pair, recursive decent

	t = p.sc.Next() // space
	if t.kind != SPACE_TOKEN {
		return NewErrorNode(fmt.Sprintf("Expected a space. Error at column %d\n", p.sc.head_index))
	}

	rnode := p.Parse() // right side of the pair, recursive decent

	t = p.sc.Next() // rparen
	if t.kind != RPAREN_TOKEN {
		return NewErrorNode(fmt.Sprintf("Expected a ) . Error at column %d\n", p.sc.head_index))
	}

	return NewPairNode(*lnode, *rnode)
}
