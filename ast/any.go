package ast

var _ Type = (*Any)(nil)
var _ Named = (*Any)(nil)

// Any представление типа golang.protobuf.Any
type Any struct {
	unique

	File *File
}

func (*Any) GetName() string {
	return "Any"
}

func (*Any) GetFullName() string {
	return "golang.protobuf.Any"
}

func (*Any) genericType() {}
func (*Any) node()        {}
