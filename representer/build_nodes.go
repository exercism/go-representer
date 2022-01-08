package representer

import (
	"github.com/tehsphinx/astrav"
)

type skipType int

const (
	skipNone skipType = iota
	skipNode
	skipAll
)

type node struct {
	NodeType string `json:"0_ast_type"`
	Name     string `json:"1_name,omitempty"`
	GoType   string `json:"2_type,omitempty"`
	Value    string `json:"3_value,omitempty"`
	Children []node `json:"7_children,omitempty"`
}

func (s *Representation) buildNode(n astrav.Node) (node, skipType, []astrav.Node) {
	var rNode node
	// add more node types to handle them specifically as needed
	children := n.Children()
	switch nde := n.(type) {
	default:
		rNode = s.defaultNode(n, true)
	case *astrav.FuncDecl:
		rNode = s.funcDecl(nde)
		// continue with children of function block
		children = n.ChildByNodeType(astrav.NodeTypeBlockStmt).Children()
	case *astrav.CallExpr:
		rNode, children = s.callExpr(nde)
	case *astrav.Comment, *astrav.CommentGroup:
		// ignore node and all children
		return node{}, skipAll, nil
	case *astrav.File:
		children = n.ChildNodes(func(n astrav.Node) bool { return !n.IsNodeType(astrav.NodeTypeIdent) })
		// ignore node, process children
		return node{}, skipNode, children
	}
	rNode = s.appendChildren(rNode, children)
	return rNode, skipNone, nil
}

func (s *Representation) appendChildren(parent node, children []astrav.Node) node {
	for _, n := range children {
		chNode, skip, chs := s.buildNode(n)
		switch skip {
		case skipAll:
			continue
		case skipNode:
			parent = s.appendChildren(parent, chs)
			continue
		}
		parent.Children = append(parent.Children, chNode)
	}
	return parent
}

func (s *Representation) defaultNode(n astrav.Node, translateName bool) node {
	reprNode := node{NodeType: formatNodeType(n.NodeType())}
	if named, ok := n.(astrav.Named); ok {
		name := named.NodeName()
		if translateName {
			name = s.getPlaceHolder(name)
		}
		reprNode.Name = name
	}
	if t := n.ValueType(); t != nil {
		reprNode.GoType = t.String()
	}

	// simple node specifics. more complex ones get a switch case in buildNode method.
	if nde, ok := n.(*astrav.BasicLit); ok {
		reprNode.Value = nde.Value
	}
	if nde, ok := n.(*astrav.BinaryExpr); ok {
		reprNode.Value = nde.Op.String()
	}

	return reprNode
}

func (s *Representation) funcDecl(n *astrav.FuncDecl) node {
	reprNode := s.defaultNode(n, true)

	funcType, ok := n.ChildByNodeType(astrav.NodeTypeFuncType).(*astrav.FuncType)
	if !ok {
		return reprNode
	}

	s.replaceFieldNames(funcType.FindByNodeType(astrav.NodeTypeField))
	reprNode.GoType = s.buildCode(funcType.FuncType)
	return reprNode
}

func (s *Representation) callExpr(n *astrav.CallExpr) (node, []astrav.Node) {
	reprNode := s.defaultNode(n, false)
	if nde, ok := n.SelExpr().(*astrav.SelectorExpr); ok {
		reprNode.Name = nde.NodeName()
		if t := nde.ValueType(); t != nil {
			reprNode.GoType = t.String()
		}
	}

	// ignore ident children
	children := n.ChildNodes(func(nde astrav.Node) bool { return nde != n.SelExpr() })
	return reprNode, children
}
