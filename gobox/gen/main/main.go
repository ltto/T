package main

import (
	"database/sql"
	"fmt"
	"regexp"

	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo"
	"github.com/ltto/T/gobox/gen"
)

func main() {
	e := echo.New()
	e.File("/", "gobox/gen/index.html")
	e.POST("/db2struct", func(c echo.Context) error {
		param := new(gen.Param)
		if err := c.Bind(param); err != nil {
			return err
		}
		result, err := gen.ScanTable(param)
		if err != nil {
			return err
		}
		return c.String(200, result)
	})

	e.POST("/tables", func(c echo.Context) error {
		param := new(gen.Param)
		if err := c.Bind(param); err != nil {
			return err
		}
		if param.Regexp == "" {
			param.Regexp = ".*"
		}
		compile, err := regexp.Compile(param.Regexp)
		if err != nil {
			return err
		}
		if !(param.Password == "") {
			param.Password = ":" + param.Password
		}
		url := fmt.Sprintf("%v%v@tcp(%v:%v)/%v?charset=utf8mb4&collation=utf8mb4_bin&loc=Local&parseTime=true",
			param.User,
			param.Password,
			param.Host,
			param.Port,
			param.Database,
		)
		db, err := sql.Open("mysql", url)
		if err != nil {
			return err
		}
		rows, err := db.Query("SHOW TABLES ")
		if err != nil {
			return err
		}
		var tables []string
		var tab string
		for rows.Next() {
			if err := rows.Scan(&tab); err != nil {
				return err
			}
			if compile.Match([]byte(tab)) {
				tables = append(tables, tab)
			}
		}
		return c.JSON(200, tables)
	})

	e.Start(":8080")
}
