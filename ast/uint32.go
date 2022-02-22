package ast

var _ Hashable = (*Uint32)(nil)
var _ Named = (*Uint32)(nil)

// Uint32 представление для типа uint32
type Uint32 struct {
	unique
}

func (*Uint32) GetName() string {
	return "uint32"
}

func (*Uint32) GetFullName() string {
	return "uint32"
}

func (*Uint32) equivalent(v ScalarNode) bool {
	_, ok := v.(*Uint32)
	return ok
}

func (*Uint32) genericType() {}
func (*Uint32) hashable()    {}
func (*Uint32) node()        {}
func (*Uint32) scalar()      {}
