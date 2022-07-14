package context

import (
	"duck/ling/parser/ast/node"
	"duck/ling/syntax/type"
)

// Realm :
// https://tc39.es/ecma262/#sec-code-realms
type Realm struct {
	Intrinsics        *IntrinsicsRecord
	GlobalObject      *_type.ECMAObject
	GlobalEnvironment *GlobalEnvironmentRecord
	TemplateMap       []*TemplateMapRecord
	HostDefined       interface{}
}

// NewRealm :
// Create a new realm.
// https://tc39.es/ecma262/#sec-createrealm
func NewRealm() *Realm {
	return &Realm{
		Intrinsics:        &IntrinsicsRecord{},
		GlobalObject:      nil,
		GlobalEnvironment: nil,
		TemplateMap:       []*TemplateMapRecord{},
		HostDefined:       nil,
	}
}

// IntrinsicsRecord :
// https://tc39.es/ecma262/#table-well-known-intrinsic-objects
// TODO
type IntrinsicsRecord struct {
	//AggregateError
}

// GlobalEnvironmentRecord TODO
type GlobalEnvironmentRecord struct{}

// TemplateMapRecord TODO
type TemplateMapRecord struct {
	Site  *node.ASTNode
	Array *_type.ECMAObject
}
