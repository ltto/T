package creat

import (
	"io/ioutil"
	"os"
	"path"
)

func createMapperXml(list map[string][]tableDesc, dest string, nullAble bool) {
	os.MkdirAll(path.Join(dest), 0777)
	for table, descs := range list {
		pri := "id"
		for _, v := range descs {
			if v.Key == "PRI" {
				pri = v.SQLField
			}
		}
		xml := toUp(table) + "Mapper.xml"
		xmlByte := os.Expand(mapper, func(mapperParam string) string {
			switch mapperParam {
			case ".Columns":
				var columns = ""
				for _, desc := range descs {
					columns += "`" + desc.SQLField + "`, "
				}
				return columns[0 : len(columns)-2]
			case ".TableB":
				return "`" + table + "`"
			case ".Save":
				return os.Expand(save, func(save string) string {
					switch save {
					case ".Inserts":
						var insers = ""
						for _, desc := range descs {
							if pri == desc.SQLField {
								insers += " null,"
							} else {
								insers += " #{obj." + toUp(desc.SQLField) + "},"
							}
						}
						return insers[0 : len(insers)-1]
					case ".Table":
						return table
					default:
						return ""
					}
				})
			case ".SelectByID":
				return os.Expand(selectByID, func(query string) string {
					switch query {
					case ".Table":
						return table
					case ".PRI":
						return pri
					default:
						return ""
					}
				})
			case ".UpdateByID":
				return os.Expand(updateByID, func(s string) string {
					switch s {
					case ".Table":
						return table
					case ".PRI":
						return pri
					case ".Updates":
						var updates = ""
						tmp := Updates
						if nullAble {
							tmp = UpdatesNull
						}
						for _, desc := range descs {
							updates += os.Expand(tmp, func(s string) string {
								switch s {
								case ".SqlField":
									return desc.SQLField
								case ".GoField":
									return "obj." + toUp(desc.SQLField)
								default:
									return ""
								}
							})
						}
						return updates
					default:
						return ""
					}
				})
			case ".DeleteByID":
				return os.Expand(deleteByID, func(s string) string {
					switch s {
					case ".Table":
						return table
					case ".PRI":
						return pri
					default:
						return ""
					}
				})
			case ".DeleteByIDs":
				return os.Expand(deleteByIDs, func(delete string) string {
					switch delete {
					case ".Table":
						return table
					case ".PRI":
						return pri
					default:
						return ""
					}
				})
			case ".ResultMap":
				return os.Expand(resultMap, func(resultMapParam string) string {
					if resultMapParam == ".Columns" {
						var Columns = ""
						for _, desc := range descs {
							Columns += os.Expand(column, func(result string) string {
								switch result {
								case ".TagName":
									if desc.SQLField == pri {
										return id
									}
									return noID
								case ".SqlField":
									return desc.SQLField
								case ".GoField":
									return toUp(desc.SQLField)
								case ".GoType":
									return sqlTypeMap[desc.SQLType]
								default:
									return ""
								}
							})
						}
						return Columns
					}
					return ""
				})
			default:
				return ""
			}
		})
		os.MkdirAll(path.Join(dest, "mapper"), 0777)
		if err := ioutil.WriteFile(path.Join(dest, "mapper", xml), []byte(xmlByte), 0777); err != nil {
			panic(err)
		}
	}
}

const save string = "INSERT INTO `${.Table}`(<include refid=\"base\"/>) VALUES(${.Inserts})"
const selectByID string = "SELECT\n" +
	"        <include refid=\"base\"/>\n" +
	"        FROM `${.Table}` WhERE `${.PRI}` = #{${.PRI}}"

const selectLimit string = "SELECT\n" +
	"        <include refid=\"base\"/>\n" +
	"        FROM `${.Table}` WhERE `${.PRI}` = #{${.PRI}}"

const updateByID string = "UPDATE `${.Table}` SET `${.PRI}`=#{obj.${.PRI}} \n" +
	"${.Updates}" +
	"        WHERE `${.PRI}`=#{obj.${.PRI}}"
const Updates string = "        <if test=\"${.GoField} != nil\">,`${.SqlField}` = #{${.GoField}}</if>\n"
const UpdatesNull string = "        <if test=\"${.GoField}.Valid\">,`${.SqlField}` = #{${.GoField}}</if>\n"
const deleteByID string = "delete FROM `${.Table}` WHERE `${.PRI}`=#{${.PRI}}"
const deleteByIDs string = "delete FROM `${.Table}` WHERE `${.PRI}` in\n" +
	"        <foreach item=\"item\" index=\"index\" collection=\"ids\" open=\"(\" separator=\",\" close=\")\">#{item}</foreach>"

const mapper string = `<?xml version="1.0" encoding="UTF-8"?>
<!DOCTYPE mapper PUBLIC "-//mybatis.org//DTD Mapper 3.0//EN"
        "https://raw.githubusercontent.com/ltto/GoMybatis/master/mybatis-3-mapper.dtd">
<mapper>
    <!--logic_enable 逻辑删除字段-->
    <!--logic_deleted 逻辑删除已删除字段-->
    <!--logic_undelete 逻辑删除 未删除字段-->
    <!--version_enable 乐观锁版本字段,支持int,int8,int16,int32,int64-->
    ${.ResultMap}
    <sql id="base">
        ${.Columns}
    </sql>
    <!-- insert -->
    <insert id="Save" useGeneratedKeys="true">
        ${.Save}
    </insert>
    <!-- selectOne -->
    <select id="SelectByID">
        ${.SelectByID}
    </select>

    <select id="SelectLimit">
        SELECT
        <include refid="base"/>
        FROM ${.TableB}
        limit #{o},#{l}
    </select>

	<select id="SelectCount">
        SELECT count(1) FROM ${.TableB}
    </select>

    <!-- update -->
    <update id="UpdateByID">
        ${.UpdateByID}
    </update>

    <!--delete one-->
    <delete id="DeleteByID">
        ${.DeleteByID}
    </delete>

    <!-- delete list -->
    <delete id="DeleteByIDs">
        ${.DeleteByIDs}
    </delete>
</mapper>
`
const resultMap string = `<resultMap id="BaseResultMap" tables="biz_user_address">
${.Columns}
    </resultMap>
`
const column = "        <${.TagName} column=\"${.SqlField}\" property=\"${.GoField}\" langType=\"${.GoType}\"/>\n"
const id = "id"
const noID = "result"
