package mybatis

import (
	"github.com/beevik/etree"
	"github.com/ltto/T/mybatis/node"
	"strconv"
)

func NewNodeRoot(root *etree.Element, Sql map[string]*etree.Element) *node.DMLRoot {
	id := root.SelectAttrValue("id", "")
	method := root.Tag
	child := node.PareChildXML(root.Child)
	UseGeneratedKeys := false
	if method == "insert" {
		UseGeneratedKeys, _ = strconv.ParseBool(root.SelectAttrValue("useGeneratedKeys", "false"))
	}
	dmlRoot := &node.DMLRoot{
		Child:            child,
		ID:               id,
		SQLInclude:       make(map[string]*node.DMLRoot),
		Method:           method,
		UseGeneratedKeys: UseGeneratedKeys,
	}
	if len(Sql) != 0 {
		for k, element := range Sql {
			dmlRoot.SQLInclude[k] = NewNodeRoot(element, nil)
		}
	}

	return dmlRoot
}
