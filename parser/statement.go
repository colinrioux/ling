package parser

import (
	"duck/ling/lexer/token"
	"duck/ling/parser/ast"
)

// parseStatement :
// Parses a statement.
//
// 	 Statement : BlockStatement
//	           | VariableStatement
//	           | EmptyStatement
//	           | ExpressionStatement
//	           | IfStatement
//	           | BreakableStatement
//	           | ContinueStatement
//	           | BreakStatement
//	           | ReturnStatement
//	           | WithStatement
//	           | LabelledStatement
//	           | ThrowStatement
//	           | TryStatement
//	           | DebuggerStatement
// https://tc39.es/ecma262/#prod-Statement
func (parser *Parser) parseStatement() *ast.Node {
	if parser.CurrentToken.Type == token.LBRACE {
		return parser.parseBlockStatement()
	}
	return parser.parseEmptyStatement()
}

// parseEmptyStatement :
// Parses an empty statement.
//
// 	EmptyStatement : ;
// https://tc39.es/ecma262/#prod-EmptyStatement
func (parser *Parser) parseEmptyStatement() *ast.Node {
	parser.eat(token.SEMICOLON)
	return (*ast.Node)(&ast.EmptyNode{})
}

// parseBlockStatement :
// Parses a block statement.
//
// 	BlockStatement : Block
// https://tc39.es/ecma262/#prod-BlockStatement
func (parser *Parser) parseBlockStatement() *ast.Node {
	return parser.parseBlock()
}

// parseBlock :
// Parses a block.
//
// 	Block : { StatementList }
// https://tc39.es/ecma262/#prod-Block
func (parser *Parser) parseBlock() *ast.Node {
	parser.eat(token.LBRACE)
	nodes := parser.parseStatementList()
	parser.eat(token.RBRACE)

	// Build the AST node for the block
	root := ast.NewBlockNode()
	for _, node := range nodes {
		root.Children = append(root.Children, node)
	}

	return (*ast.Node)(root)
}

// parseStatementList :
// Parses a statement list.
//
// 	StatementList : StatementListItem
//	              | StatementList StatementListItem
// https://tc39.es/ecma262/#prod-StatementList
func (parser *Parser) parseStatementList() []*ast.Node {
	node := parser.parseStatementListItem()
	nodes := []*ast.Node{node}

	// Parse additional items if there are any.
	for parser.CurrentToken.Type != token.RBRACE {
		nodes = append(nodes, parser.parseStatementListItem())
	}

	return nodes
}

// parseStatementListItem :
// Parses a statement list item.
//
// 	StatementListItem : Statement
//	                  | Declaration
// https://tc39.es/ecma262/#prod-StatementListItem
func (parser *Parser) parseStatementListItem() *ast.Node {
	// TODO declaration
	return parser.parseStatement()
}

// parseVariableStatement :
// "var" variables are scoped to the running execution context.
// If declared, they are instantiated as undefined.
// https://tc39.es/ecma262/#sec-variable-statement
// TODO update comment
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
