package main

import (
	"duck/ling/js/parser"
	"fmt"
)

func main() {
	//t := [...]int{1, 2, 3}
	//s1 := make(_type.ECMASet, len(t))
	//for i, v := range t {
	//	s1[i] = v
	//}
	//t = [...]int{4, 5, 6}
	//s2 := make(_type.ECMASet, len(t))
	//for i, v := range t {
	//	s2[i] = v
	//}
	////fmt.Println(s1)
	////fmt.Println(s2)
	//rel := _type.NewRelation(&s1, &s1)
	//if rel.I() != nil {
	//	fmt.Println(rel.I().GetGraph())
	//}
	parser.Text = "undefined"
	fmt.Println(parser.Text)
	parser.CurrentChar = parser.Text[parser.Pos]
	parser.CurrentToken = parser.GetNextToken()
	fmt.Println(parser.CurrentToken)
	//var nde node.IASTNode = *(parser.Expr())
	//fmt.Println(nde)
	//fmt.Println(nde.Visit())
}
