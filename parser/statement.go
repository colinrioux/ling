package parser

import (
	"fmt"
	"ling/lexer/keyword"
	"ling/lexer/token"
	"ling/parser/ast"
	"log"
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
		log.Println("invalid syntax ps")
		return nil
	}

	if parser.CurrentToken.Type == token.LBRACE {
		return parser.parseBlockStatement()
	} else if parser.CurrentToken.Type == token.KEYWORD && parser.CurrentToken.Value == keyword.VAR {
		return parser.parseVariableStatement()
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
		log.Println("invalid syntax pes")
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
	root := ast.NewBlockNode2()
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
		parser.CurrentToken, _ = parser.Lexer.GetNextToken()
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
	node := ast.NewVariableDeclarationNode()
	if parser.CurrentToken.Type == token.KEYWORD && parser.CurrentToken.Value == keyword.VAR {
		parser.eat(token.KEYWORD)
		nodes := parser.parseVariableDeclarationList()
		for _, nde := range nodes {
			node.AddDeclaration(nde)
		}
	}
	return (*ast.Node)(node)
}

// parseExpressionStatement :
// Parses an expression statement.
//
//	ExpressionStatement :
//		[lookahead not in set({, function, async [no line terminator] function, class, let, [)]? Expression ;
// https://tc39.es/ecma262/#prod-ExpressionStatement
func (parser *Parser) parseExpressionStatement() *ast.Node {
	//lookAhead := parser.CurrentToken
	return nil
}

// parseIfStatement :
// Parses an if statement.
//
//	IfStatement : if ( Expression ) Statement else Statement
//	            | if ( Expression ) Statement [lookahead not equal else]
// https://tc39.es/ecma262/#prod-IfStatement
func (parser *Parser) parseIfStatement() *ast.Node {
	return nil
}

// parseIterationStatement :
// Parses an iteration statement.
//
//	IterationStatement : DoWhileStatement
//	                   | WhileStatement
//	                   | ForStatement
//	                   | ForInOfStatement
// https://tc39.es/ecma262/#prod-IterationStatement
func (parser *Parser) parseIterationStatement() *ast.Node {
	return nil
}

// parseDoWhileStatement :
// Parses a do while statement.
//
//	DoWhileStatement : do Statement while ( Expression ) ;
// https://tc39.es/ecma262/#prod-DoWhileStatement
func (parser *Parser) parseDoWhileStatement() *ast.Node {
	return nil
}

// parseWhileStatement :
// Parses a while statement.
//
//	WhileStatement : while ( Expression ) Statement
// https://tc39.es/ecma262/#prod-WhileStatement
func (parser *Parser) parseWhileStatement() *ast.Node {
	return nil
}

// parseForStatement :
// Parses a for statement.
//
//	ForStatement : for ( [lookahead not equal let [] Expression? ; Expression? ; Expression? ) Statement
//	             | for ( var VariableDeclarationList ; Expression? ; Expression? ) Statement
//	             | for ( LexicalDeclaration Expression? ; Expression? ) Statement
// https://tc39.es/ecma262/#prod-ForStatement
func (parser *Parser) parseForStatement() *ast.Node {
	return nil
}

// parseForInOfStatement :
// Parses a for-in, for-of, or for-await-of statement.
//
//	ForInOfStatement : for ( [lookahead not equal let [] LeftHandSideExpression in Expression ) Statement
//	                 | for ( var ForBinding in Expression ) Statement
//	                 | for ( ForDeclaration in Expression ) Statement
//	                 | for ( [lookahead not in set(let, async of)] LeftHandSideExpression of AssignmentExpression ) Statement
//	                 | for ( var ForBinding of AssignmentExpression ) Statement
//	                 | for ( ForDeclaration of AssignmentExpression ) Statement
//	                 | for await ( [lookahead not equal let] LeftHandSideExpression of AssignmentExpression ) Statement
//	                 | for await ( var ForBinding of AssignmentExpression ) Statement
//	                 | for ( ForDeclaration in Expression ) Statement
// https://tc39.es/ecma262/#prod-ForInOfStatement
func (parser *Parser) parseForInOfStatement() *ast.Node {
	return nil
}

// parseContinueStatement :
// Parses a continue statement.
//
//	ContinueStatement : continue ;
//	                  | continue [no LineTerminator] LabelIdentifier ;
// https://tc39.es/ecma262/#prod-ContinueStatement
func (parser *Parser) parseContinueStatement() *ast.Node {
	return nil
}

// parseBreakStatement :
// Parses a break statement.
//
//	BreakStatement : break ;
//	               | break [no LineTerminator] LabelIdentifier ;
// https://tc39.es/ecma262/#prod-BreakStatement
func (parser *Parser) parseBreakStatement() *ast.Node {
	return nil
}

// parseReturnStatement :
// Parses a return statement.
//
//	ReturnStatement : return ;
//	                | return [no LineTerminator] Expression ;
// https://tc39.es/ecma262/#prod-ReturnStatement
func (parser *Parser) parseReturnStatement() *ast.Node {
	return nil
}

// parseWithStatement :
// Parses a with statement.
//
//	WithStatement : with ( Expression ) Statement
// https://tc39.es/ecma262/#sec-with-statement LEGACY
func (parser *Parser) parseWithStatement() *ast.Node {
	return nil
}

// parseSwitchStatement :
// Parses a switch statement.
//
//	SwitchStatement : switch ( Expression ) CaseBlock
// https://tc39.es/ecma262/#prod-SwitchStatement
func (parser *Parser) parseSwitchStatement() *ast.Node {
	return nil
}
