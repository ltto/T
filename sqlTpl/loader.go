package sqlTpl

import (
	"database/sql"
	"errors"
	"io/ioutil"
	"path"
	"path/filepath"
	"reflect"
	"strings"
	"text/template"
	"unsafe"
)

func (t TplEngine) LoadMappers(ptr ...interface{}) {
	for i := range ptr {
		t.LoadMapper(ptr[i])
	}
}

func (t TplEngine) LoadMapper(ptr interface{}) {
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
		outTag := structField.Tag.Get("mapperOut")
		name := pt.Name()
		field.Set(t.makeFunc(field.Type(), t.m[name+".gohtml"].TplMap[structField.Name], tag, outTag))
	}
}

func (t *TplEngine) Scanner(dest string) {
	if t.scan {
		return
	}
	dir, err := ioutil.ReadDir(dest)
	if err != nil {
		panic(err)
	}
	m := make(map[string][]byte, len(dir))
	for _, info := range dir {
		if info.IsDir() || !strings.HasSuffix(info.Name(), ".gohtml") {
			continue
		}
		bytes, err := ioutil.ReadFile(path.Join(dest, info.Name()))
		if err != nil {
			panic(err)
		}
		m[info.Name()] = bytes
	}
	t.ScannerByBytes(m)
	t.scan = true
}

func (t *TplEngine) ScannerByBytes(b map[string][]byte) {
	if t.scan {
		return
	}
	t.TPLS = b
	t.m = make(map[string]SqlTplFile)
	for s, bytes := range b {
		fileTpl := LoadTpl(s, bytes, t.db)
		t.m[s] = fileTpl
	}
	t.scan = true
}

func LoadTpl(info string, bytes []byte, DB *sql.DB) SqlTplFile {
	filet := template.New(info)
	filet.Funcs(funcMap)
	files, err := filet.Parse(string(bytes))
	if err != nil {
		panic(err)
	}
	name := reflect.ValueOf(files).Elem().FieldByName("common")
	mapSet := name.Elem().FieldByName("tmpl")
	keys := mapSet.MapKeys()
	fileTpl := SqlTplFile{}
	tplMap := make(map[string]*SqlTpl)
	for _, key := range keys {
		pointer := mapSet.MapIndex(key).Pointer()
		tpl := (*template.Template)(unsafe.Pointer(pointer))
		tpl.Funcs(funcMap)
		keyStr := key.String()
		if filepath.Base(info) == keyStr {
			fileTpl.t = tpl
		} else {
			tplMap[keyStr] = NewSqlTpl(tpl, keyStr, DB)
		}
	}
	fileTpl.TplMap = tplMap
	return fileTpl
}
