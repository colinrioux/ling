package parser

import (
	"ling/lexer/token"
	"ling/parser/ast"
)

// parseInitializer :
// Parses an initializer.
//
// 	Initializer : = AssignmentExpression
// https://tc39.es/ecma262/#prod-Initializer
func (parser *Parser) parseInitializer() *ast.Node {
	parser.eat(token.ASSIGN)
	return parser.parseAssignmentExpression()
}
