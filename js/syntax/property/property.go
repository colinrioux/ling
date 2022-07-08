package property

import (
	"fmt"
)

type SimpleAttribute struct {
	Writable     bool
	Enumerable   bool
	Configurable bool
}

func (a *SimpleAttribute) String() string {
	return fmt.Sprintf("{Writable:%v,Enumerable:%v,Configurable:%v}", a.Writable, a.Enumerable, a.Configurable)
}
