package creat

import (
	"go/format"
	"io/ioutil"
	"os"
	"path"

	_ "github.com/guregu/null"
)

func createStruct(list map[string][]tableDesc, pkgName, dest string, nullAble bool) {
	os.MkdirAll(path.Join(dest), 0777)
	goFile := os.Expand(goFile, func(s string) string {
		switch s {
		case "packageName":
			return pkgName
		case "imports":
			if nullAble {
				return `import ("github.com/guregu/null")`
			} else {
				return ``
			}
		case "structs":
			var structs = ""
			for k, descList := range list {
				structs += os.Expand(structDesc, func(s string) string {
					switch s {
					case "structName":
						return toUp(k)
					case "fields":
						var fields = ""
						for _, desc := range descList {
							fields += os.Expand(fieldDesc, func(s string) string {
								switch s {
								case "fieldName":
									return toUp(desc.SQLField)
								case "typeName":
									return desc.GoField + "`json:\"" + desc.SQLField + "\"`"
								default:
									return ""
								}
							})
						}
						return fields
					default:
						return ""
					}
				})
			}
			return structs
		default:
			return ""
		}
	})
	fileName := path.Join(dest, "obj.go")
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

const goFile = `package ${packageName}
${imports}
${init}
${vars}
${structs}

`
const structDesc = `
type ${structName} struct {
	${fields}
}`
const fieldDesc = `${fieldName}   ${typeName}
`
