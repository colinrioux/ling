package _type

// DataProperty vs AccessorProperty :
// Data properties vs accessor properties are two different means to define properties.
// Useful link: https://tc39.es/ecma262/#sec-object-type
// Useful link: https://stackoverflow.com/questions/29592110/difference-between-accessor-property-and-data-property-in-ecmascript

// ObjectAttribute :
// Base attribute type that are shared between data and accessor properties.
type ObjectAttribute struct {
	Enumerable   bool
	Configurable bool
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
	value      interface{}
	attributes *ObjectAttributeData
}

// NewDataProperty :
// Create a new data property.
func NewDataProperty(value interface{}) *DataProperty {
	return &DataProperty{
		value:      value,
		attributes: NewObjectAttributeData(),
	}
}

// AccessorProperty :
// Accessor property type.
type AccessorProperty struct {
	value      interface{}
	attributes *ObjectAttributeAccessor
}

// NewAccessorProperty :
// Create a new accessor property.
func NewAccessorProperty(value interface{}) *AccessorProperty {
	return &AccessorProperty{
		value:      value,
		attributes: NewObjectAttributeAccessor(),
	}
}

// MethodProperty :
// Method properties refer to internal methods defined by the object. This is implementation specific.
// The value must be of type function.
// https://tc39.es/ecma262/#sec-object-internal-methods-and-internal-slots
type MethodProperty struct {
	value interface{}
}

// NewMethodProperty :
// Creates a new method property. TODO
func NewMethodProperty(value interface{}) *MethodProperty {
	return &MethodProperty{
		value: value,
	}
}

// ECMAObject :
// An object is a collection of properties. Most types in ECMA are children of this base object type.
// https://tc39.es/ecma262/#sec-object-type
// TODO symbols can be used as keys to these properties. We need to have a means to convert them to appropriate strings
type ECMAObject struct {
	*ECMAPrimitive
	DataProperties     map[string]*DataProperty
	AccessorProperties map[string]*AccessorProperty
	InternalMethods    map[string]*MethodProperty
	// https://tc39.es/ecma262/#sec-object-internal-methods-and-internal-slots
	internalSlots struct {
		PrivateElements []*PrivateElementRecord
		Prototype       *ECMAObject
	}
}

// NewECMAObject :
// Create a new ordinary ECMA object.
func NewECMAObject() *ECMAObject {
	return &ECMAObject{
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

// GetPrototypeOf :
// Determine the object that provides inherited properties for this object.
// A nil value indicates that there are no inherited properties.
// https://tc39.es/ecma262/#table-essential-internal-methods
func GetPrototypeOf() *CompletionRecord {
	return nil
}

// SetPrototypeOf :
// Associate this object with another object that provides inherited properties.
// Passing nil indicates that there are no inherited properties.
// Returns true indicating that the operation was completed successfully
// or false indicating that the operation was not successful.
// https://tc39.es/ecma262/#table-essential-internal-methods
func SetPrototypeOf(obj *ECMAObject) *CompletionRecord {
	return nil
}

// IsExtensible :
// Determine whether it is permitted to add additional properties to this object.
// https://tc39.es/ecma262/#table-essential-internal-methods
func IsExtensible() *CompletionRecord {
	return nil
}

// PreventExtensions :
// Control whether new properties may be added to this object.
// Returns true if the operation was successful or false if the operation was unsuccessful.
// https://tc39.es/ecma262/#table-essential-internal-methods
func PreventExtensions() *CompletionRecord {
	return nil
}

// GetOwnProperty :
// Return a Property Descriptor for the own property of this object whose key is propertyKey,
// or nil if no such property exists.
// https://tc39.es/ecma262/#table-essential-internal-methods
func GetOwnProperty(propertyKey string) *CompletionRecord {
	return nil
}

// DefineOwnProperty :
// Create or alter the own property, whose key is propertyKey,
// to have the state described by PropertyDescriptor.
// Return true if that property was successfully created/updated
// or false if the property could not be created or updated.
// https://tc39.es/ecma262/#table-essential-internal-methods
func DefineOwnProperty(propertyKey string, propertyDescriptor *PropertyDescriptor) *CompletionRecord {
	return nil
}

// HasProperty :
// Return a Boolean value indicating whether this object already has either an own
// or inherited property whose key is propertyKey.
// https://tc39.es/ecma262/#table-essential-internal-methods
func HasProperty(propertyKey string) *CompletionRecord {
	return nil
}

// Get :
// Return the value of the property whose key is propertyKey from this object.
// If any ECMAScript code must be executed to retrieve the property value,
// Receiver is used as the "this" value when evaluating the code.
// https://tc39.es/ecma262/#table-essential-internal-methods
func Get(propertyKey string, receiver *ECMAObject) *CompletionRecord {
	return nil
}

// Set :
// Set the value of the property whose key is propertyKey to value.
// If any ECMAScript code must be executed to set the property value,
// Receiver is used as the "this" value when evaluating the code.
// Returns true if the property value was set or false if it could not be set.
// https://tc39.es/ecma262/#table-essential-internal-methods
func Set(propertyKey string, receiver *ECMAObject) *CompletionRecord {
	return nil
}

// Delete :
// Remove the own property whose key is propertyKey from this object.
// Return false if the property was not deleted and is still present.
// Return true if the property was deleted or is not present.
// https://tc39.es/ecma262/#table-essential-internal-methods
func Delete(propertyKey string) *CompletionRecord {
	return nil
}

// OwnPropertyKeys :
// Return a List whose elements are all the property keys for the object.
// https://tc39.es/ecma262/#table-essential-internal-methods
func OwnPropertyKeys() *CompletionRecord {
	return nil
}
