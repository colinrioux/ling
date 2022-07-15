package main

import (
	"duck/ling/parser"
	"fmt"
)

func main() {
	//parser.Text = "null"
	psr := parser.NewParser("var x = \"hi\"")
	fmt.Println(psr.Text)
	fmt.Println(psr.CurrentToken)

	//nde := parser.expression()
	//fmt.Println(nde.Visit())
	//fmt.Printf("%v bytes\n", unsafe.Sizeof(nde))
}
