package representer

import (
	"fmt"
	"net/http"
	"path"

	"github.com/pkg/errors"
	"github.com/tehsphinx/astrav"
)

// Extract extracts the Representation from a given solution folder.
func Extract(path string) (*Representation, error) {
	pkg, err := LoadPackage(path)
	if err != nil {
		return nil, fmt.Errorf("failed to parse solution AST: %w", err)
	}

	repr := NewRepresentation()
	repr.Process(pkg)
	return repr, nil
}

// LoadPackage loads a go package from a folder
func LoadPackage(dir string) (*astrav.Package, error) {
	root := http.Dir(".")
	if path.IsAbs(dir) {
		root = "/"
	}

	folder := astrav.NewFolder(root, dir)
	_, err := folder.ParseFolder()
	if err != nil {
		return nil, errors.WithStack(err)
	}
	return folder.Package(folder.Pkg.Name()), nil
}
