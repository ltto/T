package ref

import (
	"fmt"
	"reflect"
	"strconv"
	"time"
)

type SqlType string

func (s SqlType) NewPoint() interface{} {
	switch s {
	//数值类型
	case "TINYINT",       //8(bit)
		"SMALLINT",       //16(bit)
		"MEDIUMINT",      //34(bit)
		"INT", "INTEGER": //32(bit)
		var i int
		return &i
	case "BIGINT": //64(bit)
		var i int
		return &i
	case "FLOAT": //32(bit)
		var i float32
		return &i
	case "DOUBLE", "DECIMAL":
		var i float64
		return &i
	//字符串类型
	case "CHAR",      //0-255(byte)
		"VARCHAR",    //0-65535(byte)
		"TINYBLOB",   //0-255(byte)
		"TINYTEXT",   //0-255(byte)
		"BLOB",       //0-65535(byte)
		"TEXT",       //0-65535(byte)
		"MEDIUMBLOB", //0-16777215(byte)
		"MEDIUMTEXT", //0-16777215(byte)
		"LONGBLOB",   //0-4294967295(byte)
		"LONGTEXT":   //0-4294967295(byte)
		var i string
		return &i
	//日期和时间类型
	case "DATE",     //YYYY-MM-DD
		"TIME",      //HH:MM:SS
		"YEAR",      //YYYY
		"DATETIME",  //YYYY-MM-DD HH:MM:SS
		"TIMESTAMP": //YYYY-MM-DD HH:MM:SS
		var i time.Time
		return &i
	}
	return nil
}

type Val struct {
	src  interface{}
	data interface{}
	nil  bool
	Deep int
}

func NewVal(src interface{}) Val {
	typeVal, deep := SrcTypeVal(src)
	val := Val{src: src, data: typeVal, Deep: deep, nil: typeVal == nil}
	return val
}
func (s Val) IsNil() bool {
	return s.nil
}
func (s Val) Interface() interface{} {
	return s.src
}

func (s Val) Data() interface{} {
	return s.data
}
func (s Val) Bool() bool {
	p := s.BoolP()
	if p == nil {
		return false
	}
	return *p
}
func (s Val) BoolP() *bool {
	var d *bool
	switch reflect.TypeOf(s.data).Kind() {
	case reflect.Int:
		b := int64(s.data.(int)) > 0
		d = &b
	case reflect.Int8:
		b := int64(s.data.(int8)) > 0
		d = &b
	case reflect.Int16:
		b := int64(s.data.(int16)) > 0
		d = &b
	case reflect.Int32:
		b := int64(s.data.(int32)) > 0
		d = &b
	case reflect.Int64:
		b := int64(s.data.(int64)) > 0
		d = &b
	case reflect.Uint:
		b := int64(s.data.(uint)) > 0
		d = &b
	case reflect.Uint8:
		b := int64(s.data.(uint8)) > 0
		d = &b
	case reflect.Uint16:
		b := int64(s.data.(uint16)) > 0
		d = &b
	case reflect.Uint32:
		b := int64(s.data.(uint32)) > 0
		d = &b
	case reflect.Uint64:
		b := int64(s.data.(uint64)) > 0
		d = &b
	case reflect.Float32:
		b := int64(s.data.(float32)) > 0
		d = &b
	case reflect.Float64:
		b := int64(s.data.(float64)) > 0
		d = &b
	case reflect.String:
		b, _ := strconv.ParseBool(s.String())
		d = &b
	}
	return d
}

