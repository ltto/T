package mybatis

import (
	"database/sql"
	"github.com/ltto/T/gobox/ref"
	"reflect"
	"strings"
)

func bindReturn(ft reflect.Type, result *SQLResult, e error) (results []reflect.Value) {
	results = make([]reflect.Value, ft.NumOut())
	var errIndex []int
	var objIndex []int
	//results[i] = bindReturnErr(out, e)
	//return bindReturnValue(out, result)
	for i := 0; i < ft.NumOut(); i++ {
		if out := ft.Out(i); ref.IsError(out) {
			errIndex = append(errIndex, i)
		} else {
			objIndex = append(objIndex, i)
		}
	}
	var bindErr error
	for _, index := range objIndex {
		out := ft.Out(index)
		if bindErr == nil {
			value, err := bindReturnValue(out, result)
			if err != nil {
				bindErr = err
				value = reflect.New(out).Elem()
			}
			results[index] = value
		} else {
			results[index] = reflect.New(out).Elem()
		}
	}
	if bindErr != nil {
		e = bindErr
	}
	for _, index := range errIndex {
		results[index] = bindReturnErr(ft.Out(index), e)
	}
	return
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
func NewValue(t reflect.Type) (returnValue, seter reflect.Value) {
	returnValue = reflect.New(t)
	seter = returnValue
	outT := t
	for outT.Kind() == reflect.Ptr {
		outT = outT.Elem()
		value := reflect.New(outT)
		seter.Elem().Set(value)
		seter = value
	}
	return
}
func bindReturnValue(t reflect.Type, result *SQLResult) (v reflect.Value, err error) {
	if result == nil {
		return reflect.Value{}, nil
	}
	returnValue, seter := NewValue(t)
	outT := seter.Elem().Type()
	switch (outT).Kind() {
	case reflect.Map: //必须是map[string]interface{}
		seter.Elem().Set(reflect.MakeMap(outT))
		if err = result.scanOBJ(&seter, outT); err != nil {
			return
		}
	case reflect.Interface:
		typeOf := reflect.TypeOf(map[string]interface{}{})
		makeMap := reflect.MakeMap(typeOf)
		if err = result.scanOBJ(&seter, typeOf); err != nil {
			return
		}
		seter.Elem().Set(makeMap)
	case reflect.Slice:
		slice := reflect.MakeSlice(outT, 0, 0)
		if err = result.scanOBJ(&slice, slice.Type(), outT.Elem()); err != nil {
			return
		}
		seter.Elem().Set(slice)
	case reflect.Struct:
		elem := reflect.New(outT).Elem()
		if err = result.scanOBJ(&elem, outT); err != nil {
			return
		}
		seter.Elem().Set(elem)
	}
	returnValue = returnValue.Elem()
	return returnValue, err
}

func bindInKey(args []reflect.Value, src interface{}, tag string) error {
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
