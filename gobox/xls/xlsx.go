package xls

import (
	"errors"
	"fmt"
	"github.com/tealeg/xlsx"
	"reflect"
	"strconv"
	"time"
)

var xErr = errors.New("need slice struct")
var nilErr = errors.New("is nil")

type Marshaler interface {
	MarshalExcel() interface{}
}
type FieldHandler func(interface{}) interface{}

func Marshal(i interface{}, heads []string, field map[string]FieldHandler) (f *xlsx.File, err error) {
	if i == nil {
		return nil, nilErr
	}
	v := getV(reflect.ValueOf(i))
	t := v.Type()
	switch t.Kind() {
	case reflect.Slice:
		elemT := t.Elem()
		if elemT.Kind() != reflect.Struct {
			return nil, xErr
		}
		if v.Len() == 0 {
			return f, nil
		}
		f = xlsx.NewFile()
		sheet, _ := f.AddSheet("sheet1")
		for i := 0; i < v.Len(); i++ {
			if i == 0 {
				addHaed(heads, sheet)
			}
			addRow(sheet, heads, v.Index(i), field)
		}
		return f, nil
	case reflect.Struct:
		f := xlsx.NewFile()
		sheet, _ := f.AddSheet("sheet1")
		addHaed(heads, sheet)
		addRow(sheet, heads, v, field)
		return f, nil
	default:
		return nil, xErr
	}
}

func addRow(sheet *xlsx.Sheet, heads []string, v reflect.Value, field map[string]FieldHandler) {
	row := sheet.AddRow()
	for i, head := range heads {
		if head == "" || head == "-" {
			continue
		}
		cell := row.AddCell()
		handler, ok := field[head]
		if ok {
			val := handler(v.Field(i).Interface())
			setValue(cell, val)
		}
		setValue(cell, v.Field(i).Interface())
	}
}

func addHaed(heads []string, sheet *xlsx.Sheet) {
	row := sheet.AddRow()
	for _, head := range heads {
		if head == "" || head == "-" {
			continue
		}
		cell := row.AddCell()
		cell.SetString(head)
	}
}

func getV(v reflect.Value) reflect.Value {
	for reflect.Ptr == v.Type().Kind() {
		v = v.Elem()
	}
	return v
}
func GetHeads(i interface{}) (m []string) {
	if i == nil {
		return nil
	}
	v := getV(reflect.ValueOf(i))
	t := v.Type()
	fields := t.NumField()
	m = make([]string, fields)
	for i := 0; i < fields; i++ {
		name := t.Field(i).Tag.Get("excel")
		m[i] = name
	}
	return m
}

// SetInt sets a cell's value to an integer.
func setValue(c *xlsx.Cell, n interface{}) {
	marshaler, ok := n.(Marshaler)
	if ok {
		n = marshaler.MarshalExcel()
	}
	switch t := n.(type) {
	case time.Time:
		c.SetDateTime(t)
		return
	case int, int8, int16, int32, int64:
		c.SetNumeric(fmt.Sprintf("%d", n))
	case float64:
		// When formatting floats, do not use fmt.Sprintf("%v", n), this will cause numbers below 1e-4 to be printed in
		// scientific notation. Scientific notation is not a valid way to store numbers in XML.
		// Also not not use fmt.Sprintf("%f", n), this will cause numbers to be stored as X.XXXXXX. Which means that
		// numbers will lose precision and numbers with fewer significant digits such as 0 will be stored as 0.000000
		// which causes tests to fail.
		c.SetNumeric(strconv.FormatFloat(t, 'f', -1, 64))
	case float32:
		c.SetNumeric(strconv.FormatFloat(float64(t), 'f', -1, 32))
	case string:
		c.SetString(t)
	case []byte:
		c.SetString(string(t))
	case nil:
		c.SetString("")
	default:
		c.SetString(fmt.Sprintf("%v", n))
	}
}
