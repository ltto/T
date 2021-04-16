package mybatis

import (
	"github.com/ltto/T/mybatis/node"
	"strconv"
)

func NewNodeRoot(root node.Token, includes map[string]node.Token) *node.DMLRoot {
	UseGeneratedKeys, _ := strconv.ParseBool(root.Attr("useGeneratedKeys"))
	dmlRoot := &node.DMLRoot{
		Child:            root.Child(),
		ID:               root.Attr("id"),
		SQLInclude:       make(map[string]*node.DMLRoot),
		Method:           root.Tag(),
		UseGeneratedKeys: UseGeneratedKeys,
	}
	for k, element := range includes {
		dmlRoot.SQLInclude[k] = NewNodeRoot(element, nil)
	}
	return dmlRoot
}
