{{define "base"}}
id,name,uuid,pc_link,h5_link,remark,version,create_time,delete_flag
{{end}}
{{define "Save"}}
INSERT INTO `biz_activity`({{template "base"}})
VALUES(null,#{obj.Name},#{obj.Uuid},#{obj.PcLink},#{obj.H5Link},#{obj.Remark},#{obj.Version},#{obj.CreateTime},#{obj.DeleteFlag})
{{end}}
{{define "SelectByID"}}
	select {{template "base"}} from  biz_activity where `id`=#{Id}
{{end}}
{{define "SelectLimit"}}
	SELECT {{template "base"}} FROM `biz_activity` limit #{o},#{l}
{{end}}
{{define "SelectCount"}}
	SELECT count(1) FROM `biz_activity`
{{end}}
{{define "UpdateByID"}}
	UPDATE `biz_activity` SET `id`=#{obj.Id}
	{{if unBlank .obj.Name }},`name` = #{obj.Name}{{end}}
{{if unBlank .obj.Uuid }},`uuid` = #{obj.Uuid}{{end}}
{{if unBlank .obj.PcLink }},`pc_link` = #{obj.PcLink}{{end}}
{{if unBlank .obj.H5Link }},`h5_link` = #{obj.H5Link}{{end}}
{{if unBlank .obj.Remark }},`remark` = #{obj.Remark}{{end}}
{{if unBlank .obj.Version }},`version` = #{obj.Version}{{end}}
{{if unBlank .obj.CreateTime }},`create_time` = #{obj.CreateTime}{{end}}
{{if unBlank .obj.DeleteFlag }},`delete_flag` = #{obj.DeleteFlag}{{end}}

	WHERE `id`=#{obj.Id}
{{end}}
{{define "DeleteByID"}}
	delete FROM `biz_activity` WHERE `id`=#{Id}
{{end}}
{{define "DeleteByIDs"}}
	delete FROM `biz_activity` WHERE `id` in{{tplfor .ids "(" ")" "," "ids"}}
{{end}}