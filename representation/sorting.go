package representation

import (
	"fmt"
	"go/ast"
	"sort"
)

// constants for top level declarations in a file.
// The order of the constants defines the sorting order in the representation.
const (
	typeImport = iota + 1
	typeConst
	typeVar
	typeFunc
)

var declTypes = map[string]int{
	"import":        typeImport,
	"const":         typeConst,
	"var":           typeVar,
	"*ast.FuncDecl": typeFunc,
}

func (s *Representation) sortInFile(f *ast.File) {
	sort.Slice(f.Decls, func(i, j int) bool {
		typeI := typeString(f.Decls[i])
		typeJ := typeString(f.Decls[j])
		if typeI == typeJ {
			sizeI := f.Decls[i].End() - f.Decls[i].Pos()
			sizeJ := f.Decls[j].End() - f.Decls[j].Pos()
			return sizeJ < sizeI
		}
		return typeI < typeJ
	})
}

func typeString(decl ast.Decl) int {
	str := fmt.Sprintf("%T", decl)
	if str != "*ast.GenDecl" {
		return declTypes[str]
	}
	return declTypes[decl.(*ast.GenDecl).Tok.String()]
}
