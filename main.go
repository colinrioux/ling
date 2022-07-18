package main

import (
	"fmt"
	"ling/parser"
)

func main() {
	//psr := parser.NewParserString("var x = \"hi\"")
	psr := parser.NewParserString("{}")
	fmt.Println(psr.Lexer.Text)
	fmt.Printf("Len of text: %d\n", len(psr.Lexer.Text))
	//fmt.Println(psr.CurrentToken)
	tree := psr.Parse()
	if tree != nil {
		fmt.Printf("Tree: %v ; Visit Result:%v\n", tree.String(), tree.Visit())
	}

	//nde := parser.expression()
	//fmt.Println(nde.Visit())
	//fmt.Printf("%v bytes\n", unsafe.Sizeof(nde))
}
