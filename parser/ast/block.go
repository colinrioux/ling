package ast

import (
	"fmt"
	"github.com/google/uuid"
)

// BlockNode :
// A BlockNode is used to represent a list of statements surrounded by {}.
// This list of statements are referenced by this node's Children.
type BlockNode Node

// NewBlockNode1 :
// Creates a new BlockNode.
func NewBlockNode1(id string) *BlockNode {
	return &BlockNode{
		Type:     BlockNodeType,
		Children: []*Node{},
		Name:     id,
	}
}

// NewBlockNode2 :
// Creates a new BlockNode with a randomly generated id.
func NewBlockNode2() *BlockNode {
	return &BlockNode{
		Type:     BlockNodeType,
		Children: []*Node{},
		Name:     uuid.NewString(),
	}
}

func (node *BlockNode) String() string {
	return fmt.Sprintf("BlockNode(%v, %v)", node.Name, (*Node)(node).ChildrenToString())
}

// Visit :
// Visit method for BlockNode.
func (node *BlockNode) Visit() interface{} {
	for _, child := range node.Children {
		if child != nil {
			child.Visit()
		}
	}
	return nil
}
