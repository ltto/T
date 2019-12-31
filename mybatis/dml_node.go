package mybatis

import (
	"strconv"
	"strings"

	"github.com/beevik/etree"
	"github.com/ltto/T/mybatis/node"
)

func NewNodeRoot(root *etree.Element, Sql *map[string]*etree.Element) *DMLRoot {
	id := root.SelectAttrValue("id", "")
	tag := root.Tag
	child := node.PareChild(root.Child)
	UseGeneratedKeys := false
	if tag == "insert" {
		UseGeneratedKeys, _ = strconv.ParseBool(root.SelectAttrValue("useGeneratedKeys", "false"))
	}
	return &DMLRoot{
		Child:            child,
		ID:               id,
		Tag:              tag,
		Sql:              Sql,
		UseGeneratedKeys: UseGeneratedKeys,
	}
}

type DMLRoot struct {
	ID               string
	Tag              string
	UseGeneratedKeys bool
	Sql              *map[string]*etree.Element
	Child            []node.Node
}

func (n *DMLRoot) Pare(m map[string]interface{}) (s string, err error) {
	newM := m
	newM["conv"] = map[string]string{}
	newM["_sql"] = n.Sql
	nodes, err := node.PareNodes(newM, n.Child)
	return strings.TrimSpace(nodes), err
}
