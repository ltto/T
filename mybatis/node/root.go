package node

import (
	"strings"

	"github.com/beevik/etree"
)

func NewNodeRoot(root *etree.Element, Sql *map[string]*etree.Element) *DMLRoot {
	id := root.SelectAttrValue("id", "")
	tag := root.Tag
	child := PareChild(root.Child)
	return &DMLRoot{
		Child: child,
		ID:    id,
		Tag:   tag,
		Sql:   Sql,
	}
}

type DMLRoot struct {
	ID    string
	Tag   string
	Sql   *map[string]*etree.Element
	Child []Node
}

func (n *DMLRoot) Pare(m map[string]interface{}) (s string, err error) {
	newM := m
	newM["conv"] = map[string]string{}
	newM["_sql"] = n.Sql
	nodes, err := PareNodes(newM, n.Child)
	return strings.TrimSpace(nodes), err
}
