package sqlTpl

import (
	"database/sql"
	"fmt"
	"reflect"
	"strings"

	"github.com/ltto/T/Tsql"
	"github.com/ltto/T/gobox/ref"
)

type TplError error

func (t TplEngine) makeFunc(typ reflect.Type, tpl *SqlTpl, tag, outTag string) reflect.Value {
	if tpl == nil {
		panic("tpl is nil")
	}
	if typ.Kind() != reflect.Func {
		panic("reflect: call of MakeFunc with non-Func type")
	}
	numOut := typ.NumOut()
	if numOut < 1 && numOut > 2 {
		panic("NumOut 1 or 2 ")
	}
	if typ.Out(numOut-1).String() != "error" {
		panic("last Out need error ")
	}
	tags := strings.Split(tag, ",")
	if len(tags) == 1 && tags[0] == "" {
		tags = tags[0:0]
	}
	if len(tags) != typ.NumIn() {
		panic(fmt.Sprintf("func need %v but fund %d in", tags, typ.NumIn()))
	}

	txi := -1
	for i := 0; i < typ.NumIn(); i++ {
		if typ.In(i).String() == "*sql.Tx" {
			txi = i
			break
		}
	}

	return reflect.MakeFunc(typ, func(in []reflect.Value) []reflect.Value {
		var (
			result      Tsql.QueryResult
			err         error
			tx          *sql.Tx
			returnValue reflect.Value
		)
		if txi > 0 {
			tx = in[txi].Interface().(*sql.Tx)
		}
		out := make([]reflect.Value, numOut)

		m := make(map[string]interface{})
		for i := range in {
			m[tags[i]] = in[i].Interface()
			ref.BreakDataVal(in[i].Interface(), m, tags[i], ".")
		}
		result, err = tpl.ExecSQL(m, tx)

		if numOut == 2 {
			returnValue = reflect.New(typ.Out(0))
			if err == nil {
				outT := typ.Out(0)
				if outT.Kind() == reflect.Ptr {
					outT = outT.Elem()
				}
				switch (outT).Kind() {
				case reflect.Map, reflect.Interface:
					returnValue.Elem().Set(reflect.MakeMap(typ.Out(0)))
				case reflect.Slice:
					returnValue.Elem().Set(reflect.MakeSlice(typ.Out(0), 0, 0))
				}
				err = result.DecodePtr(returnValue, outTag)
			}
			returnValue = returnValue.Elem()
		}
		bindReturn(typ, out, returnValue, err)
		return out
	})
}

func bindReturn(typ reflect.Type, out []reflect.Value, result reflect.Value, e error) {
	size := len(out)
	switch size {
	case 1:
		if e != nil {
			out[0] = reflect.New(typ.Out(0))
			out[0].Elem().Set(reflect.ValueOf(e))
			out[0] = out[0].Elem()
		}
	case 2:
		out[0] = result
		if e != nil {
			out[1] = reflect.New(typ.Out(1))
			out[1].Elem().Set(reflect.ValueOf(e))
			out[1] = out[1].Elem()
		} else {
			out[1] = reflect.Zero(typ.Out(1))
		}
	}
}
