package node

import (
	"errors"
	"fmt"
	"reflect"

	"github.com/beevik/etree"
)

func NewNodeForEach(es *etree.Element) *ForEach {
	child := PareChild(es.Child)
	return &ForEach{
		Child:      child,
		Close:      es.SelectAttrValue("close", ""),
		Collection: es.SelectAttrValue("collection", ""),
		Index:      es.SelectAttrValue("index", ""),
		Item:       es.SelectAttrValue("item", ""),
		Open:       es.SelectAttrValue("open", ""),
		Separator:  es.SelectAttrValue("separator", ""),
	}
}

type ForEach struct {
	Child      []Node
	Item       string
	Index      string
	Collection string
	Open       string
	Separator  string
	Close      string
}

func (n *ForEach) Pare(m map[string]interface{}) (s string, err error) {
	coll := m[n.Collection]
	nodes := ""
	if coll == nil {
		return n.Open + nodes + n.Close, nil
	}
	v := reflect.ValueOf(coll)
	for v.Kind() == reflect.Ptr {
		v = v.Elem()
	}
	temp := m["_temp"].(map[string]string)
	switch v.Kind() {
	case reflect.Slice:
		for i := 0; i < v.Cap(); i++ {
			temp[n.Index] = fmt.Sprint(i)
			m[n.Index] = i
			temp[n.Item] = fmt.Sprintf("#{%s.%d}", n.Collection, i)
			m[fmt.Sprintf("%s.%d", n.Collection, i)] = v.Index(i).Interface()

			if ps, err := PareNodes(m, n.Child); err != nil {
				return s, err
			} else {
				nodes += ps + n.Separator
			}
		}
	case reflect.Map:
		iter := v.MapRange()
		for iter.Next() {
			m[n.Index] = fmt.Sprint(iter.Key().Interface())
			m[n.Index] = iter.Key().Interface()
			m[n.Item] = fmt.Sprintf("#{%s.%s}", n.Collection, iter.Key())
			m[fmt.Sprintf("%s.%s", n.Collection, iter.Key())] = iter.Value().Interface()
			if ps, err := PareNodes(m, n.Child); err != nil {
				return s, err
			} else {
				nodes += ps + n.Separator
			}
		}
	default:
		return s, errors.New("ForEach need slice or map")
	}
	defer func() {
		delete(m, n.Index)
		delete(temp, n.Index)
		delete(temp, n.Item)
	}()
	return n.Open + nodes[:len(nodes)-len(n.Separator)] + n.Close, nil
}
