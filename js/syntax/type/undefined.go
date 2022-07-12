package _type

// ECMAUndefined :
// Primitive ECMA undefined.
type ECMAUndefined struct {
	*ECMAPrimitive
}

// NewPrimitiveUndefined :
// Create a new primitive undefined.
func NewPrimitiveUndefined() *ECMAUndefined {
	return &ECMAUndefined{
		ECMAPrimitive: NewPrimitive3(UndefinedType_, "undefined", nil),
	}
}

// ToBool :
// Convert this primitive into a boolean value.
// https://developer.mozilla.org/en-US/docs/Glossary/Falsy
// https://developer.mozilla.org/en-US/docs/Glossary/Truthy
func (p *ECMAUndefined) ToBool() (bool, error) {
	return false, nil
}
