package mybatis

import (
	"database/sql"
	"github.com/ltto/T/mybatis/node"
	"reflect"
	"strings"
)

type Field struct {
	DBName string
	Field  reflect.Value
	Struct reflect.StructField
}
type SQLResult struct {
	Method       node.SQLOperate
	LastInsertId int64
	RowsAffected int64
	Rows         *sql.Rows
}

func (s *SQLResult) scanOBJ(returnValue *reflect.Value, types ...reflect.Type) (err error) {
	var columns []string
	if columns, err = s.Rows.Columns(); err != nil {
		return err
	}
	switch types[0].Kind() {
	case reflect.Map:
		if rows := s.Rows; rows.Next() {
			if err := scanMap(columns, rows, *returnValue); err != nil {
				return err
			}
		}
	case reflect.Struct:
		if rows := s.Rows; rows.Next() {
			if err := scanStruct(columns, rows, *returnValue); err != nil {
				return err
			}
		}
	case reflect.Slice:
		rows := s.Rows
		value, seter := NewValue(types[1])
		for rows.Next() {
			switch seter.Elem().Kind() {
			case reflect.Interface:
				makeMap := reflect.MakeMap(reflect.TypeOf(map[string]interface{}{}))
				if err = scanMap(columns, rows, makeMap); err != nil {
					return
				}
				seter.Elem().Set(makeMap)
				*returnValue = reflect.Append(*returnValue, value.Elem())
			case reflect.Map:
				makeMap := reflect.MakeMap(seter.Elem().Type())
				if err = scanMap(columns, rows, makeMap); err != nil {
					return
				}
				seter.Elem().Set(makeMap)
				*returnValue = reflect.Append(*returnValue, value.Elem())
			case reflect.Struct:
				if err = scanStruct(columns, rows, seter.Elem()); err != nil {
					return
				}
				*returnValue = reflect.Append(*returnValue, value.Elem())
			}
		}
	}

	return nil
}

func scanMap(columns []string, rows *sql.Rows, v reflect.Value) error {
	var vs = make([]interface{}, len(columns))
	for i := range columns {
		var temp interface{}
		vs[i] = &temp
	}
	if err := rows.Scan(vs...); err != nil {
		return err
	}
	for i, column := range columns {
		if ip, ok := vs[i].(*interface{}); ok {
			vs[i] = *ip
		}
		if bs, ok := vs[i].([]byte); ok {
			vs[i] = string(bs)
		}
		if vs[i] == nil {
			v.SetMapIndex(reflect.ValueOf(column), reflect.ValueOf(&vs[i]).Elem())
		} else {
			v.SetMapIndex(reflect.ValueOf(column), reflect.ValueOf(vs[i]))
		}
	}
	return nil
}

func scanStruct(columns []string, rows *sql.Rows, v reflect.Value) error {
	fm := map[string]interface{}{}
	t := v.Type()
	for i := 0; i < v.NumField(); i++ {
		field := t.Field(i)
		value := v.Field(i)
		get := field.Tag.Get("json")
		split := strings.Split(get, ";")
		if field.Type.Kind() == reflect.Ptr {
			fm[split[0]] = value.Addr().Interface()
		} else {
			reflectValue := reflect.New(reflect.PtrTo(field.Type))
			reflectValue.Elem().Set(value.Addr())
			fm[split[0]] = reflectValue.Elem().Interface()
		}
	}
	var vs = make([]interface{}, len(columns))
	for i, column := range columns {
		if obj, ok := fm[column]; ok {
			vs[i] = obj
		}
	}

	if err := rows.Scan(vs...); err != nil {
		return err
	}
	return nil
}
