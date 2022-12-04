package representation

import (
	"go/ast"
	"go/token"
	"strconv"
	"strings"

	"golang.org/x/tools/go/ast/astutil"
)

func (s *Representation) normalizeDeclarations(pkg *ast.Package) {
	astutil.Apply(pkg, func(cursor *astutil.Cursor) bool {
		node := cursor.Node()
		if node == nil {
			return true
		}

		switch n := node.(type) {
		default:
			return true
		case *ast.DeclStmt:
			genDecl, ok := n.Decl.(*ast.GenDecl)
			if !ok {
				return true
			}
			if genDecl.Tok.String() != "var" {
				return true
			}

			var keep bool
			for _, spec := range genDecl.Specs {
				specNode, ok := spec.(*ast.ValueSpec)
				if !ok {
					keep = true
					continue
				}

				for i, name := range specNode.Names {
					val := getValue(specNode.Values, i)
					if val == nil {
						val = getDefaultLiteralForExpression(specNode.Type)
					}
					val = addTypeAsNeeded(val, specNode.Type)
					newDecl := newAssignment(name.Name, val)
					cursor.InsertBefore(newDecl)
				}
			}
			if !keep {
				cursor.Delete()
			}
		}

		return true
	}, nil)
}

func addTypeAsNeeded(val ast.Expr, typeExpr ast.Expr) ast.Expr {
	if _, ok := val.(*ast.BasicLit); !ok {
		return val
	}
	typeIdent, ok := typeExpr.(*ast.Ident)
	if !ok {
		return val
	}
	if !needsType(typeIdent.Name) {
		return val
	}

	return &ast.CallExpr{
		Fun:  typeIdent,
		Args: []ast.Expr{val},
	}
}

func needsType(ident string) bool {
	switch ident {
	case "string", "int", "float64", "complex128", "bool":
		return false
	}
	return true
}

func newAssignment(name string, value ast.Expr) ast.Node {
	return &ast.AssignStmt{
		Lhs: []ast.Expr{&ast.Ident{Name: name}},
		Tok: token.DEFINE,
		Rhs: []ast.Expr{value},
	}
}

func getValue(values []ast.Expr, i int) ast.Expr {
	if len(values) > i {
		return values[i]
	}
	return nil
}

func getDefaultLiteralForExpression(typeExpr ast.Expr) ast.Expr {
	typeIdent, ok := typeExpr.(*ast.Ident)
	if !ok {
		return &ast.BasicLit{
			Kind:  token.STRING,
			Value: `"unimplemented case in representer: getDefaultLiteralForExpression: typeExpr is not *ast.Ident"`,
		}
	}

	kind := typeLookup(typeIdent.Name)
	switch kind {
	case token.IDENT:
		return &ast.CallExpr{
			Fun: typeIdent,
			Args: []ast.Expr{
				&ast.BasicLit{
					Kind:  token.INT,
					Value: `0`,
				},
			},
		}
	}
	return &ast.BasicLit{
		Kind:  kind,
		Value: getDefaultLiteral(kind),
	}
}

func typeLookup(name string) token.Token {
	switch strings.ToLower(name) {
	case "int":
		return token.INT
	case "float":
		return token.FLOAT
	case "imag":
		return token.IMAG
	case "char":
		return token.CHAR
	case "string":
		return token.STRING
	}
	return token.IDENT
}

func getDefaultLiteral(kind token.Token) string {
	switch kind {
	case token.INT:
		return `0`
	case token.FLOAT:
		return `0.0`
	case token.IMAG:
		return `0i`
	case token.CHAR:
		return `0`
	case token.STRING:
		return `""`
	}
	return `"unimplemented case in representer: getDefaultLiteral: unknown kind: ` + strconv.Itoa(int(kind)) + `"`
}
