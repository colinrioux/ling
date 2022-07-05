package _type

type ECMARecord map[string]ECMAType

// UnicodeRecord TODO
//type UnicodeRecord struct {
//	CodePoint           rune
//	CodeUnitCount       uint
//	IsUnpairedSurrogate bool
//}
type UnicodeRecord ECMARecord

type CompletionType string

const (
	NormalCompletion   CompletionType = "normal"
	BreakCompletion    CompletionType = "break"
	ContinueCompletion CompletionType = "continue"
	ReturnCompletion   CompletionType = "return"
	ThrowCompletion    CompletionType = "throw"
)

// CompletionRecord :
// https://tc39.es/ecma262/#sec-completion-record-specification-type
type CompletionRecord ECMARecord

//func (r *CompletionRecord) IsAbrupt() bool {
//	return (*r)["Type"] != NormalCompletion
//}
//
//func NewCompletionRecord(_type interface{}, value interface{}, target interface{}) *CompletionRecord {
//	return &CompletionRecord{
//		"Type":   _type,
//		"Value":  value,
//		"Target": target,
//	}
//}

//func NewNormalCompletion(value interface{}) *CompletionRecord {
//	return NewCompletionRecord(NormalCompletion, value, "")
//}
//
//func NewThrowCompletion(value interface{}) *CompletionRecord {
//	return NewCompletionRecord(ThrowCompletion, value, "")
//}

func (r *CompletionRecord) UpdateEmpty(value interface{}) *CompletionRecord {
	// TODO
	//if (cr.Type == ReturnCompletion || cr.Type == ThrowCompletion) && len(value)
	return nil
}

// ReferenceRecord TODO
type ReferenceRecord struct {
	Base          *ECMAType
	ReferenceName *ECMAType
	Strict        bool
	ThisValue     *ECMAType
}
