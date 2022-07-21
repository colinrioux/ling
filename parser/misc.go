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

// parseCaseBlock :
// Parses a case block.
//
//	CaseBlock : { CaseClauses? }
//	          | { CaseClauses? DefaultClause CaseClauses? }
// https://tc39.es/ecma262/#prod-CaseBlock
func (parser *Parser) parseCaseBlock() *ast.Node {
	return nil
}

// parseCaseClauses :
// Parses case clauses.
//
//	CaseClauses : CaseClause
//	            | CaseClauses CaseClause
// https://tc39.es/ecma262/#prod-CaseClauses
func (parser *Parser) parseCaseClauses() []*ast.Node {
	return nil
}

// parseCaseClause :
// Parses a case clause.
//
//	CaseClause : case Expression : StatementList?
// https://tc39.es/ecma262/#prod-CaseClause
func (parser *Parser) parseCaseClause() *ast.Node {
	return nil
}

// parseDefaultClause :
// Parses a default clause.
//
// 	DefaultClause : default : StatementList?
// https://tc39.es/ecma262/#prod-DefaultClause
func (parser *Parser) parseDefaultClause() *ast.Node {
	return nil
}

// parseLabelledItem :
// Parses a labelled item.
//
//	LabelledItem : Statement
//	             | FunctionDeclaration
// https://tc39.es/ecma262/#prod-LabelledItem
func (parser *Parser) parseLabelledItem() *ast.Node {
	return nil
}

// parseCatch :
// Parses a catch.
//
//	Catch : catch ( CatchParameter ) Block
//	      | catch Block
// https://tc39.es/ecma262/#prod-Catch
func (parser *Parser) parseCatch() *ast.Node {
	return nil
}

// parseFinally :
// Parses a finally.
//
//	Finally : finally Block
// https://tc39.es/ecma262/#prod-Finally
func (parser *Parser) parseFinally() *ast.Node {
	return nil
}

// parseCatchParameter :
// Parses a catch parameter.
//
//	CatchParameter : BindingIdentifier
//	               | BindingPattern
// https://tc39.es/ecma262/#prod-CatchParameter
func (parser *Parser) parseCatchParameter() *ast.Node {
	return nil
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
