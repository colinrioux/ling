package _type

import (
	"fmt"
	"math"
)

// ECMABigInt
// Primitive ECMA big int.
type ECMABigInt ECMAPrimitive

// NewPrimitiveBigInt :
// Create a new primitive number.
func NewPrimitiveBigInt(identifier string, value interface{}) *ECMABigInt {
	return &ECMABigInt{
		Type:       BigIntType_,
		Identifier: identifier,
		Value:      value,
	}
}

func (p *ECMABigInt) String() string {
	return fmt.Sprintf("BigInt(%v,%v)", p.Identifier, p.Value)
}

// ToBool :
// Convert this primitive into a boolean value.
// https://developer.mozilla.org/en-US/docs/Glossary/Falsy
// https://developer.mozilla.org/en-US/docs/Glossary/Truthy
func (p *ECMABigInt) ToBool() (bool, error) {
	v := p.Value.(float64)
	return v != 0 && !math.IsNaN(v), nil
}
