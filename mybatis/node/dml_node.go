package node

import (
	"strings"
)

type DMLRoot struct {
	ID               string
	SQLInclude       map[string]*DMLRoot
	Child            []Node
	Method           string
	UseGeneratedKeys bool
}

func (n *DMLRoot) Pare(args map[string]interface{}) (s string, err error) {
	newArgs := args
	//use temp for foreach
	newArgs["_temp"] = map[string]string{}
	//use sql for include tag
	newArgs["_sql"] = n.SQLInclude
	nodes, err := PareNodes(newArgs, n.Child)
	return strings.TrimSpace(nodes), err
}
