package ast

import "strconv"

var _ Unique = (*Option)(nil)
var _ Named = (*Option)(nil)

// Option опция
type Option struct {
	unique

	Name      string
	Value     OptionValue
	Extension *Extension
}

func (o *Option) GetName() string {
	return o.Name
}

// OptionValue значение опции
type OptionValue interface {
	Unique

	isOptionValue()
}

var _ OptionValue = (*EmbeddedOption)(nil)
var _ Valuable = (*EmbeddedOption)(nil)

// EmbeddedOption представление встроенной опции
type EmbeddedOption struct {
	unique
	Value string
}

func (o *EmbeddedOption) GetValue() interface{} {
	return o.Value
}

func (*EmbeddedOption) isOptionValue() {}

var _ OptionValue = (*EnumOption)(nil)
var _ Valuable = (*EnumOption)(nil)

// EnumOption представление опций типа Enum
type EnumOption struct {
	unique
	Value *EnumValue
}

func (o *EnumOption) GetValue() interface{} {
	return o.Value.Name
}

func (*EnumOption) isOptionValue() {}

var _ OptionValue = (*IntOption)(nil)
var _ Valuable = (*IntOption)(nil)

// IntOption branch of OptionValue
type IntOption struct {
	unique
	Value int64
}

func (o *IntOption) GetValue() interface{} {
	return int(o.Value)
}

func (*IntOption) isOptionValue() {}

func (o *IntOption) String() string {
	return strconv.FormatInt(o.Value, 10)
}

var _ OptionValue = (*UintOption)(nil)
var _ Valuable = (*UintOption)(nil)

// UintOption branch of OptionValue
type UintOption struct {
	unique
	Value uint64
}

func (o *UintOption) GetValue() interface{} {
	return int(o.Value)
}

func (*UintOption) isOptionValue() {}

func (o *UintOption) String() string {
	return strconv.FormatUint(o.Value, 10)
}

var _ OptionValue = (*FloatOption)(nil)
var _ Valuable = (*FloatOption)(nil)

// FloatOption branch of OptionValue
type FloatOption struct {
	unique
	Value float64
}

func (o *FloatOption) GetValue() interface{} {
	return o.Value
}

func (*FloatOption) isOptionValue() {}

func (o *FloatOption) String() string {
	return strconv.FormatFloat(o.Value, ' ', 5, 64)
}

var _ OptionValue = (*StringOption)(nil)
var _ Valuable = (*StringOption)(nil)

// StringOption branch of OptionValue
type StringOption struct {
	unique
	Value string
}

func (o *StringOption) GetValue() interface{} {
	return o.Value
}

func (*StringOption) isOptionValue() {}

func (o *StringOption) String() string {
	return o.Value
}

var _ OptionValue = (*BoolOption)(nil)
var _ Valuable = (*BoolOption)(nil)

// BoolOption branch of OptionValue
type BoolOption struct {
	unique
	Value bool
}

func (o *BoolOption) GetValue() interface{} {
	return o.Value
}

func (o *BoolOption) String() string {
	return strconv.FormatBool(o.Value)
}

func (*BoolOption) isOptionValue() {}

// ArrayOption branch of OptionValue
type ArrayOption struct {
	unique
	Value []OptionValue
}

func (*ArrayOption) isOptionValue() {}

// MapOption branch of OptionValue
type MapOption struct {
	unique
	Value map[string]OptionValue
}

func (*MapOption) isOptionValue() {}

type optChecker struct {
	bearer  OptionsBearer
	optName string
}

func CheckOption(bearer OptionsBearer, optName string) *optChecker {
	return &optChecker{
		bearer:  bearer,
		optName: optName,
	}
}

func (c *optChecker) Exists() bool {
	for _, opt := range c.bearer.GetOptions() {
		if opt.Name == c.optName {
			return true
		}
	}

	return false
}

func (c *optChecker) CheckValue(cond func(v OptionValue) bool) bool {
	for _, opt := range c.bearer.GetOptions() {
		if opt.Name == c.optName {
			return cond(opt.Value)
		}
	}

	return false
}

func (c *optChecker) Equal(v interface{}) bool {
	for _, opt := range c.bearer.GetOptions() {
		if opt.Name == c.optName {
			if valuable, ok := opt.Value.(Valuable); ok {
				return valuable.GetValue() == v
			}
		}
	}

	return false
}
