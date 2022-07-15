package parser

import (
	"duck/ling/lexer"
	"duck/ling/lexer/token"
	"duck/ling/parser/ast"
	"fmt"
)

// Parser :
// A type to represent an ECMA parser.
type Parser struct {
	CurrentToken *token.Token
	Lexer        *lexer.Lexer
}

// NewParser :
// Create a new ECMA parser.
func NewParser(text string) *Parser {
	return &Parser{
		CurrentToken: nil,
		Lexer:        lexer.NewLexer(text),
	}
}

// eat :
// Instruct the parser to eat the CurrentToken if it matches tokenType and get the next token from the lexer.
func (parser *Parser) eat(tokenType token.Type) {
	if tokenType == parser.CurrentToken.Type {
		parser.CurrentToken = parser.Lexer.GetNextToken()
	}
}

// GetNext :
// Get the next token.
func (parser *Parser) GetNext() {
	parser.CurrentToken = parser.Lexer.GetNextToken()
}

func Visit(root *ast.Node) {
	if root == nil {
		return
	}
	Visit((*root).Left)
	Visit((*root).Right)
	fmt.Println((*root).Token)
}
