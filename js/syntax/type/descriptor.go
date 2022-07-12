package _type

type PropertyDescriptorType uint8

const (
	DataPropertyDescriptor PropertyDescriptorType = iota
	AccessorPropertyDescriptor
	MethodPropertyDescriptor //unused
	GenericPropertyDescriptor
)

// PropertyDescriptor :
// The PropertyDescriptor type is used to explain the manipulation and reification of ECMAObject property attributes.
// A PropertyDescriptor is a ECMARecord with zero or more fields,
// where each field's name is an attribute name and its value is a corresponding attribute value.
// https://tc39.es/ecma262/#sec-property-descriptor-specification-type
type PropertyDescriptor struct {
	*ECMARecord
	_type PropertyDescriptorType
}

// NewPropertyDescriptor :
// Creates a new property descriptor.
func NewPropertyDescriptor(_type PropertyDescriptorType, value interface{}) *PropertyDescriptor {
	return &PropertyDescriptor{
		ECMARecord: NewECMARecord(value),
		_type:      _type,
	}
}

func NewDataPropertyDescriptor(value *DataProperty) *PropertyDescriptor {
	return &PropertyDescriptor{
		ECMARecord: NewECMARecord(value),
		_type:      DataPropertyDescriptor,
	}
}

func NewAccessorPropertyDescriptor(value *AccessorProperty) *PropertyDescriptor {
	return &PropertyDescriptor{
		ECMARecord: NewECMARecord(value),
		_type:      AccessorPropertyDescriptor,
	}
}

func NewMethodPropertyDescriptor(value *MethodProperty) *PropertyDescriptor {
	return &PropertyDescriptor{
		ECMARecord: NewECMARecord(value),
		_type:      MethodPropertyDescriptor,
	}
}

// FromObject :
// Converts an object to a property descriptor.
// https://tc39.es/ecma262/#sec-topropertydescriptor
// TODO
func FromObject(obj *ECMAObject) *PropertyDescriptor {
	return nil
}

// IsAccessor :
// Checks if this descriptor is an access descriptor.
// https://tc39.es/ecma262/#sec-isaccessordescriptor
func (d *PropertyDescriptor) IsAccessor() bool {
	return d._type == AccessorPropertyDescriptor
}

// IsData :
// Checks if this descriptor is a data descriptor.
// https://tc39.es/ecma262/#sec-isdatadescriptor
func (d *PropertyDescriptor) IsData() bool {
	return d._type == DataPropertyDescriptor
}

// IsGeneric :
// Checks if this descriptor is a generic descriptor.
// https://tc39.es/ecma262/#sec-isgenericdescriptor
func (d *PropertyDescriptor) IsGeneric() bool {
	return d._type == GenericPropertyDescriptor
}

// IsMethod :
// Checks if this descriptor is a method descriptor.
func (d *PropertyDescriptor) IsMethod() bool {
	return d._type == MethodPropertyDescriptor
}

// ToObject :
// Converts a property descriptor to an object.
// https://tc39.es/ecma262/#sec-frompropertydescriptor
// TODO
func (d *PropertyDescriptor) ToObject() *ECMAObject {
	return nil
}

// Complete :
// Tags a property descriptor as complete.
// https://tc39.es/ecma262/#sec-completepropertydescriptor
// TODO
func (d *PropertyDescriptor) Complete() {

}
