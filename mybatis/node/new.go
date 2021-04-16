package node

import (
	"regexp"
	"strconv"
)

func NewNodeRoot(root Token, includes map[string]Token) *DMLRoot {
	UseGeneratedKeys, _ := strconv.ParseBool(root.Attr("useGeneratedKeys"))
	dmlRoot := &DMLRoot{
		Child:            root.Child(),
		ID:               root.Attr("id"),
		SQLInclude:       make(map[string]*DMLRoot),
		Method:           root.Tag(),
		UseGeneratedKeys: UseGeneratedKeys,
	}
	for k, element := range includes {
		dmlRoot.SQLInclude[k] = NewNodeRoot(element, nil)
	}
	return dmlRoot
}

var reg, _ = regexp.Compile("[ |\n\t]+")

func NewNodeText(str string) *Text {
	return &Text{Text: reg.ReplaceAllString(str, " ")}
}

func NewNodeIF(test string, child []Node) *IF {
	return &IF{Test: test, Child: child}
}

func NewNodeForeach(item, index, collection, open, separator, close string, child []Node) *Foreach {
	return &Foreach{
		Child:      child,
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

func TokenNewNodeIf(x Token) *IF {
	return NewNodeIF(x.Attr("test"), x.Child())
}

func TokenNewNodeForeach(x Token) *Foreach {
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

func TokenNewNodeInclude(x Token) *Include {
	return NewNodeInclude(x.Attr("refid"))
}

func TokenNewNodeText(x Token) *Text {
	return NewNodeText(x.Data())
}
