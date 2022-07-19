package lexer

import (
	"bufio"
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
func (lexer *Lexer) GetNextToken() (*token.Token, int) {
	advanceAmount := 0
	for lexer.CurrentChar != 0 {
		if unicode.IsWhitespace(lexer.CurrentChar) {
			advanceAmount += lexer.whitespace()
			continue
		}

		// Identifiers & Keywords
		if literal.IsAlpha(lexer.CurrentChar) || lexer.CurrentChar == '_' {
			return lexer.getIdentifierOrKeyword()
		}

		// ASSIGN
		if lexer.CurrentChar == '=' {
			advanceAmount += 1
			lexer.advance(advanceAmount)

			// EQUALITY
			if lexer.CurrentChar == '=' {
				advanceAmount += 1
				lexer.advance(advanceAmount)

				// STRICT EQUALITY
				if lexer.CurrentChar == '=' {
					advanceAmount += 1
					lexer.advance(advanceAmount)
					return token.NewToken(token.STRICT_EQUALITY, "==="), advanceAmount
				}

				return token.NewToken(token.EQUALITY, "=="), advanceAmount
			}

			return token.NewToken(token.ASSIGN, "="), advanceAmount
		}

		// SEMICOLON
		if lexer.CurrentChar == ';' {
			advanceAmount += 1
			lexer.advance(advanceAmount)
			return token.NewToken(token.SEMICOLON, ";"), advanceAmount
		}

		// COMMA
		if lexer.CurrentChar == ',' {
			advanceAmount += 1
			lexer.advance(advanceAmount)
			return token.NewToken(token.COMMA, ","), advanceAmount
		}

		// DOT
		if lexer.CurrentChar == '.' && !literal.IsDecimalDigit(lexer.peek(1)) {
			advanceAmount += 1
			lexer.advance(advanceAmount)

			// SPREAD
			if lexer.CurrentChar == '.' && lexer.peek(1) == '.' {
				advanceAmount += 2
				lexer.advance(advanceAmount)
				return token.NewToken(token.SPREAD, "..."), advanceAmount
			}

			return token.NewToken(token.DOT, "."), advanceAmount
		}

		// Numbers
		if literal.IsDecimalDigit(lexer.CurrentChar) || lexer.CurrentChar == '.' {
			return lexer.getNumber()
		}

		// ADDITION
		if lexer.CurrentChar == '+' {
			advanceAmount += 1
			lexer.advance(advanceAmount)

			// INCREMENT
			if lexer.CurrentChar == '+' {
				advanceAmount += 1
				lexer.advance(advanceAmount)
				return token.NewToken(token.INCREMENT, "++"), advanceAmount
			}

			// ADDITION ASSIGN
			if lexer.CurrentChar == '=' {
				advanceAmount += 1
				lexer.advance(advanceAmount)
				return token.NewToken(token.ADDITION_ASSIGN, "+="), advanceAmount
			}

			return token.NewToken(token.ADDITION, "+"), advanceAmount
		}

		// SUBTRACTION
		if lexer.CurrentChar == '-' {
			advanceAmount += 1
			lexer.advance(advanceAmount)

			// DECREMENT
			if lexer.CurrentChar == '-' {
				advanceAmount += 1
				lexer.advance(advanceAmount)
				return token.NewToken(token.DECREMENT, "--"), advanceAmount
			}

			// SUBTRACTION ASSIGN
			if lexer.CurrentChar == '=' {
				advanceAmount += 1
				lexer.advance(advanceAmount)
				return token.NewToken(token.SUBTRACTION_ASSIGN, "-="), advanceAmount
			}

			return token.NewToken(token.SUBTRACTION, "-"), advanceAmount
		}

		// MULTIPLICATION
		if lexer.CurrentChar == '*' {
			advanceAmount += 1
			lexer.advance(advanceAmount)

			// EXPONENT
			if lexer.CurrentChar == '*' {
				advanceAmount += 1
				lexer.advance(advanceAmount)

				// EXPONENT ASSIGN
				if lexer.CurrentChar == '=' {
					advanceAmount += 1
					lexer.advance(advanceAmount)
					return token.NewToken(token.EXPONENT_ASSIGN, "**="), advanceAmount
				}

				return token.NewToken(token.EXPONENT, "**"), advanceAmount
			}

			// RIGHT MULTI LINE COMMENT
			if lexer.CurrentChar == '/' {
				advanceAmount += 1
				lexer.advance(advanceAmount)
				return token.NewToken(token.RMULTI_LINE_COMMENT, "*/"), advanceAmount
			}

			// MULTIPLICATION ASSIGN
			if lexer.CurrentChar == '=' {
				advanceAmount += 1
				lexer.advance(advanceAmount)
				return token.NewToken(token.MULTIPLY_ASSIGN, "*="), advanceAmount
			}

			return token.NewToken(token.MULTIPLICATION, "*"), advanceAmount
		}

		// DIVISION
		if lexer.CurrentChar == '/' {
			advanceAmount += 1
			lexer.advance(advanceAmount)

			// SINGLE LINE COMMENT
			if lexer.CurrentChar == '/' {
				advanceAmount += 1
				lexer.advance(advanceAmount)
				return token.NewToken(token.SINGLE_LINE_COMMENT, "//"), advanceAmount
			}

			// LEFT MULTI LINE COMMENT
			if lexer.CurrentChar == '*' {
				advanceAmount += 1
				lexer.advance(advanceAmount)
				return token.NewToken(token.LMULTI_LINE_COMMENT, "/*"), advanceAmount
			}

			// DIVISION ASSIGN
			if lexer.CurrentChar == '=' {
				advanceAmount += 1
				lexer.advance(advanceAmount)
				return token.NewToken(token.DIVISION_ASSIGN, "/="), advanceAmount
			}

			return token.NewToken(token.DIVISION, "/"), advanceAmount
		}

		// LESS THAN
		if lexer.CurrentChar == '<' {
			advanceAmount += 1
			lexer.advance(advanceAmount)

			// LEFT SHIFT
			if lexer.CurrentChar == '<' {
				advanceAmount += 1
				lexer.advance(advanceAmount)

				// LEFT SHIFT ASSIGN
				if lexer.CurrentChar == '=' {
					advanceAmount += 1
					lexer.advance(advanceAmount)
					return token.NewToken(token.LSHIFT_ASSIGN, "<<="), advanceAmount
				}

				return token.NewToken(token.LSHIFT, "<<"), advanceAmount
			}

			// LESS THAN OR EQUAL
			if lexer.CurrentChar == '=' {
				advanceAmount += 1
				lexer.advance(advanceAmount)
				return token.NewToken(token.LESS_THAN_OR_EQUAL, "<="), advanceAmount
			}

			return token.NewToken(token.LESS_THAN, "<"), advanceAmount
		}

		// GREATER THAN
		if lexer.CurrentChar == '>' {
			advanceAmount += 1
			lexer.advance(advanceAmount)

			// RIGHT SHIFT
			if lexer.CurrentChar == '>' {
				advanceAmount += 1
				lexer.advance(advanceAmount)

				// RIGHT SHIFT ASSIGN
				if lexer.CurrentChar == '=' {
					advanceAmount += 1
					lexer.advance(advanceAmount)
					return token.NewToken(token.RSHIFT_ASSIGN, ">>="), advanceAmount
				}

				// UNSIGNED RIGHT SHIFT
				if lexer.CurrentChar == '>' {
					advanceAmount += 1
					lexer.advance(advanceAmount)

					// UNSIGNED RIGHT SHIFT ASSIGN
					if lexer.CurrentChar == '=' {
						advanceAmount += 1
						lexer.advance(advanceAmount)
						return token.NewToken(token.URSHIFT_ASSIGN, ">>>="), advanceAmount
					}

					return token.NewToken(token.URSHIFT, ">>>"), advanceAmount
				}

				return token.NewToken(token.RSHIFT, ">>"), advanceAmount
			}

			// GREATER THAN OR EQUAL
			if lexer.CurrentChar == '=' {
				advanceAmount += 1
				lexer.advance(advanceAmount)
				return token.NewToken(token.GREATER_THAN_OR_EQUAL, ">="), advanceAmount
			}

			return token.NewToken(token.GREATER_THAN, ">"), advanceAmount
		}

		// MODULO
		if lexer.CurrentChar == '%' {
			advanceAmount += 1
			lexer.advance(advanceAmount)

			// MODULO ASSIGN
			if lexer.CurrentChar == '=' {
				advanceAmount += 1
				lexer.advance(advanceAmount)

				return token.NewToken(token.MODULO_ASSIGN, "%="), advanceAmount
			}

			return token.NewToken(token.MODULO, "%"), advanceAmount
		}

		// BITWISE AND
		if lexer.CurrentChar == '&' {
			advanceAmount += 1
			lexer.advance(advanceAmount)

			// LOGICAL AND
			if lexer.CurrentChar == '&' {
				advanceAmount += 1
				lexer.advance(advanceAmount)

				// LOGICAL AND ASSIGN
				if lexer.CurrentChar == '=' {
					advanceAmount += 1
					lexer.advance(advanceAmount)
					return token.NewToken(token.LOGICAL_AND_ASSIGN, "&&="), advanceAmount
				}

				return token.NewToken(token.LOGICAL_AND, "&&"), advanceAmount
			}

			// BITWISE AND ASSIGN
			if lexer.CurrentChar == '=' {
				advanceAmount += 1
				lexer.advance(advanceAmount)
				return token.NewToken(token.BITWISE_AND_ASSIGN, "&="), advanceAmount
			}

			return token.NewToken(token.BITWISE_AND, "&"), advanceAmount
		}

		// BITWISE OR
		if lexer.CurrentChar == '|' {
			advanceAmount += 1
			lexer.advance(advanceAmount)

			// LOGICAL OR
			if lexer.CurrentChar == '|' {
				advanceAmount += 1
				lexer.advance(advanceAmount)

				// LOGICAL OR ASSIGN
				if lexer.CurrentChar == '=' {
					advanceAmount += 1
					lexer.advance(advanceAmount)
					return token.NewToken(token.LOGICAL_OR_ASSIGN, "||="), advanceAmount
				}

				return token.NewToken(token.LOGICAL_OR, "||"), advanceAmount
			}

			// BITWISE OR ASSIGN
			if lexer.CurrentChar == '=' {
				advanceAmount += 1
				lexer.advance(advanceAmount)
				return token.NewToken(token.BITWISE_OR_ASSIGN, "|="), advanceAmount
			}

			return token.NewToken(token.BITWISE_OR, "|"), advanceAmount
		}

		// BITWISE XOR
		if lexer.CurrentChar == '^' {
			advanceAmount += 1
			lexer.advance(advanceAmount)

			// BITWISE XOR ASSIGN
			if lexer.CurrentChar == '=' {
				advanceAmount += 1
				lexer.advance(advanceAmount)
				return token.NewToken(token.BITWISE_XOR_ASSIGN, "^="), advanceAmount
			}

			return token.NewToken(token.BITWISE_XOR, "|"), advanceAmount
		}

		// LOGICAL NOT
		if lexer.CurrentChar == '!' {
			advanceAmount += 1
			lexer.advance(advanceAmount)

			// INEQUALITY
			if lexer.CurrentChar == '=' {
				advanceAmount += 1
				lexer.advance(advanceAmount)

				// STRICT INEQUALITY
				if lexer.CurrentChar == '=' {
					advanceAmount += 1
					lexer.advance(advanceAmount)
					return token.NewToken(token.STRICT_INEQUALITY, "!=="), advanceAmount
				}

				return token.NewToken(token.INEQUALITY, "!="), advanceAmount
			}

			return token.NewToken(token.LOGICAL_NOT, "!"), advanceAmount
		}

		// NULLISH
		if lexer.CurrentChar == '?' && lexer.peek(1) == '?' {
			advanceAmount += 2
			lexer.advance(advanceAmount)

			// LOGICAL NULLISH ASSIGN
			if lexer.CurrentChar == '=' {
				advanceAmount += 1
				lexer.advance(advanceAmount)
				return token.NewToken(token.LOGICAL_NULLISH_ASSIGN, "??="), advanceAmount
			}

			return token.NewToken(token.NULLISH, "??"), advanceAmount
		}

		// OPTIONAL CHAINING
		if lexer.CurrentChar == '?' && lexer.peek(1) == '.' {
			advanceAmount += 2
			lexer.advance(advanceAmount)
			return token.NewToken(token.OPTIONAL_CHAINING, "?."), advanceAmount
		}

		// BITWISE NOT
		if lexer.CurrentChar == '~' {
			advanceAmount += 1
			lexer.advance(advanceAmount)
			return token.NewToken(token.BITWISE_NOT, "~"), advanceAmount
		}

		if lexer.CurrentChar == '(' {
			advanceAmount += 1
			lexer.advance(advanceAmount)
			return token.NewToken(token.LPAREN, "("), advanceAmount
		}

		if lexer.CurrentChar == ')' {
			advanceAmount += 1
			lexer.advance(advanceAmount)
			return token.NewToken(token.RPAREN, ")"), advanceAmount
		}

		if lexer.CurrentChar == '[' {
			advanceAmount += 1
			lexer.advance(advanceAmount)
			return token.NewToken(token.LBRACKET, "["), advanceAmount
		}

		if lexer.CurrentChar == ']' {
			advanceAmount += 1
			lexer.advance(advanceAmount)
			return token.NewToken(token.RBRACKET, "]"), advanceAmount
		}

		if lexer.CurrentChar == '{' {
			advanceAmount += 1
			lexer.advance(advanceAmount)
			return token.NewToken(token.LBRACE, "{"), advanceAmount
		}

		if lexer.CurrentChar == '}' {
			advanceAmount += 1
			lexer.advance(advanceAmount)
			return token.NewToken(token.RBRACE, "}"), advanceAmount
		}

		if lexer.CurrentChar == '"' {
			advanceAmount += 1
			lexer.advance(advanceAmount)
			return token.NewToken(token.DOUBLE_QUOTE, "\""), advanceAmount
		}

		if lexer.CurrentChar == '\'' {
			advanceAmount += 1
			lexer.advance(advanceAmount)
			return token.NewToken(token.SINGLE_QUOTE, "'"), advanceAmount
		}

		if lexer.CurrentChar == '`' {
			advanceAmount += 1
			lexer.advance(advanceAmount)
			return token.NewToken(token.BACKTICK, "`"), advanceAmount
		}

		if lexer.CurrentChar == '#' {
			advanceAmount += 1
			lexer.advance(advanceAmount)
			return token.NewToken(token.POUND, "#"), advanceAmount
		}

		if lexer.CurrentChar == '$' {
			advanceAmount += 1
			lexer.advance(advanceAmount)
			return token.NewToken(token.DOLLAR, "$"), advanceAmount
		}
	}
	return token.NewToken(token.EOF, 0), advanceAmount
}

// whitespace :
// Skips any whitespace picked up by the lexer.
func (lexer *Lexer) whitespace() int {
	advanceAmount := 0
	for lexer.CurrentChar != 0 && unicode.IsWhitespace(lexer.CurrentChar) {
		advanceAmount += 1
		lexer.advance(1)
	}
	return advanceAmount
}

// advance :
// Advance the lexer to the next dth character in Text.
func (lexer *Lexer) advance(d int) {
	lexer.Pos += d

	if lexer.IsFile {
		var r rune
		//var err error
		for i := 1; i <= d; i++ {
			r, _, _ = lexer.File.ReadRune()
			//fmt.Printf("HI COLIN %v\n", r)
			//if err != nil {
			//	// TODO error handling
			//	log.Fatal(fmt.Sprintf("%v", err))
			//}
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
		for i := 1; i <= d; i++ {
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
