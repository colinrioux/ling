package parser

import (
	"duck/ling/lexer/token"
	"duck/ling/parser/ast"
)

// parseIdentifierReference :
// Parses an identifier reference.
//
//	IdentifierReference : Identifier
//	                    | yield
//	                    | await
// https://tc39.es/ecma262/#prod-IdentifierReference
func (parser *Parser) parseIdentifierReference() *ast.Node {
	if parser.CurrentToken.Type == token.KEYWORD {
		// TODO handle await and yield
	}
	return parser.parseIdentifier()
}

// parseBindingIdentifier :
// Parses a binding identifier.
//
// 	BindingIdentifier : Identifier
//	                  | yield
//	                  | await
// https://tc39.es/ecma262/#prod-BindingIdentifier
func (parser *Parser) parseBindingIdentifier() *ast.Node {
	if parser.CurrentToken.Type == token.KEYWORD {
		// TODO handle await and yield
	}
	return parser.parseIdentifier()
}

// parseLabelIdentifier :
// Parses a label identifier.
//
//	LabelIdentifier : Identifier
//	               | yield
//	               | await
// https://tc39.es/ecma262/#prod-LabelIdentifier
func (parser *Parser) parseLabelIdentifier() *ast.Node {
	if parser.CurrentToken.Type == token.KEYWORD {
		// TODO handle await and yield
	}
	return parser.parseIdentifier()
}

// parseIdentifier :
// Parses an identifier.
//
// 	Identifier : IdentifierName but not ReservedKeyword
// https://tc39.es/ecma262/#prod-Identifier
func (parser *Parser) parseIdentifier() *ast.Node {
	return nil
}
