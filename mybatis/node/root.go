package node

import (
	"strings"

	"github.com/beevik/etree"
)


func NewNodeRoot(root *etree.Element) *Root {
	id := root.SelectAttrValue("id", "")
	tag := root.Tag
	child := PareChild(root.Child)
	return &Root{
		Child: child,
		ID:    id,
		Tag:   tag,
	}
}

type Root struct {
	ID    string
	Tag   string
	Child []Node
}

func (n *Root) Pare(m map[string]interface{}) (s string, err error) {
	newM := m
	newM["conv"] = map[string]string{}
	nodes, err := PareNodes(newM, n.Child)
	return strings.TrimSpace(nodes), err
}
