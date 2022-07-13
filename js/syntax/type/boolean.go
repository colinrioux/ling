package _type

// ECMABoolean :
// Primitive ECMA boolean.
type ECMABoolean struct {
	*ECMAPrimitive
}

// NewPrimitiveBoolean :
// Create a new primitive boolean.
func NewPrimitiveBoolean(identifier string, value interface{}) *ECMABoolean {
	return &ECMABoolean{
		ECMAPrimitive: NewPrimitive3(BooleanType_, identifier, value),
	}
}

// ToBool :
// Convert this primitive into a boolean value.
// https://developer.mozilla.org/en-US/docs/Glossary/Falsy
// https://developer.mozilla.org/en-US/docs/Glossary/Truthy
func (p *ECMABoolean) ToBool() (bool, error) {
	return p.ECMAPrimitive.Value.(bool), nil
}

type ECMABooleanObject struct {
	*ECMAObject
}
