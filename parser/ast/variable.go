package ast

import (
	"fmt"
	"ling/lexer/token"
)

// VariableNode :
// A VariableNode is used to represent a variable as var/let/const x.
type VariableNode Node

// NewVariableNode :
// Create a new VariableNode.
func NewVariableNode(token *token.Token, name string) *VariableNode {
	return &VariableNode{
		Type:  VariableNodeType,
		Token: token,
		Name:  name,
	}
}

func (node *VariableNode) String() string {
	return fmt.Sprintf("VariableNode(%v, %v)", *node.Token, node.Name)
}
