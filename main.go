package main

import (
	"fmt"
	"ling/parser"
	"log"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	psr := parser.NewParserFile("test.js")
	tree := psr.Parse()
	if tree != nil {
		fmt.Printf("Tree: %v ; Visit Result:%v\n", tree.String(), tree.Visit())
	}

	//nde := parser.expression()
	//fmt.Println(nde.Visit())
	//fmt.Printf("%v bytes\n", unsafe.Sizeof(nde))
}
