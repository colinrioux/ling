package ast

import (
	"fmt"
	"github.com/google/uuid"
	"ling/lexer/token"
)

// VariableNode :
// A VariableNode is used to represent a variable as var x.
type VariableNode Node

// NewVariableNode :
// Create a new VariableNode.
func NewVariableNode(token *token.Token, name string) *VariableNode {
	return &VariableNode{
		Type:  VariableNodeType,
		Token: token,
		Name:  name,
	}
}

func (node *VariableNode) String() string {
	return fmt.Sprintf("VariableNode(%v, %v)", *node.Token, node.Name)
}

// VariableDeclarationNode :
// A VariableDeclarationNode is used to represent 1 or more declarations.
type VariableDeclarationNode Node

// NewVariableDeclarationNode :
// Create a new VariableDeclarationNode.
func NewVariableDeclarationNode() *VariableDeclarationNode {
	return &VariableDeclarationNode{
		Children: []*Node{},
		Name:     uuid.NewString(),
	}
}

func (node *VariableDeclarationNode) String() string {
	return fmt.Sprintf("VariableDeclarationNode(%v)", (*Node)(node).ChildrenToString())
}

// AddDeclaration :
// Adds a declaration to the declaration list.
func (node *VariableDeclarationNode) AddDeclaration(declaration *Node) {
	node.Children = append(node.Children, declaration)
}
