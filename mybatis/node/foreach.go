package node

import (
	"errors"
	"fmt"
	"reflect"
)

type Foreach struct {
	Index      string
	Item       string
	Collection string
	Open       string
	Separator  string
	Close      string
	Child      []Node
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
