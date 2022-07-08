package _type

import "math"

// ECMABigInt
// Primitive ECMA big int.
type ECMABigInt struct {
	*ECMAPrimitive
}

// NewPrimitiveBigInt :
// Create a new primitive number.
func NewPrimitiveBigInt(value interface{}) *ECMABigInt {
	return &ECMABigInt{
		ECMAPrimitive: NewPrimitive(value),
	}
}

// ToBool :
// Convert this primitive into a boolean value.
// https://developer.mozilla.org/en-US/docs/Glossary/Falsy
func (p *ECMABigInt) ToBool() (bool, error) {
	v := p.ECMAPrimitive.GetValue().(float64)
	return v != 0 && !math.IsNaN(v), nil
}
