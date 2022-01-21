package representation

import (
	"encoding/json"
	"fmt"
	"go/ast"
	"go/printer"
	"go/token"
	"strings"
)

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
	code, err := s.buildCode(s.represent.Files[defaultFileName])
	if err != nil {
		return nil, fmt.Errorf("failed to build representation: %w", err)
	}

	return []byte(code), nil
}

func (s *Representation) buildCode(n ast.Node) (string, error) {
	var (
		sb = &strings.Builder{}
		fs = token.NewFileSet()
	)
	// make sure to use new fileset here to lose positions, e.g. extra whitespace
	err := printer.Fprint(sb, fs, n)
	if err != nil {
		return "", fmt.Errorf("failed to build code: %w", err)
	}
	return sb.String(), nil
}

func toJSON(res interface{}) ([]byte, error) {
	return json.MarshalIndent(res, "", "\t")
}

var builtIn = []string{"", "_", "nil",
	// constants
	"true", "false", "iota",
	// functions
	"append", "cap", "close", "complex", "copy", "delete", "imag", "len", "make", "new", "panic", "print", "println", "real", "recover",
	// types
	"bool", "byte", "complex64", "complex128", "error", "float32", "float64",
	"int", "int8", "int16", "int32", "int64", "rune", "string",
	"uint", "uint8", "uint16", "uint32", "uint64", "uintptr",
}

func isBuiltIn(name string) bool {
	for _, word := range builtIn {
		if name == word {
			return true
		}
	}
	return false
}
