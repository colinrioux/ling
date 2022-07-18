package lexer

import (
	"ling/lexer/token"
	"ling/syntax/literal"
	"strconv"
)

// getNumber :
// Get a token corresponding to a number.
// TODO BIG INT
func (lexer *Lexer) getNumber() *token.Token {
	var res string = ""
	for lexer.CurrentChar != 0 && (literal.IsDecimalDigit(lexer.CurrentChar) ||
		lexer.CurrentChar == '.' || literal.IsScientific(lexer.CurrentChar, lexer.peek(1))) {

		if literal.IsScientific(lexer.CurrentChar, lexer.peek(1)) {
			res += string(lexer.CurrentChar)
			lexer.advance(1)
		}

		res += string(lexer.CurrentChar)
		lexer.advance(1)
	}
	resI, _ := strconv.ParseFloat(res, 64)

	return token.NewToken(token.NUMBER, resI)
}
