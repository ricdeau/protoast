package ast

var _ ScalarNode = (*Bool)(nil)
var _ Named = (*Bool)(nil)

// Bool представление булевского типа
type Bool struct {
	unique
}

func (*Bool) GetName() string {
	return "bool"
}

func (*Bool) GetFullName() string {
	return "bool"
}

func (*Bool) equivalent(v ScalarNode) bool {
	_, ok := v.(*Bool)
	return ok
}

func (*Bool) genericType() {}
func (*Bool) node()        {}
func (*Bool) scalar()      {}
