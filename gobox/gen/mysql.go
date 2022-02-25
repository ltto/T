package gen

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/dave/jennifer/jen"
	"github.com/jinzhu/gorm"
	"os"
	"path"
	"errors"
	"io/ioutil"
)

type Null bool

func (n *Null) Scan(src interface{}) error {
	uint8s, ok := src.([]uint8)
	if !ok {
		return nil
	}
	*n = strings.ToUpper(string(uint8s)) == "YES"
	return nil
}

type Type struct {
	Name     string
	Len      int
	Unsigned bool
}

func (t Type) Field() (fuc FieldFuc) {
	defer func() {
		if fuc == nil {
			fuc = Interface
		}
	}()
	if t.Unsigned {
		fuc = u[strings.ToUpper(t.Name)]
		return fuc
	} else {
		fuc := m[strings.ToUpper(t.Name)]
		return fuc
	}
}

func (t *Type) Scan(src interface{}) (err error) {
	uint8s, ok := src.([]uint8)
	if !ok {
		return nil
	}
	s := string(uint8s)
	index := strings.Index(s, " unsigned")
	if index > 0 {
		t.Unsigned = true
		s = s[:len(s)-len(" unsigned")]
	}
	b := strings.Index(s, "(")
	e := strings.Index(s, ")")
	if b != -1 {
		t.Name = s[:b]
		t.Len, _ = strconv.Atoi(s[b+1 : e])
	} else {
		t.Name = s
	}
	return nil
}

type DescTab struct {
	Field   string      `gorm:"column:Field"`
	Type    Type        `gorm:"column:Type"`
	Null    Null        `gorm:"column:Null"`
	Key     string      `gorm:"column:Key"`
	Default interface{} `gorm:"column:Default"`
	Extra   string      `gorm:"column:Extra"`
}

func addStruct(f *jen.File, tab, name string, d []DescTab) {
	var codes []jen.Code
	codes = append(codes, jen.Qual("gorm.io/gorm", "Model"))

	for i := range d {
		if strings.EqualFold(d[i].Field, "id") ||
			strings.EqualFold(d[i].Field, "Created_At") ||
			strings.EqualFold(d[i].Field, "Updated_At") ||
			strings.EqualFold(d[i].Field, "Deleted_At") {
			continue
		}
		field := d[i].Type.Field()
		codes = append(codes, field(d[i]))
	}
	f.Type().Id(name).Struct(codes...).Line().
		Func().
		Params(jen.Id(string(strings.ToLower(name)[0]) + " *" + name)).
		Id("TableName").Params().String().
		Block(
			jen.Return(jen.Id("\"" + tab + "\"")),
		)
}

