package main

import (
	"duck/ling/js/parser"
	"fmt"
)

func main() {
	parser.Text = "null"
	//parser.Text = "var x = \"hi\""
	fmt.Println(parser.Text)
	parser.CurrentChar = rune(parser.Text[parser.Pos])
	parser.CurrentToken = parser.GetNextToken()
	fmt.Println(parser.CurrentToken)

	//nde := parser.Expr()
	//fmt.Println(nde.Visit())
	//fmt.Printf("%v bytes\n", unsafe.Sizeof(nde))
}
