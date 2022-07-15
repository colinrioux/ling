package ast

import "fmt"

// BlockNode :
// A BlockNode is used to represent a list of statements surrounded by {}.
// This list of statements are referenced by this node's Children.
type BlockNode Node

// NewBlockNode :
// Creates a new BlockNode.
func NewBlockNode() *BlockNode {
	return &BlockNode{
		Type:     BlockNodeType,
		Children: []*Node{},
	}
}

func (node *BlockNode) String() string {
	return fmt.Sprintf("BlockNode(%v)", (*Node)(node).ChildrenToString())
}
