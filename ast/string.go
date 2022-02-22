package ast

var _ Hashable = (*String)(nil)
var _ Named = (*String)(nil)

// String представление для стрового типа
type String struct {
	unique
}

func (*String) GetName() string {
	return "string"
}

func (*String) GetFullName() string {
	return "string"
}

func (*String) equivalent(v ScalarNode) bool {
	_, ok := v.(*String)
	return ok
}

func (*String) genericType() {}
func (*String) hashable()    {}
func (*String) node()        {}
func (*String) scalar()      {}
