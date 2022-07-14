package parser

import (
	"duck/ling/lexer/token"
	node2 "duck/ling/parser/ast/node"
)

func Factor() *node2.ASTNode {
	tok := CurrentToken
	if tok.Type == token.ADD {
		eat(token.ADD)
		return (*node2.ASTNode)(node2.NewUnaryOperatorNode(tok, Factor()))
	} else if tok.Type == token.SUB {
		eat(token.SUB)
		return (*node2.ASTNode)(node2.NewUnaryOperatorNode(tok, Factor()))
	} else if tok.Type == token.LPAREN {
		eat(token.LPAREN)
		nde := Expr()
		eat(token.RPAREN)
		return nde
	} else if tok.Type == token.NUMBER {
		eat(token.NUMBER)
		return (*node2.ASTNode)(node2.NewNumberNode(tok))
	}
	return node2.NewASTNode(nil, nil, tok)
}

func Term() *node2.ASTNode {
	var nde = Factor()
	for CurrentToken.Type == token.MUL || CurrentToken.Type == token.DIV {
		tok := CurrentToken
		if CurrentToken.Type == token.MUL {
			eat(token.MUL)
		} else if CurrentToken.Type == token.DIV {
			eat(token.DIV)
		}
		nde = (*node2.ASTNode)(node2.NewBinaryOperatorNode(nde, Term(), tok))
	}
	return nde
}

func Expr() *node2.ASTNode {
	var nde = Term()
	for CurrentToken.Type == token.ADD || CurrentToken.Type == token.SUB {
		tok := CurrentToken
		if tok.Type == token.ADD {
			eat(token.ADD)
		} else if tok.Type == token.SUB {
			eat(token.SUB)
		}
		nde = (*node2.ASTNode)(node2.NewBinaryOperatorNode(nde, Term(), tok))
	}

	return nde
}
