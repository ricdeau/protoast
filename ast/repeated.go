package ast

var _ Type = (*Repeated)(nil)

// Repeated представление для полей с repeeated
type Repeated struct {
	unique

	Type Type
}

func (*Repeated) genericType() {}
func (*Repeated) node()        {}
