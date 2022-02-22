package ast

import (
	"strings"
)

var _ Type = (*OneOf)(nil)
var _ Compound = (*OneOf)(nil)

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

func (o *OneOf) GetFullName() string {
	return o.String()
}

func (o *OneOf) GetParentMsg() *Message {
	return o.ParentMsg
}

// String референс-имя сообщения, включает в себя название пакета,
// имена родительских сообщений, в пространстве имён которых оно определено.
func (m *OneOf) String() string {
	var buf strings.Builder
	buf.WriteString(m.ParentMsg.String())
	buf.WriteByte('.')
	buf.WriteString(m.Name)

	return buf.String()
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

func (o *OneOfBranch) GetFullName() string {
	return o.Name
}

func (o *OneOfBranch) GetComment() *Comment {
	return o.Comment
}

func (o *OneOfBranch) isField() (string, Type, []*Option, int) {
	return o.Name, o.Type, o.Options, o.Sequence
}
