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
	var p *Parser = &Parser{
		Lexer: lexer.NewLexer(text),
	}
	p.CurrentToken = p.Lexer.GetNextToken()
	return p
}

// eat :
// Instruct the parser to eat the CurrentToken if it matches tokenType and get the next token from the lexer.
func (parser *Parser) eat(tokenType token.Type) {
	if tokenType == parser.CurrentToken.Type {
		parser.CurrentToken = parser.Lexer.GetNextToken()
	}
}

// Parse :
// Get the next token.
func (parser *Parser) Parse() *ast.Node {
	return parser.parseExpression()
}

func Visit(root *ast.Node) {
	if root == nil {
		return
	}
	Visit((*root).Children[0])
	Visit((*root).Children[1])
	fmt.Println((*root).Token)
}
