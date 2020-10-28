package mybatis

import (
	"errors"
	"github.com/ltto/T/mybatis/node"
	"path"
	"reflect"
	"regexp"

	"github.com/beevik/etree"
)

type DML struct {
	e   *Engine
	Cmd map[string]*node.DMLRoot
}

func (D *DML) BindPtr(ptr interface{}, conf *LoadConf) (err error) {
	v := reflect.ValueOf(ptr)
	for v.Kind() == reflect.Ptr {
		v = v.Elem()
	}
	if v.Kind() != reflect.Struct {
		return errors.New("need struct")
	}
	t := v.Type()
	for i := 0; i < v.NumField(); i++ {
		field := v.Field(i)
		structField := t.Field(i)
		if field.Kind() != reflect.Func {
			continue
		}
		dmlRoot := D.Cmd[structField.Name]
		tag := structField.Tag.Get("mapperParams")
		makeFunc, err := makeFunc(dmlRoot, field.Type(), tag, func() SqlCmd {
			return D.e.GetDB()
		}, conf)
		if err != nil {
			return err
		}
		field.Set(makeFunc)
	}
	return nil
}

type LoadConf struct {
	Tag        string
	PathPrefix string
}

func (e *Engine) LoadAndBindMap(conf *LoadConf, mapping map[string]interface{}) (err error) {
	for xmlStr, obj := range mapping {
		if conf != nil && conf.PathPrefix != "" {
			xmlStr = path.Join(conf.PathPrefix, xmlStr)
		}
		if err = e.LoadAndBind(xmlStr, obj, conf); err != nil {
			return err
		}
	}
	return err
}

func (e *Engine) LoadAndBind(XMLPath string, ptr interface{}, conf *LoadConf) (err error) {
	load, err := e.Load(XMLPath)
	if err != nil {
		return err
	}
	return load.BindPtr(ptr, conf)
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
	if err = loadDMLS(m, &dml, sqlTag); err != nil {
		return
	}
	dml.e = e
	e.DmlM[XMLPath] = &dml
	return
}

func loadDMLS(m map[string][]*etree.Element, dml *DML, sqlTag map[string]*etree.Element) (err error) {
	if err = loadDml(m, "insert", dml, sqlTag); err != nil {
		return
	}
	if err = loadDml(m, "select", dml, sqlTag); err != nil {
		return
	}
	if err = loadDml(m, "update", dml, sqlTag); err != nil {
		return
	}
	if err = loadDml(m, "delete", dml, sqlTag); err != nil {
		return
	}
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
			dml.Cmd = make(map[string]*node.DMLRoot, 0)
		}
		dml.Cmd[id] = NewNodeRoot(elements[i], sqlTag)
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
