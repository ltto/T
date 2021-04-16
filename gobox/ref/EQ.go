package ref

import "reflect"

func IsPtr(kind reflect.Kind) bool {
	return kind == reflect.Ptr
}

func IsBool(kind reflect.Kind) bool {
	return kind == reflect.Bool
}
func IsInt(kind reflect.Kind) bool {
	return kind == reflect.Int || kind == reflect.Int8 || kind == reflect.Int16 || kind == reflect.Int32 || kind == reflect.Int64
}
func IsUint(kind reflect.Kind) bool {
	return kind == reflect.Uint || kind == reflect.Uint8 || kind == reflect.Uint16 || kind == reflect.Uint32 || kind == reflect.Uint64
}
func IsIntUint(kind reflect.Kind) bool {
	return IsUint(kind) || IsInt(kind)
}
func IsFloat(kind reflect.Kind) bool {
	return kind == reflect.Float32 || kind == reflect.Float64
}

func IsNumber(kind reflect.Kind) bool {
	return IsUint(kind) || IsInt(kind) || IsFloat(kind)
}

func IsString(kind reflect.Kind) bool {
	return kind == reflect.String
}

func IsBase(kind reflect.Kind) bool {
	return IsIntUint(kind) || IsBool(kind) || IsFloat(kind) || IsString(kind)
}

func IsTime(t reflect.Type) bool {
	return t.Kind() == reflect.Struct && t.String() == "time.Time"
}
func IsBaseTime(t reflect.Type) bool {
	return IsBase(t.Kind()) || IsTime(t)
}

func IsInterface(kind reflect.Kind) bool {
	return kind == reflect.Interface
}
func IsMap(kind reflect.Kind) bool {
	return kind == reflect.Map
}
func IsStruct(kind reflect.Kind) bool {
	return kind == reflect.Struct
}
func IsSlice(kind reflect.Kind) bool {
	return kind == reflect.Slice
}

func IsValuer(t reflect.Type) bool {
	if method, ok := t.MethodByName("Value"); ok {
		mt := method.Type
		return mt.NumIn() == 0 && mt.NumOut() == 2 && mt.Out(0).String() == "driver.Value" && mt.Out(1).String() == "error"
	}
	return false
}

var et = reflect.ValueOf(map[error]error{}).Type().Elem()

func IsError(t reflect.Type) bool {
	return t.AssignableTo(et)
}
