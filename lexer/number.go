package lexer

import (
	"ling/lexer/token"
	"ling/syntax/literal"
	"strconv"
)

// getNumber :
// Get a token corresponding to a number.
// TODO BIG INT
func (lexer *Lexer) getNumber() (*token.Token, int) {
	var res string = ""
	advanceAmount := 0
	for lexer.CurrentChar != 0 && (literal.IsDecimalDigit(lexer.CurrentChar) ||
		lexer.CurrentChar == '.' || literal.IsScientific(lexer.CurrentChar, lexer.peek(1))) {

		if literal.IsScientific(lexer.CurrentChar, lexer.peek(1)) {
			res += string(lexer.CurrentChar)
			advanceAmount += 1
			lexer.advance(advanceAmount)
		}

		res += string(lexer.CurrentChar)
		advanceAmount += 1
		lexer.advance(advanceAmount)
	}
	resI, _ := strconv.ParseFloat(res, 64)

	return token.NewToken(token.NUMBER, resI), advanceAmount
}
