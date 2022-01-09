package representer

import (
	"fmt"
	"net/http"
	"path"

	"github.com/pkg/errors"
)

// Extract extracts the Representation from a given solution folder.
func Extract(path string) (*Representation, error) {
	repr, err := GetRepresentation(path)
	if err != nil {
		return nil, fmt.Errorf("failed to parse solution AST: %w", err)
	}
	return repr, nil
}

// GetRepresentation loads a go package from a folder
func GetRepresentation(dir string) (*Representation, error) {
	root := http.Dir(".")
	if path.IsAbs(dir) {
		root = "/"
	}

	repr := NewRepresentation(root, dir)
	_, err := repr.ParseFolder()
	if err != nil {
		return nil, errors.WithStack(err)
	}
	if repr.Package() == nil {
		return nil, errors.New("no Go package found")
	}

	repr.Process()

	return repr, nil
}