func (s Val) Uint64() uint64 {
	p := s.Uint64P()
	if p == nil {
		return 0
	}
	return *p
}
func (s Val) Uint64P() *uint64 {
	var d *uint64
	T := reflect.TypeOf(s.data)
	switch T.Kind() {
	case reflect.Int:
		u := uint64(s.data.(int))
		d = &u
	case reflect.Int8:
		u := uint64(s.data.(int8))
		d = &u
	case reflect.Int16:
		u := uint64(s.data.(int16))
		d = &u
	case reflect.Int32:
		u := uint64(s.data.(int32))
		d = &u
	case reflect.Int64:
		u := uint64(s.data.(int64))
		d = &u
	case reflect.Uint:
		u := uint64(s.data.(uint))
		d = &u
	case reflect.Uint8:
		u := uint64(s.data.(uint8))
		d = &u
	case reflect.Uint16:
		u := uint64(s.data.(uint16))
		d = &u
	case reflect.Uint32:
		u := uint64(s.data.(uint32))
		d = &u
	case reflect.Uint64:
		u := uint64(s.data.(uint64))
		d = &u
	case reflect.Float32:
		u := uint64(s.data.(float32))
		d = &u
	case reflect.Float64:
		u := uint64(s.data.(float64))
		d = &u
	case reflect.String:
		i, _ := strconv.ParseInt(s.data.(string), 10, 64)
		u := uint64(i)
		d = &u
	default:
		if T.Kind() == reflect.Struct && T.String() == "time.Time" {
			u := uint64((s.data.(time.Time)).Unix())
			d = &u
		}
	}
	return d
}

func (s Val) Int64() int64 {
	p := s.Int64P()
	if p == nil {
		return 0
	}
	return *p
}
func (s Val) Int64P() *int64 {
	var d *int64
	T := reflect.TypeOf(s.data)
	switch T.Kind() {
	case reflect.Int:
		i := int64(s.data.(int))
		d = &i
	case reflect.Int8:
		i := int64(s.data.(int8))
		d = &i
	case reflect.Int16:
		i := int64(s.data.(int16))
		d = &i
	case reflect.Int32:
		i := int64(s.data.(int32))
		d = &i
	case reflect.Int64:
		i := int64(s.data.(int64))
		d = &i
	case reflect.Uint:
		i := int64(s.data.(uint))
		d = &i
	case reflect.Uint8:
		i := int64(s.data.(uint8))
		d = &i
	case reflect.Uint16:
		i := int64(s.data.(uint16))
		d = &i
	case reflect.Uint32:
		i := int64(s.data.(uint32))
		d = &i
	case reflect.Uint64:
		i := int64(s.data.(uint64))
		d = &i
	case reflect.Float32:
		i := int64(s.data.(float32))
		d = &i
	case reflect.Float64:
		i := int64(s.data.(float64))
		d = &i
	case reflect.String:
		i, _ := strconv.ParseInt(s.data.(string), 10, 64)
		d = &i
	default:
		if T.Kind() == reflect.Struct && T.String() == "time.Time" {
			unix := (s.data.(time.Time)).Unix()
			d = &unix
		}
	}
	return d
}

func (s Val) Float64() float64 {
	p := s.Float64P()
	if p == nil {
		return 0
	}
	return *p
}
func (s Val) Float64P() *float64 {
	var d *float64
	T := reflect.TypeOf(s.data)
	switch T.Kind() {
	case reflect.Int:
		f := float64(s.data.(int))
		d = &f
	case reflect.Int8:
		f := float64(s.data.(int8))
		d = &f
	case reflect.Int16:
		f := float64(s.data.(int16))
		d = &f
	case reflect.Int32:
		f := float64(s.data.(int32))
		d = &f
	case reflect.Int64:
		f := float64(s.data.(int64))
		d = &f
	case reflect.Uint:
		f := float64(s.data.(uint))
		d = &f
	case reflect.Uint8:
		f := float64(s.data.(uint8))
		d = &f
	case reflect.Uint16:
		f := float64(s.data.(uint16))
		d = &f
	case reflect.Uint32:
		f := float64(s.data.(uint32))
		d = &f
	case reflect.Uint64:
		f := float64(s.data.(uint64))
		d = &f
	case reflect.Float32:
		f := float64(s.data.(float32))
		d = &f
	case reflect.Float64:
		f := float64(s.data.(float64))
		d = &f
	case reflect.String:
		f, _ := strconv.ParseFloat(s.data.(string), 10)
		d = &f
	default:
		if T.Kind() == reflect.Struct && T.String() == "time.Time" {
			f := float64((s.data.(time.Time)).Unix())
			d = &f
		}
	}
	return d
}

