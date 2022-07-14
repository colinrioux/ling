package _type

import "fmt"

// ECMAUndefined :
// Primitive ECMA undefined.
type ECMAUndefined ECMAPrimitive

// NewPrimitiveUndefined :
// Create a new primitive undefined.
func NewPrimitiveUndefined() *ECMAUndefined {
	return &ECMAUndefined{
		Type:       UndefinedType_,
		Identifier: "undefined",
		Value:      nil,
	}
}

func (p *ECMAUndefined) String() string {
	return fmt.Sprintf("Undefined(%v,%v)", p.Identifier, p.Value)
}

// ToBool :
// Convert this primitive into a boolean value.
// https://developer.mozilla.org/en-US/docs/Glossary/Falsy
// https://developer.mozilla.org/en-US/docs/Glossary/Truthy
func (p *ECMAUndefined) ToBool() (bool, error) {
	return false, nil
}
