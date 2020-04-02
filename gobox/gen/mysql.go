package gen

import (
	"fmt"
	"github.com/dave/jennifer/jen"
	"github.com/jinzhu/gorm"
	"strconv"
	"strings"
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

func addSource(f *jen.File, name string, d []DescTab) {
	codes := make([]jen.Code, len(d))
	for i := range d {
		field := d[i].Type.Field()
		codes[i] = field(d[i])
	}
	f.Type().Id(name).Struct(codes...)
}

func ScanTable(param *Param) (result string, err error) {
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
	file := jen.NewFile(param.Package)
	table2struct := param.Table2struct
	for tab, structName := range table2struct {
		var tabField []DescTab
		err = db.Raw("DESC " + tab).Scan(&tabField).Error
		if err != nil {
			return result, err
		}
		addSource(file, structName, tabField)
	}
	src := fmt.Sprintf("%#v", file)
	return src, nil
}
