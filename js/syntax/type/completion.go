package _type

type CompletionType string

const (
	NormalCompletion   CompletionType = "normal"
	BreakCompletion    CompletionType = "break"
	ContinueCompletion CompletionType = "continue"
	ReturnCompletion   CompletionType = "return"
	ThrowCompletion    CompletionType = "throw"
)

// CompletionRecord :
// The Completion Record specification type is used to explain the runtime propagation
// of values and control flow such as the behaviour of statements (break, continue, return and throw)
// that perform nonlocal transfers of control.
// https://tc39.es/ecma262/#sec-completion-record-specification-type
type CompletionRecord struct {
	Type   CompletionType
	Value  interface{}
	Target string
}

// NewCompletionRecord :
// Create a new completion record.
func NewCompletionRecord(_type CompletionType, value interface{}, target string) *CompletionRecord {
	return &CompletionRecord{
		Type:   _type,
		Value:  value,
		Target: target,
	}
}

// CopyCompletionRecord :
// Create a new completion record as a copy of another one.
func CopyCompletionRecord(other *CompletionRecord) *CompletionRecord {
	return NewCompletionRecord(other.Type, other.Value, other.Target)
}

// Await :
// https://tc39.es/ecma262/#await
// TODO
func Await(value interface{}) *CompletionRecord {
	return nil
}

// UpdateEmpty :
// Return a copy of a completion record with an updated value if not a return or throw type.
// https://tc39.es/ecma262/#sec-updateempty
func UpdateEmpty(completionRecord *CompletionRecord, value interface{}) *CompletionRecord {
	var record *CompletionRecord = CopyCompletionRecord(completionRecord)
	if record.Type == ReturnCompletion || record.Type == ThrowCompletion {
		return completionRecord
	}

	record.Value = value
	return record
}
