package node

import (
	"duck/ling/js/ast/token"
	"fmt"
)

type INumberNode interface {
	IASTNode
}

type NumberNode struct {
	*ASTNode
	Value interface{}
}

func NewNumberNode(token *token.Token) *NumberNode {
	base := NewASTNode(nil, nil, token)
	return &NumberNode{ASTNode: base, Value: token.Value}
}

func (node NumberNode) String() string {
	return fmt.Sprintf("NumberNode(%v,%s)", node.ASTNode, node.Value)
}

func (node NumberNode) Visit() interface{} {
	return node.Value
}
