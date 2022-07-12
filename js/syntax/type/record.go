package _type

// ECMARecord :
// The Record type is used to describe data aggregations within the algorithms of this specification.
// A Record type value consists of one or more named fields.
// https://tc39.es/ecma262/#sec-list-and-record-specification-type
type ECMARecord struct {
	value interface{}
}

// NewECMARecord :
// Creates a new ECMA record.
func NewECMARecord(value interface{}) *ECMARecord {
	return &ECMARecord{
		value: value,
	}
}

func (r *ECMARecord) GetValue() interface{} {
	return r.value
}

// UnicodeRecord TODO
//type UnicodeRecord struct {
//	CodePoint           rune
//	CodeUnitCount       uint
//	IsUnpairedSurrogate bool
//}
type UnicodeRecord ECMARecord

//func (r *CompletionRecord) IsAbrupt() bool {
//	return (*r)["Type"] != NormalCompletion
//}
//

//func (r *CompletionRecord) UpdateEmpty(value interface{}) *CompletionRecord {
//	// TODO
//	//if (cr.Type == ReturnCompletion || cr.Type == ThrowCompletion) && len(value)
//	return nil
//}

// ReferenceRecord TODO
type ReferenceRecord struct {
	Base          *ECMAType
	ReferenceName *ECMAType
	Strict        bool
	ThisValue     *ECMAType
}

// PrivateElementRecord TODO
type PrivateElementRecord struct{}
