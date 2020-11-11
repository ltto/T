package mybatis

import (
	"errors"
	"fmt"
	"github.com/ltto/T/gobox/ref"
	"github.com/ltto/T/mybatis/node"
	"reflect"
	"strings"
)

func makeFunc(n *node.DMLRoot, ft reflect.Type, tagStr string, db func() SqlCmd, conf *LoadConf) (val reflect.Value, err error) {
	if ft == nil || ft.Kind() != reflect.Func {
		return val, errors.New("你看看你传的参数是个啥")
	}
	mappings := strings.Split(tagStr, ",")
	if len(mappings) == 1 && mappings[0] == "" {
		mappings = mappings[0:0]
	}
	if len(mappings) != ft.NumIn() {
		return val, errors.New(fmt.Sprintf("func params len(%v) but fund(%d)", mappings, ft.NumIn()))
	}
	return reflect.MakeFunc(ft, func(args []reflect.Value) (returns []reflect.Value) {
		var (
			result *SQLResult
			err    error
		)
		defer func() {
			//if i := recover(); i != nil {
			//	returns = bindReturn(ft, returnValue, errors.New(fmt.Sprint("recover():", i, "\r\n", string(debug.Stack()))))
			//}
		}()
		sqlExc, err := PareSQL(pareArgs(args, mappings), n)
		if err != nil {
			return bindReturn(ft, nil, err)
		}

		if result, err = sqlExc.ExecSQL(db()); err != nil {
			return bindReturn(ft, result, err)
		}
		//if n.UseGeneratedKeys {
		//	switch sqlExc.Operate {
		//	case node.INSERT:
		//	case node.UPDATE:
		//	}
		//	var tag = "json"
		//	if conf != nil && conf.Tag != "" {
		//		tag = conf.Tag
		//	}
		//	if err := bindInKey(args, val.Data(), tag); err != nil {
		//		return bindReturn(ft, result, err)
		//	}
		//}
		return bindReturn(ft, result, err)
	}), nil
}
func pareArgs(args []reflect.Value, mappings []string) (m map[string]interface{}) {
	m = make(map[string]interface{})
	for i := range args {
		mapping := mappings[i]
		m[mapping] = args[i].Interface()
		ref.BreakDataVal(args[i].Interface(), m, mapping, ".")
	}
	return
}