package ast

import (
	"duck/ling/lexer/token"
	"duck/ling/syntax/type"
	"fmt"
)

// NumberNode :
// A leaf node used to represent numbers as the token of the node.
type NumberNode Node

// NewNumberNode :
// Create a new NumberNode.
func NewNumberNode(value *token.Token) *NumberNode {
	return &NumberNode{
		Type:  NumberNodeType,
		Left:  nil,
		Right: nil,
		Token: value,
	}
}

func (node NumberNode) String() string {
	return fmt.Sprintf("NumberNode(%v)", *node.Token)
}

// Visit :
// Visit method for a NumberNode.
func (node NumberNode) Visit() interface{} {
	return node.Token.Value.(*_type.ECMANumber).Value
}
