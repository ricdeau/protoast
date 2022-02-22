package ast

var _ ScalarNode = (*Bytes)(nil)
var _ Named = (*Bytes)(nil)

// Bytes представление типа bytes
type Bytes struct {
	unique
}

func (*Bytes) GetName() string {
	return "bytes"
}

func (*Bytes) GetFullName() string {
	return "bytes"
}

func (*Bytes) equivalent(v ScalarNode) bool {
	_, ok := v.(*Bytes)
	return ok
}

func (*Bytes) genericType() {}
func (*Bytes) node()        {}
func (*Bytes) scalar()      {}
