package mybatis

import (
	"errors"
	"fmt"
	"reflect"
	"strings"

	"github.com/ltto/T/Tsql"
	"github.com/ltto/T/gobox/ref"
)

func (n *DMLRoot) BindFunc(ptr interface{}, tx SqlCmd) (err error) {
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
		if field.Kind() != reflect.Func ||
			structField.Name != n.ID {
			continue
		}
		tag := structField.Tag.Get("mapperParams")
		makeFunc, err := n.makeFunc(field.Type(), tx, tag, )
		if err != nil {
			return err
		}
		field.Set(makeFunc)
	}
	return nil
}

func (n *DMLRoot) makeFunc(ft reflect.Type, tx SqlCmd, tagStr string) (val reflect.Value, err error) {
	if ft == nil || ft.Kind() != reflect.Func {
		return val, errors.New("你看看你传的参数是个啥")
	}
	if ft.NumOut() > 2 {
		return val, errors.New("return need 1 or 2")
	}
	if ft.Out(ft.NumOut()-1).String() != "error" {
		return val, errors.New("last return need `error`")
	}
	tags := strings.Split(tagStr, ",")
	if len(tags) == 1 && tags[0] == "" {
		tags = tags[0:0]
	}
	if len(tags) != ft.NumIn() {
		return val, errors.New(fmt.Sprintf("func need %v but fund %d in", tags, ft.NumIn()))
	}
	return reflect.MakeFunc(ft, func(args []reflect.Value) (results []reflect.Value) {
		var (
			result      Tsql.QueryResult
			err         error
			returnValue reflect.Value
		)
		results = make([]reflect.Value, ft.NumOut())

		m := make(map[string]interface{})
		for i := range args {
			m[tags[i]] = args[i].Interface()
			ref.BreakDataVal(args[i].Interface(), m, tags[i], ".")
		}
		sqlExc, err := PareSQL(m, n)
		if err != nil {
			bindReturn(ft, results, returnValue, err)
			return
		}
		result, err = sqlExc.ExecSQL(tx)
		if err != nil {
			bindReturn(ft, results, returnValue, err)
			return
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
		bindReturn(ft, results, returnValue, err)
		return results
	}), nil
}

func bindReturn(ft reflect.Type, results []reflect.Value, result reflect.Value, e error) {
	size := len(results)
	switch size {
	case 1:
		results[0] = reflect.New(ft.Out(0)).Elem()
		if e != nil {
			results[0].Elem().Set(reflect.ValueOf(e))
			results[0] = results[0].Elem()
		}
	case 2:
		results[0] = result
		results[1] = reflect.New(ft.Out(1)).Elem()
		if e != nil {
			results[1].Elem().Set(reflect.ValueOf(e))
			results[1] = results[1].Elem()
		} else {
			results[1] = reflect.Zero(ft.Out(1))
		}
	}
}
