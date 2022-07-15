package ast

import (
	"duck/ling/lexer/token"
	"fmt"
)

// VariableNode :
// A variable node is used to represent a variable as var/let x.
type VariableNode Node

// NewVariableNode :
// Create a new VariableNode.
func NewVariableNode(token *token.Token) *VariableNode {
	return &VariableNode{
		Type:  VariableNodeType,
		Left:  nil,
		Right: nil,
		Token: token,
	}
}

func (node *VariableNode) String() string {
	return fmt.Sprintf("VariableNode(%v)", *node.Token)
}
