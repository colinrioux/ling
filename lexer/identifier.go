package lexer

import (
	"ling/lexer/keyword"
	"ling/lexer/token"
	"ling/syntax/literal"
)

// getIdentifierOrKeyword :
// Get a token corresponding to an identifier or keyword.
func (lexer *Lexer) getIdentifierOrKeyword() *token.Token {
	result := ""
	for lexer.CurrentChar != 0 && literal.IsAlphaNumeric(lexer.CurrentChar) {
		result += string(lexer.CurrentChar)
		lexer.advance(1)
	}

	k1 := keyword.GetReservedKeyword(result)
	if k1 != keyword.RKNOTFOUND {
		return token.NewToken(token.KEYWORD, k1)
	}

	k2 := keyword.GetStrictReservedKeyword(result)
	if k2 != keyword.SRKNOTFOUND {
		return token.NewToken(token.KEYWORD, k2)
	}

	return token.NewToken(token.IDENTIFIER, result)
}
