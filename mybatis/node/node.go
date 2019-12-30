package node

import (
	"errors"
	"reflect"
	"strconv"
	"strings"

	"github.com/beevik/etree"
)

//配置解析
type Node interface {
	Pare(m map[string]interface{}) (s string, err error)
}

func PareChild(es []etree.Token) (nodes []Node) {
	for i := range es {
		elem := es[i]
		switch e := elem.(type) {
		case *etree.Element:
			switch strings.ToLower(e.Tag) {
			case "if":
				nodes = append(nodes, NewNodeIf(e))
			case "foreach":
				nodes = append(nodes, NewNodeForEach(e))
			case "include":
				nodes = append(nodes, NewNodeInclude(e))
			}
		case *etree.CharData:
			nodes = append(nodes, NewNodeText(e.Data))
		}
	}
	return
}

func PareNodes(m map[string]interface{}, nodes []Node) (s string, err error) {
	for _, node := range nodes {
		pare, err := node.Pare(m)
		if err != nil {
			return s, err
		}
		s += pare
	}
	return
}

func FindStr(m map[string]interface{}, str string) (interface{}, error) {
	s := str
	var objs = []string{s}
	var apps = []string{""}
	count := strings.Count(s, ".")
	for i := 0; i < count; i++ {
		objs = append(objs, s[:strings.LastIndex(s, ".")])
		apps = append(apps, str[strings.LastIndex(s, ".")+1:])
		s = s[:strings.LastIndex(s, ".")]
	}
	var app []string
	var inter interface{}
	for idx, o := range objs {
		if i, ok := m[o]; ok {
			inter = i
			app = strings.Split(apps[idx], ".")
			break
		}
	}
	if inter == nil {
		return str, nil
	}
	if len(app) == 0 {
		return inter, nil
	}
	return find(reflect.ValueOf(inter), app)
}
func find(v reflect.Value, app []string) (interface{}, error) {
	for _, k := range app {
		for v.Kind() == reflect.Ptr {
			v = v.Elem()
		}
		switch v.Kind() {
		case reflect.Struct:
			v = v.FieldByName(k)
		case reflect.Map:
			v = v.MapIndex(reflect.ValueOf(k))
		case reflect.Slice:
			i, err := strconv.Atoi(k)
			if err != nil {
				return nil, err
			}
			v = v.Index(i)
		default:
			return nil, errors.New("type error")
		}
	}
	return v.Interface(), nil
}
