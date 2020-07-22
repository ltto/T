package xls

import (
	"errors"
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"time"

	"github.com/tealeg/xlsx"
)

var xErr = errors.New("need slice struct")
var nilErr = errors.New("is nil")

type Marshaler interface {
	MarshalExcel() interface{}
}
type Jsoner interface {
	MarshalJSONParam(map[string]interface{}) ([]byte, error)
}

type Heads struct {
	Name      string
	FieldName string
}

func getHeads(i interface{}, lang string) (m []Heads) {
	if i == nil {
		return nil
	}
	t := reflect.TypeOf(i)
	for t.Kind() == reflect.Ptr {
		t = t.Elem()
	}
	if t.Kind() == reflect.Slice {
		t = t.Elem()
	}
	fields := getFields(t)
	for _, f := range fields {
		name := f.Tag.Get("excel")
		if name == "" || name == "-" {
			continue
		}
		split := strings.Split(name, ",")
		heads := Heads{Name: f.Name, FieldName: f.Name}
		for _, ss := range split {
			sp := strings.Split(ss, ":")
			if len(sp) == 2 && strings.EqualFold(lang, sp[0]) {
				heads.Name = sp[1]
				break
			}
		}
		m = append(m, heads)
	}
	return m
}

func getFields(t reflect.Type) (list []reflect.StructField) {
	fields := t.NumField()
	for i := 0; i < fields; i++ {
		f := t.Field(i)
		if f.Anonymous {
			structFields := getFields(f.Type)
			list = append(list, structFields...)
		} else {
			list = append(list, f)
		}
	}
	return
}

func Marshal(ptr interface{}, lang string) (f *xlsx.File, err error) {
	if ptr == nil {
		return nil, nilErr
	}
	heads := getHeads(ptr, lang)
	v := getV(reflect.ValueOf(ptr))
	t := v.Type()
	switch t.Kind() {
	case reflect.Slice:
		elemT := t.Elem()
		if elemT.Kind() != reflect.Struct {
			return nil, xErr
		}
		f = xlsx.NewFile()
		sheet, _ := f.AddSheet("sheet1")
		addHaed(heads, sheet)
		for i := 0; i < v.Len(); i++ {
			addRow(sheet, heads, v.Index(i), lang)
		}
		return f, nil
	case reflect.Struct:
		f := xlsx.NewFile()
		sheet, _ := f.AddSheet("sheet1")
		addHaed(heads, sheet)
		addRow(sheet, heads, v, lang)
		return f, nil
	default:
		return nil, xErr
	}
}

func addRow(sheet *xlsx.Sheet, heads []Heads, v reflect.Value, lang string) {
	row := sheet.AddRow()
	for _, head := range heads {
		if head.Name == "" || head.Name == "-" {
			continue
		}
		cell := row.AddCell()
		setValue(cell, v.FieldByName(head.FieldName).Interface(), lang)
	}
}

func addHaed(heads []Heads, sheet *xlsx.Sheet) {
	row := sheet.AddRow()
	for _, head := range heads {
		if head.Name == "" || head.Name == "-" {
			continue
		}
		cell := row.AddCell()
		cell.SetString(head.Name)
	}
}

func getV(v reflect.Value) reflect.Value {
	for reflect.Ptr == v.Type().Kind() {
		v = v.Elem()
	}
	return v
}

// SetInt sets a cell's value to an integer.
func setValue(c *xlsx.Cell, n interface{}, lang string) {
	if marshaler, ok := n.(Marshaler); ok {
		n = marshaler.MarshalExcel()
	} else if j, ok := n.(Jsoner); ok {
		param, err := j.MarshalJSONParam(map[string]interface{}{"lang": lang})
		if err != nil {
			n = "err"
		} else {
			n = strings.Trim(string(param), "\"")
		}
	}

	switch t := n.(type) {
	case time.Time:
		c.SetDateTime(t)
		return
	case int, int8, int16, int32, int64:
		c.SetValue(fmt.Sprintf("%d", n))
	case float64:
		// When formatting floats, do not use fmt.Sprintf("%v", n), this will cause numbers below 1e-4 to be printed in
		// scientific notation. Scientific notation is not a valid way to store numbers in XML.
		// Also not not use fmt.Sprintf("%f", n), this will cause numbers to be stored as X.XXXXXX. Which means that
		// numbers will lose precision and numbers with fewer significant digits such as 0 will be stored as 0.000000
		// which causes tests to fail.
		c.SetValue(strconv.FormatFloat(t, 'f', -1, 64))
	case float32:
		c.SetValue(strconv.FormatFloat(float64(t), 'f', -1, 32))
	case string:
		c.SetString(t)
	case []byte:
		c.SetString(string(t))
	case nil:
		c.SetString("")
	default:
		v := fmt.Sprintf("%v", n)
		if v == "<nil>" {
			v = ""
		}
		c.SetString(v)
	}
}
