package ast

var _ Node = (*Method)(nil)
var _ Named = (*Method)(nil)
var _ Commented = (*Method)(nil)
var _ OptionsBearer = (*Method)(nil)

// MethodType тип метода: унарный, входящий/исходящий/двунаправленный потоковый.
type MethodType byte

// Типы методов.
const (
	_ MethodType = iota
	Unary
	InputStream
	OutputStream
	BidirectionalStream
)

// Method представление для метода
type Method struct {
	unique

	File    *File
	Service *Service

	Name    string
	Comment *Comment
	Input   Type
	Output  Type

	Options []*Option
}

func (m *Method) GetOptions() []*Option {
	return m.Options
}

func (m *Method) GetName() string {
	return m.Name
}

func (m *Method) GetFullName() string {
	return m.File.Package + "." + m.Service.Name + "/" + m.Name
}

func (m *Method) GetComment() *Comment {
	return m.Comment
}

func (m *Method) node() {}

// InputMessage возвращает структуру запроса (минуя оборачивающий Stream, если нужно)
func (m *Method) InputMessage() *Message {
	return getMessage(m.Input)
}

// OutputMessage аналогично InputMessage, возвращает структуру ответа, при необходимости снимая stream
func (m *Method) OutputMessage() *Message {
	return getMessage(m.Output)
}

// GetMethodType возвращает тип метода.
func (m *Method) GetMethodType() MethodType {
	_, in := m.Input.(*Stream)
	_, out := m.Output.(*Stream)
	switch {
	case in && out:
		return BidirectionalStream
	case in:
		return InputStream
	case out:
		return OutputStream
	default:
		return Unary
	}
}

func getMessage(m Type) *Message {
	v, ok := m.(*Message)
	if ok {
		return v
	}

	return m.(*Stream).Type.(*Message)
}
