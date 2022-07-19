package ast

import (
	"fmt"
	"ling/lexer/token"
)

// Node :
// Base Node type for all AST nodes.
type Node struct {
	Type     NodeType
	Children []*Node
	Token    *token.Token
	Name     string
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
	case BlockNodeType:
		return (*BlockNode)(node).Visit()
	}
	return nil
}

func (node *Node) String() string {
	switch node.Type {
	case NumberNodeType:
		return (*NumberNode)(node).String()
	case BinaryOperatorNodeType:
		return (*BinaryOperatorNode)(node).String()
	case UnaryOperatorNodeType:
		return (*UnaryOperatorNode)(node).String()
	case VariableNodeType:
		return (*VariableNode)(node).String()
	case AssignmentNodeType:
		return (*AssignmentNode)(node).String()
	case BlockNodeType:
		return (*BlockNode)(node).String()
	case VariableDeclarationNodeType:
		return (*VariableDeclarationNode)(node).String()
	}
	return fmt.Sprintf("ASTNode(%v,%v)", node.ChildrenToString(), node.Token)
}

func (node *Node) ChildrenToString() string {
	m := len(node.Children)
	str := "Children("
	for i, child := range node.Children {
		if child == nil {
			str += fmt.Sprintf("%d=nil", i)
			continue
		}
		str += fmt.Sprintf("%d=%v", i, child)
		if i != m-1 {
			str += ","
		}
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
	BlockNodeType
	VariableDeclarationNodeType
)

// EmptyNode :
// Used to represent useless nodes (with default initialization).
type EmptyNode Node
