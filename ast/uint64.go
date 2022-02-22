package ast

var _ Hashable = (*Uint64)(nil)
var _ Named = (*Uint64)(nil)

// Uint64 представление для типа uint64
type Uint64 struct {
	unique
}

func (*Uint64) GetName() string {
	return "uint64"
}

func (*Uint64) GetFullName() string {
	return "uint64"
}

func (*Uint64) equivalent(v ScalarNode) bool {
	_, ok := v.(*Uint64)
	return ok
}

func (*Uint64) genericType() {}
func (*Uint64) hashable()    {}
func (*Uint64) node()        {}
func (*Uint64) scalar()      {}
