package Tsql

import (
	"errors"
	"fmt"
	"reflect"

	"github.com/bndr/gotabulate"
	"github.com/ltto/T/gobox/ref"
)

type QueryResult struct {
	Data []map[string][]ref.Val
}

func (q QueryResult) String() string {
	var td [][]interface{}
	var th []string
	for _, datum := range q.Data {
		var row []interface{}
		for k := range datum {
			if len(datum) != len(th) {
				th = append(th, k)
			}
			row = append(row, datum[k])
		}
		td = append(td, row)
	}
	create := gotabulate.Create(td)
	create.SetHeaders(th)
	return "\n" + create.Render("simple")
}

func (q QueryResult) Index(index int) map[string][]ref.Val {
	return q.Data[index]
}

func (q QueryResult) Rows() int {
	return len(q.Data)
}

func (q QueryResult) IsBlank() bool {
	return q.Data == nil || len(q.Data) == 0
}

func (q *QueryResult) Append(cell map[string][]ref.Val) {
	q.Data = append(q.Data, cell)
}

func (q QueryResult) Scan(ptr interface{}, outTag string) error {
	err := q.DecodePtr(reflect.ValueOf(ptr), outTag)
	return err
}

func (q QueryResult) DecodePtr(ptr reflect.Value, outTag string) error {
	if q.IsBlank() {
		return nil
	}
	for ptr.Kind() == reflect.Ptr {
		ptr = ptr.Elem()
	}
	if outTag != "" {
		if len(q.Data) != 1 {
			i := len(q.Data)
			return errors.New(fmt.Sprintf("need 1 fund %d", i))
		}
		vals := q.Data[0][outTag]
		vals[0].BindData(ptr, false)
	}
	if ptr.Kind() == reflect.Slice {
		slice := ptr
		for i := range q.Data {
			newV := reflect.New(slice.Type().Elem())
			if err := ref.BindDataVal(newV.Interface(), q.Data[i], "json"); err != nil {
				return err
			}
			slice = reflect.Append(slice, newV.Elem())
		}
		ptr.Set(slice)
	} else if len(q.Data) >= 1 {
		return ref.BindDataVal(ptr.Interface(), q.Data[0], "json")
	}
	return nil
}
