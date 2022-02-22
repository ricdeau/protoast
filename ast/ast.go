package ast

// Node представление базовой ноды
type Node interface {
	Unique
	node()
}

// Type представление типа
type Type interface {
	Node
	genericType()
}

// ScalarNode скалярные типы
type ScalarNode interface {
	Type
	scalar()
	equivalent(v ScalarNode) bool
}

// Hashable типы могущие являться ключами словарей
type Hashable interface {
	ScalarNode
	hashable()
}

// Named типы имеющие имя.
type Named interface {
	GetName() string
	GetFullName() string
}

// Valuable типы имеющие значение.
type Valuable interface {
	GetValue() interface{}
}

// Compound типы, входящие в состав других типов.
type Compound interface {
	Named
	GetParentMsg() *Message
}

// OptionsBearer типы, которые могут содержать proto-опции.
type OptionsBearer interface {
	GetOptions() []*Option
}

type Commented interface {
	GetComment() *Comment
}
