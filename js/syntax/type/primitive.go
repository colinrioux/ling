package _type

//var KeywordPrimitiveMap = map[string]interface{}{
//	"true":      NewPrimitiveBoolean("true", true),
//	"false":     NewPrimitiveBoolean("false", false),
//	"undefined": NewPrimitiveUndefined(),
//	"null":      NewPrimitiveNull(),
//}

type IECMAPrimitive interface {
	ToBool() (bool, error)
	IsSameAs(other *ECMAPrimitive) (bool, error)
	IsSameAsNonNumeric(other *ECMAPrimitive) (bool, error)
}

// ECMAPrimitive :
// Base primitive type that all other primitives inherit from.
type ECMAPrimitive struct {
	_type      ECMAPrimitiveType // the _type for such a primitive
	identifier string            // the variable identifier
	value      interface{}       // value held by the identifier
}

// NewPrimitive1 :
// Create a new base primitive type with a name.
func NewPrimitive1(identifier string) *ECMAPrimitive {
	return &ECMAPrimitive{
		identifier: identifier,
	}
}

// NewPrimitive2 :
// Create a new base primitive type with an identifier and value.
func NewPrimitive2(identifier string, value interface{}) *ECMAPrimitive {
	return &ECMAPrimitive{
		identifier: identifier,
		value:      value,
	}
}

// NewPrimitive3 :
// Create a new base primitive type with a type, identifier, and value.
func NewPrimitive3(_type ECMAPrimitiveType, identifier string, value interface{}) *ECMAPrimitive {
	return &ECMAPrimitive{
		_type:      _type,
		identifier: identifier,
		value:      value,
	}
}

// GetType :
// Get the type for this primitive
func (p *ECMAPrimitive) GetType() ECMAPrimitiveType {
	return p._type
}

// GetIdentifier :
// Get the identifier for this primitive.
func (p *ECMAPrimitive) GetIdentifier() string {
	return p.identifier
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

// IsSameAs :
// Check if this primitive is the same as the other primitive.
// https://tc39.es/ecma262/#sec-samevalue
func (p *ECMAPrimitive) IsSameAs(other *ECMAPrimitive) (bool, error) {
	if p.GetType() != other.GetType() {
		return false, nil
	}
	if p.GetType() == NumberType_ {
		// TODO
		return false, nil
	}

	if p.GetType() == BigIntType_ {
		// TODO
		return false, nil
	}
	return p.IsSameAsNonNumeric(other)
}

// IsSameAsNonNumeric :
// Check if this primitive is the same as the other primitive.
// https://tc39.es/ecma262/#sec-samevaluenonnumeric
func (p *ECMAPrimitive) IsSameAsNonNumeric(other *ECMAPrimitive) (bool, error) {
	// handle undefined & null types
	if p.GetType() == UndefinedType_ || p.GetType() == NullType_ {
		return true, nil
	}

	// handle string type
	if p.GetType() == StringType_ {
		// TODO : does go string allow ecma like strings
		return p.GetValue().(string) == other.GetValue().(string), nil
	}

	// handle boolean type
	if p.GetType() == BooleanType_ {
		return p.GetValue().(bool) == other.GetValue().(bool), nil
	}

	// handle symbol type
	if p.GetType() == SymbolType_ {
		// TODO
		return false, nil
	}

	// otherwise: handle object type
	// TODO
	return false, nil
}

type ECMAPrimitiveType uint8

const (
	UndefinedType_ ECMAPrimitiveType = iota
	NullType_
	BooleanType_
	StringType_
	SymbolType_
	NumberType_
	BigIntType_
	ObjectType_
)
