package _type

import "errors"

// ECMASymbol
// Primitive ECMA symbol.
type ECMASymbol struct {
	*ECMAPrimitive
}

// NewPrimitiveSymbol :
// Create a new primitive symbol.
func NewPrimitiveSymbol(identifier string, value interface{}) *ECMASymbol {
	return &ECMASymbol{
		ECMAPrimitive: NewPrimitive3(SymbolType_, identifier, value),
	}
}

// ToBool :
// Convert this primitive into a boolean value.
// https://developer.mozilla.org/en-US/docs/Glossary/Falsy
// https://developer.mozilla.org/en-US/docs/Glossary/Truthy
func (p *ECMASymbol) ToBool() (bool, error) {
	return false, errors.New("symbols cannot be cast to boolean")
}
