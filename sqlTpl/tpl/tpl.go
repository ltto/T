//go:generate tpl2go --name TPLS  --pkg $GOPACKAGE --file $GOFILE --dir sql/ --t ./bytes_tpl.go
package tpl

import (
	"database/sql"

	"github.com/ltto/T/sqlTpl"
)

var (
	Engine *sqlTpl.TplEngine
	TPLS   map[string][]byte
)

func Initialization(db *sql.DB, mappers ...interface{}) {
	Engine = sqlTpl.NewTplEngine(db)
	Engine.ScannerByBytes(TPLS)
	Engine.LoadMappers(mappers...)
}
