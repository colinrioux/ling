package _type

import (
	"duck/ling/js/ast/token"
	"strings"
)

type ECMAType int
type ECMATypeClass bool

type ECMAValue struct {
	class ECMATypeClass
	_type ECMAType
	value interface{}
	token *token.Token
}

func NewValue(value string) *ECMAValue {
	var t ECMAType = UndefinedType
	var c ECMATypeClass = LanguageType
	var tokType token.Type = token.UNDEFINED
	switch strings.ToLower(value) {
	case "null":
		t = NullType
		c = LanguageType
		tokType = token.NULL
	case "true":
	case "false":
		t = BooleanType
		c = LanguageType
		tokType = token.BOOLEAN
	case "undefined":
		t = UndefinedType
		c = LanguageType
		tokType = token.UNDEFINED
	}

	return &ECMAValue{
		class: c,
		_type: t,
		value: value,
		token: token.NewToken(tokType, value),
	}
}

func (v *ECMAValue) Type() ECMAType {
	return v._type
}

func (v *ECMAValue) Value() interface{} {
	return v.value
}

func (v *ECMAValue) Token() *token.Token {
	return v.token
}

const (
	UndefinedType ECMAType = iota
	NullType
	BooleanType
	StringType
	SymbolType
	NumberType
	BigIntType
	ObjectType
)

const (
	LanguageType      ECMATypeClass = false
	SpecificationType ECMATypeClass = true
)
