package parser

import (
	"duck/ling/lexer"
	"duck/ling/parser/ast"
	"fmt"
)

// Parser :
// A type to represent an ECMA parser.
type Parser struct {
	Lexer *lexer.Lexer
}

// NewParser :
// Create a new ECMA parser.
func NewParser(text string) *Parser {
	return &Parser{
		Lexer: lexer.NewLexer(text),
	}
}

func Visit(root *ast.Node) {
	if root == nil {
		return
	}
	Visit((*root).Left)
	Visit((*root).Right)
	fmt.Println((*root).Token)
}
