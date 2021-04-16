package mybatis

import (
	"database/sql"
	"github.com/ltto/T/mybatis/node"
	"reflect"
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
