package _type

type ECMAType uint32

const (
	UNDEFINED ECMAType = iota
	NULL
	BOOLEAN
	STRING
	SYMBOL
	NUMBER
	BIGINT
	OBJECT
)

var typeMap = map[ECMAType]interface{}{
	UNDEFINED: nil,
	NULL:      nil,
	BOOLEAN:   func() {},
	STRING:    func() {},
	SYMBOL:    func() {},
	NUMBER:    func() {},
	BIGINT:    func() {},
	OBJECT:    func() {},
}

// IsType :
// Check if a type is valid
func IsType(_type ECMAType) bool {
	_, ok := typeMap[_type]
	return ok
}

// New :
// Create a new ECMAType
func New(_type ECMAType) {

}
