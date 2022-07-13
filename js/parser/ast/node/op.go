package node

import (
	"duck/ling/js/lexer/token"
	"fmt"
)

type IBinaryOperatorNode interface {
	IASTNode
	GetOperator() *token.Token
}

type BinaryOperatorNode struct {
	*ASTNode
	operator *token.Token
}

func NewBinaryOperatorNode(left *IASTNode, right *IASTNode, operator *token.Token) *BinaryOperatorNode {
	base := NewASTNode(left, right, operator)
	return &BinaryOperatorNode{ASTNode: base, operator: operator}
}

func (node BinaryOperatorNode) GetOperator() *token.Token {
	return node.operator
}

func (node BinaryOperatorNode) Visit() any {
	if node.operator.Type == token.ADD {
		return ((*node.GetLeft()).Visit()).(int) + ((*node.GetRight()).Visit()).(int)
	} else if node.operator.Type == token.SUB {
		return ((*node.GetLeft()).Visit()).(int) - ((*node.GetRight()).Visit()).(int)
	} else if node.operator.Type == token.MUL {
		return ((*node.GetLeft()).Visit()).(int) * ((*node.GetRight()).Visit()).(int)
	}
	// DIV
	return ((*node.GetLeft()).Visit()).(int) + ((*node.GetRight()).Visit()).(int)
}

func (node BinaryOperatorNode) String() string {
	return fmt.Sprintf("BinaryOperatorNode(%v,%s)", node.ASTNode, node.operator)
}

type IUnaryOperatorNode interface {
	IASTNode
	GetOperator() *token.Token
	GetExpression() *IASTNode
}

type UnaryOperatorNode struct {
	*ASTNode
	operator   *token.Token
	expression *IASTNode
}

func NewUnaryOperatorNode(left *IASTNode, right *IASTNode, operator *token.Token, expression *IASTNode) *UnaryOperatorNode {
	base := NewASTNode(left, right, operator)
	return &UnaryOperatorNode{ASTNode: base, operator: operator, expression: expression}
}

func (node UnaryOperatorNode) GetOperator() *token.Token {
	return node.operator
}

func (node UnaryOperatorNode) GetExpression() *IASTNode {
	return node.expression
}

func (node UnaryOperatorNode) Visit() any {
	op := node.operator.Type
	if op == token.ADD {
		return +((*node.GetExpression()).Visit()).(int)
	}
	// MINUS
	return -((*node.GetExpression()).Visit()).(int)
}

func (node UnaryOperatorNode) String() string {
	return fmt.Sprintf("UnaryOperatorNode(%v,%s,%v)", node.ASTNode, node.operator, node.expression)
}

type INoOperationNode interface {
	IASTNode
}

type NoOperationNode struct {
	*ASTNode
}
