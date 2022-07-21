package ast

import (
	"fmt"
)

// DebuggerNode :
// A DebuggerNode is used to represent a debugger statement.
type DebuggerNode Node

// NewDebuggerNode :
// Creates a new DebuggerNode.
func NewDebuggerNode() *DebuggerNode {
	return &DebuggerNode{
		Type: DebuggerNodeType,
	}
}

func (node *DebuggerNode) String() string {
	return fmt.Sprintf("DebuggerNode()")
}

// TryNode :
// A TryNode is used to represent a try statement.
type TryNode Node

// NewTryNode :
// Creates a new TryNode.
func NewTryNode(block *Node, catch *Node, finally *Node) *TryNode {
	return &TryNode{
		Type:     TryNodeType,
		Children: []*Node{block, catch, finally},
	}
}

func (node *TryNode) String() string {
	return fmt.Sprintf("TryNode(%v)", (*Node)(node).ChildrenToString())
}
