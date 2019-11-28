package ginT

import (
	"path/filepath"
	"reflect"
	"strconv"
	"strings"

	"github.com/ltto/T/ginT/ctx"
)

type Func struct {
	name   string
	in     []Field
	out    []Field
	outErr bool
	v      *reflect.Value
}

func (f *Func) Call(ctx *ctx.Context) interface{} {
	var inv = make([]reflect.Value, len(f.in))
	if len(f.in) > 0 {
		if f.in[0].Ptr {
			inv[0] = f.InOne()
		} else {
			inv[0] = f.InOne().Elem()
		}
	}
	if len(f.in) == 2 {
		value := reflect.New(f.in[1].t)
		err := ctx.JsonBody(value.Interface())
		if err != nil {
			panic(err)
		}
		if f.in[1].Ptr {
			inv[1] = value
		} else {
			inv[1] = value.Elem()
		}
	}

	call := f.v.Call(inv)
	if len(f.out) == 2 {
		if err, ok := call[1].Interface().(error); ok && err != nil {
			return err
		}
	}
	if len(f.out) > 0 {
		return call[0].Interface()
	}
	return nil
}

func (f *Func) InOne() reflect.Value {
	structV := reflect.New(f.in[0].t)
	for i := 0; i < f.in[0].t.NumField(); i++ {
		field := f.in[0].t.Field(i)
		if path, pathOK := field.Tag.Lookup("path"); pathOK {
			fv := parseT(field.Type, path)
			structV.Field(i).Set(reflect.ValueOf(fv))
		} else if form, formOK := field.Tag.Lookup("form"); formOK {
			fv := parseT(field.Type, form)
			structV.Field(i).Set(reflect.ValueOf(fv))
		} else if query, queryOK := field.Tag.Lookup("query"); queryOK {
			fv := parseT(field.Type, query)
			structV.Field(i).Set(reflect.ValueOf(fv))
		} else if head, headOK := field.Tag.Lookup("head"); headOK {
			fv := parseT(field.Type, head)
			structV.Field(i).Set(reflect.ValueOf(fv))
		}
	}
	return structV
}

func parseT(t reflect.Type, v string) reflect.Value {
	if t.Kind() == reflect.Ptr {
		panic("滚")
	}
	value := reflect.New(t)
	switch t.Kind() {
	case reflect.Bool:
		b, _ := strconv.ParseBool(v)
		value.Elem().SetBool(b)
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		b, _ := strconv.ParseInt(v, 10, 64)
		value.Elem().SetInt(b)
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		b, _ := strconv.ParseUint(v, 10, 64)
		value.Elem().SetUint(b)
	case reflect.Float32, reflect.Float64:
		b, _ := strconv.ParseFloat(v, 10)
		value.Elem().SetFloat(b)
	case reflect.String:
		value.Elem().SetString(v)
	}
	return value.Elem()
}

func NewFunc(v reflect.Value) *Func {
	if v.Kind() != reflect.Func {
		panic("滚")
	}
	t := v.Type()
	f := Func{name: v.String(), v: &v}
	if t.NumIn() > 2 {
		panic("入参不能大于2个")
	}

	f.in = make([]Field, t.NumIn())
	for i := 0; i < t.NumIn(); i++ {
		in := t.In(i)
		if in.Kind() != reflect.Struct {
			panic("非Struct 不能入参")
		}
		f.in[i] = getField(in)
	}

	if t.NumOut() > 2 {
		panic("返回值不能超过2个")
	}
	if t.NumOut() == 2 {
		if !(t.Out(1).PkgPath() == "" && t.Out(1).String() == "error") {
			panic("如果2个返回值 最后一个必须为error")
		}
		f.outErr = true
	}

	f.out = make([]Field, t.NumOut())
	for i := 0; i < t.NumOut(); i++ {
		f.out[i] = getField(t.Out(i))
	}
	return &f
}

func getField(t reflect.Type) Field {
	field := Field{}
	ptr := t.Kind() == reflect.Ptr
	if ptr {
		t = t.Elem()
	}
	path := t.PkgPath()
	if path == "" {
		path = "builtin"
	}
	field.Pkg = path
	field.Type = t.String()
	field.Ptr = ptr
	field.t = t
	return field
}

type Field struct {
	Pkg  string
	Type string
	Ptr  bool
	t    reflect.Type
}

func (f *Field) String() string {
	pkg := f.Pkg
	if f.Ptr {
		pkg = "*" + pkg
	}

	if strings.Contains(f.Type, ".") {
		return pkg + filepath.Ext(f.Type)
	}
	return pkg + "." + f.Type
}
