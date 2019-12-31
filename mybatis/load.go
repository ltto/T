package mybatis

import (
	"fmt"
	"regexp"

	"github.com/beevik/etree"
	"github.com/ltto/T/mybatis/node"
)

type DML struct {
	Namespace string
	Insert    map[string]*node.DMLRoot
	Select    map[string]*node.DMLRoot
	Update    map[string]*node.DMLRoot
	Delete    map[string]*node.DMLRoot
}

func MainLoad() {
	load := Load("/Users/ltt/go/src/github.com/ltto/T/mybatis/AlbumsMapper.xml")
	sql, _ := PareSQL(map[string]interface{}{}, load.Insert["Save"])
	fmt.Println(sql)
}

func Load(XMLPath string) (dml DML) {
	doc := etree.NewDocument()
	if err := doc.ReadFromFile(XMLPath); err != nil {
		panic(err)
	}
	element := doc.SelectElement("mapper")
	m := findSQLTPL(element, element.ChildElements())

	sqlTag := make(map[string]*etree.Element, len(m["sql"]))
	for i, v := range m["sql"] {
		id := v.SelectAttrValue("id", "")
		sqlTag[id] = m["sql"][i]
	}
	dml.Insert = make(map[string]*node.DMLRoot, len(m["insert"]))
	for i := range m["insert"] {
		dml.Insert[m["insert"][i].SelectAttrValue("id", "")] = node.NewNodeRoot(m["insert"][i], &sqlTag)
	}
	dml.Select = make(map[string]*node.DMLRoot, len(m["select"]))
	for i := range m["select"] {
		dml.Select[m["select"][i].SelectAttrValue("id", "")] = node.NewNodeRoot(m["select"][i], &sqlTag)
	}
	dml.Update = make(map[string]*node.DMLRoot, len(m["update"]))
	for i := range m["update"] {
		dml.Update[m["update"][i].SelectAttrValue("id", "")] = node.NewNodeRoot(m["update"][i], &sqlTag)
	}
	dml.Delete = make(map[string]*node.DMLRoot, len(m["delete"]))
	for i := range m["d"] {
		dml.Delete[m["delete"][i].SelectAttrValue("id", "")] = node.NewNodeRoot(m["delete"][i], &sqlTag)
	}
	return
}

func findSQLTPL(root *etree.Element, list []*etree.Element) map[string][]*etree.Element {
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
