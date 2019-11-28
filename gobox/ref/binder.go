package ref

import (
	"database/sql"
	"encoding"
	"errors"
	"reflect"
	"strings"
)

type BindUnmarshaler interface {
	// UnmarshalParam decodes and assigns a value from an form or query param.
	UnmarshalParam(param string) error
}

func BindDataStr(ptr interface{}, data map[string][]string, tag string, noTagBind bool) error {
	return BindDataVal(ptr, Str2Val(data), tag, noTagBind)
}
func Str2Val(data map[string][]string) map[string][]Val {
	maps := make(map[string][]Val)
	for k, v := range data {
		var list = make([]Val, len(v))
		for i, str := range v {
			list[i] = NewVal(str)
		}
		maps[k] = list
	}
	return maps
}

func BindDataVal(ptr interface{}, data map[string][]Val, tag string, noTagBind bool) error {
	typ := reflect.TypeOf(ptr).Elem()
	val := reflect.ValueOf(ptr).Elem()
	if typ.Kind() == reflect.Ptr {
		typ = typ.Elem()
		newV := reflect.New(typ)
		val.Set(newV)
		val = val.Elem()
	}

	if typ.Kind() != reflect.Struct && !(typ.Kind() == reflect.Map && typ.Elem().Key().Kind() == reflect.String && typ.Elem().Kind() == reflect.Interface) {
		return errors.New("binding element must be a struct or map[string]interface")
	}

	if typ.Kind() == reflect.Map {
		if map_, ok := val.Interface().(map[string]interface{}); ok {
			for k := range data {
				if len(data[k]) == 1 {
					map_[k] = data[k][0]
				} else {
					map_[k] = data[k]
				}
			}
			return nil
		}
	}

	for i := 0; i < typ.NumField(); i++ {
		typeField := typ.Field(i)
		structField := val.Field(i)
		if !structField.CanSet() {
			continue
		}
		structFieldKind := structField.Kind()
		inputFieldName := typeField.Tag.Get(tag)

		if inputFieldName == "" && noTagBind {
			inputFieldName = typeField.Name
			// If tag is nil, we inspect if the field is a struct.
			if _, ok := bindUnmarshaler(structField); !ok && (structFieldKind == reflect.Struct) {
				if err := BindDataVal(structField.Addr().Interface(), data, tag, noTagBind); err != nil {
					return err
				}
				continue
			}
		}

		inputValue, exists := data[inputFieldName]
		if !exists {
			// Go json.Unmarshal supports case insensitive binding.  However the
			// url params are bound case sensitive which is inconsistent.  To
			// fix this we must check all of the map values in a
			// case-insensitive search.
			inputFieldName = strings.ToLower(inputFieldName)
			for k, v := range data {
				if strings.ToLower(k) == inputFieldName {
					inputValue = v
					exists = true
					break
				}
			}
		}

		if !exists {
			continue
		}

		// Call this first, in case we're dealing with an alias to an array type
		if ok, err := unmarshalField(typeField.Type.Kind(), inputValue[0], structField); ok {
			if err != nil {
				return err
			}
			continue
		}

		numElems := len(inputValue)
		if structFieldKind == reflect.Slice && numElems > 0 {
			sliceOf := structField.Type().Elem().Kind()
			slice := reflect.MakeSlice(structField.Type(), numElems, numElems)
			for j := 0; j < numElems; j++ {
				if err := setWithProperType(sliceOf, inputValue[j], slice.Index(j)); err != nil {
					return err
				}
			}
			val.Field(i).Set(slice)
		} else if structFieldKind == reflect.Map && numElems > 0 {
			mapOf := structField.Type().Elem().Kind()
			map_ := reflect.MakeMap(structField.Type())
			for j := 0; j < numElems; j++ {
				if err := setWithProperType(mapOf, inputValue[j], map_.MapIndex(reflect.ValueOf(j))); err != nil {
					return err
				}
			}
			val.Field(i).Set(map_)
		} else if err := setWithProperType(typeField.Type.Kind(), inputValue[0], structField); err != nil {
			return err
		}
	}
	return nil
}

