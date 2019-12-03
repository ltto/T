package main

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/guregu/null"
	"github.com/ltto/T/sqlTpl"
)

type AlbumsMapper struct {
	Save        func(obj Albums) error           `mapperParams:"obj"`
	SelectByID  func(id int) (*Albums, error)    `mapperParams:"Cid"`
	SelectLimit func(o, l int) ([]Albums, error) `mapperParams:"o,l"`
	SelectCount func() (int, error)
	UpdateByID  func(obj Albums) error `mapperParams:"obj"`
	DeleteByID  func(id int) error     `mapperParams:"Cid"`
	DeleteByIDs func(ids []int) error  `mapperParams:"ids"`
}

func main() {
	open, err := sqlTpl.Open("mysql", "root@tcp(127.0.0.1:3306)/im?charset=utf8mb4&collation=utf8mb4_bin&loc=Local&parseTime=true")
	if err != nil {
		panic(err)
	}
	open.Scanner("/Users/ltt/code/sqlTpl/main/")
	mapper := AlbumsMapper{}
	open.LocalMapper(&mapper)
	fmt.Println(mapper.SelectByID(6))
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
