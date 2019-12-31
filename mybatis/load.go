package mybatis

import (
	"errors"
	"regexp"

	"github.com/beevik/etree"
)

type DML struct {
	Namespace string
	Cmd       map[string]*DMLRoot
}

func Load(XMLPath string) (dml DML, err error) {
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
	dml.Cmd = make(map[string]*DMLRoot, len(m["insert"])+len(m["select"])+len(m["update"])+len(m["delete"]))
	for i := range m["insert"] {
		id := m["insert"][i].SelectAttrValue("id", "")
		if dml.Cmd[id] != nil {
			return dml, errors.New("重复的ID:" + id)
		}
		dml.Cmd[id] = NewNodeRoot(m["insert"][i], &sqlTag)
	}
	for i := range m["select"] {
		id := m["select"][i].SelectAttrValue("id", "")
		if dml.Cmd[id] != nil {
			return dml, errors.New("重复的ID:" + id)
		}
		dml.Cmd[id] = NewNodeRoot(m["select"][i], &sqlTag)
	}
	for i := range m["update"] {
		id := m["update"][i].SelectAttrValue("id", "")
		if dml.Cmd[id] != nil {
			return dml, errors.New("重复的ID:" + id)
		}
		dml.Cmd[id] = NewNodeRoot(m["update"][i], &sqlTag)
	}
	for i := range m["delete"] {
		id := m["delete"][i].SelectAttrValue("id", "")
		if dml.Cmd[id] != nil {
			return dml, errors.New("重复的ID:" + id)
		}
		dml.Cmd[id] = NewNodeRoot(m["delete"][i], &sqlTag)
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
