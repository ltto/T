package sqlTpl

import (
	"database/sql"
	"errors"
	"fmt"
	"reflect"
	"strings"

	"github.com/ltto/T/gobox/ref"
)

func (t TplEngine) makeFunc(typ reflect.Type, tpl *SqlTpl, tag string) reflect.Value {
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
			result QueryResult
			err    error
			tx     *sql.Tx
		)
		if txi > 0 {
			tx = in[txi].Interface().(*sql.Tx)
		}

		out := make([]reflect.Value, numOut)
		var errRT error = errors.New("")
		if numOut == 2 {
			out[0] = reflect.New(typ.Out(0)).Elem()
			out[1] = reflect.Zero(reflect.TypeOf(errRT))
		} else {
			out[0] = reflect.Zero(reflect.TypeOf(errRT))
		}
		m := make(map[string]interface{})
		for i := range in {
			m[tags[i]] = in[i].Interface()
			ref.BreakDataVal(in[i].Interface(), m, tags[i], ".")
		}

		if result, err = tpl.ExecSQL(m, tx); err != nil {
			out[numOut-1] = reflect.ValueOf(err)
			return out
		}
		if numOut == 2 {
			outT := typ.Out(0)
			if outT.Kind() == reflect.Ptr {
				outT = outT.Elem()
			}
			var returnValue *reflect.Value = nil
			//build return Type
			var returnV = reflect.New(typ.Out(0))
			switch (outT).Kind() {
			case reflect.Map, reflect.Interface:
				returnV.Elem().Set(reflect.MakeMap(typ.Out(0)))
			case reflect.Slice:
				returnV.Elem().Set(reflect.MakeSlice(typ.Out(0), 0, 0))
			}
			returnValue = &returnV
			if err = result.DecodePtr(returnValue.Interface()); err != nil {
				out[1] = reflect.ValueOf(err)
				return out
			}
			out[0] = returnValue.Elem()
		}
		return out
	})
}
