package ast

import (
	"duck/ling/lexer/token"
	"fmt"
)

// Node :
// Base Node type for all AST nodes.
type Node struct {
	Type  NodeType
	Left  *Node
	Right *Node
	Token *token.Token
}

// NewNode :
// Create a new Node.
func NewNode(left *Node, right *Node, token *token.Token) *Node {
	return &Node{
		Type:  UnknownNodeType,
		Left:  left,
		Right: right,
		Token: token,
	}
}

// Visit :
// Visitation per node on AST traversal.
func (node *Node) Visit() interface{} {
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

func (node *Node) String() string {
	return fmt.Sprintf("ASTNode(%v,%v,%v)", node.Left, node.Right, *node.Token)
}

type NodeType uint16

const (
	UnknownNodeType NodeType = iota
	NumberNodeType
	BinaryOperatorNodeType
	UnaryOperatorNodeType
	VariableNodeType
	AssignmentNodeType
)
