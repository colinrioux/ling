package _type

// ECMANull :
// Primitive ECMA null.
type ECMANull struct {
	*ECMAPrimitive
}

// NewPrimitiveNull :
// Create a new primitive null.
func NewPrimitiveNull() *ECMANull {
	return &ECMANull{
		ECMAPrimitive: NewPrimitive3(NullType_, "null", nil),
	}
}

// ToBool :
// Convert this primitive into a boolean value.
// https://developer.mozilla.org/en-US/docs/Glossary/Falsy
// https://developer.mozilla.org/en-US/docs/Glossary/Truthy
func (p *ECMANull) ToBool() (bool, error) {
	return false, nil
}
