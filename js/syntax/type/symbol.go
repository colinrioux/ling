package _type

import (
	"errors"
	"fmt"
)

// ECMASymbol
// Primitive ECMA symbol.
type ECMASymbol ECMAPrimitive

// NewPrimitiveSymbol :
// Create a new primitive symbol.
func NewPrimitiveSymbol(identifier string, value interface{}) *ECMASymbol {
	return &ECMASymbol{
		Type:       SymbolType_,
		Identifier: identifier,
		Value:      value,
	}
}

func (p *ECMASymbol) String() string {
	return fmt.Sprintf("Symbol(%v,%v)", p.Identifier, p.Value)
}

// ToBool :
// Convert this primitive into a boolean value.
// https://developer.mozilla.org/en-US/docs/Glossary/Falsy
// https://developer.mozilla.org/en-US/docs/Glossary/Truthy
func (p *ECMASymbol) ToBool() (bool, error) {
	return false, errors.New("symbols cannot be cast to boolean")
}
