package parser

import (
	"duck/ling/js/ast/keyword"
	"duck/ling/js/ast/node"
	"duck/ling/js/ast/token"
	"fmt"
	"github.com/scott-ainsworth/go-ascii"
	"strconv"
)

var Pos = 0
var CurrentToken *token.Token = nil
var CurrentChar byte
var Text string

func Peek() byte {
	peekPos := Pos + 1
	if peekPos > len(Text)-1 {
		return 0
	}
	return Text[peekPos]
}

func SkipWhitespace() {
	for CurrentChar != 0 && ascii.IsSpace(CurrentChar) {
		Advance()
	}
}

func GetInt() int {
	var res string = ""
	for CurrentChar != 0 && ascii.IsDigit(CurrentChar) {
		res += string(CurrentChar)
		Advance()
	}
	resI, _ := strconv.Atoi(res)
	return resI
}

func Id() *token.Token {
	result := ""
	for CurrentChar != 0 && ascii.IsAlnum(CurrentChar) {
		result += string(CurrentChar)
		Advance()
	}

	_, tok := keyword.IsIdentifier(result)
	return tok
}

func GetNextToken() *token.Token {
	for CurrentChar != 0 {
		if ascii.IsSpace(CurrentChar) {
			SkipWhitespace()
			continue
		}

		if ascii.IsLetter(CurrentChar) || CurrentChar == '_' {
			return Id()
		}

		if CurrentChar == '=' {
			Advance()
			return token.NewToken(token.ASSIGN, '=')
		}

		if CurrentChar == ';' {
			Advance()
			return token.NewToken(token.SEMICOLON, ';')
		}

		if ascii.IsDigit(CurrentChar) {
			return token.NewToken(token.INTEGER, GetInt())
		}

		if CurrentChar == '+' {
			Advance()
			return token.NewToken(token.ADD, '+')
		}

		if CurrentChar == '-' {
			Advance()
			return token.NewToken(token.SUB, '-')
		}

		if CurrentChar == '*' {
			Advance()
			return token.NewToken(token.MUL, '*')
		}

		if CurrentChar == '/' {
			Advance()
			return token.NewToken(token.DIV, '/')
		}

		if CurrentChar == '(' {
			Advance()
			return token.NewToken(token.LPAREN, '(')
		}

		if CurrentChar == ')' {
			Advance()
			return token.NewToken(token.RPAREN, ')')
		}
	}
	return token.NewToken(token.EOF, 0)
}

func eat(tokenType token.Type) {
	if tokenType == CurrentToken.GetType() {
		CurrentToken = GetNextToken()
	}
}

func Advance() {
	Pos++
	if Pos > len(Text)-1 {
		CurrentChar = 0
	} else {
		CurrentChar = Text[Pos]
	}
}

func Factor() *node.IASTNode {
	tok := CurrentToken
	var nde node.IASTNode
	if tok.GetType() == token.ADD {
		eat(token.ADD)
		nde = node.NewUnaryOperatorNode(nil, nil, tok, Factor())
		return &nde
	} else if tok.GetType() == token.SUB {
		eat(token.SUB)
		nde = node.NewUnaryOperatorNode(nil, nil, tok, Factor())
		return &nde
	} else if tok.GetType() == token.LPAREN {
		eat(token.LPAREN)
		nde := Expr()
		eat(token.RPAREN)
		return nde
	} else if tok.GetType() == token.INTEGER {
		eat(token.INTEGER)
		nde = node.NewNumberNode(tok)
		return &nde
	}
	nde = node.NewASTNode(nil, nil, tok)
	return &nde
}

func Term() *node.IASTNode {
	var nde *node.IASTNode = Factor()
	for CurrentToken.GetType() == token.MUL || CurrentToken.GetType() == token.DIV {
		tok := CurrentToken
		if CurrentToken.GetType() == token.MUL {
			eat(token.MUL)
		} else if CurrentToken.GetType() == token.DIV {
			eat(token.DIV)
		}
		var tmp node.IASTNode = node.NewBinaryOperatorNode(nde, Term(), tok)
		nde = &tmp
	}
	return nde
}

func Expr() *node.IASTNode {
	var nde *node.IASTNode = Term()
	for CurrentToken.GetType() == token.ADD || CurrentToken.GetType() == token.SUB {
		tok := CurrentToken
		if tok.GetType() == token.ADD {
			eat(token.ADD)
		} else if tok.GetType() == token.SUB {
			eat(token.SUB)
		}
		var tmp node.IASTNode = node.NewBinaryOperatorNode(nde, Term(), tok)
		nde = &tmp
	}

	return nde
}

func Visit(root *node.IASTNode) {
	if root == nil {
		return
	}
	Visit((*root).GetLeft())
	Visit((*root).GetRight())
	fmt.Println((*root).GetToken())
}
