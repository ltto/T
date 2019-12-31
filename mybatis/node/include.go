package node

import "github.com/beevik/etree"

func NewNodeInclude(es *etree.Element) *Include {
	return &Include{
		RefId: es.SelectAttrValue("refid", ""),
	}
}

type Include struct {
	RefId string
}

func (n *Include) Pare(m map[string]interface{}) (s string, err error) {
	sqlM := m["_sql"].(*map[string]*etree.Element)
	sql := (*sqlM)[n.RefId]
	sqlChild := PareChild(sql.Child)
	return PareNodes(m, sqlChild)
}
