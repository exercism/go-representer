package representation

import (
	"go/ast"
	"go/importer"
	"go/parser"
	"go/token"
	"go/types"
	"net/http"
	"os"
	"path"
	"strings"

	"github.com/pkg/errors"
	"github.com/tehsphinx/astrav"
)

const defaultFileName = "solution.go"

// NewRepresentation creates a new folder with given path. Use ParseAST to parse ast from go files in path.
// The pkgPath is the import path of the package to be used by types.parseInfo.
func NewRepresentation(dir string) *Representation {
	root := http.Dir(".")
	if path.IsAbs(dir) {
		root = "/"
	}

	return &Representation{
		dir:  dir,
		root: root,
		fSet: token.NewFileSet(),
		pkgs: map[string]*ast.Package{},

		mapping: map[string]string{},
	}
}

// Representation represents a go package folder
type Representation struct {
	dir  string
	root http.FileSystem

	info     *types.Info
	fSet     *token.FileSet
	pkgs     map[string]*ast.Package
	typesPkg *types.Package

	plhInc      int
	mapping     map[string]string // key: variable name, value: placeholder name
	represent   *ast.Package
	importNames []string
}

// ParseAST will parse all to files in folder. It skips test files.
func (s *Representation) ParseAST() (map[string]*ast.Package, error) {
	pkgs, _, err := astrav.Parse(s.fSet, s.root, s.dir, func(info os.FileInfo) bool {
		return !strings.HasSuffix(info.Name(), "_test.go") && info.Name() != "embed.go"
	}, parser.AllErrors)
	if err != nil {
		return nil, err
	}

	for name, pkg := range pkgs {
		s.pkgs[name] = pkg
	}

	if s.typesPkg, err = s.parseInfo(s.dir, s.fSet, s.getFiles()); err != nil {
		return nil, err
	}

	return s.pkgs, nil
}

// parseInfo parses all files for type information which is then available from the Nodes.
func (s *Representation) parseInfo(path string, fSet *token.FileSet, files []*ast.File) (*types.Package, error) {
	s.info = &types.Info{
		Types:      map[ast.Expr]types.TypeAndValue{},
		Defs:       map[*ast.Ident]types.Object{},
		Uses:       map[*ast.Ident]types.Object{},
		Scopes:     map[ast.Node]*types.Scope{},
		Implicits:  map[ast.Node]types.Object{},
		Selections: map[*ast.SelectorExpr]*types.Selection{},
	}
	var conf = types.Config{
		Importer: importer.Default(),
	}

	pkg, err := conf.Check(path, fSet, files, s.info)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return pkg, nil
}

func (s *Representation) getPackage() *ast.Package {
	for _, pkg := range s.pkgs {
		return pkg
	}
	return nil
}

func (s *Representation) getFiles() []*ast.File {
	var files []*ast.File
	for _, pkg := range s.pkgs {
		for _, file := range pkg.Files {
			files = append(files, file)
		}
	}
	return files
}

func (s *Representation) getType(node ast.Node) types.Type {
	if node == nil {
		return nil
	}
	ident, ok := node.(*ast.Ident)
	if !ok {
		return nil
	}
	return s.getIdentType(ident)
}

func (s *Representation) getIdentType(ident *ast.Ident) types.Type {
	obj := s.info.ObjectOf(ident)
	if obj == nil {
		return nil
	}
	return s.info.ObjectOf(ident).Type()
}
