package token

import "fmt"

type Token struct {
	Type  Type
	Value interface{}
}

func (tok Token) String() string {
	return fmt.Sprintf("Token(%s, %v)", tok.Type, tok.Value)
}

func NewToken(_type Type, value interface{}) *Token {
	return &Token{Type: _type, Value: value}
}

func ValueToInt(tok *Token) int {
	res, _ := tok.Value.(int)
	return res
}

type Type uint16

const (
	ILLEGAL Type = iota
	EOF
	COMMENT
	UNDEFINED
	NULL
	BOOLEAN
	STRING
	SYMBOL
	OBJECT
	NUMBER
	ADD
	SUB
	MUL
	DIV
	LPAREN
	RPAREN
	IDENTIFIER
	KEYWORD
	ASSIGN
	SEMICOLON
)

var asString = [...]string{
	"UNKNOWN",
	"EOF",
	"COMMENT",
	"UNDEFINED",
	"NULL",
	"BOOLEAN",
	"STRING",
	"SYMBOL",
	"OBJECT",
	"NUMBER",
	"ADD",
	"SUB",
	"MUL",
	"DIV",
	"LPAREN",
	"RPAREN",
	"IDENTIFIER",
	"KEYWORD",
	"ASSIGN",
	"SEMICOLON",
}

func (tok Type) String() string {
	return asString[tok]
}
