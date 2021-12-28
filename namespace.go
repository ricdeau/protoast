package protoast

import (
	"github.com/ricdeau/protoast/ast"
)

type Namespace interface {
	GetType(name string) ast.Type
	GetService(name string) *ast.Service
}
