package _type

import "math"

// ECMABigInt
// Primitive ECMA big int.
type ECMABigInt struct {
	*ECMAPrimitive
}

// NewPrimitiveBigInt :
// Create a new primitive number.
func NewPrimitiveBigInt(identifier string, value interface{}) *ECMABigInt {
	return &ECMABigInt{
		ECMAPrimitive: NewPrimitive3(BigIntType_, identifier, value),
	}
}

// ToBool :
// Convert this primitive into a boolean value.
// https://developer.mozilla.org/en-US/docs/Glossary/Falsy
// https://developer.mozilla.org/en-US/docs/Glossary/Truthy
func (p *ECMABigInt) ToBool() (bool, error) {
	v := p.ECMAPrimitive.GetValue().(float64)
	return v != 0 && !math.IsNaN(v), nil
}
