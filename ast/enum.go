package ast

import "strings"

var _ Type = &Enum{}
var _ Named = &Enum{}

// Enum представление типа enum
type Enum struct {
	unique

	File      *File
	ParentMsg *Message

	Name    string
	Options []*Option
	Values  []*EnumValue
}

func (e *Enum) GetName() string {
	return e.Name
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

var _ Unique = &EnumValue{}
var _ Named = &EnumValue{}

// EnumValue представление поля для Enum-а
type EnumValue struct {
	unique

	Name    string
	Integer int
	Options []*Option
}

func (e *EnumValue) GetName() string {
	return e.Name
}
