{{define "base"}}
id,albumID,URL,createdAt,updatedAt,deletedAt
{{end}}
{{define "Save"}}
INSERT INTO `photos`({{template "base"}})
VALUES(null,#{obj.AlbumID},#{obj.URL},#{obj.CreatedAt},#{obj.UpdatedAt},#{obj.DeletedAt})
{{end}}
{{define "SelectByID"}}
	select {{template "base"}} from  photos where `id`=#{Id}
{{end}}
{{define "SelectLimit"}}
	SELECT {{template "base"}} FROM `photos` limit #{o},#{l}
{{end}}
{{define "SelectCount"}}
	SELECT count(1) FROM `photos`
{{end}}
{{define "UpdateByID"}}
	UPDATE `photos` SET `id`=#{obj.Id}
	{{if unBlank .obj.AlbumID }},`albumID` = #{obj.AlbumID}{{end}}
{{if unBlank .obj.URL }},`URL` = #{obj.URL}{{end}}
{{if unBlank .obj.CreatedAt }},`createdAt` = #{obj.CreatedAt}{{end}}
{{if unBlank .obj.UpdatedAt }},`updatedAt` = #{obj.UpdatedAt}{{end}}
{{if unBlank .obj.DeletedAt }},`deletedAt` = #{obj.DeletedAt}{{end}}

	WHERE `id`=#{obj.Id}
{{end}}
{{define "DeleteByID"}}
	delete FROM `photos` WHERE `id`=#{Id}
{{end}}
{{define "DeleteByIDs"}}
	delete FROM `photos` WHERE `id` in{{tplfor .ids "(" ")" "," "ids"}}
{{end}}