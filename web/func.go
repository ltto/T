package web

import (
	"reflect"
)

type Func struct {
	inLen        int
	in           *reflect.Type
	inIn         int
	ctx          int
	ginCtx       bool
	out          *reflect.Type
	InterfaceMap InterfaceMap
	outLen       int
	outErr       bool
	v            reflect.Value
}

func (f Func) Call(ctx *Context) (interface{}, error) {
	var ins = make([]reflect.Value, f.inLen)
	if f.ctx >= 0 {
		if f.ginCtx {
			ins[f.ctx] = reflect.ValueOf(ctx.Context)
		} else {
			ins[f.ctx] = reflect.ValueOf(ctx)
		}
	}
	if f.inIn >= 0 {
		newIn := reflect.New(*f.in)
		err := ctx.Bind(newIn.Interface())
		if err != nil {
			return nil, err
		}
		ins[f.inIn] = newIn.Elem()
	}
	out := f.v.Call(ins)
	if len(out) == 2 {
		return out[0].Interface(), out[1].Interface().(error)
	}
	if len(out) == 1 {
		return out[0].Interface(), nil
	}

	return nil, nil
}

func NewFunc(f interface{}) *Func {
	check(f)
	fff := &Func{ctx: -1, inIn: -1}
	fv := reflect.ValueOf(f)
	ft := fv.Type()
	for i := 0; i < ft.NumIn(); i++ {
		in := ft.In(i)
		for in.Kind() == reflect.Ptr {
			in = in.Elem()
		}
		if in.String() == "Context" {
			fff.ctx = i
		} else if in.String() == "gin.Context" {
			fff.ctx = i
			fff.ginCtx = true
		} else {
			fff.inIn = i
			fff.in = &in
		}
	}
	for i := 0; i < ft.NumOut(); i++ {
		out := ft.Out(i)
		if ft.Out(i).PkgPath() == "" && ft.Out(i).String() == "error" {
			fff.outErr = true
		} else {
			fff.out = &out
		}
	}
	fff.v = reflect.ValueOf(f)
	fff.inLen = ft.NumIn()
	fff.outLen = ft.NumOut()
	return fff
}

func check(f interface{}) {
	fv := reflect.ValueOf(f)
	ft := fv.Type()
	if fv.Kind() != reflect.Func || ft.NumIn() > 2 || ft.NumOut() > 2 {
		panic("need func in<=2 out<=2")
	}
	for i := 0; i < ft.NumIn(); i++ {
		in := ft.In(i)
		if in.Kind() == reflect.Ptr {
			in = in.Elem()
		}
		if in.Kind() != reflect.Struct && in.String() != "echo.Context" {
			panic("param need struct or echo.Context")
		}
	}
	if ft.NumOut() == 2 {
		if !(ft.Out(1).PkgPath() == "" && ft.Out(1).String() == "error") {
			panic("如果2个返回值 最后一个必须为error")
		}
	}
}
