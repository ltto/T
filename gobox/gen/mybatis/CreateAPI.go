package creat

import (
	"go/format"
	"io/ioutil"
	"os"
	"path"
	"strings"

	_ "github.com/guregu/null"
)

func createAPI(list map[string][]tableDesc, pkgName, dest string) {
	os.MkdirAll(path.Join(dest), 0777)
	goFile := os.Expand(goFile, func(s string) string {
		switch s {
		case "packageName":
			return pkgName
		case "imports":
			return `
import (
	"net/http"
	"reflect"

	"github.com/labstack/echo/v4"
	"github.com/ltto/T/echoT"
	"github.com/ltto/T/echoT/vo"
)
`
		case "init":
			init := `func init(){
`
			for table, v := range list {

				f := func(s string) string {
					switch s {
					case "structName":
						return strings.ToLower(toUp(table))
					case "struct":
						return toUp(table)
					case "PRI":
						for _, v2 := range v {
							if v2.PRIKey() {
								return toUp(v2.SQLField)
							}
						}
						return "id"
					}
					return ""
				}
				init += os.Expand(line, f) + os.Expand(One, f) + os.Expand(List, f) + os.Expand(Save, f) + os.Expand(Update, f) +
					"//------------------------\n\n"
			}
			return init + "}"
		default:
			return ""
		}
	})
	fileName := path.Join(dest, "api.go")
	os.Remove(fileName)
	if bytes, err := format.Source([]byte(goFile)); err == nil {
		if err := ioutil.WriteFile(fileName, bytes, 0777); err != nil {
			panic(err)
		}
	} else {
		if err := ioutil.WriteFile(fileName, []byte(goFile), 0777); err != nil {
			panic(err)
		}
	}

}

var line = "//******************${struct}******************//\n"
var One = "//${struct} get one\n" +
	"echoT.R(echoT.RouterInfo{Mapping: \"/${structName}\", HttpMethod: http.MethodGet,\n" +
	"	InterfaceMap: echoT.InterfaceMap{\"data\": reflect.TypeOf(${struct}{})},\n" +
	"	Do: func(c echo.Context, res struct{ ID int `query:\"${PRI}\"` }) vo.Result {\n" +
	"	if one, err := ${struct}MP.SelectByID(res.ID); err != nil {\n" +
	"		return *vo.Err(err)\n" +
	"	} else {\n" +
	"		return *vo.Success(one)\n" +
	"	}\n" +
	"	},\n" +
	"})\n"

var List = "//${struct} get list\n" +
	"echoT.R(echoT.RouterInfo{Mapping: \"/${structName}s\", HttpMethod: http.MethodGet,\n" +
	"	InterfaceMap: echoT.InterfaceMap{\"data\": reflect.TypeOf([]${struct}{})},\n" +
	"	Do: func(c echo.Context, res *vo.Page) vo.Result {\n" +
	"	list, err := ${struct}MP.SelectLimit(res.Limit())\n" +
	"	if err != nil {\n" +
	"		return *vo.Err(err)\n" +
	"	}\n" +
	"	count, err := ${struct}MP.SelectCount()\n" +
	"	if err != nil {\n" +
	"		return *vo.Err(err)\n" +
	"	}\n" +
	"	res.Count=count\n" +
	"	return *vo.List(list, res)\n" +
	"	},\n" +
	"})\n"
var Save = "//${struct} save\n" +
	"echoT.R(echoT.RouterInfo{Mapping: \"/${structName}/save\", HttpMethod: http.MethodPost,\n" +
	"	InterfaceMap: echoT.InterfaceMap{\"data\": reflect.TypeOf(${struct}{})},\n" +
	"	Do: func(c echo.Context, res ${struct}) vo.Result {\n" +
	"	if err := ${struct}MP.Save(res); err != nil {\n" +
	"		return *vo.Err(err)\n" +
	"			}\n" +
	"		return *vo.Success(res)\n" +
	"	},\n" +
	"})\n"
var Update = "//${struct} update\n" +
	"echoT.R(echoT.RouterInfo{Mapping: \"/${structName}\", HttpMethod: http.MethodPost,\n" +
	"	Auth:         false,\n" +
	"	InterfaceMap: echoT.InterfaceMap{\"data\": reflect.TypeOf(${struct}{})},\n" +
	"	Do: func(c echo.Context, res ${struct}) vo.Result {\n" +
	"	if err := ${struct}MP.UpdateByID(res); err != nil {\n" +
	"		return *vo.Err(err)\n" +
	"			}\n" +
	"		return *vo.Success(res)\n" +
	"	},\n" +
	"})\n"