func addCRUD(f *jen.File, tab, name string, d []DescTab) {
	DB := jen.Id("DB")
	sqlName := ""
	model := jen.Id("&" + name).Block()
	sqlName = "Get" + name + "ByID"
	var cols = make([]string, len(d))
	for i, descTab := range d {
		cols[i] = "`" + descTab.Field + "`"

	}
	selectStr := strings.Join(cols, ",")

	f.Comment(sqlName)
	f.Func().Id(sqlName).Params(jen.Id("id").Int()).Params(jen.Id("result").Id(name), jen.Err().Error()).Block(
		jen.Return(jen.Id("TX"+sqlName).Call(DB, jen.Id("id"))),
	)
	f.Func().Id("TX"+sqlName).Params(jen.Id("tx").Id("*gorm.DB"), jen.Id("id").Int()).Params(jen.Id("result").Id(name), jen.Err().Error()).Block(
		jen.Err().Op("=").Id("tx").Op(".").Id("Model").Call(model).Op(".").Id("Select").Call(jen.Lit(selectStr)).Op(".").Id("Find").Call(jen.Id("&result"), jen.Lit("`id`=?"), jen.Id("id")).Op(".Error"),
		jen.Return(),
	)

	sqlName = "GetOne" + name
	f.Comment(sqlName)
	f.Func().Id(sqlName).Params(jen.Id("cond").Id("...interface{}")).Params(jen.Id("result").Id(name), jen.Err().Error()).Block(
		jen.Return(jen.Id("TX"+sqlName).Call(DB, jen.Id("cond..."))),
	)
	f.Func().Id("TX"+sqlName).Params(jen.Id("tx").Id("*gorm.DB"), jen.Id("cond").Id("...interface{}")).Params(jen.Id("result").Id(name), jen.Err().Error()).Block(
		jen.Err().Op("=").Id("tx").Op(".").Id("Model").Call(model).Op(".").Id("Select").Call(jen.Lit(selectStr)).Op(".").Id("Find").Call(jen.Id("&result"), jen.Id("cond...")).Op(".Error"),
		jen.Return(),
	)
	sqlName = "GetList" + name
	f.Comment(sqlName)
	f.Func().Id(sqlName).Params(jen.Id("offset"), jen.Id("limit").Int(), jen.Id("order").Map(jen.String()).Bool(), jen.Id("cond").Id("...interface{}")).Params(jen.Id("result").Id("[]"+name), jen.Id("count").Int64(), jen.Err().Error()).Block(
		jen.Return(jen.Id("TX"+sqlName).Call(DB, jen.Id("offset"), jen.Id("limit"), jen.Id("order"), jen.Id("cond..."))),
	)
	f.Func().Id("TX"+sqlName).Params(jen.Id("tx").Id("*gorm.DB"), jen.Id("offset"), jen.Id("limit").Int(), jen.Id("order").Map(jen.String()).Bool(), jen.Id("cond").Id("...interface{}")).Params(jen.Id("result").Id("[]"+name), jen.Id("count").Int64(), jen.Err().Error()).Block(
		jen.Id("temp").Op(":=").Id("tx").Op(".").Id("Model").Call(model).Op(".").Id("Select").Call(jen.Lit(selectStr)),
		jen.If(jen.Id("len").Call(jen.Id("cond")).Op(">").Id("0")).Block(
			jen.Id("temp").Op(".").Id("Where").Call(jen.Id("cond[0]"), jen.Id("cond[1:]...")),
			jen.Return(),
		),
		jen.Id("temp").Op("=").Id("temp").Op(".").Id("Count").Call(jen.Id("&count")).Op(".").Id("Offset").Call(jen.Id("offset")).Op(".").Id("Limit").Call(jen.Id("limit")),
		jen.If(jen.Len(jen.Id("order")).Op(">").Id("0")).Block(
			jen.Id("column").Op(":=").Qual("gorm.io/gorm/clause", "OrderByColumn").Block(),
			jen.For(jen.Id("s, b := range order")).Block(
				jen.Id("column.Column.Name = s"),
				jen.Id("column.Desc = b"),
				jen.Break(),
			),
			jen.Id("temp.Order(column)"),
		),
		jen.Err().Op("=").Id("temp").Op(".").Id("Scan").Call(jen.Id("&result")).Op(".Error"),
		jen.Return(),
	)

	sqlName = "Save" + name
	f.Comment(sqlName)
	f.Func().Id(sqlName).Params(jen.Id("obj").Id("*" + name)).Params(jen.Err().Error()).Block(
		jen.Return(jen.Id("TX"+sqlName).Call(DB, jen.Id("obj"))),

	)
	f.Func().Id("TX"+sqlName).Params(jen.Id("tx").Id("*gorm.DB"), jen.Id("obj").Id("*"+name)).Params(jen.Err().Error()).Block(
		jen.Return(jen.Id("tx").Op(".").Id("Model").Call(model)).Op(".").Id("Create").Call(jen.Id("&obj")).Op(".Error"),

	)

	sqlName = "Save" + name + "List"
	f.Comment(sqlName)
	f.Func().Id(sqlName).Params(jen.Id("obj").Id("..." + name)).Params(jen.Err().Error()).Block(
		jen.Return(jen.Id("TX"+sqlName).Call(DB, jen.Id("obj..."))),
	)
	f.Func().Id("TX"+sqlName).Params(jen.Id("tx").Id("*gorm.DB"), jen.Id("obj").Id("..."+name)).Params(jen.Err().Error()).Block(
		jen.Return(jen.Id("tx").Op(".").Id("Model").Call(model)).Op(".").Id("Create").Call(jen.Id("&obj")).Op(".Error"),
	)

	sqlName = "Update" + name
	f.Comment(sqlName)
	f.Func().Id(sqlName).Params(jen.Id("updates").Map(jen.String()).Interface(), jen.Id("where").String(), jen.Id("args").Id("...interface{}")).Error().Block(
		jen.Return(jen.Id("TX"+sqlName).Call(DB, jen.Id("updates"), jen.Id("where"), jen.Id("args..."))),
	)
	f.Func().Id("TX"+sqlName).Params(jen.Id("tx").Id("*gorm.DB"), jen.Id("updates").Map(jen.String()).Interface(), jen.Id("where").String(), jen.Id("args").Id("...interface{}")).Error().Block(
		jen.Return(jen.Id("tx").Op(".").Id("Model").Call(model).Op(".").Id("Where").Call(jen.Id("where"), jen.Id("args...")).Op(".").Id("Updates").Call(jen.Id("updates")).Op(".").Id("Error")),
	)

	sqlName = "Update" + name + "Model"
	f.Comment(sqlName)
	f.Func().Id(sqlName).Params(jen.Id("obj").Id("*" + name)).Params(jen.Error()).Block(
		jen.Return(jen.Id("TX"+sqlName).Call(DB, jen.Id("obj"))),
	)
	f.Func().Id("TX"+sqlName).Params(jen.Id("tx").Id("*gorm.DB"), jen.Id("obj").Id("*"+name)).Params(jen.Error()).Block(
		jen.Return(jen.Id("tx").Op(".").Id("Model").Call(model).Op(".").Id("Where").Call(jen.Lit("id=?"), jen.Id("obj.ID")).Op(".").Id("Updates").Call(jen.Id("&obj")).Op(".").Id("Error")),
	)

	sqlName = "Delete" + name + "ByID"
	f.Comment(sqlName)
	f.Func().Id(sqlName).Params(jen.Id("id").Id("...int")).Error().Block(
		jen.Return(jen.Id("TX"+sqlName).Call(DB, jen.Id("id..."))),
	)
	f.Func().Id("TX"+sqlName).Params(jen.Id("tx").Id("*gorm.DB"), jen.Id("id").Id("...int")).Error().Block(
		jen.Return(jen.Id("tx").Op(".").Id("Delete").Call(model, jen.Id("id")).Op(".").Id("Error")),
	)
}

