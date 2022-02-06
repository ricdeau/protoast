package ast

var _ Type = &OneOf{}
var _ Compound = &OneOf{}

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

var _ Unique = (*OneOfBranch)(nil)
var _ Field = (*OneOfBranch)(nil)
var _ Commented = (*OneOfBranch)(nil)

// OneOfBranch представление для ветви
type OneOfBranch struct {
	unique

	Name     string
	Comment  *Comment
	Type     Type
	ParentOO *OneOf
	Sequence int
	Options  []*Option
}

func (o *OneOfBranch) GetOptions() []*Option {
	return o.Options
}

func (o *OneOfBranch) GetName() string {
	return o.Name
}

func (o *OneOfBranch) GetComment() *Comment {
	return o.Comment
}

func (o *OneOfBranch) isField() (string, Type, []*Option, int) {
	return o.Name, o.Type, o.Options, o.Sequence
}
