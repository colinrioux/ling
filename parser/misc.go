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
