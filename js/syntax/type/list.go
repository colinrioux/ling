package _type

type IECMAList interface {
	Concat(other *ECMAList) *ECMAList
}

type ECMAList []interface{}

// Concat :
// Returns a new list of two lists concatenated with one another.
func (l *ECMAList) Concat(other *ECMAList) *ECMAList {
	newList := append(*l, *other...)
	return &newList
}
