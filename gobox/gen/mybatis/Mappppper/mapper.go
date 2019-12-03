package Mappppper

import (
	"io/ioutil"

	_ "github.com/go-sql-driver/mysql"
	"github.com/ltto/GoMybatis"
)

func init() {
	const MysqlUri = "root@tcp(127.0.0.1:3306)/im?charset=utf8mb4&collation=utf8mb4_bin&loc=Local&parseTime=true"
	engine := GoMybatis.GoMybatisEngine{}.New()
	if _, err := engine.Open("mysql", MysqlUri); err != nil {
		panic(err.Error())
	}
	//读取mapper xml文件
	if bytes, err := ioutil.ReadFile("/Users/ltt/go/src/github.com/ltto/T/gobox/gen/mybatis/Mappppper/mapper/AlbumsMapper.xml"); err != nil {
		panic(err)
	} else {
		engine.WriteMapperPtr(&AlbumsMP, bytes)
	}
	if bytes, err := ioutil.ReadFile("/Users/ltt/go/src/github.com/ltto/T/gobox/gen/mybatis/Mappppper/mapper/BizPropertyMapper.xml"); err != nil {
		panic(err)
	} else {
		engine.WriteMapperPtr(&BizPropertyMP, bytes)
	}
	if bytes, err := ioutil.ReadFile("/Users/ltt/go/src/github.com/ltto/T/gobox/gen/mybatis/Mappppper/mapper/PhotosMapper.xml"); err != nil {
		panic(err)
	} else {
		engine.WriteMapperPtr(&PhotosMP, bytes)
	}
	if bytes, err := ioutil.ReadFile("/Users/ltt/go/src/github.com/ltto/T/gobox/gen/mybatis/Mappppper/mapper/BizActivityMapper.xml"); err != nil {
		panic(err)
	} else {
		engine.WriteMapperPtr(&BizActivityMP, bytes)
	}
	if bytes, err := ioutil.ReadFile("/Users/ltt/go/src/github.com/ltto/T/gobox/gen/mybatis/Mappppper/mapper/UsersMapper.xml"); err != nil {
		panic(err)
	} else {
		engine.WriteMapperPtr(&UsersMP, bytes)
	}

}

var (
	BizActivityMP = BizActivityMapper{}
	UsersMP       = UsersMapper{}
	AlbumsMP      = AlbumsMapper{}
	BizPropertyMP = BizPropertyMapper{}
	PhotosMP      = PhotosMapper{}
)

type BizActivityMapper struct {
	Save        func(obj BizActivity) error           `mapperParams:"obj"`
	SelectByID  func(id int) (BizActivity, error)     `mapperParams:"id"`
	SelectLimit func(o, l int) ([]BizActivity, error) `mapperParams:"o,l"`
	SelectCount func() (int, error)
	UpdateByID  func(obj BizActivity) error `mapperParams:"obj"`
	DeleteByID  func(id int) error          `mapperParams:"id"`
	DeleteByIDs func(ids []int) error       `mapperParams:"ids"`
}
type UsersMapper struct {
	Save        func(obj Users) error           `mapperParams:"obj"`
	SelectByID  func(id int) (Users, error)     `mapperParams:"id"`
	SelectLimit func(o, l int) ([]Users, error) `mapperParams:"o,l"`
	SelectCount func() (int, error)
	UpdateByID  func(obj Users) error `mapperParams:"obj"`
	DeleteByID  func(id int) error    `mapperParams:"id"`
	DeleteByIDs func(ids []int) error `mapperParams:"ids"`
}
type AlbumsMapper struct {
	Save        func(obj Albums) error           `mapperParams:"obj"`
	SelectByID  func(id int) (Albums, error)     `mapperParams:"cid"`
	SelectLimit func(o, l int) ([]Albums, error) `mapperParams:"o,l"`
	SelectCount func() (int, error)
	UpdateByID  func(obj Albums) error `mapperParams:"obj"`
	DeleteByID  func(id int) error     `mapperParams:"cid"`
	DeleteByIDs func(ids []int) error  `mapperParams:"ids"`
}
type BizPropertyMapper struct {
	Save        func(obj BizProperty) error           `mapperParams:"obj"`
	SelectByID  func(id int) (BizProperty, error)     `mapperParams:"id"`
	SelectLimit func(o, l int) ([]BizProperty, error) `mapperParams:"o,l"`
	SelectCount func() (int, error)
	UpdateByID  func(obj BizProperty) error `mapperParams:"obj"`
	DeleteByID  func(id int) error          `mapperParams:"id"`
	DeleteByIDs func(ids []int) error       `mapperParams:"ids"`
}
type PhotosMapper struct {
	Save        func(obj Photos) error           `mapperParams:"obj"`
	SelectByID  func(id int) (Photos, error)     `mapperParams:"id"`
	SelectLimit func(o, l int) ([]Photos, error) `mapperParams:"o,l"`
	SelectCount func() (int, error)
	UpdateByID  func(obj Photos) error `mapperParams:"obj"`
	DeleteByID  func(id int) error     `mapperParams:"id"`
	DeleteByIDs func(ids []int) error  `mapperParams:"ids"`
}
