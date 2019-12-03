{{define "base"}}
id,nick,openID,createdAt,updatedAt,deletedAt
{{end}}
{{define "Save"}}
INSERT INTO `users`({{template "base"}})
VALUES(null,#{obj.Nick},#{obj.OpenID},#{obj.CreatedAt},#{obj.UpdatedAt},#{obj.DeletedAt})
{{end}}
{{define "SelectByID"}}
	select {{template "base"}} from  users where `id`=#{Id}
{{end}}
{{define "SelectLimit"}}
	SELECT {{template "base"}} FROM `users` limit #{o},#{l}
{{end}}
{{define "SelectCount"}}
	SELECT count(1) FROM `users`
{{end}}
{{define "UpdateByID"}}
	UPDATE `users` SET `id`=#{obj.Id}
	{{if unBlank .obj.Nick }},`nick` = #{obj.Nick}{{end}}
{{if unBlank .obj.OpenID }},`openID` = #{obj.OpenID}{{end}}
{{if unBlank .obj.CreatedAt }},`createdAt` = #{obj.CreatedAt}{{end}}
{{if unBlank .obj.UpdatedAt }},`updatedAt` = #{obj.UpdatedAt}{{end}}
{{if unBlank .obj.DeletedAt }},`deletedAt` = #{obj.DeletedAt}{{end}}

	WHERE `id`=#{obj.Id}
{{end}}
{{define "DeleteByID"}}
	delete FROM `users` WHERE `id`=#{Id}
{{end}}
{{define "DeleteByIDs"}}
	delete FROM `users` WHERE `id` in{{tplfor .ids "(" ")" "," "ids"}}
{{end}}