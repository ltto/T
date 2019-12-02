package creat

import (
	"database/sql"
	"strings"
)

type tableDesc struct {
	Table    string
	SQLField string
	SQLType  string
	GoField  string
	Null     string
	Key      string
	dDefault *string
	extra    *string
}

func (t tableDesc) PRIKey() bool {
	return t.Key == "PRI"
}

var sqlTypeMap = map[string]string{
	//8(bit)
	"TINYINT": "int",
	//16(bit)
	"SMALLINT": "int",
	//34(bit)
	"MEDIUMINT": "int",
	//32(bit)
	"INT": "int", "INTEGER": "int",
	//64(bit)
	"BIGINT": "int64",
	//32(bit)
	"FLOAT": "float32",
	//64(bit)
	"DOUBLE": "float64", "DECIMAL": "float64",
	"CHAR":       "string", //0-255(byte)
	"VARCHAR":    "string", //0-65535(byte)
	"TINYBLOB":   "string", //0-255(byte)
	"TINYTEXT":   "string", //0-255(byte)
	"BLOB":       "string", //0-65535(byte)
	"TEXT":       "string", //0-65535(byte)
	"MEDIUMBLOB": "string", //0-16777215(byte)
	"MEDIUMTEXT": "string", //0-16777215(byte)
	"LONGBLOB":   "string", //0-4294967295(byte)
	"LONGTEXT":   "string", //0-4294967295(byte)

	"DATE":      "time.Time", //YYYY-MM-DD
	"TIME":      "time.Time", //HH:MM:SS
	"YEAR":      "time.Time", //YYYY
	"DATETIME":  "time.Time", //YYYY-MM-DD HH:MM:SS
	"TIMESTAMP": "time.Time", //YYYY-MM-DD HH:MM:SS
}

var sqlNullTypeMap = map[string]string{
	//8(bit)
	"TINYINT": "null.Int",
	//16(bit)
	"SMALLINT": "null.Int",
	//34(bit)
	"MEDIUMINT": "null.Int",
	//32(bit)
	"INT": "null.Int", "INTEGER": "null.Int",
	//64(bit)
	"BIGINT": "null.Int",
	//32(bit)
	"FLOAT": "null.Float",
	//64(bit)
	"DOUBLE": "null.Float", "DECIMAL": "null.Float",
	"CHAR":       "null.String", //0-255(byte)
	"VARCHAR":    "null.String", //0-65535(byte)
	"TINYBLOB":   "null.String", //0-255(byte)
	"TINYTEXT":   "null.String", //0-255(byte)
	"BLOB":       "null.String", //0-65535(byte)
	"TEXT":       "null.String", //0-65535(byte)
	"MEDIUMBLOB": "null.String", //0-16777215(byte)
	"MEDIUMTEXT": "null.String", //0-16777215(byte)
	"LONGBLOB":   "null.String", //0-4294967295(byte)
	"LONGTEXT":   "null.String", //0-4294967295(byte)

	"DATE":      "null.Time", //YYYY-MM-DD
	"TIME":      "null.Time", //HH:MM:SS
	"YEAR":      "null.Time", //YYYY
	"DATETIME":  "null.Time", //YYYY-MM-DD HH:MM:SS
	"TIMESTAMP": "null.Time", //YYYY-MM-DD HH:MM:SS
}

func ListParam(db *sql.DB, nullAble bool) map[string][]tableDesc {
	var list = make(map[string][]tableDesc, 10)
	rows, err := db.Query("SHOW TABLES")
	if err != nil {
		panic(err)
	}
	for rows.Next() {
		var table string
		if err := rows.Scan(&table); err != nil {
			panic(err)
		}
		desc, err := db.Query("desc " + table)
		if err != nil {
			panic(err)
		}
		list[table] = make([]tableDesc, 0, 16)
		for desc.Next() {
			d := tableDesc{Table: table}
			err = desc.Scan(&d.SQLField, &d.SQLType, &d.Null, &d.Key, &d.dDefault, &d.extra)
			if err != nil {
				panic(err)
			}
			index := strings.Index(d.SQLType, "(")
			if index != -1 {
				d.SQLType = d.SQLType[:index]
			}
			d.SQLType = strings.ToUpper(d.SQLType)
			if nullAble {
				d.GoField = sqlNullTypeMap[d.SQLType]
			} else {
				d.GoField = sqlTypeMap[d.SQLType]
			}
			list[table] = append(list[table], d)
		}
	}
	return list

}
func DoCreat(pkgName, dest, url string, list map[string][]tableDesc) {
	nullAble := true
	//createMapperXml(list, dest, nullAble)
	createStruct(list, pkgName, dest, nullAble)
	//createMapper(list, pkgName, dest, url)
	//createAPI(list, pkgName, dest)
	//createGormF(list, pkgName, dest, url)
	createGormM(list, pkgName, dest, url)
}

func toUp(s string) string {
	newStr := ""
	Up := false
	for _, v := range s {
		if v == '_' {
			Up = true
			continue
		}
		strCell := ""
		if Up && v != '_' {
			strCell = strings.ToUpper(string(v))
			Up = false
		} else {
			strCell = string(v)
		}
		newStr += strCell
	}
	return strings.ToUpper(newStr[0:1]) + newStr[1:]
}
