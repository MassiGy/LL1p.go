package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	read_prompt_and_parse(
		func(exp string) {
			fmt.Println("Expression         : ", exp)
			fmt.Println("Tree representation: ")

			sc := NewScanner([]rune(exp), -1)
			parser := NewParser(sc)

			for sc.HasNext() {
				node := parser.Parse()
				visit(*node, 0)
			}
			fmt.Println(";")
		},
	)
}

func read_prompt_and_parse(parsefn func(exp string)) {
	promptScanner := bufio.NewScanner(os.Stdin)
	for promptScanner.Scan() {
		parsefn(promptScanner.Text())
	}
}

func visit(n Node, n_spaces int) {
	indent(n_spaces)

	switch n.kind {

	case CHAR_NODE:
		fmt.Printf("Char('%c')\n", n.val.(CharValue).char)

	case PAIR_NODE:
		fmt.Printf("Pair(\n")

		visit(n.val.(PairValue).left, n_spaces+2)
		visit(n.val.(PairValue).right, n_spaces+2)

		indent(n_spaces)
		fmt.Printf(")\n")

	case ERROR_NODE:
		fmt.Printf(n.val.(string) + "\n")
	}

}

func indent(n_spaces int) {
	for i := 0; i < n_spaces; i++ {
		fmt.Printf(" ")
	}
}
