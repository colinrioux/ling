package parser

import "duck/ling/parser/ast"

// parseDeclaration :
// Parses a declaration.
//
// 	Declaration : HoistableDeclaration
//	            | ClassDeclaration
//	            | LexicalDeclaration
// https://tc39.es/ecma262/#prod-Declaration
func (parser *Parser) parseDeclaration() *ast.Node {
	return nil
}
