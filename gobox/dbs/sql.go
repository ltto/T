package dbs

import (
	"database/sql"
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"time"

	"github.com/jinzhu/gorm"
)

func QueryMap(db *gorm.DB, sqlStr string, param ...interface{}) ([]map[string]interface{}, error) {
	return QueryMapAlas(db, nil, sqlStr, param...)
}

type Bind func(string) interface{}

func QueryMapAlas(db *gorm.DB, colAlas map[string]string, sqlStr string, param ...interface{}) ([]map[string]interface{}, error) {
	var (
		query *sql.Rows
		err   error
	)
	if query, err = db.Raw(sqlStr, param...).Rows(); err != nil {
		return nil, err
	}
	maps, err := rows2maps(query, colAlas)
	if err != nil {
		return nil, err
	}
	return maps.Maps(), nil
}

func rows2maps(rows *sql.Rows, colAlas map[string]string) (resultsSlice QueryResult, err error) {
	for i := 0; rows.Next(); i++ {
		result, err := row2map(rows, colAlas)
		if err != nil {
			return QueryResult{}, err
		}
		resultsSlice.Append(result)
	}
	return resultsSlice, nil
}
func row2map(rows *sql.Rows, colAlas map[string]string) (resultsMap map[string]Val, err error) {
	var fields []*sql.ColumnType
	if fields, err = rows.ColumnTypes(); err != nil {
		return nil, err
	}
	var scanResultContainers = make([]interface{}, len(fields))
	for i := 0; i < len(fields); i++ {
		var scanResultContainer interface{}
		scanResultContainers[i] = &scanResultContainer
	}
	if err := rows.Scan(scanResultContainers...); err != nil {
		return nil, err
	}
	var result = make(map[string]Val)
	for ii, key := range fields {
		vv, _ := toNumber(key.DatabaseTypeName(), scanResultContainers[ii])
		var fieldName = key.Name()
		if alas, ok := colAlas[key.Name()]; ok {
			fieldName = alas
		}
		if fieldName != "*exclude" {
			result[fieldName] = NewVal(vv)
		}
	}
	return result, nil
}

func toNumber(sqlType string, val interface{}) (interface{}, bool) {
	var s string
	var ok bool
	var b []uint8
	if s, ok = val.(string); !ok {
		if b, ok = val.([]uint8); !ok {
			return val, false
		} else {
			s = string(b)
		}
	}
	switch strings.ToUpper(sqlType) {
	case "TINYINT",
		"SMALLINT",
		"MEDIUMINT",
		"INT",
		"INTEGER",
		"BIGINT":
		i, _ := strconv.Atoi(s)
		return i, true
	case "FLOAT",
		"DOUBLE",
		"DECIMAL":
		f, _ := strconv.ParseFloat(s, 64)
		return f, true
	}
	return val, false
}

type QueryResult struct {
	Data []map[string]Val
}

func (q QueryResult) Maps() (m []map[string]interface{}) {
	m = make([]map[string]interface{}, len(q.Data))
	for i, datum := range q.Data {
		m[i] = make(map[string]interface{}, len(datum))
		for k := range datum {
			m[i][k] = datum[k].data
		}
	}
	return
}

func (q *QueryResult) Append(cell map[string]Val) {
	q.Data = append(q.Data, cell)
}

type Val struct {
	src  interface{}
	data interface{}
	nil  bool
}

func NewVal(src interface{}) Val {
	typeVal, _ := srcTypeVal(src, 0)
	if b, ok := typeVal.([]uint8); ok {
		typeVal = string(b)
	}
	val := Val{src: src, data: typeVal, nil: typeVal == nil}
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
		if s.IsTime() {
			i := s.data.(time.Time)
			d = &i
		}
	}
	return d
}

func (s Val) IsTime() bool {
	T := reflect.TypeOf(s.data)
	return T.Kind() == reflect.Struct && T.String() == "time.Time"
}
func (s Val) StringP() *string {
	if s.IsNil() {
		return nil
	}
	str := s.String()
	return &str
}

func (s Val) MarshalJSON() ([]byte, error) {
	if s.IsNil() {
		return []byte("null"), nil
	}
	if s.IsTime() {
		return []byte(fmt.Sprintf("%d", s.Uint64())), nil
	}
	s2 := s.String()
	return []byte(s2), nil
}
func (s Val) String() string {
	if s.IsNil() {
		return "null"
	}
	if str, ok := s.data.(string); ok {
		return fmt.Sprintf("\"%s\"", str)
	} else if str, ok := s.data.([]uint8); ok {
		return fmt.Sprintf("\"%s\"", string(str))
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
func srcType(t reflect.Type, deep int) (reflect.Type, int) {
	if t.Kind() == reflect.Ptr {
		return srcType(t.Elem(), deep+1)
	} else {
		return t, deep
	}
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
