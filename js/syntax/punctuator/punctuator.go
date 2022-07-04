package punctuator

// Punctuator. https://tc39.es/ecma262/#sec-punctuators
const (
	RightBrace = "[}]"
	Div        = "[/|/=]"
	Other      = "[{|(|)|[|]|.|...|;|,|>|<|>=|<=|==|!=|===|!==|+|-|*|%|**" +
		"|++|--|<<|>>|>>>|&|\u007C|\u2038|!|~|&&|\u007C\u007C|??|?|:|=|+=" +
		"|-=|*=|%=|**=|<<=|>>=|>>>=|&=|\u007C=|\u2038=|&&=|\u007C\u007C=|??=|=>]"
)

// OptionalChaining : https://tc39.es/ecma262/#prod-OptionalChainingPunctuator
//func OptionalChaining() string {
//	rx := pcre.MustCompile(literal.DecimalDigit, 0)
//	matcher := rx.Matcher([]byte{lookahead}, 0)
//	return "?."
//}
