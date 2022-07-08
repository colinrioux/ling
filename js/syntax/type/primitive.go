package _type

var KeywordPrimitiveMap = map[string]interface{}{
	"true":      NewPrimitiveBoolean(true),
	"false":     NewPrimitiveBoolean(false),
	"undefined": NewPrimitiveUndefined(),
	"null":      NewPrimitiveNull(),
}

type IECMAPrimitive interface {
	ToBool() (bool, error)
}

// ECMAPrimitive :
// Base primitive type that all other primitives inherit from.
type ECMAPrimitive struct {
	value interface{}
}

// NewPrimitive :
// Create a new base primitive type with a value.
func NewPrimitive(value interface{}) *ECMAPrimitive {
	return &ECMAPrimitive{
		value: value,
	}
}

// GetValue :
// Get the value associated with this primitive, if there is one.
func (p *ECMAPrimitive) GetValue() interface{} {
	return p.value
}

// ToBool :
// Convert this primitive into a boolean value.
// https://developer.mozilla.org/en-US/docs/Glossary/Falsy
// https://developer.mozilla.org/en-US/docs/Glossary/Truthy
func (p *ECMAPrimitive) ToBool() (bool, error) {
	return false, nil
}
