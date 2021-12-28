package namespace

import (
	"strings"
	"text/scanner"

	"github.com/ricdeau/protoast/internal/errors"

	"github.com/ricdeau/protoast/ast"
)

type nodeTuple struct {
	item ast.Node
	pos  scanner.Position
}

func newPlain(name string, builder *Builder) Namespace {
	return &plain{
		name:    name,
		builder: builder,
		ns:      map[string]nodeTuple{},
	}
}

type plain struct {
	name string
	pkg  string

	builder *Builder
	ns      map[string]nodeTuple

	final bool
}

func (n *plain) GetService(name string) *ast.Service {
	res, _ := n.getNode(name)
	if res == nil {
		return nil
	}
	srv, _ := res.(*ast.Service)
	return srv
}

func (n *plain) getNode(name string) (ast.Node, scanner.Position) {
	var pos scanner.Position
	res, ok := n.ns[name]
	if !ok {
		items := strings.Split(name, ".")
		if len(items) != 2 {
			return nil, pos
		}
		ns := n.WithScope(items[0])
		if ns == nil {
			return nil, pos
		}
		return ns.getNode(items[1])
	}

	return res.item, res.pos
}

func (n *plain) WithImport(pkgNamespace Namespace) (Namespace, error) {
	return newImport(n, pkgNamespace, n.builder), nil
}

func (n *plain) WithScope(name string) Namespace {
	return newScope(name, n, n.builder)
}

func (n *plain) GetType(name string) ast.Type {
	res, ok := n.ns[name]
	if !ok {
		return nil
	}

	v, ok := res.item.(ast.Type)
	if !ok {
		return nil
	}

	return v
}

func (n *plain) SetNode(name string, def ast.Node, defPos scanner.Position) error {
	prev, ok := n.ns[name]
	if ok {
		return errors.Newf("%s duplicate type %s declartion, the previous one was %s", defPos, name, prev.pos)
	}

	n.ns[name] = nodeTuple{
		item: def,
		pos:  defPos,
	}
	return nil
}

func (n *plain) Finalized() bool { return n.final }
func (n *plain) Finalize()       { n.final = true }
func (n *plain) String() string  { return n.name }

func (n *plain) SetPkgName(pkg string) error {
	n.pkg = pkg
	return nil
}

func (n *plain) PkgName() string { return n.pkg }
