package ast

import "strings"

var _ Type = (*Enum)(nil)
var _ Named = (*Enum)(nil)
var _ Commented = (*Enum)(nil)

// Enum представление типа enum
type Enum struct {
	unique

	File      *File
	ParentMsg *Message

	Name    string
	Comment *Comment
	Options []*Option
	Values  []*EnumValue
}

func (e *Enum) GetName() string {
	return e.Name
}

func (e *Enum) GetFullName() string {
	return e.String()
}

func (e *Enum) GetComment() *Comment {
	return e.Comment
}

func (*Enum) genericType() {}
func (*Enum) node()        {}

// Enum референс-имя перечисления, включает в себя название пакета,
// имена родительских сообщений, в пространстве имён которых оно определено.
func (e *Enum) String() string {
	var buf strings.Builder
	if e.ParentMsg == nil {
		buf.WriteString(e.File.Package)
	} else {
		buf.WriteString(e.ParentMsg.String())
	}
	buf.WriteByte('.')
	buf.WriteString(e.Name)

	return buf.String()
}

var _ Unique = (*EnumValue)(nil)
var _ Named = (*EnumValue)(nil)
var _ Commented = (*EnumValue)(nil)

// EnumValue представление поля для Enum-а
type EnumValue struct {
	unique

	Name    string
	Comment *Comment
	Integer int
	Options []*Option
}

func (e *EnumValue) GetName() string {
	return e.Name
}

func (e *EnumValue) GetFullName() string {
	return e.Name
}

func (e *EnumValue) GetComment() *Comment {
	return e.Comment
}
