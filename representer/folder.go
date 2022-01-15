package representer

import (
	"go/ast"
	"go/importer"
	"go/parser"
	"go/token"
	"go/types"
	"net/http"
	"os"
	"strings"

	"github.com/pkg/errors"
	"github.com/tehsphinx/astrav"
)

const defaultFileName = "solution.go"

// NewRepresentation creates a new folder with given path. Use ParseFolder to parse ast from go files in path.
// The pkgPath is the import path of the package to be used by types.ParseInfo.
func NewRepresentation(root http.FileSystem, dir string) *Representation {
	return &Representation{
		dir:  dir,
		root: root,
		FSet: token.NewFileSet(),
		Pkgs: map[string]*ast.Package{},

		mapping: map[string]string{},
	}
}

// Representation represents a go package folder
type Representation struct {
	dir  string
	root http.FileSystem

	Info *types.Info
	FSet *token.FileSet
	Pkgs map[string]*ast.Package
	Pkg  *types.Package

	plhInc      int
	mapping     map[string]string // key: variable name, value: placeholder name
	represent   *ast.Package
	importNames []string
}

// ParseFolder will parse all to files in folder. It skips test files.
func (s *Representation) ParseFolder() (map[string]*ast.Package, error) {
	pkgs, _, err := astrav.Parse(s.FSet, s.root, s.dir, func(info os.FileInfo) bool {
		return !strings.HasSuffix(info.Name(), "_test.go") && info.Name() != "embed.go"
	}, parser.AllErrors)
	if err != nil {
		return nil, err
	}

	for name, pkg := range pkgs {
		s.Pkgs[name] = pkg
	}

	if s.Pkg, err = s.ParseInfo(s.dir, s.FSet, s.getFiles()); err != nil {
		return nil, err
	}

	return s.Pkgs, nil
}

// ParseInfo parses all files for type information which is then available
// from the Nodes. When using Representation.ParseFolder, this is done automatically.
func (s *Representation) ParseInfo(path string, fSet *token.FileSet, files []*ast.File) (*types.Package, error) {
	s.Info = &types.Info{
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

	pkg, err := conf.Check(path, fSet, files, s.Info)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return pkg, nil
}

// Package returns a package by name
func (s *Representation) Package() *ast.Package {
	for _, pkg := range s.Pkgs {
		return pkg
	}
	return nil
}

func (s *Representation) getFiles() []*ast.File {
	var files []*ast.File
	for _, pkg := range s.Pkgs {
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
	obj := s.Info.ObjectOf(ident)
	if obj == nil {
		return nil
	}
	return s.Info.ObjectOf(ident).Type()
}
