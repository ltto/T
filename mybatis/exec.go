package mybatis

import (
	"database/sql"
	"fmt"

	"github.com/ltto/T/Tsql"
	"github.com/ltto/T/gobox/ref"
	"github.com/ltto/T/gobox/str"
)

type SqlCmd interface {
	Query(query string, args ...interface{}) (*sql.Rows, error)
	Exec(query string, args ...interface{}) (sql.Result, error)
}
type SqlExec struct {
	db     *sql.DB
	SQL    string
	params []interface{}
}

func PareSQL(m map[string]interface{}, root *DMLRoot) (ex SqlExec, err error) {
	pare, err := root.Pare(m)
	if err != nil {
		return ex, err
	}
	ex.SQL = str.Expand('#', pare, func(s string) string {
		ex.params = append(ex.params, m[s])
		return "?"
	})
	return
}

func (s SqlExec) Query(tx SqlCmd) (Tsql.QueryResult, error) {
	var (
		rows *sql.Rows
		err  error
	)
	if tx == nil {
		rows, err = s.db.Query(s.SQL, s.params...)
	} else {
		rows, err = tx.Query(s.SQL, s.params...)
	}
	if err != nil {
		return Tsql.QueryResult{}, err
	}
	return rows2maps(rows)
}
func (s SqlExec) Exec(tx SqlCmd) (Tsql.QueryResult, error) {
	var (
		err    error
		result sql.Result
	)
	rt := Tsql.QueryResult{}
	if tx == nil {
		result, err = s.db.Exec(s.SQL, s.params...)
	} else {
		result, err = tx.Exec(s.SQL, s.params...)
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

func (s SqlExec) ExecSQL(tx SqlCmd) (result Tsql.QueryResult, err error) {
	fmt.Println("SQLCmd:::", s.SQL, s.params)
	if Operate(s.SQL) == SELECT {
		if result, err = s.Query(tx); err != nil {
			return result, err
		}
	} else {
		if result, err = s.Exec(tx); err != nil {
			return result, err
		}
	}
	return result, nil
}

func rows2maps(rows *sql.Rows) (resultsSlice Tsql.QueryResult, err error) {
	for i := 0; rows.Next(); i++ {
		result, err := row2map(rows)
		if err != nil {
			return Tsql.QueryResult{}, err
		}
		resultsSlice.Append(result)
	}
	return resultsSlice, nil
}

func row2map(rows *sql.Rows) (resultsMap map[string][]ref.Val, err error) {
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
	var result = make(map[string][]ref.Val)
	for ii, key := range fields {
		result[key.Name()] = []ref.Val{ref.NewVal(scanResultContainers[ii])}
	}
	return result, nil
}
