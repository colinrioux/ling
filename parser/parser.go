package parser

import (
	"duck/ling/lexer/token"
	"duck/ling/parser/ast"
	"duck/ling/syntax/literal"
	"duck/ling/syntax/unicode"
	"fmt"
)

// Parser :
// A type to represent an ECMA parser.
type Parser struct {
	Pos          int
	CurrentToken *token.Token
	CurrentChar  rune
	Text         string
}

// NewParser :
// Create a new ECMA parser.
func NewParser(text string) *Parser {
	return &Parser{
		Pos:          0,
		CurrentToken: nil,
		CurrentChar:  rune(text[0]),
		Text:         text,
	}
}

func (parser *Parser) parseWhitespace() {
	for parser.CurrentChar != 0 && unicode.IsWhitespace(parser.CurrentChar) {
		parser.advance()
	}
}

func (parser *Parser) parseNextToken() *token.Token {
	for parser.CurrentChar != 0 {
		if unicode.IsWhitespace(parser.CurrentChar) {
			parser.parseWhitespace()
			continue
		}

		// Identifiers
		if literal.IsAlpha(parser.CurrentChar) || parser.CurrentChar == '_' {
			return parser.parseIdentifier()
		}

		if parser.CurrentChar == '=' {
			parser.advance()
			return token.NewToken(token.ASSIGN, "=")
		}

		if parser.CurrentChar == ';' {
			parser.advance()
			return token.NewToken(token.SEMICOLON, ";")
		}

		// Numbers
		if literal.IsDecimalDigit(parser.CurrentChar) || parser.CurrentChar == '.' {
			// TODO have parseNumber return a token instead.
			return token.NewToken(token.NUMBER, parser.parseNumber())
		}

		if parser.CurrentChar == '+' {
			parser.advance()
			return token.NewToken(token.ADD, "+")
		}

		if parser.CurrentChar == '-' {
			parser.advance()
			return token.NewToken(token.SUB, "-")
		}

		if parser.CurrentChar == '*' {
			parser.advance()
			return token.NewToken(token.MUL, "*")
		}

		if parser.CurrentChar == '/' {
			parser.advance()
			return token.NewToken(token.DIV, "/")
		}

		if parser.CurrentChar == '(' {
			parser.advance()
			return token.NewToken(token.LPAREN, "(")
		}

		if parser.CurrentChar == ')' {
			parser.advance()
			return token.NewToken(token.RPAREN, ")")
		}
	}
	return token.NewToken(token.EOF, 0)
}

// Start :
// Start parsing the text.
func (parser *Parser) Start() {

}

// eat :
// Parse the next token.
func (parser *Parser) eat(tokenType token.Type) {
	if tokenType == parser.CurrentToken.Type {
		parser.CurrentToken = parser.parseNextToken()
	}
}

// advance :
// Parse the next character in Text.
func (parser *Parser) advance() {
	parser.Pos++
	if parser.Pos > len(parser.Text)-1 {
		parser.CurrentChar = 0
	} else {
		parser.CurrentChar = rune(parser.Text[parser.Pos])
	}
}

// peek :
// Parse the next character in Text without incrementing the iterator.
func (parser *Parser) peek() rune {
	peekPos := parser.Pos + 1
	if peekPos > len(parser.Text)-1 {
		return 0
	}
	return rune(parser.Text[peekPos])
}

func Visit(root *ast.Node) {
	if root == nil {
		return
	}
	Visit((*root).Left)
	Visit((*root).Right)
	fmt.Println((*root).Token)
}
