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

func (r *CompletionRecord) GetType() CompletionType {
	return r._type
}

func (r *CompletionRecord) GetTarget() string {
	return r.target
}
