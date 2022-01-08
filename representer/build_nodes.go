package representer

import (
	"strconv"
	"strings"

	"github.com/tehsphinx/astrav"
)

type skipType int

const (
	skipNone skipType = iota
	skipNode
	skipAll
)

type node struct {
	T        string `json:"0_t"`
	PlHolder string `json:"1_plh,omitempty"`
	Children []node `json:"7_ch"`
}

func (s *Representation) buildNode(n astrav.Node) (node, skipType) {
	var rNode node
	// add more node types to handle them specifically as needed
	switch n.NodeType() {
	default:
		rNode = s.defaultNode(n)
	case astrav.NodeTypeComment, astrav.NodeTypeCommentGroup:
		// ignore node and all children
		return node{}, skipAll
	case astrav.NodeTypeFile:
		// ignore node, process children
		return node{}, skipNode
	}
	rNode = s.buildChildren(rNode, n)
	return rNode, skipNone
}

func (s *Representation) buildChildren(rNode node, n astrav.Node) node {
	for _, n := range n.Children() {
		chNode, skip := s.buildNode(n)
		switch skip {
		case skipAll:
			continue
		case skipNode:
			rNode = s.buildChildren(rNode, n)
			continue
		}
		rNode.Children = append(rNode.Children, chNode)
	}
	return rNode
}

func (s *Representation) defaultNode(n astrav.Node) node {
	reprNode := node{T: formatNodeType(n.NodeType())}
	if named, ok := n.(astrav.Named); ok {
		reprNode.PlHolder = s.getPlaceHolder(named.NodeName())
	}
	return reprNode
}

func (s *Representation) getPlaceHolder(name string) string {
	if plh, ok := s.mapping[name]; ok {
		return plh
	}

	s.plhInc++
	plh := "PLACEHOLDER_" + strconv.Itoa(s.plhInc)
	s.mapping[name] = plh
	return plh
}

func formatNodeType(t astrav.NodeType) string {
	return strings.TrimPrefix(string(t), "*astrav.")
}
