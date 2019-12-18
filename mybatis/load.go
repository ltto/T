package mybatis

import (
	"errors"
	"fmt"
	"reflect"
	"regexp"
	"strconv"
	"strings"

	"github.com/beevik/etree"
	"github.com/ltto/T/gobox/str"
)



func MainLoad() {
	doc := etree.NewDocument()
	if err := doc.ReadFromFile("/Users/ltt/go/src/github.com/ltto/GoMybatis/Mappppper/mapper/AlbumsMapper.xml"); err != nil {
		panic(err)
	}
	element := doc.SelectElement("mapper")
	m := deep(element, element.ChildElements())
	root := NewNodeRoot(m["delete"][1])
	m2 := map[string]interface{}{
		"ids": []int{0, 1, 2, 3, 4, 5, 6, 7, 8},
	}
	pare, _ := root.Pare(m2)
	fmt.Println()
	fmt.Println(pare)
	fmt.Println()
	PareSQL(m2, pare)
	fmt.Println()
}

func deep(root *etree.Element, list []*etree.Element) map[string][]*etree.Element {
	m := make(map[string][]*etree.Element)
	for i := range list {
		elem := list[i]
		var reg, _ = regexp.Compile("[ |\n|\t]+")
		elem.SetText(reg.ReplaceAllString(elem.Text(), " "))
		if _, ok := m[elem.Tag]; !ok {
			m[elem.Tag] = root.SelectElements(elem.Tag)
		}
	}
	return m
}

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

func NewNodeRoot(root *etree.Element) *NodeRoot {
	id := root.SelectAttrValue("id", "")
	tag := root.Tag
	child := PareChild(root.Child)
	return &NodeRoot{
		Child: child,
		ID:    id,
		Tag:   tag,
	}
}

func NewNodeIf(es *etree.Element) *NodeIf {
	child := PareChild(es.Child)
	test := es.SelectAttrValue("test", "")
	return &NodeIf{
		Child: child,
		Test:  test,
	}
}
func NewNodeForEach(es *etree.Element) *NodeForEach {
	child := PareChild(es.Child)
	return &NodeForEach{
		Child:      child,
		Close:      es.SelectAttrValue("close", ""),
		Collection: es.SelectAttrValue("collection", ""),
		Index:      es.SelectAttrValue("index", ""),
		Item:       es.SelectAttrValue("item", ""),
		Open:       es.SelectAttrValue("open", ""),
		Separator:  es.SelectAttrValue("separator", ""),
	}
}
func NewNodeInclude(es *etree.Element) *NodeSQl {
	child := PareChild(es.Child)
	return &NodeSQl{
		Child: child,
		Id:    es.SelectAttrValue("refid", ""),
	}
}
func NewNodeText(str string) *NodeText {
	var reg, _ = regexp.Compile("[ |\n|\t]+")
	return &NodeText{Text: reg.ReplaceAllString(str, " ")}
}

type NodeRoot struct {
	ID    string
	Tag   string
	Child []Node
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

func (n *NodeRoot) Pare(m map[string]interface{}) (s string, err error) {
	newM := m
	newM["conv"] = map[string]string{}
	nodes, err := PareNodes(newM, n.Child)
	return strings.TrimSpace(nodes), err
}

type NodeIf struct {
	Child []Node
	Test  string
}

func (n *NodeIf) pareIF(m map[string]interface{}) (bool, error) {
	parseBool, err := strconv.ParseBool(n.Test)
	if err == nil {
		return parseBool, nil
	}
	if strings.Contains(n.Test, "=") {
		split := strings.Split(n.Test, "=")
		if len(split) != 2 {
			return false, errors.New("bad if " + n.Test)
		}
		str0, err := FindStr(m, split[0])
		if err != nil {
			return false, err
		}
		str1, err := FindStr(m, split[1])
		if err != nil {
			return false, err
		}
		return fmt.Sprintf("%v", str0) == fmt.Sprintf("%v", str1), nil
	} else {
		str, err := FindStr(m, n.Test)
		if err != nil {
			return false, nil
		}
		if b, ok := str.(bool); ok {
			return b, nil
		} else {
			return false, errors.New("bad if test not bool")
		}
	}
}
func FindStr(m map[string]interface{}, str string) (interface{}, error) {
	s := str
	var objs = []string{s}
	var apps = []string{""}
	count := strings.Count(s, ".")
	for i := 0; i < count; i++ {
		objs = append(objs, s[:strings.LastIndex(s, ".")])
		apps = append(objs, str[strings.LastIndex(s, ".")+1:])
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
		return inter, nil
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

func (n *NodeIf) Pare(m map[string]interface{}) (s string, err error) {
	pareIF, err := n.pareIF(m)
	if err != nil {
		return s, err
	}
	if pareIF {
		return PareNodes(m, n.Child)
	} else {
		return "", nil
	}
}

type NodeForEach struct {
	Child      []Node
	Item       string
	Index      string
	Collection string
	Open       string
	Separator  string
	Close      string
}

func (n *NodeForEach) Pare(m map[string]interface{}) (s string, err error) {
	coll := m[n.Collection]
	nodes := ""
	if coll == nil {
		return n.Open + nodes + n.Close, nil
	}
	v := reflect.ValueOf(coll)
	for v.Kind() == reflect.Ptr {
		v = v.Elem()
	}
	conv := m["conv"].(map[string]string)
	switch v.Kind() {
	case reflect.Slice:
		for i := 0; i < v.Cap(); i++ {
			conv[n.Index] = fmt.Sprint(i)
			m[n.Index] = i
			conv[n.Item] = fmt.Sprintf("#{%s.%d}", n.Collection, i)
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
		delete(conv, n.Index)
		delete(conv, n.Item)
	}()
	return n.Open + nodes[:len(nodes)-len(n.Separator)] + n.Close, nil
}

type NodeSQl struct {
	Child []Node
	Id    string
}

func (n *NodeSQl) Pare(m map[string]interface{}) (s string, err error) {
	return PareNodes(m, n.Child)
}

type NodeText struct {
	Text string
}

func (n *NodeText) Pare(m map[string]interface{}) (s string, err error) {
	conv := m["conv"].(map[string]string)
	expand := str.Expand('#', n.Text, func(s string) string {
		s2, ok := conv[s]
		if ok {
			return s2
		}
		return "#{" + s + "}"
	})
	return expand, nil
}
