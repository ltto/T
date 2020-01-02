package mybatis

import (
	"database/sql"

	"github.com/ltto/T/gobox/utils"
)

type Engine struct {
	DB   *sql.DB
	DmlM map[string]*DML
}

func (e Engine) GetDB() SqlCmd {
	tx := txs[utils.GetGID()]
	if tx != nil {
		return tx
	}
	return e.DB
}

func Open(driverName, dataSourceName string) (*Engine, error) {
	DB, err := sql.Open(driverName, dataSourceName)
	if err != nil {
		return nil, err
	}
	engine := Engine{DB: DB, DmlM: make(map[string]*DML)}
	return &engine, nil
}
