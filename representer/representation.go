package representer

import (
	"encoding/json"
	"fmt"
	"go/ast"
	"go/printer"
	"go/token"
	"path"
	"strconv"
	"strings"
)

// Process processes the solutions AST and extracts the representation.
func (s *Representation) Process() {
	pkg := s.Package()
	s.normalize(pkg)
	s.represent = pkg
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
	var (
		pkgCode    string
		filesCount = len(s.represent.Files)
	)
	for _, file := range s.represent.Files {
		if 1 < filesCount {
			pkgCode += fmt.Sprintf("\n\n// ----- File: %s -----\n\n", file.Name.String())
		}

		code, err := s.buildCode(file)
		if err != nil {
			return nil, fmt.Errorf("failed to build representation: %w", err)
		}
		pkgCode += code
	}
	return []byte(pkgCode), nil
}

func (s *Representation) getPlaceHolder(name string) string {
	if isBuiltIn(name) {
		return name
	}
	if plh, ok := s.mapping[name]; ok {
		return plh
	}

	s.plhInc++
	plh := "PLACEHOLDER_" + strconv.Itoa(s.plhInc)
	s.mapping[name] = plh
	return plh
}

func (s *Representation) buildCode(n ast.Node) (string, error) {
	sb := &strings.Builder{}
	err := printer.Fprint(sb, token.NewFileSet(), n)
	if err != nil {
		return "", fmt.Errorf("failed to build code: %w", err)
	}
	return sb.String(), nil
}

func (s *Representation) collectImport(node ast.Node) {
	n, ok := node.(*ast.ImportSpec)
	if !ok {
		return
	}
	name := n.Path.Value
	name = strings.Trim(name, "\"")
	if strings.Contains(name, "/") {
		_, name = path.Split(name)
	}
	if n.Name != nil {
		name = n.Name.Name
	}
	s.importNames = append(s.importNames, name)
}

func (s *Representation) isImport(name string) bool {
	for _, importName := range s.importNames {
		if importName == name {
			return true
		}
	}
	return false
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
