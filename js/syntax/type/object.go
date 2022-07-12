package _type

import (
	"errors"
	"reflect"
)

// DataProperty vs AccessorProperty :
// Data properties vs accessor properties are two different means to define properties.
// Useful link: https://tc39.es/ecma262/#sec-object-type
// Useful link: https://stackoverflow.com/questions/29592110/difference-between-accessor-property-and-data-property-in-ecmascript

// ObjectAttribute :
// Base attribute type that are shared between data and accessor properties.
type ObjectAttribute struct {
	Enumerable   bool
	Configurable bool
	Observed     bool
}

// NewObjectAttribute :
// Create a new base attribute type.
func NewObjectAttribute() *ObjectAttribute {
	return &ObjectAttribute{
		Enumerable:   false,
		Configurable: false,
	}
}

// ObjectAttributeData :
// Attribute type specific to data properties.
type ObjectAttributeData struct {
	*ObjectAttribute
	Value    interface{}
	Writable bool
}

// NewObjectAttributeData :
// Create a new data specific attribute type.
func NewObjectAttributeData() *ObjectAttributeData {
	return &ObjectAttributeData{
		ObjectAttribute: NewObjectAttribute(),
		Value:           NewPrimitiveUndefined(),
		Writable:        false,
	}
}

// ObjectAttributeAccessor :
// Attribute type specific to accessor properties.
type ObjectAttributeAccessor struct {
	*ObjectAttribute
	Get interface{}
	Set interface{}
}

// NewObjectAttributeAccessor :
// Create a new accessor specific attribute type.
func NewObjectAttributeAccessor() *ObjectAttributeAccessor {
	return &ObjectAttributeAccessor{
		ObjectAttribute: NewObjectAttribute(),
		Get:             NewPrimitiveUndefined(),
		Set:             NewPrimitiveUndefined(),
	}
}

// DataProperty :
// Data property type.
type DataProperty struct {
	Value      interface{}
	Attributes *ObjectAttributeData
}

// NewDataProperty :
// Create a new data property.
func NewDataProperty(value interface{}) *DataProperty {
	return &DataProperty{
		Value:      value,
		Attributes: NewObjectAttributeData(),
	}
}

// AccessorProperty :
// Accessor property type.
type AccessorProperty struct {
	Value      interface{}
	Attributes *ObjectAttributeAccessor
}

// NewAccessorProperty :
// Create a new accessor property.
func NewAccessorProperty(value interface{}) *AccessorProperty {
	return &AccessorProperty{
		Value:      value,
		Attributes: NewObjectAttributeAccessor(),
	}
}

// MethodProperty :
// Method properties refer to internal methods defined by the object. This is implementation specific.
// The value must be of type function.
// https://tc39.es/ecma262/#sec-object-internal-methods-and-internal-slots
type MethodProperty struct {
	Value interface{}
}

// NewMethodProperty :
// Creates a new method property. TODO
func NewMethodProperty(value interface{}) *MethodProperty {
	return &MethodProperty{
		Value: value,
	}
}

// ECMAObject :
// An object is a primitive type that is a collection of properties.
// Most types in ECMA are children of this base object type.
// https://tc39.es/ecma262/#sec-object-type
// TODO symbols can be used as keys to these properties. We need to have a means to convert them to appropriate strings
type ECMAObject struct {
	*ECMAPrimitive
	DataProperties     map[string]*DataProperty
	AccessorProperties map[string]*AccessorProperty
	InternalMethods    map[string]*MethodProperty
	// https://tc39.es/ecma262/#sec-object-internal-methods-and-internal-slots
	InternalSlots struct {
		PrivateElements []*PrivateElementRecord
		Prototype       *ECMAObject
		Extensible      bool
	}
}

// NewECMAObject :
// Create a new ordinary ECMA object.
func NewECMAObject(identifier string) *ECMAObject {
	return &ECMAObject{
		ECMAPrimitive: NewPrimitive3(ObjectType_, identifier, nil),
		InternalMethods: map[string]*MethodProperty{
			"getPrototypeOf":    NewMethodProperty(GetPrototypeOf),
			"setPrototypeOf":    NewMethodProperty(SetPrototypeOf),
			"isExtensible":      NewMethodProperty(IsExtensible),
			"preventExtensions": NewMethodProperty(PreventExtensions),
			"getOwnProperty":    NewMethodProperty(GetOwnProperty),
			"defineOwnProperty": NewMethodProperty(DefineOwnProperty),
			"hasProperty":       NewMethodProperty(HasProperty),
			"get":               NewMethodProperty(Get),
			"set":               NewMethodProperty(Set),
			"delete":            NewMethodProperty(Delete),
			"ownPropertyKeys":   NewMethodProperty(OwnPropertyKeys),
		},
	}
}

