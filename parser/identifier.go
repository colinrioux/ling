package parser

import (
	"duck/ling/lexer/keyword"
	"duck/ling/lexer/token"
	"duck/ling/syntax/literal"
)

func ParseIdentifier() *token.Token {
	result := ""
	for CurrentChar != 0 && literal.IsAlphaNumeric(CurrentChar) {
		result += string(CurrentChar)
		Advance()
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
