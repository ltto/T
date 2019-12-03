package ref

import (
	"database/sql/driver"
	"fmt"
	"reflect"
)

func BreakData(ptr interface{}, tag, gap string) map[string]interface{} {
	m := make(map[string]interface{})
	BreakDataVal(ptr, m, tag, gap)
	return m
}

func BreakDataVal(ptr interface{}, m map[string]interface{}, tag, gap string) {
	val := reflect.ValueOf(ptr)
	for val.Kind() == reflect.Ptr {
		val = val.Elem()
	}
	typ := val.Type()

	if typ.Kind() == reflect.Map {
		if tag != "" {
			tag = tag + "."
		}
		iter := val.MapKeys()
		for _, k := range iter {
			v := val.MapIndex(k)
			BreakDataVal(v.Interface(), m, fmt.Sprintf("%s%s%s", tag, gap, k.String()), gap)
		}
	} else if typ.Kind() == reflect.Slice {
		for i := 0; i < val.Cap(); i++ {
			BreakDataVal(val.Index(i).Interface(), m, fmt.Sprintf("%s[%d]", tag, i), gap)
		}
	} else if IsBaseTime(typ) {
		m[tag] = breakProperType(typ, val)
	} else if IsStruct(typ.Kind()) {
		if tag != "" {
			tag = tag + "."
		}
		for i := 0; i < typ.NumField(); i++ {
			structField := val.Field(i)
			for structField.Kind() == reflect.Ptr {
				structField = structField.Elem()
			}
			typeField := typ.Field(i)
			if !structField.CanInterface() {
				continue
			}

			structFieldType := structField.Type()
			inputFieldName := tag + typeField.Name

			if valuer, ok := bindValuer(structField); ok {
				value, _ := valuer.Value()
				m[inputFieldName] = value
			} else if obj, ok := breakField(structField.Type(), structField); ok {
				m[inputFieldName] = obj
			} else if structFieldType.Kind() == reflect.Struct && !IsTime(structFieldType) {
				BreakDataVal(structField.Addr().Interface(), m, inputFieldName, gap)
			} else if structFieldType.Kind() == reflect.Slice {
				for j := 0; j < structField.Cap(); j++ {
					BreakDataVal(structField.Index(j).Interface(), m, fmt.Sprintf("%s[%d]", inputFieldName, j), gap)
				}
			} else if structFieldType.Kind() == reflect.Map {
				keys := structField.MapKeys()
				for _, key := range keys {
					v := structField.MapIndex(key)
					BreakDataVal(v.Interface(), m, fmt.Sprintf("%s%s%s", inputFieldName, gap, key.String()), gap)
				}
			} else if IsBaseTime(structFieldType) {
				m[inputFieldName] = breakProperType(structField.Type(), structField)
			}
		}
	}
}

func breakProperType(T reflect.Type, structField reflect.Value) interface{} {
	if obj, ok := breakField(T, structField); ok {
		return obj
	}
	switch T.Kind() {
	case reflect.Ptr:
		return breakProperType(structField.Elem().Type(), structField.Elem())
	default:
		return structField.Interface()
	}
}

func breakField(T reflect.Type, field reflect.Value) (interface{}, bool) {
	switch T.Kind() {
	case reflect.Ptr:
		return breakFieldPtr(field)
	default:
		return breakFieldNonPtr(field)
	}
}

func bindValuer(field reflect.Value) (driver.Valuer, bool) {
	ptr := reflect.New(field.Type())
	if ptr.CanInterface() {
		iface := ptr.Interface()
		if valuer, ok := iface.(driver.Valuer); ok {
			return valuer, true
		}
	}
	return nil, false
}

func breakFieldNonPtr(field reflect.Value) (interface{}, bool) {
	if unmarshaler, ok := bindValuer(field); ok {
		value, err := unmarshaler.Value()
		if err != nil {
			return value, true
		}
	}

	return nil, false
}

func breakFieldPtr(field reflect.Value) (interface{}, bool) {
	if field.IsNil() {
		// Initialize the pointer to a nil value
		field.Set(reflect.New(field.Type().Elem()))
	}
	return breakFieldNonPtr(field.Elem())
}
