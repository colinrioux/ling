package node

import (
	"duck/ling/lexer/token"
	"fmt"
)

type BinaryOperatorNode ASTNode

func NewBinaryOperatorNode(left *ASTNode, right *ASTNode, operator *token.Token) *BinaryOperatorNode {
	return &BinaryOperatorNode{
		Type:  BinaryOperatorNodeType,
		Left:  left,
		Right: right,
		Token: operator,
	}
}

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
	return fmt.Sprintf("BinaryOperatorNode(%v,%v,%s)", node.Left, node.Right, *node.Token)
}

type UnaryOperatorNode ASTNode

func NewUnaryOperatorNode(operator *token.Token, expression *ASTNode) *UnaryOperatorNode {
	return &UnaryOperatorNode{
		Type:  UnaryOperatorNodeType,
		Left:  nil,
		Right: expression,
		Token: operator,
	}
}

func (node *UnaryOperatorNode) Visit() any {
	op := node.Token.Type
	if op == token.ADD {
		return +((*node.Right).Visit()).(float64)
	}
	// MINUS
	return -((*node.Right).Visit()).(float64)
}

func (node *UnaryOperatorNode) String() string {
	return fmt.Sprintf("UnaryOperatorNode(%v,%s)", node.Right, *node.Token)
}
