package node

import "regexp"

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
