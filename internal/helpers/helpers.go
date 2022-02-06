package helpers

import (
	"github.com/emicklei/proto"
	"github.com/ricdeau/protoast/ast"
)

func FromProtoComment(c *proto.Comment) *ast.Comment {
	if c == nil {
		return nil
	}

	return &ast.Comment{
		Value: c.Message(),
		Lines: c.Lines,
	}
}
