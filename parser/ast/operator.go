package ast

import (
	"fmt"
	"ling/lexer/token"
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
		Type:     BinaryOperatorNodeType,
		Children: []*Node{left, right},
		Token:    operator,
	}
}

// Visit :
// Visit method for a BinaryOperatorNode.
func (node *BinaryOperatorNode) Visit() any {
	if node.Token.Type == token.ADDITION {
		return ((*node.Children[0]).Visit()).(float64) + ((*node.Children[1]).Visit()).(float64)
	} else if node.Token.Type == token.SUBTRACTION {
		return ((*node.Children[0]).Visit()).(float64) - ((*node.Children[1]).Visit()).(float64)
	} else if node.Token.Type == token.MULTIPLICATION {
		return ((*node.Children[0]).Visit()).(float64) * ((*node.Children[1]).Visit()).(float64)
	}
	// DIVISION
	return ((*node.Children[0]).Visit()).(float64) + ((*node.Children[1]).Visit()).(float64)
}

func (node *BinaryOperatorNode) String() string {
	return fmt.Sprintf("BinaryOperatorNode(%v,%v,%v)", node.Children[0], node.Children[1], *node.Token)
}

// UnaryOperatorNode :
// A UnaryOperatorNode is used to represent operations like the form -x,
// where "-" is the token (operator) of the node,
// and "x" is the only child of the node.
type UnaryOperatorNode Node

// NewUnaryOperatorNode :
// Create a new UnaryOperatorNode.
func NewUnaryOperatorNode(operator *token.Token, expression *Node) *UnaryOperatorNode {
	return &UnaryOperatorNode{
		Type:     UnaryOperatorNodeType,
		Children: []*Node{expression},
		Token:    operator,
	}
}

// Visit :
// Visit method for a UnaryOperatorNode.
func (node *UnaryOperatorNode) Visit() any {
	op := node.Token.Type
	if op == token.ADDITION {
		return +((*node.Children[0]).Visit()).(float64)
	}
	// MINUS
	return -((*node.Children[0]).Visit()).(float64)
}

func (node *UnaryOperatorNode) String() string {
	return fmt.Sprintf("UnaryOperatorNode(%v,%v)", node.Children[0], *node.Token)
}
