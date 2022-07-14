package parser

import (
	"duck/ling/lexer/token"
	"duck/ling/parser/ast"
	"duck/ling/syntax/literal"
	"duck/ling/syntax/unicode"
	"fmt"
)

var Pos = 0
var CurrentToken *token.Token = nil
var CurrentChar rune
var Text string

func Peek() rune {
	peekPos := Pos + 1
	if peekPos > len(Text)-1 {
		return 0
	}
	return rune(Text[peekPos])
}

func SkipWhitespace() {
	for CurrentChar != 0 && unicode.IsWhitespace(CurrentChar) {
		Advance()
	}
}

func GetNextToken() *token.Token {
	for CurrentChar != 0 {
		if unicode.IsWhitespace(CurrentChar) {
			SkipWhitespace()
			continue
		}

		if literal.IsAlpha(CurrentChar) || CurrentChar == '_' {
			return ParseIdentifier()
		}

		if CurrentChar == '=' {
			Advance()
			return token.NewToken(token.ASSIGN, "=")
		}

		if CurrentChar == ';' {
			Advance()
			return token.NewToken(token.SEMICOLON, ";")
		}

		if literal.IsDecimalDigit(CurrentChar) || CurrentChar == '.' {
			return token.NewToken(token.NUMBER, ParseNumber())
		}

		if CurrentChar == '+' {
			Advance()
			return token.NewToken(token.ADD, "+")
		}

		if CurrentChar == '-' {
			Advance()
			return token.NewToken(token.SUB, "-")
		}

		if CurrentChar == '*' {
			Advance()
			return token.NewToken(token.MUL, "*")
		}

		if CurrentChar == '/' {
			Advance()
			return token.NewToken(token.DIV, "/")
		}

		if CurrentChar == '(' {
			Advance()
			return token.NewToken(token.LPAREN, "(")
		}

		if CurrentChar == ')' {
			Advance()
			return token.NewToken(token.RPAREN, ")")
		}
	}
	return token.NewToken(token.EOF, 0)
}

func eat(tokenType token.Type) {
	if tokenType == CurrentToken.Type {
		CurrentToken = GetNextToken()
	}
}

func Advance() {
	Pos++
	if Pos > len(Text)-1 {
		CurrentChar = 0
	} else {
		CurrentChar = rune(Text[Pos])
	}
}

func Visit(root *ast.ASTNode) {
	if root == nil {
		return
	}
	Visit((*root).Left)
	Visit((*root).Right)
	fmt.Println((*root).Token)
}
