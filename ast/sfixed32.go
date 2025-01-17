package ast

var _ Hashable = (*Sfixed32)(nil)
var _ Named = (*Sfixed32)(nil)

// Sfixed32 представление для типа sfixed32
type Sfixed32 struct {
	unique
}

func (*Sfixed32) GetName() string {
	return "sfixed32"
}

func (*Sfixed32) GetFullName() string {
	return "sfixed32"
}

func (*Sfixed32) equivalent(v ScalarNode) bool {
	_, ok := v.(*Sfixed32)
	return ok
}

func (*Sfixed32) genericType() {}
func (*Sfixed32) hashable()    {}
func (*Sfixed32) node()        {}
func (*Sfixed32) scalar()      {}
