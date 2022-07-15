package lexer

import (
	"duck/ling/lexer/token"
	"duck/ling/syntax/literal"
	"strconv"
)

// getNumber :
// Get a token corresponding to a number.
// TODO BIG INT
func (lexer *Lexer) getNumber() *token.Token {
	var res string = ""
	for lexer.CurrentChar != 0 && (literal.IsDecimalDigit(lexer.CurrentChar) ||
		lexer.CurrentChar == '.' || literal.IsScientific(lexer.CurrentChar, lexer.peek())) {

		if literal.IsScientific(lexer.CurrentChar, lexer.peek()) {
			res += string(lexer.CurrentChar)
			lexer.advance()
		}

		res += string(lexer.CurrentChar)
		lexer.advance()
	}
	resI, _ := strconv.ParseFloat(res, 64)

	return token.NewToken(token.NUMBER, resI)
}
