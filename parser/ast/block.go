package ast

// BlockNode :
// A BlockNode is used to represent a list of statements surrounded by {}.
type BlockNode Node

// NewBlockNode :
// Creates a new BlockNode.
func NewBlockNode() *BlockNode {
	return &BlockNode{}
}
