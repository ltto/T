package node

import (
	"fmt"
	"github.com/ltto/T/gobox/str"
	"strings"
)

type DMLRoot struct {
	ID               string
	SQLInclude       map[string]*DMLRoot
	Child            []Node
	Method           string
	UseGeneratedKeys bool
}

func Select() *DMLRoot {
	return &DMLRoot{Method: "SELECT"}
}

func (n *DMLRoot) IF(test string) *IF {
	nodeIF := NewNodeIF(test, n)
	n.Child = append(n.Child, nodeIF)
	return nodeIF
}

func (n *DMLRoot) Foreach(item, index, collection, open, separator, close string) *Foreach {
	foreach := NewNodeForeach(item, index, collection, open, separator, close, n)
	n.Child = append(n.Child, foreach)
	return foreach
}

func (n *DMLRoot) Include(refId string) Cell {
	n.Child = append(n.Child, NewNodeInclude(refId))
	return n
}

func (n *DMLRoot) Text(s string) Cell {
	n.Child = append(n.Child, NewNodeText(s))
	return n
}

type PrePareSQL struct {
	SQL     string
	Params  []interface{}
	Operate SQLOperate
}

func (n *DMLRoot) pare(args map[string]interface{}) (s string, err error) {
	var newArgs map[string]interface{}
	if args == nil {
		newArgs = map[string]interface{}{}
	} else {
		newArgs = args
	}
	//use temp for foreach
	newArgs["_temp"] = map[string]string{}
	//use sql for include tag
	newArgs["_sql"] = n.SQLInclude
	nodes, err := PareNodes(newArgs, n.Child)
	return strings.TrimSpace(nodes), err
}

// 预解析SQL
func (n *DMLRoot) PareSQL(args map[string]interface{}) (preSQL *PrePareSQL, err error) {
	pare, err := n.pare(args)
	if err != nil {
		return nil, err
	}
	p := &PrePareSQL{
		SQL:     pare,
		Operate: Operate(n.Method),
	}
	p.SQL = pare
	p.SQL = str.Expand('$', p.SQL, func(s string) string {
		return fmt.Sprintf("%v", args[s])
	})
	p.SQL = str.Expand('#', p.SQL, func(s string) string {
		p.Params = append(p.Params, args[s])
		return "?"
	})
	return p, nil
}
