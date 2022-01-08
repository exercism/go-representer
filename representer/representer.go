package representer

import (
	"fmt"
	"net/http"
	"os"
	"path"
	"strings"

	"github.com/pkg/errors"
	"github.com/tehsphinx/astrav"
)

// Extract extracts the Representation from a given solution folder.
func Extract(path string) (*Representation, error) {
	pkg, err := LoadPackage(path)
	if err != nil {
		return nil, fmt.Errorf("failed to parse solution AST: %w", err)
	}
	if pkg == nil {
		return nil, errors.New("no Go package found")
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
	_, err := folder.ParseFolder(func(info os.FileInfo) bool {
		return !strings.HasSuffix(info.Name(), "_test.go") && info.Name() != "embed.go"
	})
	if err != nil {
		return nil, errors.WithStack(err)
	}
	return folder.Package(folder.Pkg.Name()), nil
}
