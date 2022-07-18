package parser

import (
	"fmt"
	"ling/lexer/keyword"
	"ling/lexer/token"
	"ling/parser/ast"
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
	if parser.CurrentToken == nil {
		// TODO error handling
		fmt.Println("invalid syntax ps")
		return nil
	}

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
	if parser.CurrentToken == nil {
		// TODO error handling
		fmt.Println("invalid syntax pes")
		return nil
	}
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
	if parser.CurrentToken == nil {
		// TODO error handling
		fmt.Println("invalid syntax psl")
		return nodes
	}

	for parser.CurrentToken.Type != token.EOF && parser.CurrentToken.Type != token.RBRACE {
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
// Parses a variable statement.
//
//	VariableStatement : var VariableDeclarationList
// https://tc39.es/ecma262/#sec-variable-statement
// https://tc39.es/ecma262/#prod-VariableStatement
func (parser *Parser) parseVariableStatement() *ast.Node {
	if parser.CurrentToken.Type == token.KEYWORD && parser.CurrentToken.Value == keyword.VAR {
		parser.eat(token.KEYWORD)
		//nodes := parser.parseVariableDeclarationList()
		// TODO push variable decls onto ast somehow?
	}
	return nil
}
