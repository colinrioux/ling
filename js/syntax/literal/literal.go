package literal

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
)

// String Literal. https://tc39.es/ecma262/#sec-ecmascript-language-lexical-grammar-literals
const (
	SingleEscapeCharacter         = "[\u0027|\u0022|\u005C|b|f|n|r|t|v]"
	ZeroToThree                   = "[0-3]"
	FourToSeven                   = "[4-7]"
	NonOctalDecimalEscapeSequence = "[8-9]"
	HexEscapeSequence             = "[x]" + HexDigit + HexDigit
	Hex4Digits                    = HexDigit + HexDigit + HexDigit + HexDigit
)
