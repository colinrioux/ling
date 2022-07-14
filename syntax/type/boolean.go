package _type

import "fmt"

// ECMABoolean :
// Primitive ECMA boolean.
type ECMABoolean ECMAPrimitive

// NewPrimitiveBoolean :
// Create a new primitive boolean.
func NewPrimitiveBoolean(identifier string, value interface{}) *ECMABoolean {
	return &ECMABoolean{
		Type:       BooleanType_,
		Identifier: identifier,
		Value:      value,
	}
}

func (p *ECMABoolean) String() string {
	return fmt.Sprintf("Boolean(%v,%v)", p.Identifier, p.Value)
}

// ToBool :
// Convert this primitive into a boolean value.
// https://developer.mozilla.org/en-US/docs/Glossary/Falsy
// https://developer.mozilla.org/en-US/docs/Glossary/Truthy
func (p *ECMABoolean) ToBool() (bool, error) {
	return p.Value.(bool), nil
}
