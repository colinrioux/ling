package _type

import "fmt"

//var KeywordPrimitiveMap = map[string]interface{}{
//	"true":      NewPrimitiveBoolean("true", true),
//	"false":     NewPrimitiveBoolean("false", false),
//	"undefined": NewPrimitiveUndefined(),
//	"null":      NewPrimitiveNull(),
//}

type IECMAPrimitive interface {
	ToBool() (bool, error)
}

// ECMAPrimitive :
// Base primitive type that all other primitives inherit from.
type ECMAPrimitive struct {
	Type       ECMAPrimitiveType // the _type for such a primitive
	Identifier string            // the variable identifier
	Value      interface{}       // value held by the identifier
}

// NewPrimitive1 :
// Create a new base primitive type with a name.
func NewPrimitive1(identifier string) *ECMAPrimitive {
	return &ECMAPrimitive{
		Identifier: identifier,
	}
}

// NewPrimitive2 :
// Create a new base primitive type with an identifier and value.
func NewPrimitive2(identifier string, value interface{}) *ECMAPrimitive {
	return &ECMAPrimitive{
		Identifier: identifier,
		Value:      value,
	}
}

// NewPrimitive3 :
// Create a new base primitive type with a type, identifier, and value.
func NewPrimitive3(_type ECMAPrimitiveType, identifier string, value interface{}) *ECMAPrimitive {
	return &ECMAPrimitive{
		Type:       _type,
		Identifier: identifier,
		Value:      value,
	}
}

func (p *ECMAPrimitive) String() string {
	return fmt.Sprintf("Primitive(%v,%v)", p.Identifier, p.Value)
}

// ToBool :
// Convert this primitive into a boolean value.
// https://developer.mozilla.org/en-US/docs/Glossary/Falsy
// https://developer.mozilla.org/en-US/docs/Glossary/Truthy
func (p *ECMAPrimitive) ToBool() (bool, error) {
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
