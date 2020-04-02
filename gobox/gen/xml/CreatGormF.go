package creat

import (
	"go/format"
	"io/ioutil"
	"os"
	"path"
	"strings"
)

func createGormF(list map[string][]tableDesc, pkgName, dest, url string) {
	os.MkdirAll(path.Join(dest), 0777)
	goFile := os.Expand(goFile, func(s string) string {
		switch s {
		case "packageName":
			return pkgName
		case "imports":
			return `
import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm" 
)
`
		case "init":
			return `
func init(){
	GormDB, e := gorm.Open("mysql", "` + url + `")
	if e != nil {
		panic(e)
	} else {
		DB = GormDB
	}
}
`
		case "vars":
			vars := "var DB *gorm.DB\n"
			for table, v := range list {
				f := func(s string) string {
					switch s {
					case "structName":
						return toUp(table)
					case "struct":
						return strings.ToLower(toUp(table))
					case "PRI":
						for _, v2 := range v {
							if v2.PRIKey() {
								return v2.SQLField
							}
						}
						return "id"
					}
					return ""
				}
				vars += os.Expand(gorm, f)
			}
			return vars
		default:
			return ""
		}
	})
	fileName := path.Join(dest, "gorm.go")
	os.Remove(fileName)
	if bytes, err := format.Source([]byte(goFile)); err != nil {
		if err := ioutil.WriteFile(fileName, []byte(goFile), 0777); err != nil {
			panic(err)
		}
	} else {
		if err := ioutil.WriteFile(fileName, bytes, 0777); err != nil {
			panic(err)
		}

	}

}

var gorm = `var ${structName}DB = struct {
	GetByID     func(id int) (${structName}, error)
	ListByLimit func(o, l int) ([]${structName}, error)
	Count       func() (int, error)
	UpdateById  func(${structName}) error
	DeleteById  func(id int) error
	DeleteByIds func(id ...int) error
}{
	GetByID: func(id int) (${structName}, error) {
		obj := ${structName}{}
		err := DB.Model(&${structName}{}).Where("${PRI}=?", id).Find(&obj).Error
		return obj, err
	},
	ListByLimit: func(o, l int) ([]${structName}, error) {
		var list []${structName}
		err := DB.Model(&${structName}{}).Offset(o).Limit(l).Find(&list).Error
		return list, err
	},
	Count: func() (int, error) {
		var (
			list  []${structName}
			count int
		)
		err := DB.Model(&${structName}{}).Find(&list).Count(&count).Error
		return count, err
	},
	UpdateById: func(up ${structName}) error {
		return DB.Model(&${structName}{}).Update(up).Error
	},
	DeleteById: func(id int) error {
		return DB.Model(&${structName}{}).Where("${PRI}=?", id).Delete(nil).Error
	},
	DeleteByIds: func(id ...int) error {
		return DB.Model(&${structName}{}).Where("${PRI} in (?)", id).Delete(nil).Error
	},
}
`
