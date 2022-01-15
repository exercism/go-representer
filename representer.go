package representer

import (
	"fmt"

	"github.com/exercism/go-representer/representation"
	"github.com/pkg/errors"
)

// Extract extracts and returns the representation and mapping from given solution path.
func Extract(solutionPath string) ([]byte, []byte, error) {
	repr := representation.NewRepresentation(solutionPath)
	if _, err := repr.ParseAST(); err != nil {
		return nil, nil, errors.WithMessage(err, "failed to parse solution AST")
	}

	if err := repr.Normalize(); err != nil {
		return nil, nil, err
	}

	reprBts, err := repr.RepresentationBytes()
	if err != nil {
		return nil, nil, fmt.Errorf("failed to serialize representation: %w", err)
	}
	mappingBts, err := repr.MappingBytes()
	if err != nil {
		return nil, nil, fmt.Errorf("failed to serialize representation: %w", err)
	}
	return reprBts, mappingBts, nil
}
