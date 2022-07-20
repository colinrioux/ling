package parser

import (
	"ling/lexer/token"
	"ling/parser/ast"
)

// parseBindingPattern :
// Parses a binding pattern.
//
//	BindingPattern : ObjectBindingPattern
//	               | ArrayBindingPattern
// https://tc39.es/ecma262/#prod-BindingPattern
func (parser *Parser) parseBindingPattern() *ast.Node {
	return nil
}

// parseObjectBindingPattern :
// Parses an object binding pattern.
//
//	ObjectBindingPattern : { }
//	                     | { BindingRestProperty }
//	                     | { BindingPropertyList }
//	                     | { BindingPropertyList , BindingRestProperty? }
// https://tc39.es/ecma262/#prod-ObjectBindingPattern
func (parser *Parser) parseObjectBindingPattern() *ast.Node {
	parser.eat(token.LBRACE)
	return nil
}

// parseArrayBindingPattern :
// Parses an array binding pattern.
//
//	ArrayBindingPattern : [ Elision? BindingRestElement? ]
//	                    | [ BindingElementList ]
//	                    | [ BindingElementList , Elision? BindingRestElement? ]
// https://tc39.es/ecma262/#prod-ArrayBindingPattern
func (parser *Parser) parseArrayBindingPattern() *ast.Node {
	return nil
}

// parseBindingRestProperty :
// Parses a binding rest property.
//
//	BindingRestProperty : ... BindingIdentifier
// https://tc39.es/ecma262/#prod-BindingRestProperty
func (parser *Parser) parseBindingRestProperty() *ast.Node {
	parser.eat(token.SPREAD)
	return parser.parseBindingIdentifier()
}

// parseBindingPropertyList :
// Parses a binding property list.
//
//	BindingPropertyList : BindingProperty
//	                    | BindingPropertyList , BindingProperty
// https://tc39.es/ecma262/#prod-BindingPropertyList
func (parser *Parser) parseBindingPropertyList() []*ast.Node {
	return nil
}

// parseBindingElementList :
// Parses a binding element list.
//
//	BindingElementList : BindingElisionElement
// 	                   | BindingElementList , BindingElisionElement
// https://tc39.es/ecma262/#prod-BindingElementList
func (parser *Parser) parseBindingElementList() []*ast.Node {
	return nil
}

// parseBindingElisionElement :
// Parses a binding elision element.
//
//	BindingElisionElement : Elision? BindingElement
// https://tc39.es/ecma262/#prod-BindingElisionElement
func (parser *Parser) parseBindingElisionElement() *ast.Node {
	return nil
}

// parseBindingProperty :
// Parses a binding property.
//
//	BindingProperty : SingleNameBinding
// 	                | PropertyName : BindingElement
// https://tc39.es/ecma262/#prod-BindingProperty
func (parser *Parser) parseBindingProperty() *ast.Node {
	return nil
}

// parseBindingElement :
// Parses a binding element.
//
//	BindingElement : SingleNameBinding
//	               | BindingPattern Initializer?
// https://tc39.es/ecma262/#prod-BindingElement
func (parser *Parser) parseBindingElement() *ast.Node {
	return nil
}

// parseSingleNameBinding :
// Parses a single name binding.
//
//	SingleNameBinding : BindingIdentifier Initializer?
// https://tc39.es/ecma262/#prod-SingleNameBinding
func (parser *Parser) parseSingleNameBinding() *ast.Node {
	return nil
}

// parseBindingRestElement :
// Parses a binding rest element.
//
//	BindingRestElement : ... BindingIdentifier
//	                   | ... BindingPattern
// https://tc39.es/ecma262/#prod-BindingRestElement
func (parser *Parser) parseBindingRestElement() *ast.Node {
	return nil
}
