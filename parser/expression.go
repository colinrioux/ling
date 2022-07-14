package parser

import (
	"duck/ling/lexer/token"
	"duck/ling/parser/ast"
)

func Factor() *ast.ASTNode {
	tok := CurrentToken
	if tok.Type == token.ADD {
		eat(token.ADD)
		return (*ast.ASTNode)(ast.NewUnaryOperatorNode(tok, Factor()))
	} else if tok.Type == token.SUB {
		eat(token.SUB)
		return (*ast.ASTNode)(ast.NewUnaryOperatorNode(tok, Factor()))
	} else if tok.Type == token.LPAREN {
		eat(token.LPAREN)
		nde := Expr()
		eat(token.RPAREN)
		return nde
	} else if tok.Type == token.NUMBER {
		eat(token.NUMBER)
		return (*ast.ASTNode)(ast.NewNumberNode(tok))
	}
	return ast.NewASTNode(nil, nil, tok)
}

func Term() *ast.ASTNode {
	var nde = Factor()
	for CurrentToken.Type == token.MUL || CurrentToken.Type == token.DIV {
		tok := CurrentToken
		if CurrentToken.Type == token.MUL {
			eat(token.MUL)
		} else if CurrentToken.Type == token.DIV {
			eat(token.DIV)
		}
		nde = (*ast.ASTNode)(ast.NewBinaryOperatorNode(nde, Term(), tok))
	}
	return nde
}

func Expr() *ast.ASTNode {
	var nde = Term()
	for CurrentToken.Type == token.ADD || CurrentToken.Type == token.SUB {
		tok := CurrentToken
		if tok.Type == token.ADD {
			eat(token.ADD)
		} else if tok.Type == token.SUB {
			eat(token.SUB)
		}
		nde = (*ast.ASTNode)(ast.NewBinaryOperatorNode(nde, Term(), tok))
	}

	return nde
}
