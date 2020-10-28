package main

import (
	"fmt"
	"github.com/guregu/null"

	_ "github.com/go-sql-driver/mysql"
	"github.com/ltto/T/mybatis"
)

func main() {
	var albumsMapper AlbumsMapper
	engine, err := mybatis.Open("mysql", "root:123@tcp(127.0.0.1:3306)/im?charset=utf8mb4&collation=utf8mb4_bin&loc=Local&parseTime=true")
	if err != nil {
		panic(err)
	}
	if err = engine.LoadAndBindMap(&mybatis.LoadConf{
		PathPrefix: "/Users/liutongtong/gocode/T/mybatis",
		Tag:        "json",
	}, map[string]interface{}{
		"AlbumsMapper.xml": &albumsMapper,
	}); err != nil {
		panic(err)
	}

	if err = engine.BeginTX(); err != nil {
		panic(err)
	}
	//albums := Albums{}
	//albums.Name.Scan("xiaoming")
	//fmt.Println(albums.Name.Value())
	//fmt.Println(albumsMapper.Save(&albums))
	fmt.Println(engine.Commit())
	//fmt.Println(albums)
	id, err := albumsMapper.SelectByID(78)
	fmt.Println(id, err)
}

type AlbumsMapper struct {
	Save        func(obj Albums) error                          `mapperParams:"obj"`
	SelectByID  func(id int) (************[]*****map[string]interface{}, error) `mapperParams:"cid"`
	UpdateByID  func(obj *Albums) error                         `mapperParams:"obj"`
	DeleteByID  func(id int) error                              `mapperParams:"cid"`
	DeleteByIDs func(ids []int) error                           `mapperParams:"ids"`
}

type Albums struct {
	Cid       null.Int    `json:"cid;primary_key"`
	UserID    null.Int    `json:"userID"`
	Name      null.String `json:"name"`
	URL       null.String `json:"URL"`
	CreatedAt null.Time   `json:"createdAt"`
	UpdatedAt null.Time   `json:"updatedAt"`
	DeletedAt null.Time   `json:"deletedAt"`
}

func (a Albums) Scan(src interface{}) error {
	panic("implement me")
}
