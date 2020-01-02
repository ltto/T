package mybatis

import (
	"errors"
	"regexp"

	"github.com/beevik/etree"
)

type DML struct {
	e   Engine
	Cmd map[string]*DMLRoot
}

func (e Engine) LoadAndBind(XMLPath string, ptr interface{}) (err error) {
	load, err := e.Load(XMLPath)
	if err != nil {
		return err
	}
	return load.BindPtr(ptr)
}
func (e Engine) Load(XMLPath string) (dml DML, err error) {
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
			err = errors.New("重复的ID:" + id)
			return
		}
		dml.Cmd[id] = NewNodeRoot(m["insert"][i], &sqlTag)
	}
	for i := range m["select"] {
		id := m["select"][i].SelectAttrValue("id", "")
		if dml.Cmd[id] != nil {
			err = errors.New("重复的ID:" + id)
			return
		}
		dml.Cmd[id] = NewNodeRoot(m["select"][i], &sqlTag)
	}
	for i := range m["update"] {
		id := m["update"][i].SelectAttrValue("id", "")
		if dml.Cmd[id] != nil {
			err = errors.New("重复的ID:" + id)
			return
		}
		dml.Cmd[id] = NewNodeRoot(m["update"][i], &sqlTag)
	}
	for i := range m["delete"] {
		id := m["delete"][i].SelectAttrValue("id", "")
		if dml.Cmd[id] != nil {
			err = errors.New("重复的ID:" + id)
			return
		}
		dml.Cmd[id] = NewNodeRoot(m["delete"][i], &sqlTag)
	}
	dml.e = e
	e.DmlM[XMLPath] = &dml
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
