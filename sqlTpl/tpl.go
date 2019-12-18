package sqlTpl

import (
	"bytes"
	"database/sql"
	"fmt"
	"regexp"
	"text/template"

	"github.com/ltto/T/Tsql"
	"github.com/ltto/T/gobox/ref"
	"github.com/ltto/T/gobox/str"
)

type SqlTplFile struct {
	t      *template.Template
	TplMap map[string]*SqlTpl
}

type SqlTpl struct {
	t        *template.Template
	funcName string
	db       *sql.DB
}


func NewSqlTpl(t *template.Template, funcName string, DB *sql.DB) *SqlTpl {
	return &SqlTpl{t: t, funcName: funcName, db: DB}
}
func (s SqlTpl) Query(tx *sql.Tx, sqlStr string, param ...interface{}) (Tsql.QueryResult, error) {
	var (
		rows *sql.Rows
		err  error
	)
	if tx == nil {
		rows, err = s.db.Query(sqlStr, param...)
	} else {
		rows, err = tx.Query(sqlStr, param...)
	}
	if err != nil {
		return Tsql.QueryResult{}, err
	}
	return rows2maps(rows)
}
func (s SqlTpl) Exec(tx *sql.Tx, sqlStr string, param ...interface{}) (Tsql.QueryResult, error) {
	var (
		err    error
		result sql.Result
	)
	rt := Tsql.QueryResult{}
	if tx == nil {
		result, err = s.db.Exec(sqlStr, param...)
	} else {
		result, err = tx.Exec(sqlStr, param...)
	}
	if err != nil {
		return rt, err
	}
	id, _ := result.LastInsertId()
	affected, _ := result.RowsAffected()
	rt.Data = make([]map[string][]ref.Val, 1)
	rt.Data[0] = make(map[string][]ref.Val)
	rt.Data[0]["sql.insert"] = []ref.Val{ref.NewVal(id)}
	rt.Data[0]["sql.update"] = []ref.Val{ref.NewVal(affected)}
	return rt, nil
}

func (s SqlTpl) ExecSQL(data map[string]interface{}, tx *sql.Tx) (Tsql.QueryResult, error) {
	SQL, SQLParams, err := s.ParseSQL(data)
	var result Tsql.QueryResult
	fmt.Println("ExecSQL:", SQL, )
	fmt.Println("SQLParams", SQLParams)
	if Operate(SQL) == SELECT {
		if result, err = s.Query(tx, SQL, SQLParams...); err != nil {
			return result, err
		}
	} else {
		if result, err = s.Exec(tx, SQL, SQLParams...); err != nil {
			return result, err
		}
	}
	fmt.Println("SQL RESULT", result)
	return result, nil
}

func (s SqlTpl) ParseSQL(data map[string]interface{}) (string, []interface{}, error) {
	buffer := bytes.NewBuffer([]byte{})
	if err := s.t.Execute(buffer, data); err != nil {
		return "", nil, err
	}
	param := make([]interface{}, 0, 16)
	compile, _ := regexp.Compile("[ |\r|\n]+")
	parse := buffer.Bytes()
	parse = compile.ReplaceAll(parse, []byte(" "))
	fmt.Println("ParseSQL:", string(parse))
	sql := str.Expand('#', string(parse), func(s string) string {
		param = append(param, data[s])
		return "?"
	})
	return sql, param, nil
}
