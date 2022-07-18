package parser

import (
	"ling/lexer/keyword"
	"ling/lexer/token"
	"ling/parser/ast"
)

func (parser *Parser) factor() *ast.Node {
	tok := parser.CurrentToken
	if tok.Type == token.ADDITION {
		parser.eat(token.ADDITION)
		return (*ast.Node)(ast.NewUnaryOperatorNode(tok, parser.factor()))
	} else if tok.Type == token.SUBTRACTION {
		parser.eat(token.SUBTRACTION)
		return (*ast.Node)(ast.NewUnaryOperatorNode(tok, parser.factor()))
	} else if tok.Type == token.LPAREN {
		parser.eat(token.LPAREN)
		nde := parser.parseExpression1()
		parser.eat(token.RPAREN)
		return nde
	} else if tok.Type == token.NUMBER {
		parser.eat(token.NUMBER)
		return (*ast.Node)(ast.NewNumberNode(tok))
	}
	return ast.NewNode(nil, nil, tok)
}

func (parser *Parser) term() *ast.Node {
	var nde = parser.factor()
	for parser.CurrentToken.Type == token.MULTIPLICATION || parser.CurrentToken.Type == token.DIVISION {
		tok := parser.CurrentToken
		if parser.CurrentToken.Type == token.MULTIPLICATION {
			parser.eat(token.MULTIPLICATION)
		} else if parser.CurrentToken.Type == token.DIVISION {
			parser.eat(token.DIVISION)
		}
		nde = (*ast.Node)(ast.NewBinaryOperatorNode(nde, parser.term(), tok))
	}
	return nde
}

func (parser *Parser) parseExpression1() *ast.Node {
	var nde = parser.term()
	for parser.CurrentToken.Type == token.ADDITION || parser.CurrentToken.Type == token.SUBTRACTION {
		tok := parser.CurrentToken
		if tok.Type == token.ADDITION {
			parser.eat(token.ADDITION)
		} else if tok.Type == token.SUBTRACTION {
			parser.eat(token.SUBTRACTION)
		}
		nde = (*ast.Node)(ast.NewBinaryOperatorNode(nde, parser.term(), tok))
	}

	return nde
}

// parsePrimaryExpression :
// Parses a primary expression.
//
//	PrimaryExpression : this
//	                  | IdentifierReference
//	                  | Literal
//	                  | ArrayLiteral
//	                  | ObjectLiteral
//	                  | FunctionExpression
//	                  | ClassExpression
//	                  | GeneratorExpression
//	                  | AsyncFunctionExpression
//	                  | AsyncGeneratorExpression
//	                  | RegularExpressionLiteral
//	                  | TemplateLiteral
//	                  | CoverParenthesizedExpressionAndArrowParameterList
// https://tc39.es/ecma262/#prod-PrimaryExpression
func (parser *Parser) parsePrimaryExpression() *ast.Node {
	if parser.CurrentToken.Type == token.KEYWORD {
		// Handle this.
		if parser.CurrentToken.Value == keyword.THIS {

		}
		// Handle partial IdentifierReference.
		if parser.CurrentToken.Value == keyword.AWAIT || parser.CurrentToken.Value == keyword.YIELD {

		}
	} else if parser.CurrentToken.Type == token.IDENTIFIER {
		// Handle rest of IdentifierReference.
	}
	return nil
}

// parseUpdateExpression :
// Parses an update expression.
//
// 	UpdateExpression : LeftHandSideExpression
//	                 | LeftHandSideExpression Whitespace(no line terminator) ++
//	                 | LeftHandSideExpression Whitespace(no line terminator) --
//	                 | ++ UnaryExpression
//	                 | -- UnaryExpression
// https://tc39.es/ecma262/#prod-UpdateExpression
func (parser *Parser) parseUpdateExpression() *ast.Node {
	return nil
}

// parseLeftHandSideExpression :
// Parses a left hand side expression.
//
// 	LeftHandSideExpression : NewExpression
//	                       | CallExpression
//	                       | OptionalExpression
// https://tc39.es/ecma262/#prod-LeftHandSideExpression
func (parser *Parser) parseLeftHandSideExpression() *ast.Node {
	return nil
}

// parseNewExpression :
// Parses a new expression.
//
// 	NewExpression : MemberExpression
//	              | new NewExpression
// https://tc39.es/ecma262/#prod-NewExpression
func (parser *Parser) parseNewExpression() *ast.Node {
	if parser.CurrentToken.Type == token.KEYWORD && parser.CurrentToken.Value == keyword.NEW {
		parser.eat(token.KEYWORD)
		return parser.parseNewExpression()
	}
	return parser.parseMemberExpression()
}

// parseMemberExpression :
// Parses a member expression.
//
//	MemberExpression : PrimaryExpression
//                 	 | MemberExpression [ Expression ]
//                	 | MemberExpression . IdentifierName
//               	 | MemberExpression TemplateLiteral
//              	 | SuperProperty
//               	 | MetaProperty
//              	 | new MemberExpression Arguments
//              	 | MemberExpression . PrivateIdentifier
// https://tc39.es/ecma262/#prod-MemberExpression
func (parser *Parser) parseMemberExpression() *ast.Node {
	return nil
}

// parseExpression :
// Parses an expression.
//
// 	Expression : AssignmentExpression
//	           | Expression , AssignmentExpression
// https://tc39.es/ecma262/#prod-Expression
func (parser *Parser) parseExpression() *ast.Node {
	return nil
}

// parseAssignmentExpression :
// Parses an assignment expression.
//
//	AssignmentExpression : ConditionalExpression
//	                     | YieldExpression
//	                     | ArrowFunction
//	                     | AsyncArrowFunction
//	                     | LeftHandSideExpression = AssignmentExpression
//	                     | LeftHandSideExpression AssignmentOperator AssignmentExpression
//	                     | LeftHandSideExpression &&= AssignmentExpression
//	                     | LeftHandSideExpression ||= AssignmentExpression
//	                     | LeftHandSideExpression ??= AssignmentExpression
// https://tc39.es/ecma262/#prod-AssignmentExpression
func (parser *Parser) parseAssignmentExpression() *ast.Node {
	return nil
}
