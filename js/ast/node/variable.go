package node

type IAssignmentNode interface {
	IASTNode
}

type AssignmentNode struct {
	*ASTNode
}

type IVariableNode interface {
	IASTNode
}

type VariableNode struct {
	*ASTNode
}
