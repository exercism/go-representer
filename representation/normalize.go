package representation

import (
	"go/ast"
	"go/types"

	"github.com/pkg/errors"
	"golang.org/x/tools/go/ast/astutil"
)

// Normalize processes the solutions AST to normalize the representation.
func (s *Representation) Normalize() error {
	pkg := s.getPackage()
	if pkg == nil {
		return errors.New("no Go package found")
	}

	pkg = s.normalize(pkg)
	s.represent = pkg
	return nil
}

func (s *Representation) normalize(pkg *ast.Package) *ast.Package {
	f := ast.MergePackageFiles(pkg, ast.FilterImportDuplicates+ast.FilterUnassociatedComments)
	pkg.Files = map[string]*ast.File{defaultFileName: f}

	s.sortInFile(f)

	astutil.Apply(pkg, func(cursor *astutil.Cursor) bool {
		node := cursor.Node()
		if node == nil {
			return true
		}

		s.collectImport(node)
		s.rename(node, cursor)
		return true
	}, nil)
	return pkg
}

func (s *Representation) rename(node ast.Node, cursor *astutil.Cursor) {
	switch n := node.(type) {
	default:
		return
	case *ast.SelectorExpr:
	case *ast.Ident:
		switch parent := cursor.Parent().(type) {
		case *ast.File:
			return
		case *ast.Field:
			// skip if current node is the type of the field
			if parent.Type == n {
				return
			}
		case *ast.SelectorExpr:
			// call to a different package is usually a call to stdlib
			if x, ok := parent.X.(*ast.Ident); ok {
				// check if this is a type or func from an import and don't rename.
				if s.isImport(x.Name) {
					return
				}

				// check for methods on stdlib types and don't rename.
				identType := s.getIdentType(x)
				switch t := identType.(type) {
				case *types.Named:
					// type is from a different package. Don't rename methods (Sel).
					if t.Obj().Pkg() != nil && node == parent.Sel {
						return
					}
				}
			}
		}
		n.Name = s.getPlaceHolder(n.Name)
	}
}
