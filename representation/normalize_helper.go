package representation

import (
	"go/ast"
	"path"
	"strconv"
	"strings"
)

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
