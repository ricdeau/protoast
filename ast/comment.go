package ast

var _ Unique = &Comment{}

// Comment представление комментария
type Comment struct {
	unique

	Value string
	Lines []string
}

func (c *Comment) GetValue() string {
	if c == nil {
		return ""
	}

	return c.Value
}

func (c *Comment) GetLines() []string {
	if c == nil {
		return nil
	}

	return c.Lines
}
