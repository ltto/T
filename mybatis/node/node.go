package node

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"
)

// Node 配置解析
type Node interface {
	pare(args map[string]interface{}) (s string, err error)
}

func PareNodes(args map[string]interface{}, nodes []Node) (s string, err error) {
	for _, node := range nodes {
		pare, err := node.pare(args)
		if err != nil {
			return s, err
		}
		s += pare
	}
	return
}

func FindStr(args map[string]interface{}, str string) (interface{}, error) {
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
		if i, ok := args[o]; ok {
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
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			return v.Int(), nil
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
			return v.Uint(), nil
		case reflect.Float64, reflect.Float32:
			return v.Float(), nil
		case reflect.String:
			return v.String(), nil
		case reflect.Bool:
			return v.Bool(), nil
		default:
			return nil, fmt.Errorf("%v type not supported", v.Kind())
		}
	}
	return v.Interface(), nil
}
