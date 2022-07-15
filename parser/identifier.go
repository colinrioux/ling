package parser

import (
	"duck/ling/lexer/token"
)

// parseIdentifier :
// Parses an ECMA identifier appropriately. Keywords are also managed here.
func (parser *Parser) parseIdentifier() *token.Token {
	result := ""
	//for parser.CurrentChar != 0 && literal.IsAlphaNumeric(parser.CurrentChar) {
	//	result += string(parser.CurrentChar)
	//	parser.advance()
	//}
	//
	//k1 := keyword.GetReservedKeyword(result)
	//if k1 != keyword.RKNOTFOUND {
	//	if k1 == keyword.VAR {
	//		// If we have a var keyword, we want to parse a variable statement.
	//		return parser.parseVariableStatement()
	//	}
	//	return token.NewToken(token.KEYWORD, k1)
	//}
	//
	//k2 := keyword.GetStrictReservedKeyword(result)
	//if k2 != keyword.SRKNOTFOUND {
	//	return token.NewToken(token.KEYWORD, k2)
	//}

	return token.NewToken(token.IDENTIFIER, result)
}
