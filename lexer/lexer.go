package lexer

import (
	"duck/ling/lexer/token"
	"duck/ling/syntax/literal"
	"duck/ling/syntax/unicode"
)

// Lexer :
// A type to represent an ECMA lexer. The lexer scans the
// input to produce tokens.
type Lexer struct {
	Pos         int
	CurrentChar rune
	Text        string
}

// NewLexer :
// Creates a new lexer.
func NewLexer(text string) *Lexer {
	return &Lexer{
		Pos:         0,
		CurrentChar: rune(text[0]),
		Text:        text,
	}
}

// GetNextToken :
// Get the next token in the Text.
func (lexer *Lexer) GetNextToken() *token.Token {
	for lexer.CurrentChar != 0 {
		if unicode.IsWhitespace(lexer.CurrentChar) {
			lexer.whitespace()
			continue
		}

		// Identifiers & Keywords
		if literal.IsAlpha(lexer.CurrentChar) || lexer.CurrentChar == '_' {
			return lexer.getIdentifierOrKeyword()
		}

		// TODO
		if lexer.CurrentChar == '=' {
			lexer.advance()
			return token.NewToken(token.ASSIGN, "=")
		}

		// TODO
		if lexer.CurrentChar == ';' {
			lexer.advance()
			return token.NewToken(token.SEMICOLON, ";")
		}

		// Numbers
		// TODO big int
		if literal.IsDecimalDigit(lexer.CurrentChar) || lexer.CurrentChar == '.' {
			return lexer.getNumber()
		}

		// TODO
		if lexer.CurrentChar == '+' {
			lexer.advance()
			return token.NewToken(token.ADD, "+")
		}

		if lexer.CurrentChar == '-' {
			lexer.advance()
			return token.NewToken(token.SUB, "-")
		}

		if lexer.CurrentChar == '*' {
			lexer.advance()
			return token.NewToken(token.MUL, "*")
		}

		if lexer.CurrentChar == '/' {
			lexer.advance()
			return token.NewToken(token.DIV, "/")
		}

		if lexer.CurrentChar == '(' {
			lexer.advance()
			return token.NewToken(token.LPAREN, "(")
		}

		if lexer.CurrentChar == ')' {
			lexer.advance()
			return token.NewToken(token.RPAREN, ")")
		}
	}
	return token.NewToken(token.EOF, 0)
}

// whitespace :
// Skips any whitespace picked up by the lexer.
func (lexer *Lexer) whitespace() {
	for lexer.CurrentChar != 0 && unicode.IsWhitespace(lexer.CurrentChar) {
		lexer.advance()
	}
}

// advance :
// Advance the lexer to the next character in Text.
func (lexer *Lexer) advance() {
	lexer.Pos++
	if lexer.Pos > len(lexer.Text)-1 {
		lexer.CurrentChar = 0
	} else {
		lexer.CurrentChar = rune(lexer.Text[lexer.Pos])
	}
}

// peek :
// Get the next character in Text without incrementing the iterator.
func (lexer *Lexer) peek() rune {
	peekPos := lexer.Pos + 1
	if peekPos > len(lexer.Text)-1 {
		return 0
	}
	return rune(lexer.Text[peekPos])
}
