package _type

type CompletionType string

const (
	NormalCompletion   CompletionType = "normal"
	BreakCompletion    CompletionType = "break"
	ContinueCompletion CompletionType = "continue"
	ReturnCompletion   CompletionType = "return"
	ThrowCompletion    CompletionType = "throw"
)

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

// CompletionRecord :
// The Completion Record specification type is used to explain the runtime propagation
// of values and control flow such as the behaviour of statements (break, continue, return and throw)
// that perform nonlocal transfers of control.
// https://tc39.es/ecma262/#sec-completion-record-specification-type
type CompletionRecord struct {
	*ECMARecord
	_type  CompletionType
	target string
}

// NewCompletionRecord :
// Create a new completion record.
func NewCompletionRecord(_type CompletionType, value interface{}, target string) *CompletionRecord {
	return &CompletionRecord{
		ECMARecord: NewECMARecord(value),
		_type:      _type,
		target:     target,
	}
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
