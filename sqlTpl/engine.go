package sqlTpl

import (
	"database/sql"
)

type TplEngine struct {
	db   *sql.DB
	scan bool
	load bool
	m    map[string]SqlTplFile
	TPLS map[string][]byte
}

func NewTplEngine(db *sql.DB) *TplEngine {
	return &TplEngine{db: db}
}

func Open(driverName, dataSourceName string) (TplEngine, error) {
	db, err := sql.Open(driverName, dataSourceName)
	return TplEngine{db: db}, err
}
