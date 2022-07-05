package grammar

import "duck/ling/js/syntax/unicode"

type NonTerminalType uint32

const (
	SourceCharacter NonTerminalType = iota
	Whitespace
)

var NonTerminalMap = map[NonTerminalType]interface{}{
	SourceCharacter: unicode.SourceCharacter,
	Whitespace:      unicode.Whitespace,
}

//type NonTerminal func()
//
//func GetNonTerminal(_type NonTerminalType) NonTerminal {
//	//switch _type {
//	//case Whitespace:
//	//	NonTerminalMap[Whitespace].(func(rune))
//	//}
//}
