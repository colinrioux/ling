package ast

import (
	"fmt"
	"ling/lexer/token"
)

// NumberNode :
// A leaf node used to represent numbers as the token of the node.
type NumberNode Node

// NewNumberNode :
// Create a new NumberNode.
func NewNumberNode(value *token.Token) *NumberNode {
	return &NumberNode{
		Type:  NumberNodeType,
		Token: value,
	}
}

func (node NumberNode) String() string {
	return fmt.Sprintf("NumberNode(%v)", *node.Token)
}

// Visit :
// Visit method for a NumberNode.
func (node NumberNode) Visit() float64 {
	return node.Token.Value.(float64)
}
