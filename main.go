package main

import (
	"duck/ling/js/parser"
	"fmt"
)

func main() {
	parser.Text = "2 + 3"
	//parser.Text = "var x = \"hi\""
	fmt.Println(parser.Text)
	parser.CurrentChar = rune(parser.Text[parser.Pos])
	parser.CurrentToken = parser.GetNextToken()
	//fmt.Println(parser.CurrentToken)
	//parser.CurrentToken = parser.GetNextToken()
	//fmt.Println(parser.CurrentToken)
	//parser.CurrentToken = parser.GetNextToken()
	//fmt.Println(parser.CurrentToken)
	//parser.CurrentToken = parser.GetNextToken()
	//fmt.Println(parser.CurrentToken)
	nde := parser.Expr()
	//fmt.Println(nde)
	fmt.Println(nde.Visit())
}
