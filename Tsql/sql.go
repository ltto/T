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

func (q QueryResult) DecodePtr(ptr reflect.Value, outTag string) error {
	if q.IsBlank() {
		return nil
	}
	v := ptr
	var ptred bool
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}
	if outTag != "" {
		if len(q.Data) != 1 {
			i := len(q.Data)
			return errors.New(fmt.Sprintf("need 1 fund %d", i))
		}
		vals := q.Data[0][outTag]
		vals[0].BindData(ptr, ptred)
	}
	ptrV := v.Addr().Interface()
	if ptrV == nil {
		return nil
	}
	if v.Kind() == reflect.Slice {
		slice := reflect.ValueOf(ptrV).Elem()
		for i := range q.Data {
			newV := reflect.New(v.Type().Elem())
			if err := ref.BindDataVal(newV.Interface(), q.Data[i], "sql", false); err != nil {
				return err
			}
			slice = reflect.Append(slice, newV.Elem())
		}
		reflect.ValueOf(ptrV).Elem().Set(slice)
	} else if len(q.Data) == 1 {
		return ref.BindDataVal(ptrV, q.Data[0], "sql", false)
	}
	return nil
}
