package mybatis

import (
	"errors"
	"regexp"

	"github.com/beevik/etree"
	"github.com/ltto/T/tp"
)

type DML struct {
	e   *Engine
	Cmd map[string]*DMLRoot
}

func (e *Engine) LoadAndBindMap(m tp.H) (err error) {
	for k, v := range m {
		if err = e.LoadAndBind(k, v); err != nil {
			return err
		}
	}
	return err
}

func (e *Engine) LoadAndBind(XMLPath string, ptr interface{}) (err error) {
	load, err := e.Load(XMLPath)
	if err != nil {
		return err
	}
	return load.BindPtr(ptr)
}
func (e *Engine) Load(XMLPath string) (dml DML, err error) {
	doc := etree.NewDocument()
	if err = doc.ReadFromFile(XMLPath); err != nil {
		return
	}
	element := doc.SelectElement("mapper")
	m := findSQLTPL(element, element.ChildElements())

	sqlTag := make(map[string]*etree.Element, len(m["sql"]))
	for i, v := range m["sql"] {
		id := v.SelectAttrValue("id", "")
		sqlTag[id] = m["sql"][i]
	}

	if err = loadDml(m, "insert", &dml, sqlTag); err != nil {
		return
	}
	if err = loadDml(m, "select", &dml, sqlTag); err != nil {
		return
	}
	if err = loadDml(m, "update", &dml, sqlTag); err != nil {
		return
	}
	if err = loadDml(m, "delete", &dml, sqlTag); err != nil {
		return
	}
	dml.e = e
	e.DmlM[XMLPath] = &dml
	return
}

func loadDml(m map[string][]*etree.Element, key string, dml *DML, sqlTag map[string]*etree.Element) (err error) {
	elements := m[key]
	for i := range elements {
		id := elements[i].SelectAttrValue("id", "")
		if dml.Cmd[id] != nil {
			return errors.New("重复的ID:" + id)
		}
		if dml.Cmd == nil {
			dml.Cmd = make(map[string]*DMLRoot, 0)
		}
		dml.Cmd[id] = NewNodeRoot(elements[i], &sqlTag)
	}
	return nil
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
