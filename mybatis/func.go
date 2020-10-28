package mybatis

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/ltto/T/Tsql"
	"github.com/ltto/T/gobox/ref"
	"github.com/ltto/T/mybatis/node"
	"reflect"
	"strings"
)

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

func makeFunc(n *node.DMLRoot, ft reflect.Type, tagStr string, db func() SqlCmd, conf *LoadConf) (val reflect.Value, err error) {
	if ft == nil || ft.Kind() != reflect.Func {
		return val, errors.New("你看看你传的参数是个啥")
	}
	if ft.NumOut() > 2 {
		return val, errors.New("the max return is 2")
	}
	if ft.Out(ft.NumOut()-1).String() != "error" {
		return val, errors.New("the last return must's `error`")
	}
	tags := strings.Split(tagStr, ",")
	if len(tags) == 1 && tags[0] == "" {
		tags = tags[0:0]
	}
	if len(tags) != ft.NumIn() {
		return val, errors.New(fmt.Sprintf("func params len(%v) but fund(%d)", tags, ft.NumIn()))
	}
	return reflect.MakeFunc(ft, func(args []reflect.Value) []reflect.Value {
		var (
			result      Tsql.QueryResult
			err         error
			returnValue reflect.Value
		)

		m := make(map[string]interface{})
		for i := range args {
			m[tags[i]] = args[i].Interface()
			ref.BreakDataVal(args[i].Interface(), m, tags[i], ".")
		}
		sqlExc, err := PareSQL(m, n)
		if err != nil {
			return bindReturn(ft, returnValue, err)
		}

		if result, err = sqlExc.ExecSQL(db()); err != nil {
			return bindReturn(ft, returnValue, err)
		}
		if n.UseGeneratedKeys {
			kk := ""
			switch Operate(sqlExc.SQL) {
			case INSERT:
				kk = "sql.insert"
			case UPDATE:
				kk = "sql.update"
			}
			val := result.Data[0][kk][0]
			var tag = "json"
			if conf != nil && conf.Tag != "" {
				tag = conf.Tag
			}
			if err := bindKey(args, val.Data(), tag); err != nil {
				return bindReturn(ft, returnValue, err)
			}
		}
		if ft.NumOut() == 2 {
			returnValue = reflect.New(ft.Out(0))
			outT := ft.Out(0)
			for outT.Kind() == reflect.Ptr {
				outT = outT.Elem()
			}
			switch (outT).Kind() {
			case reflect.Map, reflect.Interface:
				returnValue.Elem().Set(reflect.MakeMap(ft.Out(0)))
			case reflect.Slice:
				returnValue.Elem().Set(reflect.MakeSlice(ft.Out(0), 0, 0))
			}
			err = result.DecodePtr(returnValue, "")
			returnValue = returnValue.Elem()
		}
		return bindReturn(ft, returnValue, err)
	}), nil
}

func bindReturn(ft reflect.Type, result reflect.Value, e error) (results []reflect.Value) {
	results = make([]reflect.Value, ft.NumOut())
	size := len(results)
	switch size {
	case 1:
		results[0] = bindReturnOne(ft.Out(0), result, e)
	case 2:
		results[0] = bindReturnOne(ft.Out(0), result, e)
		results[1] = bindReturnOne(ft.Out(1), result, e)
	}
	return
}
func bindReturnOne(fot reflect.Type, result reflect.Value, e error) (v reflect.Value) {
	if ref.IsError(fot) {
		return bindReturnErr(fot, e)
	} else {
		return bindReturnValue(fot, result)
	}
}

func bindReturnErr(t reflect.Type, e error) (v reflect.Value) {
	v = reflect.New(t).Elem()
	if e != nil {
		v.Set(reflect.ValueOf(e))
	} else {
		v = reflect.Zero(t)
	}
	return
}
func bindReturnValue(t reflect.Type, result reflect.Value) (v reflect.Value) {
	if !result.IsValid() {
		v = reflect.New(t).Elem()
		return
	}
	return result
}

func bindKey(args []reflect.Value, src interface{}, tag string) error {
	for i := range args {
		val := args[i]
		if val.Kind() != reflect.Ptr {
			return nil
		}
		for val.Kind() == reflect.Ptr {
			val = val.Elem()
		}
		typ := val.Type()
		if typ.Kind() != reflect.Struct {
			return nil
		}
		fields := ref.Fields(typ)
		for _, fieldT := range fields {
			structField := val.FieldByName(fieldT.Name)
			for structField.Kind() == reflect.Ptr {
				structField = structField.Elem()
			}
			if structField.Kind() == reflect.Invalid || !structField.CanInterface() {
				continue
			}
			split := strings.Split(fieldT.Tag.Get("json"), ";")
			key := false
			for i := range split {
				if key = split[i] == "primary_key"; key {
					break
				}
			}
			if key {
				scanner, ok := structField.Addr().Interface().(sql.Scanner)
				if ok {
					return scanner.Scan(src)
				}
				_, err := ref.UnmarshalField(structField.Kind(), ref.NewVal(src), structField)
				if err != nil {
					return err
				}
			}
		}

	}
	return nil
}
func IsValuer(v reflect.Value) bool {
	defer func() { recover() }()
	t := v.MethodByName("Value").Type()
	return t.NumOut() == 2 && t.Out(0).String() == "driver.Value" && t.Out(1).String() == "error"
}
func IsScanner(v reflect.Value) bool {
	defer func() { recover() }()
	t := v.MethodByName("Scan").Type()
	return t.NumOut() == 1 && t.Out(0).String() == "error" && t.NumIn() == 1 && t.In(0).String() == "interface"
}
