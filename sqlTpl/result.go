package sqlTpl

import (
	"errors"
	"fmt"
	"reflect"

	"github.com/bndr/gotabulate"
	"github.com/ltto/T/gobox/ref"
)

type QueryResult struct {
	data []map[string][]ref.Val
}

func (q QueryResult) String() string {
	var td [][]interface{}
	var th []string
	for _, datum := range q.data {
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
	return q.data[index]
}

func (q QueryResult) Rows() int {
	return len(q.data)
}

func (q QueryResult) IsBlank() bool {
	return q.data == nil || len(q.data) == 0
}

func (q *QueryResult) append(cell map[string][]ref.Val) {
	q.data = append(q.data, cell)
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
		if len(q.data) != 1 {
			i := len(q.data)
			return errors.New(fmt.Sprintf("need 1 fund %d", i))
		}
		vals := q.data[0][outTag]
		vals[0].BindData(ptr, ptred)
	}
	ptrV := v.Addr().Interface()
	if ptrV == nil {
		return nil
	}
	if v.Kind() == reflect.Slice {
		slice := reflect.ValueOf(ptrV).Elem()
		for i := range q.data {
			newV := reflect.New(v.Type().Elem())
			if err := ref.BindDataVal(newV.Interface(), q.data[i], "sql", false); err != nil {
				return err
			}
			slice = reflect.Append(slice, newV.Elem())
		}
		reflect.ValueOf(ptrV).Elem().Set(slice)
	} else if len(q.data) == 1 {
		return ref.BindDataVal(ptrV, q.data[0], "sql", false)
	}
	return nil
}
