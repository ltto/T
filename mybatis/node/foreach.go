package node

import (
	"errors"
	"fmt"
	"reflect"
)

type Foreach struct {
	parent     Cell
	Child      []Node
	Item       string
	Index      string
	Collection string
	Open       string
	Separator  string
	Close      string
}

func (n *Foreach) EndForeach() Cell {
	return n.parent
}

func (n *Foreach) IF(test string) *IF {
	nodeIF := NewNodeIF(test, n)
	n.Child = append(n.Child, nodeIF)
	return nodeIF
}

func (n *Foreach) Foreach(item, index, collection, open, separator, close string) *Foreach {
	foreach := NewNodeForeach(item, index, collection, open, separator, close, n)
	n.Child = append(n.Child, foreach)
	return foreach
}

func (n *Foreach) Include(refId string) Cell {
	n.Child = append(n.Child, NewNodeInclude(refId))
	return n
}

func (n *Foreach) Text(s string) Cell {
	n.Child = append(n.Child, NewNodeText(s))
	return n
}

func (n *Foreach) pare(args map[string]interface{}) (s string, err error) {
	coll := args[n.Collection]
	nodes := ""
	if coll == nil {
		return n.Open + nodes + n.Close, nil
	}
	v := reflect.ValueOf(coll)
	for v.Kind() == reflect.Ptr {
		v = v.Elem()
	}
	temp := args["_temp"].(map[string]string)
	switch v.Kind() {
	case reflect.Slice:
		for i := 0; i < v.Cap(); i++ {
			temp[n.Index] = fmt.Sprint(i)
			args[n.Index] = i
			temp[n.Item] = fmt.Sprintf("#{%s.%d}", n.Collection, i)
			args[fmt.Sprintf("%s.%d", n.Collection, i)] = v.Index(i).Interface()

			if ps, err := PareNodes(args, n.Child); err != nil {
				return s, err
			} else {
				nodes += ps + n.Separator
			}
		}
	case reflect.Map:
		iter := v.MapRange()
		for iter.Next() {
			args[n.Index] = fmt.Sprint(iter.Key().Interface())
			args[n.Index] = iter.Key().Interface()
			args[n.Item] = fmt.Sprintf("#{%s.%s}", n.Collection, iter.Key())
			args[fmt.Sprintf("%s.%s", n.Collection, iter.Key())] = iter.Value().Interface()
			if ps, err := PareNodes(args, n.Child); err != nil {
				return s, err
			} else {
				nodes += ps + n.Separator
			}
		}
	default:
		return s, errors.New("ForEach need slice or map")
	}
	defer func() {
		delete(args, n.Index)
		delete(temp, n.Index)
		delete(temp, n.Item)
	}()
	return n.Open + nodes[:len(nodes)-len(n.Separator)] + n.Close, nil
}
