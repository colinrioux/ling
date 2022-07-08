package parser

import (
	"duck/ling/js/syntax/literal"
	_type "duck/ling/js/syntax/type"
	"strconv"
)

func ParseNumber() *_type.ECMANumber {
	var res string = ""
	for CurrentChar != 0 && (literal.IsDecimalDigit(CurrentChar) ||
		CurrentChar == '.' || literal.IsScientific(CurrentChar, Peek())) {

		if literal.IsScientific(CurrentChar, Peek()) {
			res += string(CurrentChar)
			Advance()
		}

		res += string(CurrentChar)
		Advance()
	}
	resI, _ := strconv.ParseFloat(res, 64)
	resN := _type.NewECMANumber1(resI)

	return resN
}
