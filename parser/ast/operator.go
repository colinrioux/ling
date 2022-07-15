package ast

import (
	"duck/ling/lexer/token"
	"fmt"
)

// BinaryOperatorNode :
// A BinaryOperatorNode is used to represent operations like the form x * y,
// where "x" is the left child of the node,
// "*" is the token (operator) of the node,
// and "y" is the right child of the node.
type BinaryOperatorNode Node

// NewBinaryOperatorNode :
// Create a new BinaryOperatorNode.
func NewBinaryOperatorNode(left *Node, right *Node, operator *token.Token) *BinaryOperatorNode {
	return &BinaryOperatorNode{
		Type:  BinaryOperatorNodeType,
		Left:  left,
		Right: right,
		Token: operator,
	}
}

// Visit :
// Visit method for a BinaryOperatorNode.
func (node *BinaryOperatorNode) Visit() any {
	if node.Token.Type == token.ADD {
		return ((*node.Left).Visit()).(float64) + ((*node.Right).Visit()).(float64)
	} else if node.Token.Type == token.SUB {
		return ((*node.Left).Visit()).(float64) - ((*node.Right).Visit()).(float64)
	} else if node.Token.Type == token.MUL {
		return ((*node.Left).Visit()).(float64) * ((*node.Right).Visit()).(float64)
	}
	// DIV
	return ((*node.Left).Visit()).(float64) + ((*node.Right).Visit()).(float64)
}

func (node *BinaryOperatorNode) String() string {
	return fmt.Sprintf("BinaryOperatorNode(%v,%v,%v)", node.Left, node.Right, *node.Token)
}

// UnaryOperatorNode :
// A UnaryOperatorNode is used to represent operations like the form -x,
// where "-" is the token (operator) of the node,
// and "x" is the right child of the node.
type UnaryOperatorNode Node

// NewUnaryOperatorNode :
// Create a new UnaryOperatorNode.
func NewUnaryOperatorNode(operator *token.Token, expression *Node) *UnaryOperatorNode {
	return &UnaryOperatorNode{
		Type:  UnaryOperatorNodeType,
		Left:  nil,
		Right: expression,
		Token: operator,
	}
}

// Visit :
// Visit method for a UnaryOperatorNode.
func (node *UnaryOperatorNode) Visit() any {
	op := node.Token.Type
	if op == token.ADD {
		return +((*node.Right).Visit()).(float64)
	}
	// MINUS
	return -((*node.Right).Visit()).(float64)
}

func (node *UnaryOperatorNode) String() string {
	return fmt.Sprintf("UnaryOperatorNode(%v,%v)", node.Right, *node.Token)
}
