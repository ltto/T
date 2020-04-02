package creat

import (
	"io/ioutil"
	"os"
	"path"
)

func createMapperTpl(list map[string][]tableDesc, dest string) {
	os.MkdirAll(path.Join(dest), 0777)
	for table, descs := range list {
		var PRIKey tableDesc
		base := ""
		updates := ""
		saves := ""
		for _, desc := range descs {
			if desc.PRIKey() {
				PRIKey = desc
				saves += "null"
			} else {
				updates += os.Expand(tplUpdates, func(s string) string {
					switch s {
					case "field":
						return toUp(desc.SQLField)
					case "sqlField":
						return desc.SQLField
					}
					return "#{" + s + "}"
				})
				saves += "#{obj." + toUp(desc.SQLField) + "}"
			}
			saves += ","
			base += desc.SQLField + ","
		}
		saves = saves[0 : len(saves)-1]
		base = base[0 : len(base)-1]
		tpl := os.Expand(TPL, func(s string) string {
			switch s {
			case "base":
				return base
			case "saves":
				return saves
			case "tplUpdates":
				return updates
			case "table":
				return table
			case "sqlPRIKey":
				return PRIKey.SQLField
			case "PRIKey":
				return toUp(PRIKey.SQLField)

			}
			return ""
		})
		os.MkdirAll(path.Join(dest, "tpl"), 0777)
		filePath := path.Join(dest, "tpl", toUp(table)+"Mapper.gohtml")
		os.Remove(filePath)
		if err := ioutil.WriteFile(filePath, []byte(tpl), 0777); err != nil {
			panic(err)
		}
	}
}

const TPL string = "{{define \"base\"}}\n" +
	"${base}\n" +
	"{{end}}\n" +
	"{{define \"Save\"}}\n" +
	"INSERT INTO `${table}`({{template \"base\"}})\n" +
	"VALUES(${saves})\n" +
	"{{end}}\n" +
	"{{define \"SelectByID\"}}\n" +
	"	select {{template \"base\"}} from  ${table} where `${sqlPRIKey}`=#{${PRIKey}}\n" +
	"{{end}}\n" +
	"{{define \"SelectLimit\"}}\n" +
	"	SELECT {{template \"base\"}} FROM `${table}` limit #{o},#{l}\n" +
	"{{end}}\n" +
	"{{define \"SelectCount\"}}\n" +
	"	SELECT count(1) FROM `${table}`\n" +
	"{{end}}\n" +
	"{{define \"UpdateByID\"}}\n" +
	"	UPDATE `${table}` SET `${sqlPRIKey}`=#{obj.${PRIKey}}\n" +
	"	${tplUpdates}\n" +
	"	WHERE `${sqlPRIKey}`=#{obj.${PRIKey}}\n" +
	"{{end}}\n" +
	"{{define \"DeleteByID\"}}\n" +
	"	delete FROM `${table}` WHERE `${sqlPRIKey}`=#{${PRIKey}}\n" +
	"{{end}}\n" +
	"{{define \"DeleteByIDs\"}}\n" +
	"	delete FROM `${table}` WHERE `${sqlPRIKey}` in{{tplfor .ids \"(\" \")\" \",\" \"ids\"}}\n" +
	"{{end}}"
const tplUpdates = "{{if unBlank .obj.${field} }},`${sqlField}` = #{obj.${field}}{{end}}\n"
