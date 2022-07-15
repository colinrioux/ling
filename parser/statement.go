package parser

import (
	"duck/ling/parser/ast"
)

// parseVariableStatement :
// "var" variables are scoped to the running execution context.
// If declared, they are instantiated as undefined.
// https://tc39.es/ecma262/#sec-variable-statement
func (parser *Parser) parseVariableStatement() *ast.Node {
	// Get the variable name.
	//nameToken := parser.parseNextToken()
	//if nameToken.Type != token.IDENTIFIER {
	//	// TODO throw error, incorrect
	//	return nil
	//}
	//return (*ast.Node)(ast.NewVariableNode())
	return nil
}
