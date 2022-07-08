package _type

import "fmt"

type IECMAPair interface {
	Equals(other *ECMAPair) bool
}

type ECMAPair struct {
	a interface{}
	b interface{}
}

func NewPair(a interface{}, b interface{}) *ECMAPair {
	return &ECMAPair{a: a, b: b}
}

func (p *ECMAPair) Equals(other *ECMAPair) bool {
	return p.a == other.a && p.b == other.b
}

func (p *ECMAPair) String() string {
	return fmt.Sprintf("(%v,%v)", p.a, p.b)
}
