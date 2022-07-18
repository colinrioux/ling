package main

import (
	"fmt"
	"ling/parser"
)

func main() {
	//psr := parser.NewParser("var x = \"hi\"")
	psr := parser.NewParser("(3 + 5) * 5")
	fmt.Println(psr.Lexer.Text)
	tree := psr.Parse()
	if tree != nil {
		fmt.Println(tree.Visit())
	}

	//nde := parser.expression()
	//fmt.Println(nde.Visit())
	//fmt.Printf("%v bytes\n", unsafe.Sizeof(nde))
}