func ScanTable(param *Param) (result string, err error) {
	if !(param.Password == "") {
		param.Password = ":" + param.Password
	}
	url := fmt.Sprintf("%v%v@tcp(%v:%v)/%v?charset=utf8mb4&collation=utf8mb4_bin&loc=Local&parseTime=true",
		param.User,
		param.Password,
		param.Host,
		param.Port,
		param.Database,
	)
	db, err := gorm.Open("mysql", url)
	if err != nil {
		return result, err
	}

	out := param.Out
	if len(out) == 0 {
		file := jen.NewFile(param.Package)
		if param.InitAble {
			InitCode(file, url)
		}
		for tab, structName := range param.Table2struct {
			if err = StructCode(file, db, tab, structName); err != nil {
				return "", err
			}
		}
		src := fmt.Sprintf("%#v", file)
		return src, nil
	} else {
		if !path.IsAbs(out) {
			return "", errors.New("不是绝对路径")
		}
		_ = os.MkdirAll(out, 0777)
		var files = map[string]*jen.File{}
		if param.InitAble {
			file := jen.NewFile(param.Package)
			InitCode(file, url)
			files["a_init.go"] = file
		}
		for tab, structName := range param.Table2struct {
			file := jen.NewFile(param.Package)
			if err = StructCode(file, db, tab, structName); err != nil {
				return "", err
			}
			files[tab+".go"] = file
		}
		for name, file := range files {
			err := ioutil.WriteFile(path.Join(out, name), []byte(fmt.Sprintf("%#v", file)), 0777)
			if err != nil {
				return "", err
			}
		}
		return "OKKKK", nil
	}
}
func InitCode(file *jen.File, url string) {
	file.Func().Id("init").Params().Block(
		jen.Id("dsn").Op(":=").Lit(url),
		jen.Id("db").Op(",").Id("err").Op(":=").Qual("gorm.io/gorm", "Open").Call(jen.Qual("gorm.io/driver/mysql", "Open").Call(jen.Id("dsn")), jen.Id("&gorm.Config").Block(
			jen.Id("Logger").Op(":").Qual("gorm.io/gorm/logger", "Default.LogMode").Call(jen.Id("logger.Info")).Op(","),
		)),
		jen.If(jen.Id("err").Op("!=").Nil()).Block(jen.Panic(jen.Id("err"))),
		jen.Id("DB").Op("=").Id("db"),
	)
	file.Var().Id("DB").Id("*gorm.DB")
}

func StructCode(file *jen.File, db *gorm.DB, tab, structName string) (err error) {
	var tabField []DescTab
	err = db.Raw("DESC " + tab).Scan(&tabField).Error
	if err != nil {
		return err
	}
	addStruct(file, tab, structName, tabField)
	addCRUD(file, tab, structName, tabField)
	return nil
}
