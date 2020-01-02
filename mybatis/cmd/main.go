package main

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/guregu/null"
	"github.com/ltto/T/mybatis"
)

func main() {

	var m AlbumsMapper
	engine, err := mybatis.Open("mysql", "root@tcp(127.0.0.1:3306)/im?charset=utf8mb4&collation=utf8mb4_bin&loc=Local&parseTime=true")
	if err != nil {
		panic(err)
	}
	if err = engine.LoadAndBind("/Users/ltt/go/src/github.com/ltto/T/mybatis/AlbumsMapper.xml", &m); err != nil {
		panic(err)
	}

	if err = engine.BeginTX(); err != nil {
		panic(err)
	}
	err = m.Save(Albums{})
	engine.Commit()
	fmt.Println(err)
}

type AlbumsMapper struct {
	Save        func(obj Albums) error          `mapperParams:"obj"`
	SelectByID  func(id int) ([]*Albums, error) `mapperParams:"cid"`
	UpdateByID  func(obj Albums) error          `mapperParams:"obj"`
	DeleteByID  func(id int) error              `mapperParams:"cid"`
	DeleteByIDs func(ids []int) error           `mapperParams:"ids"`
}
type Albums struct {
	Cid       null.Int    `json:"cid"`
	UserID    null.Int    `json:"userID"`
	Name      null.String `json:"name"`
	URL       null.String `json:"URL"`
	CreatedAt null.Time   `json:"createdAt"`
	UpdatedAt null.Time   `json:"updatedAt"`
	DeletedAt null.Time   `json:"deletedAt"`
}
