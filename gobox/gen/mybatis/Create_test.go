package creat

import (
	"database/sql"
	"testing"

	_ "github.com/go-sql-driver/mysql"
)

func TestCreat(t *testing.T) {
	url := "root@tcp(127.0.0.1:3306)/hress2?charset=utf8mb4&collation=utf8mb4_bin&loc=Local&parseTime=true"
	db, err := sql.Open("mysql", url)
	if err != nil {
		panic(err)
	}
	defer db.Close()
	param := ListParam(db, true)
	DoCreat("Mappppper", "Mappppper", url, param)
}
