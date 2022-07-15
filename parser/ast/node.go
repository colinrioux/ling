package ast

import (
	"duck/ling/lexer/token"
	"fmt"
)

// Node :
// Base Node type for all AST nodes.
type Node struct {
	Type     NodeType
	Children []*Node
	Token    *token.Token
}

// NewNode :
// Create a new Node.
func NewNode(left *Node, right *Node, token *token.Token) *Node {
	return &Node{
		Type:     UnknownNodeType,
		Children: []*Node{left, right},
		Token:    token,
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
	return fmt.Sprintf("ASTNode(%v,%v)", node.ChildrenToString(), *node.Token)
}

func (node *Node) ChildrenToString() string {
	str := "Children("
	for i, child := range node.Children {
		str += string(rune(i)) + "=" + child.String() + ","
	}
	str += ")"
	return str
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
