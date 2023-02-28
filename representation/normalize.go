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

	f := ast.MergePackageFiles(pkg, ast.FilterImportDuplicates+ast.FilterUnassociatedComments)
	pkg.Files = map[string]*ast.File{defaultFileName: f}

	// register additional normalizations here
	s.normalizeDeclarations(pkg)

	s.sortInFile(f)
	s.normalizeNaming(pkg)

	s.represent = pkg
	return nil
}

func (s *Representation) normalizeNaming(pkg *ast.Package) {
	astutil.Apply(pkg, func(cursor *astutil.Cursor) bool {
		node := cursor.Node()
		if node == nil {
			return true
		}

		s.collectImport(node)
		s.rename(cursor)
		return true
	}, nil)
}

func (s *Representation) rename(cursor *astutil.Cursor) {
	node := cursor.Node()

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
