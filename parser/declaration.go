package parser

import (
	"ling/lexer/token"
	"ling/parser/ast"
)

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

// parseVariableDeclarationList :
// Parses a variable declaration list.
//
//	VariableDeclarationList : VariableDeclaration
//	                        | VariableDeclarationList , VariableDeclaration
// https://tc39.es/ecma262/#prod-VariableDeclarationList
func (parser *Parser) parseVariableDeclarationList() []*ast.Node {
	node := parser.parseVariableDeclaration()
	nodes := []*ast.Node{node}

	for parser.CurrentToken.Type == token.COMMA {
		nodes = append(nodes, parser.parseVariableDeclaration())
	}

	return nodes
}

// parseVariableDeclaration :
// Parses a variable declaration.
//
//	VariableDeclaration : BindingIdentifier Initializer
// 	                    | BindingIdentifier
//	                    | BindingPattern Initializer
// https://tc39.es/ecma262/#prod-VariableDeclaration
func (parser *Parser) parseVariableDeclaration() *ast.Node {
	// TODO
	//// BindingIdentifier
	//if parser.CurrentToken.Type == token.IDENTIFIER {
	//	return parser.parseBindingIdentifier()
	//}
	return nil
}

// parseForDeclaration :
// Parses a for declaration.
//
//	ForDeclaration : LetOrConst ForBinding
// https://tc39.es/ecma262/#prod-ForDeclaration
func (parser *Parser) parseForDeclaration() *ast.Node {
	return nil
}
