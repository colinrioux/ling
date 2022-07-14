package literal

import "duck/ling/util"

// Null Literal. https://tc39.es/ecma262/#sec-ecmascript-language-lexical-grammar-literals
const (
	Null = "[null]"
)

// Boolean Literal. https://tc39.es/ecma262/#sec-ecmascript-language-lexical-grammar-literals
const (
	Boolean = "[true|false]"
)

// Numeric Literal. https://tc39.es/ecma262/#sec-ecmascript-language-lexical-grammar-literals
const (
	NumericSeparator  = "[_]"
	BigIntSuffix      = "[n]"
	DecimalDigit      = "[0-9]"
	NonZeroDigit      = "[1-9]"
	ExponentIndicator = "e|E"
	BinaryDigit       = "[0-1]"
	OctalDigit        = "[0-7]"
	NonOctalDigit     = "[8-9]"
	HexDigit          = "[0-9a-fA-F]"
	AlphaNumeric      = "[0-9a-zA-Z]"
	Alpha             = "[a-zA-Z]"
	Scientific        = "[eE][-+]"
)

func IsNumericSeparator(r rune) bool {
	return util.Match(NumericSeparator, string(r))
}

func IsBigIntSuffix(r rune) bool {
	return util.Match(BigIntSuffix, string(r))
}

func IsDecimalDigit(r rune) bool {
	return util.Match(DecimalDigit, string(r))
}

func IsNonZeroDigit(r rune) bool {
	return util.Match(NonZeroDigit, string(r))
}

func IsExponentIndicator(r rune) bool {
	return util.Match(ExponentIndicator, string(r))
}

func IsBinaryDigit(r rune) bool {
	return util.Match(BinaryDigit, string(r))
}

func IsOctalDigit(r rune) bool {
	return util.Match(OctalDigit, string(r))
}

func IsNonOctalDigit(r rune) bool {
	return util.Match(NonOctalDigit, string(r))
}

func IsHexDigit(r rune) bool {
	return util.Match(HexDigit, string(r))
}

func IsAlphaNumeric(r rune) bool {
	return util.Match(AlphaNumeric, string(r))
}

func IsAlpha(r rune) bool {
	return util.Match(Alpha, string(r))
}

func IsScientific(r rune, lookahead rune) bool {
	return util.Match(Scientific, string(r)+string(lookahead))
}

// String Literal. https://tc39.es/ecma262/#sec-ecmascript-language-lexical-grammar-literals
const (
	SingleEscapeCharacter         = "[\u0027|\u0022|\u005C|b|f|n|r|t|v]"
	ZeroToThree                   = "[0-3]"
	FourToSeven                   = "[4-7]"
	NonOctalDecimalEscapeSequence = "[8-9]"
	HexEscapeSequence             = "[x]" + HexDigit + HexDigit
	Hex4Digits                    = HexDigit + HexDigit + HexDigit + HexDigit
)
