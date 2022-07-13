package node

import (
	"duck/ling/js/lexer/token"
	"fmt"
)

type IASTNode interface {
	GetToken() *token.Token
	GetLeft() *IASTNode
	GetRight() *IASTNode
	Visit() any
}

type ASTNode struct {
	left  *IASTNode
	right *IASTNode
	token *token.Token
}

func NewASTNode(left *IASTNode, right *IASTNode, token *token.Token) *ASTNode {
	return &ASTNode{left: left, right: right, token: token}
}

func (node ASTNode) GetToken() *token.Token {
	return node.token
}

func (node ASTNode) GetLeft() *IASTNode {
	return node.left
}

func (node ASTNode) GetRight() *IASTNode {
	return node.right
}

func (node ASTNode) Visit() any {
	return nil
}

func (node ASTNode) String() string {
	return fmt.Sprintf("ASTNode(%v,%v,%s)", node.left, node.right, *node.token)
}
