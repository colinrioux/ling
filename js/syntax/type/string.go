package _type

import "fmt"

// ECMAString :
// Primitive ECMA string.
type ECMAString ECMAPrimitive

// NewPrimitiveString :
// Create a new primitive string.
func NewPrimitiveString(identifier string, value interface{}) *ECMAString {
	return &ECMAString{
		Type:       StringType_,
		Identifier: identifier,
		Value:      value,
	}
}

func (p *ECMAString) String() string {
	return fmt.Sprintf("String(%v,\"%v\")", p.Identifier, p.Value)
}

// ToBool :
// Convert this primitive into a boolean value.
// https://developer.mozilla.org/en-US/docs/Glossary/Falsy
// https://developer.mozilla.org/en-US/docs/Glossary/Truthy
func (p *ECMAString) ToBool() (bool, error) {
	return len(p.Value.(string)) > 0, nil
}
