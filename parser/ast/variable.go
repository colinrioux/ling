package ast

import (
	"duck/ling/lexer/token"
)

type AssignmentNode ASTNode

type VariableNode ASTNode

func NewVariableNode(token *token.Token) *VariableNode {
	return &VariableNode{
		Type:  VariableNodeType,
		Left:  nil,
		Right: nil,
		Token: token,
	}
}
