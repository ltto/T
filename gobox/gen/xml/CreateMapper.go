package creat

import (
	"fmt"
	"go/format"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
)

func createMapper(list map[string][]tableDesc, pkgName, dest, url string) {
	os.MkdirAll(path.Join(dest), 0777)
	absPath, err := filepath.Abs(dest)
	fmt.Println(absPath)
	if err != nil {
		panic(err)
	}

	goFile := os.Expand(goFile, func(s string) string {
		switch s {
		case "packageName":
			return pkgName
		case "imports":
			return `import(
"io/ioutil"

_ "github.com/go-sql-driver/mysql"
"github.com/ltto/GoMybatis"
)`
		case "structs":
			var Mappers = ""
			for table, descs := range list {
				pri := "id"
				for _, v := range descs {
					if v.Key == "PRI" {
						pri = v.SQLField
					}
				}
				Mappers += os.Expand(mapperStr, func(s string) string {
					switch s {
					case ".Struct":
						return toUp(table)
					case ".PRI":
						return pri
					default:
						return ""
					}
				})
			}
			return Mappers
		case "vars":
			vars := "var("
			for table := range list {
				vars += toUp(table) + "MP=" + toUp(table) + "Mapper{}\n"
			}
			return vars + ")"
		case "init":
			read := ""
			for table := range list {
				mapper := path.Join(absPath, "mapper", toUp(table)+"Mapper.xml")
				vars := toUp(table) + "MP"
				read += `if bytes, err := ioutil.ReadFile("` + mapper + `");err != nil {
panic(err)
}else {
engine.WriteMapperPtr(&` + vars + `, bytes)
}
`
			}
			return os.Expand(`func init(){
const MysqlUri = "${url}"
engine := GoMybatis.GoMybatisEngine{}.New()
if _, err := engine.Open("mysql", MysqlUri); err != nil {
		panic(err.Error())
}
//读取mapper xml文件
${reads}
}`, func(s string) string {
				switch s {
				case "url":
					return url
				case "reads":
					return read
				}
				return ""
			})
		default:
			return ""
		}
	})
	fileName := path.Join(absPath, "mapper.go")
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

const mapperStr = "type ${.Struct}Mapper struct {\n" +
	"Save        func(obj ${.Struct}) error `mapperParams:\"obj\"`\n" +
	"SelectByID  func(id int) (${.Struct}, error) `mapperParams:\"${.PRI}\"`\n" +
	"SelectLimit func(o,l int) ([]${.Struct},error) `mapperParams:\"o,l\"`\n" +
	"SelectCount func() (int,error) \n" +
	"UpdateByID  func(obj ${.Struct}) error `mapperParams:\"obj\"`\n" +
	"DeleteByID  func(id int) error `mapperParams:\"${.PRI}\"`\n" +
	"DeleteByIDs func(ids []int) error `mapperParams:\"ids\"`\n" +
	"}\n"
