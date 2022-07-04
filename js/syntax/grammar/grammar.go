package grammar

type NonTerminalType uint32

const (
	Whitespace NonTerminalType = iota
)

type NonTerminal func()

//func GetNonTerminal(_type NonTerminalType) NonTerminal {
//
//}
