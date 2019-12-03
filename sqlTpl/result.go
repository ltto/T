package sqlTpl

import (
	"reflect"

	"github.com/ltto/T/gobox/ref"
)

type QueryResult struct {
	data []map[string][]ref.Val
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

func (q QueryResult) DecodePtr(ptr interface{}) error {
	if q.IsBlank() || ptr == nil {
		return nil
	}
	T, _ := ref.SrcType(ptr)
	if T.Kind() == reflect.Slice {
		slice := reflect.ValueOf(ptr).Elem()
		for i := range q.data {
			newV := reflect.New(T.Elem())
			if err := ref.BindDataVal(newV.Interface(), q.data[i], "json", false); err != nil {
				return err
			}
			slice = reflect.Append(slice, newV.Elem())
		}
		reflect.ValueOf(ptr).Elem().Set(slice)
	} else if len(q.data) == 1 {
		return ref.BindDataVal(ptr, q.data[0], "json", false)
	}
	return nil
}
