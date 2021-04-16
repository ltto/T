package node

import (
	"github.com/beevik/etree"
	"strings"
)

func NewXmlToken(token etree.Token) *XmlToken {
	x := &XmlToken{Token: token}
	switch e := token.(type) {
	case *etree.Element:
		x.e = e
		x.child = make([]Node, 0, len(e.Child))
		for i := range e.Child {
			xmlToken := NewXmlToken(e.Child[i])
			x.child = append(x.child, doPareChild(xmlToken)...)
		}
	case *etree.CharData:
		x.data = e.Data
	}
	return x
}

func doPareChild(elem Token) (nodes []Node) {
	switch strings.ToLower(elem.Tag()) {
	case "if":
		nodes = append(nodes, TokenNewNodeIf(elem))
	case "foreach":
		nodes = append(nodes, TokenNewNodeForeach(elem))
	case "include":
		nodes = append(nodes, TokenNewNodeInclude(elem))
	default:
		nodes = append(nodes, TokenNewNodeText(elem))
	}
	return nodes
}

type XmlToken struct {
	etree.Token
	e     *etree.Element
	data  string
	child []Node
}

func (x *XmlToken) Child() []Node {
	return x.child
}
func (x *XmlToken) Tag() string {
	if x.e != nil {
		return x.e.Tag
	}
	return ""
}

func (x *XmlToken) Data() string {
	return x.data
}

func (x *XmlToken) Attr(key string) string {
	if x.e != nil {
		return x.e.SelectAttrValue(key, "")
	}
	return ""
}
