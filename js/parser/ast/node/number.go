package node

import (
	"duck/ling/js/lexer/token"
	_type "duck/ling/js/syntax/type"
	"fmt"
)

type NumberNode ASTNode

func NewNumberNode(value *token.Token) *NumberNode {
	return &NumberNode{
		Type:  NumberNodeType,
		Left:  nil,
		Right: nil,
		Token: value,
	}
}

func (node NumberNode) String() string {
	return fmt.Sprintf("NumberNode(%v)", *node.Token)
}

func (node NumberNode) Visit() interface{} {
	return node.Token.Value.(*_type.ECMANumber).Value
}
