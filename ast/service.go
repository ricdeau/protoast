package ast

var _ Node = (*Service)(nil)
var _ Named = (*Service)(nil)
var _ Commented = (*Service)(nil)
var _ OptionsBearer = (*Service)(nil)

// Service представление для сервисов
type Service struct {
	unique

	File *File

	Name    string
	Comment *Comment
	Methods []*Method
	Options []*Option
}

func (s *Service) GetComment() *Comment {
	return s.Comment
}

func (s *Service) GetOptions() []*Option {
	return s.Options
}

// Method поиск метода по имени
func (s *Service) Method(name string) *Method {
	for _, method := range s.Methods {
		if method.Name == name {
			return method
		}
	}

	return nil
}

func (s *Service) GetName() string {
	return s.Name
}

func (s *Service) node() {}
