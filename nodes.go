package main

type nodeType int

const (
	_ nodeType = iota
	CHAR_NODE
	PAIR_NODE
	ERROR_NODE
)

type Node struct {
	kind nodeType
	val  any
}

type CharValue struct {
	char rune
}
type PairValue struct {
	left  Node
	right Node
}

func NewCharNode(token Token) *Node { // pointer escapes to heap
	return &Node{
		kind: CHAR_NODE,
		val: CharValue{
			char: token.lexeme,
		},
	}
}

func NewPairNode(left, right Node) *Node {
	return &Node{
		kind: PAIR_NODE,
		val:  PairValue{left, right},
	}
}

func NewErrorNode(reason string) *Node {
	return &Node{
		kind: ERROR_NODE,
		val:  reason,
	}
}
