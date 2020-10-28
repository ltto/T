package mybatis

import (
	"errors"
	"fmt"
	"github.com/ltto/T/Tsql"
	"github.com/ltto/T/gobox/ref"
	"github.com/ltto/T/mybatis/node"
	"reflect"
	"strings"
)


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
			switch sqlExc.Operate {
			case node.INSERT:
				kk = "sql.insert"
			case node.UPDATE:
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
