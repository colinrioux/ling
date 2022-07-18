package lexer

import (
	"bufio"
	"fmt"
	"ling/lexer/token"
	"ling/syntax/literal"
	"ling/syntax/unicode"
	"log"
	"os"
	"unicode/utf8"
)

const chunkSize = 10                          // TODO intuitive chunkSize selection?
const cacheSize = chunkSize + (chunkSize / 2) // TODO can we do better?

// Lexer :
// A type to represent an ECMA lexer. The lexer scans the
// input to produce tokens.
type Lexer struct {
	Pos         int
	CurrentChar rune
	Text        string
	IsFile      bool
	File        *bufio.Reader
}

// NewLexerString :
// Creates a new lexer with string input.
func NewLexerString(text string) *Lexer {
	var lex = &Lexer{
		Pos:    0,
		Text:   text,
		IsFile: false,
	}
	if len(text) > 0 {
		lex.CurrentChar = rune(text[0])
	} else {
		lex.CurrentChar = 0
	}
	return lex
}

// NewLexerFile :
// Creates a new lexer with file input.
func NewLexerFile(fileName string) *Lexer {
	var lex = &Lexer{
		Pos:    0,
		IsFile: true,
	}

	f, err := os.Open(fileName)
	if err != nil {
		// TODO error handling
		log.Fatal(err)
		return nil
	}

	//defer f.Close()

	lex.File = bufio.NewReaderSize(f, cacheSize)
	lex.CurrentChar, _, _ = lex.File.ReadRune()
	return lex
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

		// ASSIGN
		if lexer.CurrentChar == '=' {
			lexer.advance(1)

			// EQUALITY
			if lexer.CurrentChar == '=' {
				lexer.advance(1)

				// STRICT EQUALITY
				if lexer.CurrentChar == '=' {
					lexer.advance(1)
					return token.NewToken(token.STRICT_EQUALITY, "===")
				}

				return token.NewToken(token.EQUALITY, "==")
			}

			return token.NewToken(token.ASSIGN, "=")
		}

		// SEMICOLON
		if lexer.CurrentChar == ';' {
			lexer.advance(1)
			return token.NewToken(token.SEMICOLON, ";")
		}

		// COMMA
		if lexer.CurrentChar == ',' {
			lexer.advance(1)
			return token.NewToken(token.COMMA, ",")
		}

		// DOT
		if lexer.CurrentChar == '.' && !literal.IsDecimalDigit(lexer.peek(1)) {
			lexer.advance(1)

			// SPREAD
			if lexer.CurrentChar == '.' && lexer.peek(1) == '.' {
				lexer.advance(2)
				return token.NewToken(token.SPREAD, "...")
			}

			return token.NewToken(token.DOT, ".")
		}

		// Numbers
		if literal.IsDecimalDigit(lexer.CurrentChar) || lexer.CurrentChar == '.' {
			return lexer.getNumber()
		}

		// ADDITION
		if lexer.CurrentChar == '+' {
			lexer.advance(1)

			// INCREMENT
			if lexer.CurrentChar == '+' {
				lexer.advance(1)
				return token.NewToken(token.INCREMENT, "++")
			}

			// ADDITION ASSIGN
			if lexer.CurrentChar == '=' {
				lexer.advance(1)
				return token.NewToken(token.ADDITION_ASSIGN, "+=")
			}

			return token.NewToken(token.ADDITION, "+")
		}

		// SUBTRACTION
		if lexer.CurrentChar == '-' {
			lexer.advance(1)

			// DECREMENT
			if lexer.CurrentChar == '-' {
				lexer.advance(1)
				return token.NewToken(token.DECREMENT, "--")
			}

			// SUBTRACTION ASSIGN
			if lexer.CurrentChar == '=' {
				lexer.advance(1)
				return token.NewToken(token.SUBTRACTION_ASSIGN, "-=")
			}

			return token.NewToken(token.SUBTRACTION, "-")
		}

		// MULTIPLICATION
		if lexer.CurrentChar == '*' {
			lexer.advance(1)

			// EXPONENT
			if lexer.CurrentChar == '*' {
				lexer.advance(1)

				// EXPONENT ASSIGN
				if lexer.CurrentChar == '=' {
					lexer.advance(1)
					return token.NewToken(token.EXPONENT_ASSIGN, "**=")
				}

				return token.NewToken(token.EXPONENT, "**")
			}

			// RIGHT MULTI LINE COMMENT
			if lexer.CurrentChar == '/' {
				lexer.advance(1)
				return token.NewToken(token.RMULTI_LINE_COMMENT, "*/")
			}

			// MULTIPLICATION ASSIGN
			if lexer.CurrentChar == '=' {
				lexer.advance(1)
				return token.NewToken(token.MULTIPLY_ASSIGN, "*=")
			}

			return token.NewToken(token.MULTIPLICATION, "*")
		}

		// DIVISION
		if lexer.CurrentChar == '/' {
			lexer.advance(1)

			// SINGLE LINE COMMENT
			if lexer.CurrentChar == '/' {
				lexer.advance(1)
				return token.NewToken(token.SINGLE_LINE_COMMENT, "//")
			}

			// LEFT MULTI LINE COMMENT
			if lexer.CurrentChar == '*' {
				lexer.advance(1)
				return token.NewToken(token.LMULTI_LINE_COMMENT, "/*")
			}

			// DIVISION ASSIGN
			if lexer.CurrentChar == '=' {
				lexer.advance(1)
				return token.NewToken(token.DIVISION_ASSIGN, "/=")
			}

			return token.NewToken(token.DIVISION, "/")
		}

		// LESS THAN
		if lexer.CurrentChar == '<' {
			lexer.advance(1)

			// LEFT SHIFT
			if lexer.CurrentChar == '<' {
				lexer.advance(1)

				// LEFT SHIFT ASSIGN
				if lexer.CurrentChar == '=' {
					lexer.advance(1)
					return token.NewToken(token.LSHIFT_ASSIGN, "<<=")
				}

				return token.NewToken(token.LSHIFT, "<<")
			}

			// LESS THAN OR EQUAL
			if lexer.CurrentChar == '=' {
				lexer.advance(1)
				return token.NewToken(token.LESS_THAN_OR_EQUAL, "<=")
			}

			return token.NewToken(token.LESS_THAN, "<")
		}

		// GREATER THAN
		if lexer.CurrentChar == '>' {
			lexer.advance(1)

			// RIGHT SHIFT
			if lexer.CurrentChar == '>' {
				lexer.advance(1)

				// RIGHT SHIFT ASSIGN
				if lexer.CurrentChar == '=' {
					lexer.advance(1)
					return token.NewToken(token.RSHIFT_ASSIGN, ">>=")
				}

				// UNSIGNED RIGHT SHIFT
				if lexer.CurrentChar == '>' {
					lexer.advance(1)

					// UNSIGNED RIGHT SHIFT ASSIGN
					if lexer.CurrentChar == '=' {
						lexer.advance(1)
						return token.NewToken(token.URSHIFT_ASSIGN, ">>>=")
					}

					return token.NewToken(token.URSHIFT, ">>>")
				}

				return token.NewToken(token.RSHIFT, ">>")
			}

			// GREATER THAN OR EQUAL
			if lexer.CurrentChar == '=' {
				lexer.advance(1)
				return token.NewToken(token.GREATER_THAN_OR_EQUAL, ">=")
			}

			return token.NewToken(token.GREATER_THAN, ">")
		}

		// MODULO
		if lexer.CurrentChar == '%' {
			lexer.advance(1)

			// MODULO ASSIGN
			if lexer.CurrentChar == '=' {
				lexer.advance(1)

				return token.NewToken(token.MODULO_ASSIGN, "%=")
			}

			return token.NewToken(token.MODULO, "%")
		}

		// BITWISE AND
		if lexer.CurrentChar == '&' {
			lexer.advance(1)

			// LOGICAL AND
			if lexer.CurrentChar == '&' {
				lexer.advance(1)

				// LOGICAL AND ASSIGN
				if lexer.CurrentChar == '=' {
					lexer.advance(1)
					return token.NewToken(token.LOGICAL_AND_ASSIGN, "&&=")
				}

				return token.NewToken(token.LOGICAL_AND, "&&")
			}

			// BITWISE AND ASSIGN
			if lexer.CurrentChar == '=' {
				lexer.advance(1)
				return token.NewToken(token.BITWISE_AND_ASSIGN, "&=")
			}

			return token.NewToken(token.BITWISE_AND, "&")
		}

		// BITWISE OR
		if lexer.CurrentChar == '|' {
			lexer.advance(1)

			// LOGICAL OR
			if lexer.CurrentChar == '|' {
				lexer.advance(1)

				// LOGICAL OR ASSIGN
				if lexer.CurrentChar == '=' {
					lexer.advance(1)
					return token.NewToken(token.LOGICAL_OR_ASSIGN, "||=")
				}

				return token.NewToken(token.LOGICAL_OR, "||")
			}

			// BITWISE OR ASSIGN
			if lexer.CurrentChar == '=' {
				lexer.advance(1)
				return token.NewToken(token.BITWISE_OR_ASSIGN, "|=")
			}

			return token.NewToken(token.BITWISE_OR, "|")
		}

		// BITWISE XOR
		if lexer.CurrentChar == '^' {
			lexer.advance(1)

			// BITWISE XOR ASSIGN
			if lexer.CurrentChar == '=' {
				lexer.advance(1)
				return token.NewToken(token.BITWISE_XOR_ASSIGN, "^=")
			}

			return token.NewToken(token.BITWISE_XOR, "|")
		}

		// LOGICAL NOT
		if lexer.CurrentChar == '!' {
			lexer.advance(1)

			// INEQUALITY
			if lexer.CurrentChar == '=' {
				lexer.advance(1)

				// STRICT INEQUALITY
				if lexer.CurrentChar == '=' {
					lexer.advance(1)
					return token.NewToken(token.STRICT_INEQUALITY, "!==")
				}

				return token.NewToken(token.INEQUALITY, "!=")
			}

			return token.NewToken(token.LOGICAL_NOT, "!")
		}

		// NULLISH
		if lexer.CurrentChar == '?' && lexer.peek(1) == '?' {
			lexer.advance(2)

			// LOGICAL NULLISH ASSIGN
			if lexer.CurrentChar == '=' {
				lexer.advance(1)
				return token.NewToken(token.LOGICAL_NULLISH_ASSIGN, "??=")
			}

			return token.NewToken(token.NULLISH, "??")
		}

		// OPTIONAL CHAINING
		if lexer.CurrentChar == '?' && lexer.peek(1) == '.' {
			lexer.advance(2)
			return token.NewToken(token.OPTIONAL_CHAINING, "?.")
		}

		// BITWISE NOT
		if lexer.CurrentChar == '~' {
			lexer.advance(1)
			return token.NewToken(token.BITWISE_NOT, "~")
		}

		if lexer.CurrentChar == '(' {
			lexer.advance(1)
			return token.NewToken(token.LPAREN, "(")
		}

		if lexer.CurrentChar == ')' {
			lexer.advance(1)
			return token.NewToken(token.RPAREN, ")")
		}

		if lexer.CurrentChar == '[' {
			lexer.advance(1)
			return token.NewToken(token.LBRACKET, "[")
		}

		if lexer.CurrentChar == ']' {
			lexer.advance(1)
			return token.NewToken(token.RBRACKET, "]")
		}

		if lexer.CurrentChar == '{' {
			lexer.advance(1)
			return token.NewToken(token.LBRACE, "{")
		}

		if lexer.CurrentChar == '}' {
			lexer.advance(1)
			return token.NewToken(token.RBRACE, "}")
		}

		if lexer.CurrentChar == '"' {
			lexer.advance(1)
			return token.NewToken(token.DOUBLE_QUOTE, "\"")
		}

		if lexer.CurrentChar == '\'' {
			lexer.advance(1)
			return token.NewToken(token.SINGLE_QUOTE, "'")
		}

		if lexer.CurrentChar == '`' {
			lexer.advance(1)
			return token.NewToken(token.BACKTICK, "`")
		}

		if lexer.CurrentChar == '#' {
			lexer.advance(1)
			return token.NewToken(token.POUND, "#")
		}

		if lexer.CurrentChar == '$' {
			lexer.advance(1)
			return token.NewToken(token.DOLLAR, "$")
		}
	}
	return token.NewToken(token.EOF, 0)
}

