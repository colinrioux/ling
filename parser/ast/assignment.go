package ast

import (
	"duck/ling/lexer/token"
	"fmt"
)

// AssignmentNode :
// An assignment node is used to represent an expression: var x = y,
// where "var x" is the left child of this node,
// "=" is the token of this node,
// and "y" is the resulting expression node.
type AssignmentNode Node

// NewAssignmentNode :
// Create a new assignment node.
func NewAssignmentNode(left *VariableNode, operator *token.Token, right *Node) *AssignmentNode {
	return &AssignmentNode{
		Type:  AssignmentNodeType,
		Left:  (*Node)(left),
		Right: right,
		Token: operator,
	}
}

func (node *AssignmentNode) String() string {
	return fmt.Sprintf("AssignmentNode(Var=%v,Expr=%v,Tok=%v)", node.Left, node.Right, *node.Token)
}
