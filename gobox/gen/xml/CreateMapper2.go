package creat

import (
	"fmt"
	"go/format"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
)

func createMapper2(list map[string][]tableDesc, pkgName, dest, url string) {
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
						pri = toUp(v.SQLField)
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
			read := "func init(){\nengine = sqlTpl.NewTplEngine(db.DB.DB())\n" +
				"engine.Scanner(\"HipsCenter/tpl\")"
			for table := range list {
				vars := "&" + toUp(table) + "MP"
				read += "engine.LocalMapper(" + vars + ")\n"
			}
			return read + "}"
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
