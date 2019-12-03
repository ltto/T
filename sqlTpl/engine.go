package sqlTpl

import (
	"database/sql"

	"github.com/ltto/T/gobox/gid"
)

type TplEngine struct {
	db   *sql.DB
	scan bool
	m    map[string]SqlTplFile
}

func Open(driverName, dataSourceName string) (TplEngine, error) {
	db, err := sql.Open(driverName, dataSourceName)
	gid.GetGID()
	return TplEngine{db: db}, err
}