// ToBool :
// Convert this primitive into a boolean value.
// https://developer.mozilla.org/en-US/docs/Glossary/Falsy
// https://developer.mozilla.org/en-US/docs/Glossary/Truthy
func (p *ECMAObject) ToBool() (bool, error) {
	return false, errors.New("objects cannot be cast to boolean")
}

// IsNonExtensible :
// An object is non-extensible if it has been observed to return false from its
// [[IsExtensible]] internal method, or true from its [[PreventExtensions]] internal method.
// https://tc39.es/ecma262/#sec-invariants-of-the-essential-internal-methods
func (p *ECMAObject) IsNonExtensible() (bool, error) {
	var err error = nil
	// First check if IsExtensible returns false
	isExtensible, ok := p.InternalMethods["IsExtensible"].Value.(func(*ECMAObject) *CompletionRecord)
	if ok {
		if r, ok := isExtensible(p).GetValue().(bool); ok && !r {
			return true, nil
		}
		err = errors.New("object does not have valid internal method IsExtensible")
	}

	// Now check if PreventExtensions returns true
	preventExtensions, ok := p.InternalMethods["PreventExtensions"].Value.(func(*ECMAObject) *CompletionRecord)
	if ok {
		if r, ok := preventExtensions(p).GetValue().(bool); ok && r {
			return true, nil
		}
		err = errors.New("object does not have valid internal method PreventExtensions")
	}

	if err == nil {
		return false, errors.New("object does not have internal method IsExtensible nor PreventExtensions")
	}
	return false, err
}

// GetPrototypeOf :
// Determine the object that provides inherited properties for this object.
// A nil value indicates that there are no inherited properties.
// https://tc39.es/ecma262/#table-essential-internal-methods
func GetPrototypeOf(target *ECMAObject) *CompletionRecord {
	// INVARIANTS---
	if ok, _ := target.IsNonExtensible(); ok && target.InternalSlots.Prototype != nil {
		return NewCompletionRecord(NormalCompletion, target.InternalSlots.Prototype, "")
	}
	// ---INVARIANTS

	return NewCompletionRecord(ThrowCompletion, nil, "")
}

// SetPrototypeOf :
// Associate this object with another object that provides inherited properties.
// Passing nil indicates that there are no inherited properties.
// Returns true indicating that the operation was completed successfully
// or false indicating that the operation was not successful.
// https://tc39.es/ecma262/#table-essential-internal-methods
func SetPrototypeOf(target *ECMAObject, v *ECMAObject) *CompletionRecord {
	// INVARIANTS---
	if ok, _ := target.IsNonExtensible(); ok {
		if getPrototypeOf, ok2 := target.InternalMethods["GetPrototypeOf"].Value.(func(*ECMAObject) *CompletionRecord); ok2 {
			r := getPrototypeOf(target)
			return NewCompletionRecord(
				NormalCompletion,
				reflect.ValueOf(r.GetValue()).Type() == reflect.ValueOf(v).Type(),
				"",
			)
		}
	}
	// ---INVARIANTS

	return NewCompletionRecord(ThrowCompletion, nil, "")
}

// PreventExtensions :
// Control whether new properties may be added to this object.
// Returns true if the operation was successful or false if the operation was unsuccessful.
// https://tc39.es/ecma262/#table-essential-internal-methods
func PreventExtensions(target *ECMAObject) *CompletionRecord {
	// INVARIANTS---
	target.InternalSlots.Extensible = false
	return NewCompletionRecord(NormalCompletion, true, "")
	// ---INVARIANTS
}

// GetOwnProperty :
// Return a Property Descriptor for the own property of this object whose key is propertyKey,
// or nil if no such property exists.
// https://tc39.es/ecma262/#table-essential-internal-methods
func GetOwnProperty(target *ECMAObject, propertyKey string) *CompletionRecord {
	// INVARIANTS---
	// Check if propertyKey is a data property
	if prop, ok := target.DataProperties[propertyKey]; ok {
		if !prop.Attributes.Configurable && !prop.Attributes.Writable {
			prop.Attributes.Observed = true
			return NewCompletionRecord(
				NormalCompletion,
				NewPropertyDescriptor(DataPropertyDescriptor, prop.Value),
				"",
			)
		}
	}

	// Check if propertyKey is an accessor property
	if prop, ok := target.AccessorProperties[propertyKey]; ok {
		return NewCompletionRecord(
			NormalCompletion,
			NewPropertyDescriptor(AccessorPropertyDescriptor, prop.Value),
			"",
		)
	}

	// Check if propertyKey is a method property
	if prop, ok := target.InternalMethods[propertyKey]; ok {
		return NewCompletionRecord(
			NormalCompletion,
			NewPropertyDescriptor(MethodPropertyDescriptor, prop.Value),
			"",
		)
	}

	if ok, _ := target.IsNonExtensible(); ok {
		return NewCompletionRecord(NormalCompletion, NewPrimitiveUndefined(), "")
	}
	// ---INVARIANTS

	return NewCompletionRecord(ThrowCompletion, nil, "")
}