// whitespace :
// Skips any whitespace picked up by the lexer.
func (lexer *Lexer) whitespace() {
	for lexer.CurrentChar != 0 && unicode.IsWhitespace(lexer.CurrentChar) {
		lexer.advance(1)
	}
}

// advance :
// Advance the lexer to the next dth character in Text.
func (lexer *Lexer) advance(d int) {
	lexer.Pos += d

	if lexer.IsFile {
		var r rune
		var err error
		for i := 1; i < d; i++ {
			r, _, err = lexer.File.ReadRune()
			if err != nil {
				// TODO error handling
				log.Fatal(fmt.Sprintf("%v", err))
			}
		}
		lexer.CurrentChar = r
		return
	}

	if lexer.Pos > len(lexer.Text)-1 {
		lexer.CurrentChar = 0
	} else {
		lexer.CurrentChar = rune(lexer.Text[lexer.Pos])
	}
}

// peek :
// Get the dth character in Text from Pos without incrementing the iterator.
func (lexer *Lexer) peek(d int) rune {
	if lexer.IsFile {
		peekBytes := 0
		var dec rune
		var actualBytes int
		for i := 1; i < d; i++ {
			bytesRead, _ := lexer.File.Peek(actualBytes + 4) // peek max amount of utf8 bytes
			dec, actualBytes = utf8.DecodeRune(bytesRead)
			peekBytes += actualBytes
		}
		return dec
	}

	peekPos := lexer.Pos + d
	if peekPos > len(lexer.Text)-1 {
		return 0
	}
	return rune(lexer.Text[peekPos])
}
