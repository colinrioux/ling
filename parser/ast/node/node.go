package node

import (
	"duck/ling/lexer/token"
	"fmt"
)

type ASTNode struct {
	Type  ASTNodeType
	Left  *ASTNode
	Right *ASTNode
	Token *token.Token
}

func NewASTNode(left *ASTNode, right *ASTNode, token *token.Token) *ASTNode {
	return &ASTNode{
		Type:  UnknownNodeType,
		Left:  left,
		Right: right,
		Token: token,
	}
}

func (node *ASTNode) Visit() interface{} {
	switch node.Type {
	case NumberNodeType:
		return (*NumberNode)(node).Visit()
	case BinaryOperatorNodeType:
		return (*BinaryOperatorNode)(node).Visit()
	case UnaryOperatorNodeType:
		return (*UnaryOperatorNode)(node).Visit()
	}
	return nil
}

func (node *ASTNode) String() string {
	return fmt.Sprintf("ASTNode(%v,%v,%s)", node.Left, node.Right, *node.Token)
}

type ASTNodeType uint16

const (
	UnknownNodeType ASTNodeType = iota
	NumberNodeType
	BinaryOperatorNodeType
	UnaryOperatorNodeType
	VariableNodeType
)
