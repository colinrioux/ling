package grammar

import "duck/ling/js/syntax/unicode"

type NonTerminalType uint32

const (
	SourceCharacter NonTerminalType = iota
	Whitespace
)

var NonTerminalMap = map[NonTerminalType]interface{}{
	SourceCharacter: unicode.IsSourceCharacter,
	Whitespace:      unicode.IsWhitespace,
}

//type NonTerminal func()
//
//func GetNonTerminal(_type NonTerminalType) NonTerminal {
//	//switch _type {
//	//case IsWhitespace:
//	//	NonTerminalMap[IsWhitespace].(func(rune))
//	//}
//}
