package _type

import (
	"github.com/google/uuid"
)

// PrivateName :
// A type used to describe a globally unique value.
// https://tc39.es/ecma262/#sec-private-names
type PrivateName struct {
	Name        uuid.UUID
	Description string
}

// NewPrivateName :
// Creates a new private name with a description.
func NewPrivateName(description string) *PrivateName {
	return &PrivateName{
		Name:        uuid.New(),
		Description: description,
	}
}

// String :
// Converts a private name to a string.
func (n *PrivateName) String() string {
	return n.Name.String()
}

// PrivateElement :
// A type to represent private fields, methods, or accessors of objects.
// https://tc39.es/ecma262/#sec-privateelement-specification-type
type PrivateElement struct {
	Key   *PrivateName
	Kind  PrivateElementType
	Value *ECMAPrimitive // Field & Method types only
	Get   *ECMAObject    // Accessor Only. TODO function object
	Set   *ECMAObject    // Accessor Only. TODO function object
}

// NewPrivateElement :
// Creates a new private element.
func NewPrivateElement(
	description string,
	kind PrivateElementType,
	value *ECMAPrimitive,
	get *ECMAObject,
	set *ECMAObject,
) *PrivateElement {
	return &PrivateElement{
		Key:   NewPrivateName(description),
		Kind:  kind,
		Value: value,
		Get:   get,
		Set:   set,
	}
}

type PrivateElementType uint8

const (
	FieldElement PrivateElementType = iota
	MethodElement
	AccessorElement
)
