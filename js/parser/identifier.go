package parser

import (
	"duck/ling/js/ast/keyword"
	"duck/ling/js/ast/token"
	"duck/ling/js/syntax/literal"
)

func ParseIdentifier() *token.Token {
	result := ""
	for CurrentChar != 0 && literal.IsAlphaNumeric(CurrentChar) {
		result += string(CurrentChar)
		Advance()
	}

	ok, tok := keyword.IsKeyword(result)
	if !ok {
		tok = token.NewToken(token.IDENTIFIER, result)
	}
	return tok
}
