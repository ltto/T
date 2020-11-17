package node

import (
	"github.com/beevik/etree"
	"regexp"
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

func NewNodeIfXML(es *etree.Element) *IF {
	child := PareChildXML(es.Child)
	test := es.SelectAttrValue("test", "")
	return &IF{
		Child: child,
		Test:  test,
	}
}

func NewNodeForEachXML(es *etree.Element) *Foreach {
	child := PareChildXML(es.Child)
	return &Foreach{
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
func NewNodeText(str string) *Text {
	var reg, _ = regexp.Compile("[ |\n\t]+")
	return &Text{Text: reg.ReplaceAllString(str, " ")}
}

func NewNodeIF(test string, root Cell) *IF {
	return &IF{
		parent: root,
		Test:   test,
	}
}

func NewNodeForeach(item, index, collection, open, separator, close string, root Cell) *Foreach {
	return &Foreach{
		parent:     root,
		Close:      close,
		Collection: collection,
		Index:      index,
		Item:       item,
		Open:       open,
		Separator:  separator,
	}
}

func NewNodeInclude(refId string) *Include {
	return &Include{
		RefId: refId,
	}
}
