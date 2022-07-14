package _type

// ReferenceRecord :
// The ReferenceRecord type is used to explain the behaviour
// of such operators as delete, typeof, the assignment operators,
// the super keyword and other language features.
// https://tc39.es/ecma262/#sec-reference-record-specification-type
type ReferenceRecord struct {
	Base           interface{} // *ECMAPrimitive or *EnvironmentRecord
	ReferencedName interface{} // string if Base is EnvironmentRecord; else Symbol or PrivateName
	Strict         bool
	ThisValue      *ECMAPrimitive // holds value for "this" at time of creation
}

// NewReferenceRecord :
// Create a new ReferenceRecord.
// TODO
func NewReferenceRecord() *ReferenceRecord {
	return &ReferenceRecord{}
}

// NewPrivateReferenceRecord :
// Creates a new private ReferenceRecord.
// https://tc39.es/ecma262/#sec-makeprivatereference
// TODO
func NewPrivateReferenceRecord(baseValue *ECMAPrimitive, privateIdentifier string) *ReferenceRecord {
	return nil
}

// IsPropertyReference :
// Checks if this reference refers to a property.
// https://tc39.es/ecma262/#sec-ispropertyreference
// TODO
func (r *ReferenceRecord) IsPropertyReference() bool {
	if r.IsUnresolvableReference() {
		return false
	}

	// TODO if r.Base == EnvironmentRecord, return false
	return true
}

// IsUnresolvableReference :
// Checks if this reference's Base is unresolvable (nil).
// https://tc39.es/ecma262/#sec-isunresolvablereference
func (r *ReferenceRecord) IsUnresolvableReference() bool {
	return r.Base == nil
}

// IsSuperReference :
// Checks if this reference's ThisValue is not nil.
// https://tc39.es/ecma262/#sec-issuperreference
func (r *ReferenceRecord) IsSuperReference() bool {
	return r.ThisValue != nil
}

// IsPrivateReference :
// Checks if this reference's ReferencedName is a PrivateName.
// https://tc39.es/ecma262/#sec-isprivatereference
// TODO
func (r *ReferenceRecord) IsPrivateReference() bool {
	return false
}

// GetValue :
// Returns a normal completion with the value for this reference (if any).
// Otherwise, returns an abort completion.
// https://tc39.es/ecma262/#sec-getvalue
// TODO
func (r *ReferenceRecord) GetValue() *CompletionRecord {
	return nil
}

// PutValue :
// Sets a value to w and returns
// a normal completion with the new value for this reference (if any).
// Otherwise, returns an abort completion.
// https://tc39.es/ecma262/#sec-putvalue
// TODO
func (r *ReferenceRecord) PutValue(w interface{}) *CompletionRecord {
	return nil
}

// GetThisValue :
// Returns a normal completion with the value for this reference's parent (if any).
// Otherwise, returns an abort completion.
// https://tc39.es/ecma262/#sec-getthisvalue
// TODO
func (r *ReferenceRecord) GetThisValue() *CompletionRecord {
	return nil
}

// InitializeReferenceBinding :
// Initializes the Base binding if Base is an EnvironmentRecord.
// https://tc39.es/ecma262/#sec-initializereferencedbinding
// TODO
func (r *ReferenceRecord) InitializeReferenceBinding(w interface{}) *CompletionRecord {
	return nil
}
