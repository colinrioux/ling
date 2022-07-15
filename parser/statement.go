package parser

import "duck/ling/lexer/token"

// parseVariableStatement :
// "var" variables are scoped to the running execution context.
// If declared, they are instantiated as undefined.
// https://tc39.es/ecma262/#sec-variable-statement
func (parser *Parser) parseVariableStatement() *token.Token {
	// eat the var token
	parser.eat(token.KEYWORD)

	return nil
}
