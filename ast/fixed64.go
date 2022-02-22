package ast

var _ Hashable = (*Fixed64)(nil)
var _ Named = (*Fixed64)(nil)

// Fixed64 представление типа fixed64
type Fixed64 struct {
	unique
}

func (*Fixed64) GetName() string {
	return "fixed64"
}

func (*Fixed64) GetFullName() string {
	return "fixed64"
}

func (*Fixed64) equivalent(v ScalarNode) bool {
	_, ok := v.(*Fixed64)
	return ok
}

func (*Fixed64) genericType() {}
func (*Fixed64) hashable()    {}
func (*Fixed64) node()        {}
func (*Fixed64) scalar()      {}
