package parser

import (
	"duck/ling/lexer/token"
	"duck/ling/parser/ast"
)

func (parser *Parser) factor() *ast.Node {
	tok := parser.CurrentToken
	if tok.Type == token.ADD {
		parser.eat(token.ADD)
		return (*ast.Node)(ast.NewUnaryOperatorNode(tok, parser.factor()))
	} else if tok.Type == token.SUB {
		parser.eat(token.SUB)
		return (*ast.Node)(ast.NewUnaryOperatorNode(tok, parser.factor()))
	} else if tok.Type == token.LPAREN {
		parser.eat(token.LPAREN)
		nde := parser.parseExpression()
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
	for parser.CurrentToken.Type == token.MUL || parser.CurrentToken.Type == token.DIV {
		tok := parser.CurrentToken
		if parser.CurrentToken.Type == token.MUL {
			parser.eat(token.MUL)
		} else if parser.CurrentToken.Type == token.DIV {
			parser.eat(token.DIV)
		}
		nde = (*ast.Node)(ast.NewBinaryOperatorNode(nde, parser.term(), tok))
	}
	return nde
}

func (parser *Parser) parseExpression() *ast.Node {
	var nde = parser.term()
	for parser.CurrentToken.Type == token.ADD || parser.CurrentToken.Type == token.SUB {
		tok := parser.CurrentToken
		if tok.Type == token.ADD {
			parser.eat(token.ADD)
		} else if tok.Type == token.SUB {
			parser.eat(token.SUB)
		}
		nde = (*ast.Node)(ast.NewBinaryOperatorNode(nde, parser.term(), tok))
	}

	return nde
}
