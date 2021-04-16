package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/guregu/null"
	"github.com/ltto/T/mybatis"
	"github.com/ltto/T/mybatis/node"
)

func main() {
	var albumsMapper AlbumsMapper
	engine, err := mybatis.Open("mysql", "root:123@tcp(127.0.0.1:3366)/im?charset=utf8mb4&collation=utf8mb4_bin&loc=Local&parseTime=true")
	if err != nil {
		panic(err)
	}
	if err = engine.LoadAndBindMap(&mybatis.LoadConf{
		PathPrefix: "/Users/ltt/go/src/github.com/ltto/T/mybatis/",
		Tag:        "json",
	}, map[string]interface{}{
		"AlbumsMapper.xml": &albumsMapper,
	}); err != nil {
		panic(err)
	}

	if err = engine.BeginTX(); err != nil {
		panic(err)
	}

	//root := node.Select()
	//root.Text("SELECT `cid`, `userID`, `name`, `createdAt`, `updatedAt`, `deletedAt`").Text("\r\nssssss")
	//fmt.Println(root.PareSQL(nil))
	//albums := Albums{}
	//albums.Name.Scan("xiaoming")
	//fmt.Println(albums.Name.Value())
	//fmt.Println(albumsMapper.Save(&albums))
	//fmt.Println("commit", engine.Commit())
	//fmt.Println(albums)
	//id, err := albumsMapper.SelectByID(78)
	//fmt.Println(id, err)

	sql, err := mybatis.NewNodeRoot(node.Select(
		node.Text_("select * from clnts where id in"),
		node.For_(node.F{
			Item:       "item",
			Index:      "index",
			Collection: "ids",
			Open:       "(",
			Separator:  ",",
			Close:      ")",
		}, node.Text_("#{item}")),
	), nil).PareSQL(map[string]interface{}{
		"ids": []int{1, 2, 3, 4, 5, 6, 7, 8, 9},
	})
	if err != nil {
		panic(err)
	}
	fmt.Println(sql.SQL)
}

type AlbumsMapper struct {
	Save        func(obj *Albums) error      `mapperParams:"obj"`
	SelectByID  func(id int) (Albums, error) `mapperParams:"cid"`
	UpdateByID  func(obj *Albums) error      `mapperParams:"obj"`
	DeleteByID  func(id int) error           `mapperParams:"cid"`
	DeleteByIDs func(ids []int) error        `mapperParams:"ids"`
}

type Albums struct {
	Cid    null.Int    `json:"cid;primary_key"`
	UserID null.Int    `json:"userID"`
	Name   null.String `json:"name"`
	//URL       null.String `json:"URL"`
	CreatedAt null.Time `json:"createdAt"`
	UpdatedAt null.Time `json:"updatedAt"`
	DeletedAt null.Time `json:"deletedAt"`
}

func (a Albums) Scan(src interface{}) error {
	panic("implement me")
}
