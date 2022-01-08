package representer

import (
	"encoding/json"
	"fmt"

	"github.com/tehsphinx/astrav"
)

// NewRepresentation creates a new empty representation.
func NewRepresentation() *Representation {
	return &Representation{
		mapping: map[string]string{},
	}
}

// Representation contains all information of a representation for a solution.
type Representation struct {
	plhInc    int
	mapping   map[string]string // key: variable name, value: placeholder name
	represent node
}

// Process processes the solutions AST and extracts the representation.
func (s *Representation) Process(pkg astrav.Node) {
	s.represent, _ = s.buildNode(pkg)
}

// MappingBytes retrieves the correct mapping to be written to mapping.json.
func (s *Representation) MappingBytes() ([]byte, error) {
	m := make(map[string]string, len(s.mapping))
	for k, v := range s.mapping {
		// inverts the mapping since the placeholders must be the keys
		m[v] = k
	}

	bts, err := toJSON(m)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal mapping: %w", err)
	}
	return bts, nil
}

// RepresentationBytes retrieves the bytes of the representation.
func (s *Representation) RepresentationBytes() ([]byte, error) {
	bts, err := toJSON(s.represent)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal representation: %w", err)
	}
	return bts, nil
}

func toJSON(res interface{}) ([]byte, error) {
	return json.MarshalIndent(res, "", "\t")
}
