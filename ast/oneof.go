package ast

var _ Type = &OneOf{}
var _ Compaund = &OneOf{}

// OneOf представление для oneof поля message-а
type OneOf struct {
	unique

	ParentMsg *Message

	Name     string
	Branches []*OneOfBranch
}

func (o *OneOf) GetName() string {
	return o.Name
}

func (o *OneOf) GetParentMsg() *Message {
	return o.ParentMsg
}

func (*OneOf) genericType() {}
func (*OneOf) node()        {}

var _ Unique = &OneOfBranch{}
var _ Field = &OneOfBranch{}
var _ Named = &OneOfBranch{}

// OneOfBranch представление для ветви
type OneOfBranch struct {
	unique

	Name     string
	Type     Type
	ParentOO *OneOf
	Sequence int
	Options  []*Option
}

func (o *OneOfBranch) GetName() string {
	return o.Name
}

func (o *OneOfBranch) isField() (string, Type, []*Option, int) {
	return o.Name, o.Type, o.Options, o.Sequence
}
