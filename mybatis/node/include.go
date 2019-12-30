package node

import "github.com/beevik/etree"

func NewNodeInclude(es *etree.Element) *Include {
	child := PareChild(es.Child)
	return &Include{
		Child: child,
		Id:    es.SelectAttrValue("refid", ""),
	}
}

type Include struct {
	Child []Node
	Id    string
}

func (n *Include) Pare(m map[string]interface{}) (s string, err error) {
	return PareNodes(m, n.Child)
}
