package _type

import (
	"fmt"
	"math"
)

// ECMANumber
// Primitive ECMA number.
type ECMANumber ECMAPrimitive

// NewPrimitiveNumber :
// Create a new primitive number.
func NewPrimitiveNumber(identifier string, value interface{}) *ECMANumber {
	return &ECMANumber{
		Type:       NumberType_,
		Identifier: identifier,
		Value:      value,
	}
}

func (p *ECMANumber) String() string {
	return fmt.Sprintf("Number(%v,%v)", p.Identifier, p.Value)
}

// ToBool :
// Convert this primitive into a boolean value.
// https://developer.mozilla.org/en-US/docs/Glossary/Falsy
func (p *ECMANumber) ToBool() (bool, error) {
	v := p.Value.(float64)
	return v != 0 && !math.IsNaN(v), nil
}
