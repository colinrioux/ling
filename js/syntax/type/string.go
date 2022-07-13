package _type

// ECMAString :
// Primitive ECMA string.
type ECMAString struct {
	*ECMAPrimitive
}

// NewPrimitiveString :
// Create a new primitive string.
func NewPrimitiveString(identifier string, value interface{}) *ECMAString {
	return &ECMAString{
		ECMAPrimitive: NewPrimitive3(StringType_, identifier, value),
	}
}

// ToBool :
// Convert this primitive into a boolean value.
// https://developer.mozilla.org/en-US/docs/Glossary/Falsy
// https://developer.mozilla.org/en-US/docs/Glossary/Truthy
func (p *ECMAString) ToBool() (bool, error) {
	return len(p.ECMAPrimitive.Value.(string)) > 0, nil
}
