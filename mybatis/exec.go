package mybatis

import (
	"database/sql"
	"fmt"
	"github.com/ltto/T/mybatis/node"
)

type SqlCmd interface {
	Query(query string, args ...interface{}) (*sql.Rows, error)
	Exec(query string, args ...interface{}) (sql.Result, error)
}
type SqlExec struct {
	db      *sql.DB
	SQL     string
	Operate node.SQLOperate
	params  []interface{}
}

func PareSQL(m map[string]interface{}, root *node.DMLRoot) (ex SqlExec, err error) {
	pare, err := root.PareSQL(m)
	if err != nil {
		return ex, err
	}
	ex.SQL = pare.SQL
	ex.params = pare.Params
	ex.Operate = pare.Operate
	return
}

func (s SqlExec) Query(tx SqlCmd) (*sql.Rows, error) {
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
		return nil, err
	}
	return rows, nil
}

func (s SqlExec) Exec(tx SqlCmd) (LastInsertId int64, RowsAffected int64, err error) {
	var (
		result sql.Result
	)
	if tx == nil {
		result, err = s.db.Exec(s.SQL, s.params...)
	} else {
		result, err = tx.Exec(s.SQL, s.params...)
	}
	if err != nil {
		return 0, 0, err
	}
	if LastInsertId, err = result.LastInsertId(); err != nil {
		return
	}

	if RowsAffected, err = result.RowsAffected(); err != nil {
		return
	}
	return
}

func (s SqlExec) ExecSQL(tx SqlCmd) (result *SQLResult, err error) {
	result = &SQLResult{}
	fmt.Println("SQLCmd:::", s.SQL, s.params)
	if s.Operate == node.SELECT {
		if result.Rows, err = s.Query(tx); err != nil {
			return result, err
		}
	} else {
		if result.LastInsertId, result.RowsAffected, err = s.Exec(tx); err != nil {
			return result, err
		}
	}
	return result, nil
}
