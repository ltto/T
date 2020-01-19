package mybatis

import (
	"database/sql"

	"github.com/ltto/T/gobox/utils"
)

var txs map[int64]*sql.Tx

func init() {
	txs = make(map[int64]*sql.Tx)
}
func (e *Engine) BeginTX() error {
	gid := utils.GetGID()
	if txs[gid] != nil {
		return nil
	}
	begin, err := e.DB.Begin()
	if err != nil {
		return err
	}
	txs[gid] = begin
	return nil
}

func (e *Engine) Rollback() error {
	gid := utils.GetGID()
	if txs[gid] != nil {
		return txs[gid].Rollback()
	}
	return nil
}

func (e *Engine) Commit() error {
	gid := utils.GetGID()
	if txs[gid] != nil {
		return txs[gid].Commit()
	}
	return nil
}
