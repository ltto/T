package sqlTpl

import (
	"bytes"
	"database/sql"
	"fmt"
	"regexp"
	"strings"
	"text/template"

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

func IsQuery(keyStr string) bool {
	b := strings.HasPrefix(keyStr, "Get") || strings.HasPrefix(keyStr, "List") || strings.HasPrefix(keyStr, "Query") || strings.HasPrefix(keyStr, "Select")
	return b
}

func NewSqlTpl(t *template.Template, funcName string, DB *sql.DB) *SqlTpl {
	return &SqlTpl{t: t, funcName: funcName, db: DB}
}
func (s SqlTpl) Query(tx *sql.Tx, sqlStr string, param ...interface{}) (QueryResult, error) {
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
		return QueryResult{}, err
	}
	return rows2maps(rows)
}
func (s SqlTpl) Exec(tx *sql.Tx, sqlStr string, param ...interface{}) (QueryResult, error) {
	var (
		err    error
		result sql.Result
	)
	rt := QueryResult{}
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
	rt.data = make([]map[string][]ref.Val, 1)
	rt.data[0] = make(map[string][]ref.Val)
	rt.data[0]["sql.insert"] = []ref.Val{ref.NewVal(id)}
	rt.data[0]["sql.update"] = []ref.Val{ref.NewVal(affected)}
	return rt, nil
}

func (s SqlTpl) ExecSQL(data map[string]interface{}, tx *sql.Tx) (QueryResult, error) {
	SQL, SQLParams, err := s.ParseSQL(data)
	var result QueryResult
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
