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
		nodes = append(nodes, elem.NewNodeIf())
	case "foreach":
		nodes = append(nodes, elem.NewNodeForeach())
	case "include":
		nodes = append(nodes, elem.NewNodeInclude())
	default:
		nodes = append(nodes, elem.NewNodeText())
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

func (x *XmlToken) NewNodeIf() *IF {
	return NewNodeIF(x.Attr("test"), x.Child())
}

func (x *XmlToken) NewNodeForeach() *Foreach {
	return NewNodeForeach(
		x.Attr("item"),
		x.Attr("index"),
		x.Attr("collection"),
		x.Attr("open"),
		x.Attr("separator"),
		x.Attr("close"),
		x.Child(),
	)
}

func (x *XmlToken) NewNodeInclude() *Include {
	return NewNodeInclude(x.Attr("refid"))
}

func (x *XmlToken) NewNodeText() *Text {
	return NewNodeText(x.Data())
}
