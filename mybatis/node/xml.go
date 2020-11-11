package node

import (
	"github.com/beevik/etree"
	"strings"
)

func PareChildXML(es []etree.Token) (nodes []Node) {
	for i := range es {
		elem := es[i]

		nodes = append(nodes, doPareChildXML(elem)...)
	}
	return
}

func doPareChildXML(elem etree.Token) (nodes []Node) {
	switch e := elem.(type) {
	case *etree.Element:
		switch strings.ToLower(e.Tag) {
		case "if":
			nodes = append(nodes, NewNodeIfXML(e))
		case "foreach":
			nodes = append(nodes, NewNodeForEachXML(e))
		case "include":
			nodes = append(nodes, NewNodeIncludeXML(e))
		}
	case *etree.CharData:
		nodes = append(nodes, NewNodeText(e.Data))
	}
	return nodes
}

func NewNodeIfXML(es *etree.Element) *If {
	child := PareChildXML(es.Child)
	test := es.SelectAttrValue("test", "")
	return &If{
		Child: child,
		Test:  test,
	}
}

func NewNodeForEachXML(es *etree.Element) *ForEach {
	child := PareChildXML(es.Child)
	return &ForEach{
		Child:      child,
		Close:      es.SelectAttrValue("close", ""),
		Collection: es.SelectAttrValue("collection", ""),
		Index:      es.SelectAttrValue("index", ""),
		Item:       es.SelectAttrValue("item", ""),
		Open:       es.SelectAttrValue("open", ""),
		Separator:  es.SelectAttrValue("separator", ""),
	}
}

func NewNodeIncludeXML(es *etree.Element) *Include {
	return &Include{
		RefId: es.SelectAttrValue("refid", ""),
	}
}
