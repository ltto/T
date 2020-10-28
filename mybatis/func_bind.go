package mybatis

import (
	"database/sql"
	"github.com/ltto/T/gobox/ref"
	"reflect"
	"strings"
)

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
			split := strings.Split(fieldT.Tag.Get(tag), ";")
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
