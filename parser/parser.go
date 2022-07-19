package parser

import (
	"fmt"
	"ling/lexer"
	"ling/lexer/token"
	"ling/parser/ast"
)

// Parser :
// A type to represent an ECMA parser.
type Parser struct {
	CurrentToken *token.Token
	Lexer        *lexer.Lexer
}

// NewParserString :
// Create a new ECMA parser, that takes input as a string.
func NewParserString(text string) *Parser {
	var p = &Parser{
		Lexer: lexer.NewLexerString(text),
	}
	p.CurrentToken, _ = p.Lexer.GetNextToken()
	return p
}

// NewParserFile :
// Create a new ECMA parser, that takes input as a file.
func NewParserFile(fileName string) *Parser {
	var p = &Parser{
		Lexer: lexer.NewLexerFile(fileName),
	}
	p.CurrentToken, _ = p.Lexer.GetNextToken()
	return p
}

// eat :
// Instruct the parser to eat the CurrentToken if it matches tokenType and get the next token from the lexer.
func (parser *Parser) eat(tokenType token.Type) {
	if tokenType == parser.CurrentToken.Type {
		parser.CurrentToken, _ = parser.Lexer.GetNextToken()
	}
}

// Parse :
// Get the next token.
func (parser *Parser) Parse() *ast.Node {
	return parser.parseProgram()
}

// parseProgram :
// Parses a program.
// A program starts as a global statement list.
func (parser *Parser) parseProgram() *ast.Node {
	nodes := parser.parseStatementList()

	// Build the AST node for the global block
	root := ast.NewBlockNode1("GLOBAL")
	for _, node := range nodes {
		root.Children = append(root.Children, node)
	}

	return (*ast.Node)(root)
}

func Visit(root *ast.Node) {
	if root == nil {
		return
	}
	for _, child := range (*root).Children {
		Visit(child)
	}
	//Visit((*root).Children[0])
	//Visit((*root).Children[1])
	fmt.Println((*root).Token)
}
