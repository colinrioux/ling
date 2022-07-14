package _type

import "fmt"

// ECMANull :
// Primitive ECMA null.
type ECMANull ECMAPrimitive

// NewPrimitiveNull :
// Create a new primitive null.
func NewPrimitiveNull() *ECMANull {
	return &ECMANull{
		Type:       NullType_,
		Identifier: "null",
		Value:      nil,
	}
}

func (p *ECMANull) String() string {
	return fmt.Sprintf("Null(%v,%v)", p.Identifier, p.Value)
}

// ToBool :
// Convert this primitive into a boolean value.
// https://developer.mozilla.org/en-US/docs/Glossary/Falsy
// https://developer.mozilla.org/en-US/docs/Glossary/Truthy
func (p *ECMANull) ToBool() (bool, error) {
	return false, nil
}
