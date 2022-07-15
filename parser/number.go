package parser

import (
	"duck/ling/syntax/literal"
	"duck/ling/syntax/type"
	"strconv"
)

// parseNumber :
func (parser *Parser) parseNumber() *_type.ECMANumber {
	var res string = ""
	for parser.CurrentChar != 0 && (literal.IsDecimalDigit(parser.CurrentChar) ||
		parser.CurrentChar == '.' || literal.IsScientific(parser.CurrentChar, parser.peek())) {

		if literal.IsScientific(parser.CurrentChar, parser.peek()) {
			res += string(parser.CurrentChar)
			parser.advance()
		}

		res += string(parser.CurrentChar)
		parser.advance()
	}
	resI, _ := strconv.ParseFloat(res, 64)
	resN := _type.NewPrimitiveNumber("", resI)

	return resN
}
