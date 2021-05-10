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

type PrePareSQL struct {
	SQL     string
	Params  []interface{}
	Operate SQLOperate
}

func (p PrePareSQL) String() string {
	return fmt.Sprint(p.SQL)
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

// PareSQL 预解析SQL
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
