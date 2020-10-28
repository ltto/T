package node

import "strings"

type SQLOperate int

const (
	NOT SQLOperate = iota
	INSERT
	DELETE
	UPDATE
	SELECT
)

func Operate(sql string) SQLOperate {
	sql = strings.ToUpper(strings.TrimSpace(sql))
	if strings.HasPrefix(sql, "INSERT") {
		return INSERT
	} else if strings.HasPrefix(sql, "DELETE") {
		return DELETE
	} else if strings.HasPrefix(sql, "UPDATE") {
		return UPDATE
	} else if strings.HasPrefix(sql, "SELECT") {
		return SELECT
	}
	return NOT
}
