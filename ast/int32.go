package ast

var _ Hashable = (*Int32)(nil)
var _ Named = (*Int32)(nil)

// Int32 представление типа int32
type Int32 struct {
	unique
}

func (*Int32) GetName() string {
	return "int32"
}

func (*Int32) GetFullName() string {
	return "int32"
}

func (*Int32) equivalent(v ScalarNode) bool {
	_, ok := v.(*Int32)
	return ok
}

func (*Int32) genericType() {}
func (*Int32) hashable()    {}
func (*Int32) node()        {}
func (*Int32) scalar()      {}
