package ref

import (
	"fmt"
	"reflect"
)

func NewSetV(tag reflect.Type, v interface{}) reflect.Value {
	tagV := reflect.New(tag)
	srcT, deep := srcType(tag, 0)
	switch deep {
	case 0:
		if v != nil {
			if tagV.Elem().CanSet() {
				tagV.Elem().Set(reflect.ValueOf(v))
			}
		}
		return tagV.Elem()
	case 1:
		srcV := reflect.New(srcT)
		if v != nil {
			if srcV.Elem().CanSet() {
				srcV.Elem().Set(reflect.ValueOf(v))
			}
		}
		if tagV.Elem().CanSet() {
			tagV.Elem().Set(srcV)
		}
		return tagV.Elem()
	default:
		s, _ := fmt.Printf("Unsupported SQLType %v", tag)
		panic(s)
	}
}
func SrcTypeVal(data interface{}) (interface{}, int) {
	return srcTypeVal(data, 0)
}
func srcTypeVal(data interface{}, deep int) (interface{}, int) {
	v := reflect.ValueOf(data)
	t := reflect.TypeOf(data)
	if data == nil || t.Kind() == reflect.Invalid {
		return data, deep
	}
	if t.Kind() == reflect.Ptr {
		if v.Elem().Kind() != reflect.Invalid && v.Elem().CanInterface() {
			i := v.Elem().Interface()
			return srcTypeVal(i, deep+1)
		} else {
			return nil, deep
		}
	}
	return data, deep
}

func SrcType(t interface{}) (reflect.Type, int) {
	return srcType(reflect.TypeOf(t), 0)
}
func SrcTypeT(t reflect.Type) (reflect.Type, int) {
	return srcType(t, 0)
}

func SrcKind(t interface{}) (reflect.Kind, int) {
	Type, i := srcType(reflect.TypeOf(t), 0)
	return Type.Kind(), i
}
func SrcKindT(t reflect.Type) (reflect.Kind, int) {
	Type, i := srcType(t, 0)
	return Type.Kind(), i
}

func srcType(t reflect.Type, deep int) (reflect.Type, int) {
	if t.Kind() == reflect.Ptr {
		return srcType(t.Elem(), deep+1)
	} else {
		return t, deep
	}
}

func GetScrElem(v reflect.Value, deep int) reflect.Value {
	if deep == 0 {
		return v
	}
	var v1 reflect.Value
	for i := 0; i < deep; i++ {
		v1 = v.Elem()
	}
	return v1
}

func FullName(t reflect.Type) string {
	if t.PkgPath() == "" {
		return t.String()
	}
	return t.PkgPath() + "_" + t.String()
}
func PrtType(t reflect.Type) reflect.Type {
	for t.Kind() == reflect.Ptr {
		t = t.Elem()
	}
	return t
}
func PrtValue(val reflect.Value) reflect.Value {
	for val.Kind() == reflect.Ptr {
		val = val.Elem()
	}
	return val
}
func Fields(t reflect.Type) (list []reflect.StructField) {
	t = PrtType(t)
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		if field.Anonymous {
			list = append(list, Fields(field.Type)...)
		} else {
			list = append(list, field)
		}
	}
	return
}