// DefineOwnProperty :
// Create or alter the own property, whose key is propertyKey,
// to have the state described by PropertyDescriptor.
// Return true if that property was successfully created/updated
// or false if the property could not be created or updated.
// https://tc39.es/ecma262/#table-essential-internal-methods
func DefineOwnProperty(target *ECMAObject, propertyKey string, propertyDescriptor *PropertyDescriptor) *CompletionRecord {
	// INVARIANTS---
	// Check if propertyKey is a data property
	if prop, ok := target.DataProperties[propertyKey]; ok {
		if prop.Attributes.Writable {
			return NewCompletionRecord(NormalCompletion, true, "")
		}

		// Check if all attributes of prop are same as propertyDescriptor's
		if attr, ok := propertyDescriptor.GetValue().(*DataProperty); ok {
			if attr.Attributes == prop.Attributes {
				return NewCompletionRecord(NormalCompletion, true, "")
			}
		}

		if prop.Attributes.Observed {
			return NewCompletionRecord(NormalCompletion, false, "")
		}

		return NewCompletionRecord(NormalCompletion, true, "")
	}

	// Check if propertyKey is an accessor property
	if prop, ok := target.AccessorProperties[propertyKey]; ok {
		// Check if all attributes of prop are same as propertyDescriptor's
		if attr, ok := propertyDescriptor.GetValue().(*AccessorProperty); ok {
			if attr.Attributes == prop.Attributes {
				return NewCompletionRecord(NormalCompletion, true, "")
			}
		}
	}

	if ok, _ := target.IsNonExtensible(); ok {
		return NewCompletionRecord(NormalCompletion, false, "")
	}
	// ---INVARIANTS

	return NewCompletionRecord(ThrowCompletion, nil, "")
}

// HasProperty :
// Return a Boolean value indicating whether this object already has either an own
// or inherited property whose key is propertyKey.
// https://tc39.es/ecma262/#table-essential-internal-methods
func HasProperty(target *ECMAObject, propertyKey string) *CompletionRecord {
	// INVARIANTS---
	// Check if propertyKey is a data property
	if prop, ok := target.DataProperties[propertyKey]; ok {
		return NewCompletionRecord(NormalCompletion, prop.Attributes.Observed, "")
	}
	// Check if propertyKey is an accessor property
	if prop, ok := target.AccessorProperties[propertyKey]; ok {
		return NewCompletionRecord(NormalCompletion, prop.Attributes.Observed, "")
	}
	// --INVARIANTS

	return NewCompletionRecord(ThrowCompletion, nil, "")
}

// Get :
// Return the value of the property whose key is propertyKey from this object.
// If any ECMAScript code must be executed to retrieve the property value,
// Receiver is used as the "this" value when evaluating the code.
// https://tc39.es/ecma262/#table-essential-internal-methods
func Get(target *ECMAObject, propertyKey string, receiver *ECMAObject) *CompletionRecord {
	// INVARIANTS---
	// TODO FIX
	// Check if propertyKey is a data property
	if prop, ok := target.DataProperties[propertyKey]; ok && prop.Attributes.Observed {
		return NewCompletionRecord(NormalCompletion, prop.Value, "")
	}

	// Check if propertyKey is an accessor property
	if prop, ok := target.AccessorProperties[propertyKey]; ok && prop.Attributes.Observed {
		return NewCompletionRecord(NormalCompletion, prop.Value, "")
	}
	// --INVARIANTS

	return NewCompletionRecord(ThrowCompletion, nil, "")
}

// Set :
// Set the value of the property whose key is propertyKey to value.
// If any ECMAScript code must be executed to set the property value,
// Receiver is used as the "this" value when evaluating the code.
// Returns true if the property value was set or false if it could not be set.
// https://tc39.es/ecma262/#table-essential-internal-methods
func Set(target *ECMAObject, propertyKey string, receiver *ECMAObject) *CompletionRecord {
	// INVARIANTS---
	// TODO
	// --INVARIANTS

	return NewCompletionRecord(ThrowCompletion, nil, "")
}

// Delete :
// Remove the own property whose key is propertyKey from this object.
// Return false if the property was not deleted and is still present.
// Return true if the property was deleted or is not present.
// https://tc39.es/ecma262/#table-essential-internal-methods
func Delete(target *ECMAObject, propertyKey string) *CompletionRecord {
	// INVARIANTS---
	// TODO
	// --INVARIANTS

	return NewCompletionRecord(ThrowCompletion, nil, "")
}

// OwnPropertyKeys :
// Return a List whose elements are all the property keys for the object.
// https://tc39.es/ecma262/#table-essential-internal-methods
func OwnPropertyKeys(target *ECMAObject) *CompletionRecord {
	// INVARIANTS---
	// TODO
	// --INVARIANTS

	return NewCompletionRecord(ThrowCompletion, nil, "")
}
