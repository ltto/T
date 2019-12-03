package sqlTpl

import (
	"errors"
	"io/ioutil"
	"path"
	"path/filepath"
	"reflect"
	"strings"
	"text/template"
	"unsafe"
)

func (t TplEngine) LocalMapper(ptr interface{}) {
	if !t.scan {
		panic(errors.New("no scanner files"))
	}
	pv := reflect.ValueOf(ptr)
	for pv.Kind() == reflect.Ptr {
		pv = pv.Elem()
	}
	if pv.Kind() != reflect.Struct {
		panic(errors.New("need struct"))
	}
	pt := pv.Type()
	for i := 0; i < pv.NumField(); i++ {
		field := pv.Field(i)
		structField := pt.Field(i)
		if field.Kind() != reflect.Func {
			continue
		}
		tag := structField.Tag.Get("mapperParams")
		field.Set(t.makeFunc(field.Type(), t.m[pt.Name()+".tpl"].TplMap[structField.Name], tag))
	}
}

func (t *TplEngine) Scanner(dest string) {
	if t.scan {
		return
	}
	t.m = make(map[string]SqlTplFile)
	dir, err := ioutil.ReadDir(dest)
	if err != nil {
		panic(err)
	}
	for _, info := range dir {
		if info.IsDir() || !strings.HasSuffix(info.Name(), ".tpl") {
			continue
		}
		filet := template.New(info.Name())
		filet.Funcs(funcMap)
		byte, err := ioutil.ReadFile(path.Join(dest, info.Name()))
		if err != nil {
			panic(err)
		}
		files, err := filet.Parse(string(byte))
		if err != nil {
			panic(err)
		}
		name := reflect.ValueOf(files).Elem().FieldByName("common")
		mapRange := name.Elem().FieldByName("tmpl").MapRange()
		fileTpl := SqlTplFile{}
		tplMap := make(map[string]*SqlTpl)
		for mapRange.Next() {
			key := mapRange.Key()
			pointer := mapRange.Value().Pointer()
			tpl := (*template.Template)(unsafe.Pointer(pointer))
			tpl.Funcs(funcMap)
			keyStr := key.String()
			if filepath.Base(info.Name()) == keyStr {
				fileTpl.t = tpl
			} else {
				tplMap[keyStr] = NewSqlTpl(tpl, keyStr, t.db)
			}
		}
		fileTpl.TplMap = tplMap
		t.m[info.Name()] = fileTpl
	}
	t.scan = true
}
