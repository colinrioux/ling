package parser

import (
	"duck/ling/lexer/token"
	"duck/ling/parser/ast"
)

func (parser *Parser) factor() *ast.Node {
	tok := parser.Lexer.CurrentToken
	if tok.Type == token.ADD {
		parser.Lexer.Eat(token.ADD)
		return (*ast.Node)(ast.NewUnaryOperatorNode(tok, parser.factor()))
	} else if tok.Type == token.SUB {
		parser.Lexer.Eat(token.SUB)
		return (*ast.Node)(ast.NewUnaryOperatorNode(tok, parser.factor()))
	} else if tok.Type == token.LPAREN {
		parser.Lexer.Eat(token.LPAREN)
		nde := parser.parseExpression()
		parser.Lexer.Eat(token.RPAREN)
		return nde
	} else if tok.Type == token.NUMBER {
		parser.Lexer.Eat(token.NUMBER)
		return (*ast.Node)(ast.NewNumberNode(tok))
	}
	return ast.NewNode(nil, nil, tok)
}

func (parser *Parser) term() *ast.Node {
	var nde = parser.factor()
	for parser.Lexer.CurrentToken.Type == token.MUL || parser.Lexer.CurrentToken.Type == token.DIV {
		tok := parser.Lexer.CurrentToken
		if parser.Lexer.CurrentToken.Type == token.MUL {
			parser.Lexer.Eat(token.MUL)
		} else if parser.Lexer.CurrentToken.Type == token.DIV {
			parser.Lexer.Eat(token.DIV)
		}
		nde = (*ast.Node)(ast.NewBinaryOperatorNode(nde, parser.term(), tok))
	}
	return nde
}

func (parser *Parser) parseExpression() *ast.Node {
	var nde = parser.term()
	for parser.Lexer.CurrentToken.Type == token.ADD || parser.Lexer.CurrentToken.Type == token.SUB {
		tok := parser.Lexer.CurrentToken
		if tok.Type == token.ADD {
			parser.Lexer.Eat(token.ADD)
		} else if tok.Type == token.SUB {
			parser.Lexer.Eat(token.SUB)
		}
		nde = (*ast.Node)(ast.NewBinaryOperatorNode(nde, parser.term(), tok))
	}

	return nde
}