func setWithProperType(valueKind reflect.Kind, val Val, structField reflect.Value) error {
	// But also call it here, in case we're dealing with an array of BindUnmarshalers
	if ok, err := unmarshalField(valueKind, val, structField); ok {
		return err
	}

	switch valueKind {
	case reflect.Ptr:
		return setWithProperType(structField.Elem().Kind(), val, structField.Elem())
	case reflect.Int:
		return setIntField(val, structField)
	case reflect.Int8:
		return setIntField(val, structField)
	case reflect.Int16:
		return setIntField(val, structField)
	case reflect.Int32:
		return setIntField(val, structField)
	case reflect.Int64:
		return setIntField(val, structField)
	case reflect.Uint:
		return setUintField(val, structField)
	case reflect.Uint8:
		return setUintField(val, structField)
	case reflect.Uint16:
		return setUintField(val, structField)
	case reflect.Uint32:
		return setUintField(val, structField)
	case reflect.Uint64:
		return setUintField(val, structField)
	case reflect.Bool:
		return setBoolField(val, structField)
	case reflect.Float32:
		return setFloatField(val, structField)
	case reflect.Float64:
		return setFloatField(val, structField)
	case reflect.String:
		structField.SetString(val.String())
	default:
		return errors.New("unknown type")
	}
	return nil
}

func unmarshalField(valueKind reflect.Kind, val Val, field reflect.Value) (bool, error) {
	switch valueKind {
	case reflect.Ptr:
		return unmarshalFieldPtr(val, field)
	default:
		return unmarshalFieldNonPtr(val, field)
	}
}

// bindUnmarshaler attempts to unmarshal a reflect.Value into a BindUnmarshaler
func bindUnmarshaler(field reflect.Value) (BindUnmarshaler, bool) {
	ptr := reflect.New(field.Type())
	if ptr.CanInterface() {
		iface := ptr.Interface()
		if unmarshaler, ok := iface.(BindUnmarshaler); ok {
			return unmarshaler, ok
		}
	}
	return nil, false
}

// textUnmarshaler attempts to unmarshal a reflect.Value into a TextUnmarshaler
func textUnmarshaler(field reflect.Value) (encoding.TextUnmarshaler, bool) {
	ptr := reflect.New(field.Type())
	if ptr.CanInterface() {
		iface := ptr.Interface()
		if unmarshaler, ok := iface.(encoding.TextUnmarshaler); ok {
			return unmarshaler, ok
		}
	}
	return nil, false
}

func sqlScanner(field reflect.Value) (sql.Scanner, bool) {
	ptr := reflect.New(field.Type())
	if ptr.CanInterface() {
		iface := ptr.Interface()
		if scanner, ok := iface.(sql.Scanner); ok {
			return scanner, ok
		}
	}
	return nil, false
}

func unmarshalFieldNonPtr(value Val, field reflect.Value) (bool, error) {
	if unmarshaler, ok := bindUnmarshaler(field); ok {
		err := unmarshaler.UnmarshalParam(value.String())
		field.Set(reflect.ValueOf(unmarshaler).Elem())
		return true, err
	}
	if sqlScanner, ok := sqlScanner(field); ok {
		err := sqlScanner.Scan(value.data)
		field.Set(reflect.ValueOf(sqlScanner).Elem())
		return true, err
	}
	if unmarshaler, ok := textUnmarshaler(field); ok {
		err := unmarshaler.UnmarshalText([]byte(value.String()))
		field.Set(reflect.ValueOf(unmarshaler).Elem())
		return true, err
	}

	return false, nil
}

func unmarshalFieldPtr(value Val, field reflect.Value) (bool, error) {
	if field.IsNil() {
		// Initialize the pointer to a nil value
		field.Set(reflect.New(field.Type().Elem()))
	}
	return unmarshalFieldNonPtr(value, field.Elem())
}

func setIntField(val Val, field reflect.Value) error {
	if val.IsNil() {
		field.SetInt(0)
	}
	field.SetInt(val.Int64())
	return nil
}

func setUintField(val Val, field reflect.Value) error {
	if val.IsNil() {
		field.SetUint(0)
	}
	field.SetUint(val.Uint64())
	return nil
}

func setBoolField(val Val, field reflect.Value) error {
	if val.IsNil() {
		field.SetBool(false)
	}
	field.SetBool(val.Bool())
	return nil
}

func setFloatField(val Val, field reflect.Value) error {
	if val.IsNil() {
		field.SetFloat(0.0)
	}
	field.SetFloat(val.Float64())
	return nil
}
