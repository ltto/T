package gen

import (
	"github.com/dave/jennifer/jen"
	"github.com/ltto/T/gobox/str"
)

type Param struct {
	JsonAble     bool              `json:"json_able"`
	NullAble     bool              `json:"null_able"`
	GormAble     bool              `json:"gorm_able"`
	Package      string            `json:"package"`
	User         string            `json:"user"`
	Host         string            `json:"host"`
	Port         string            `json:"port"`
	Password     string            `json:"password"`
	Database     string            `json:"database"`
	Dst          string            `json:"dst"`
	Regexp       string            `json:"regexp"`
	Table2struct map[string]string `json:"table_2_struct"`
}

var m = map[string]FieldFuc{
	"TINYINT":    Int8,
	"SMALLINT":   Int16,
	"MEDIUMINT":  Int32,
	"INT":        Int32,
	"INTEGER":    Int32,
	"BIGINT":     Int64,
	"FLOAT":      Float32,
	"DOUBLE":     Float64,
	"DECIMAL":    Float64,
	"CHAR":       String,
	"VARCHAR":    String,
	"TINYBLOB":   String,
	"TINYTEXT":   String,
	"BLOB":       String,
	"TEXT":       String,
	"MEDIUMBLOB": String,
	"MEDIUMTEXT": String,
	"LOGNGBLOB":  String,
	"LONGTEXT":   String,
	"DATE":       Time,
	"TIME":       Time,
	"YEAR":       Time,
	"DATETIME":   Time,
	"TIMESTAMP":  Time,
}
var u = map[string]FieldFuc{
	"TINYINT":    Uint8,
	"SMALLINT":   Uint16,
	"MEDIUMINT":  Uint32,
	"INT":        Uint32,
	"INTEGER":    Uint32,
	"BIGINT":     Uint64,
	"FLOAT":      Float32,
	"DOUBLE":     Float64,
	"DECIMAL":    Float64,
	"CHAR":       String,
	"VARCHAR":    String,
	"TINYBLOB":   String,
	"TINYTEXT":   String,
	"BLOB":       String,
	"TEXT":       String,
	"MEDIUMBLOB": String,
	"MEDIUMTEXT": String,
	"LOGNGBLOB":  String,
	"LONGTEXT":   String,
	"DATE":       Time,
	"TIME":       Time,
	"YEAR":       Time,
	"DATETIME":   Time,
	"TIMESTAMP":  Time,
}

type FieldFuc func(DescTab) jen.Code

var jsonType = str.Underline

var gormType = func(d DescTab) string {
	return "column:" + d.Field
}

func Int8(d DescTab) jen.Code {
	return jen.Id(d.Field).Int8().Tag(map[string]string{"json": jsonType(d.Field), "gorm": gormType(d)})
}
func Int16(d DescTab) jen.Code {
	return jen.Id(d.Field).Int16().Tag(map[string]string{"json": jsonType(d.Field), "gorm": gormType(d)})
}
func Int32(d DescTab) jen.Code {
	return jen.Id(d.Field).Int32().Tag(map[string]string{"json": jsonType(d.Field), "gorm": gormType(d)})
}
func Int64(d DescTab) jen.Code {
	return jen.Id(d.Field).Int64().Tag(map[string]string{"json": jsonType(d.Field), "gorm": gormType(d)})
}
func Uint8(d DescTab) jen.Code {
	return jen.Id(d.Field).Uint8().Tag(map[string]string{"json": jsonType(d.Field), "gorm": gormType(d)})
}
func Uint16(d DescTab) jen.Code {
	return jen.Id(d.Field).Uint16().Tag(map[string]string{"json": jsonType(d.Field), "gorm": gormType(d)})
}
func Uint32(d DescTab) jen.Code {
	return jen.Id(d.Field).Uint32().Tag(map[string]string{"json": jsonType(d.Field), "gorm": gormType(d)})
}
func Uint64(d DescTab) jen.Code {
	return jen.Id(d.Field).Uint64().Tag(map[string]string{"json": jsonType(d.Field), "gorm": gormType(d)})
}
func Float32(d DescTab) jen.Code {
	return jen.Id(d.Field).Float32().Tag(map[string]string{"json": jsonType(d.Field), "gorm": gormType(d)})
}
func Float64(d DescTab) jen.Code {
	return jen.Id(d.Field).Float64().Tag(map[string]string{"json": jsonType(d.Field), "gorm": gormType(d)})
}
func String(d DescTab) jen.Code {
	return jen.Id(d.Field).String().Tag(map[string]string{"json": jsonType(d.Field), "gorm": gormType(d)})
}
func Time(d DescTab) jen.Code {
	return jen.Id(d.Field).Qual("time", "Time").Tag(map[string]string{"json": jsonType(d.Field), "gorm": gormType(d)})
}
func Interface(d DescTab) jen.Code {
	return jen.Id(d.Field).Interface().Tag(map[string]string{"json": jsonType(d.Field), "gorm": gormType(d)})
}
