package _type

import "math"

// ECMANumber
// Primitive ECMA number.
type ECMANumber struct {
	*ECMAPrimitive
}

// NewPrimitiveNumber :
// Create a new primitive number.
func NewPrimitiveNumber(identifier string, value interface{}) *ECMANumber {
	return &ECMANumber{
		ECMAPrimitive: NewPrimitive3(NumberType_, identifier, value),
	}
}

// ToBool :
// Convert this primitive into a boolean value.
// https://developer.mozilla.org/en-US/docs/Glossary/Falsy
func (p *ECMANumber) ToBool() (bool, error) {
	v := p.ECMAPrimitive.Value.(float64)
	return v != 0 && !math.IsNaN(v), nil
}
