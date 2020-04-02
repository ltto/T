package creat

import (
	"go/format"
	"io/ioutil"
	"os"
	"path"
	"strings"
)

func createGormM(list map[string][]tableDesc, pkgName, dest, url string) {
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
			vars := "var( DB *gorm.DB\n"
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
								return "`" + v2.SQLField + "`"
							}
						}
						return "`id`"
					}
					return ""
				}
				vars += os.Expand(gormVar, f)
			}
			return vars + ")"
		case "structs":
			structs := ""
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
								return "`" + v2.SQLField + "`"
							}
						}
						return "`id`"
					}
					return ""
				}
				structs += os.Expand(gormMethod, f) + "\n//------------------------" + table + "------------------------//\n"
			}
			return structs
		default:
			return ""
		}
	})
	fileName := path.Join(dest, "gormFunc.go")
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

var gormVar = "${structName}DB = ${structName}Ope{DB: DB}\n"
var gormMethod = `type ${structName}Ope struct {
	DB *gorm.DB
}
func (a ${structName}Ope) GetByID(id int) (${structName}, error) {
	obj := ${structName}{}
	err := a.DB.Model(&${structName}{}).Where("${PRI}=?", id).Find(&obj).Error
	return obj, err
}

func (a ${structName}Ope) ListByLimit(o, l int) ([]${structName}, error) {
	var list []${structName}
	err := a.DB.Model(&${structName}{}).Offset(o).Limit(l).Find(&list).Error
	return list, err
}
func (a ${structName}Ope) Count() (int, error) {
	var count int
	err := a.DB.Model(&${structName}{}).Find(&[]${structName}{}).Count(&count).Error
	return count, err
}
func (a ${structName}Ope) UpdateById(up ${structName}) error {
	return a.DB.Model(&${structName}{}).Update(up).Error
}
func (a ${structName}Ope) DeleteById(id int) error {
	return a.DB.Model(&${structName}{}).Where("${PRI}=?", id).Delete(nil).Error
}
func (a ${structName}Ope) DeleteByIds(id ...int) error {
	return a.DB.Model(&${structName}{}).Where("${PRI} in (?)", id).Delete(nil).Error
}
`
