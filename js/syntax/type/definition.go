package _type

type IDefinition interface{}
type Definition struct{}

type UndefinedDefinition struct {
	Definition
}

func NewUndefinedDefinition() *UndefinedDefinition {
	return &UndefinedDefinition{}
}

type NullDefinition struct {
	Definition
}

func NewNullDefinition() *NullDefinition {
	return &NullDefinition{}
}

type BooleanDefinition struct {
	Definition
	value bool
}

func NewBooleanDefinition(value bool) *BooleanDefinition {
	return &BooleanDefinition{value: value}
}

type StringDefinition struct {
	Definition
	value string
}

func NewStringDefinition(value string) *StringDefinition {
	return &StringDefinition{value: value}
}
