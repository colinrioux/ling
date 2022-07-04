package node

import (
	"duck/ling/js/ast/token"
	"fmt"
)

type INumberNode interface {
	IASTNode
	GetValue() any
}

type NumberNode struct {
	*ASTNode
	value any
}

func NewNumberNode(token *token.Token) *NumberNode {
	base := NewASTNode(nil, nil, token)
	return &NumberNode{ASTNode: base, value: token.GetValue()}
}

func (node NumberNode) GetValue() any {
	return node.value
}

func (node NumberNode) String() string {
	return fmt.Sprintf("NumberNode(%v,%s)", node.ASTNode, node.value)
}

func (node NumberNode) Visit() any {
	return node.value
}
