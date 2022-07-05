package token

import "fmt"

type IToken interface {
	String() string
	GetType() Type
	GetValue() any
}

type Token struct {
	_type Type
	value interface{}
}

func (tok Token) String() string {
	return fmt.Sprintf("Token(%s, %v)", tok._type, tok.value)
}

func (tok Token) GetType() Type {
	return tok._type
}

func (tok Token) GetValue() any {
	return tok.value
}

func NewToken(_type Type, value interface{}) *Token {
	return &Token{_type: _type, value: value}
}

func ValueToInt(tok *Token) int {
	res, _ := tok.GetValue().(int)
	return res
}

type Type uint32

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
	INTEGER
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
	"INTEGER",
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
