package _type

import "errors"

// ECMASymbol
// Primitive ECMA symbol.
type ECMASymbol struct {
	*ECMAPrimitive
}

// NewPrimitiveSymbol :
// Create a new primitive symbol.
func NewPrimitiveSymbol(value interface{}) *ECMASymbol {
	return &ECMASymbol{
		ECMAPrimitive: NewPrimitive(value),
	}
}

// ToBool :
// Convert this primitive into a boolean value.
// https://developer.mozilla.org/en-US/docs/Glossary/Falsy
func (p *ECMASymbol) ToBool() (bool, error) {
	return false, errors.New("symbols cannot be cast to boolean")
}
