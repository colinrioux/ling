package _type

import "fmt"

type IPair interface {
	Equals(other *Pair) bool
}

type Pair struct {
	a interface{}
	b interface{}
}

func NewPair(a interface{}, b interface{}) *Pair {
	return &Pair{a: a, b: b}
}

func (p *Pair) Equals(other *Pair) bool {
	return p.a == other.a && p.b == other.b
}

func (p *Pair) String() string {
	return fmt.Sprintf("(%v,%v)", p.a, p.b)
}