func (s Val) Time() time.Time {
	p := s.TimeP()
	if p == nil {
		return time.Time{}
	}
	return *p
}
func (s Val) TimeP() *time.Time {
	var d *time.Time
	T := reflect.TypeOf(s.data)
	switch T.Kind() {
	case reflect.Int:
		unix := time.Unix(int64(s.data.(int)), 0)
		d = &unix
	case reflect.Int8:
		unix := time.Unix(int64(s.data.(int8)), 0)
		d = &unix
	case reflect.Int16:
		unix := time.Unix(int64(s.data.(int16)), 0)
		d = &unix
	case reflect.Int32:
		unix := time.Unix(int64(s.data.(int32)), 0)
		d = &unix
	case reflect.Int64:
		unix := time.Unix(int64(s.data.(int64)), 0)
		d = &unix
	case reflect.Uint:
		unix := time.Unix(int64(s.data.(uint)), 0)
		d = &unix
	case reflect.Uint8:
		unix := time.Unix(int64(s.data.(uint8)), 0)
		d = &unix
	case reflect.Uint16:
		unix := time.Unix(int64(s.data.(uint16)), 0)
		d = &unix
	case reflect.Uint32:
		unix := time.Unix(int64(s.data.(uint32)), 0)
		d = &unix
	case reflect.Uint64:
		unix := time.Unix(int64(s.data.(uint64)), 0)
		d = &unix
	case reflect.Float32:
		unix := time.Unix(int64(s.data.(float32)), 0)
		d = &unix
	case reflect.Float64:
		unix := time.Unix(int64(s.data.(float64)), 0)
		d = &unix
	case reflect.String:
		i, _ := strconv.ParseInt(s.data.(string), 10, 64)
		unix := time.Unix(int64(i), 0)
		d = &unix
	default:
		if T.Kind() == reflect.Struct && T.String() == "time.Time" {
			i := s.data.(time.Time)
			d = &i
		}
	}
	return d
}
func (s Val) StringP() *string {
	if s.IsNil() {
		return nil
	}
	str := s.String()
	return &str
}

func (s Val) String() string {
	if str, ok := s.data.(string); ok {
		return str
	} else if str, ok := s.data.([]uint8); ok {
		return string(str)
	}
	return fmt.Sprintf("%v", s.data)
}

func (s Val) BindBool(v reflect.Value, ptr bool) {
	data := s.BoolP()
	if data == nil {
		return
	}
	if ptr {
		v.Set(reflect.ValueOf(data))
	} else {
		v.SetBool(*data)
	}
}

func (s Val) BindInt64(v reflect.Value, ptr bool) {
	data := s.Int64P()
	if data == nil {
		return
	}
	if ptr {
		v := NewSetV(v.Type(), nil)
		v.Elem().SetInt(*data)
		v.Set(v)
	} else {
		v.SetInt(*data)
	}
}

func (s Val) BindFloat64(v reflect.Value, ptr bool) {
	data := s.Float64P()
	if data == nil {
		return
	}
	if ptr {
		v := NewSetV(v.Type(), nil)
		v.Elem().SetFloat(*data)
		v.Set(v)
	} else {
		v.SetFloat(*data)
	}
}

func (s Val) BindString(v reflect.Value, ptr bool) {
	data := s.String()
	if ptr {
		v.Set(reflect.ValueOf(data))
	} else {
		v.SetString(data)
	}
}

func (s Val) BindTime(v reflect.Value, ptr bool) {
	data := s.TimeP()
	if data == nil {
		return
	}
	if ptr {
		v.Set(reflect.ValueOf(data))
	} else {
		v.Set(reflect.ValueOf(*data))
	}
}
func (s Val) BindData(v reflect.Value, ptr bool) {
	if !ptr {
		v.Set(reflect.ValueOf(s.Data()))
	}
}
