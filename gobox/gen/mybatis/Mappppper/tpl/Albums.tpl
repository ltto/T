{{define "base"}}
    cid,userID,name,URL,createdAt,updatedAt,deletedAt
{{end}}
{{define "Save"}}
    INSERT INTO `albums`({{template "base"}})
    VALUES(null,#{obj.UserID},#{obj.Name},#{obj.URL},#{obj.CreatedAt},#{obj.UpdatedAt},#{obj.DeletedAt})
{{end}}
{{define "SelectByID"}}
    select {{template "base"}} from  albums where `cid`=#{Cid}
{{end}}
{{define "SelectLimit"}}
    SELECT {{template "base"}} FROM `albums` limit #{o},#{l}
{{end}}
{{define "SelectCount"}}
    SELECT count(1) FROM `albums`
{{end}}
{{define "UpdateByID"}}
    UPDATE `albums` SET `cid`=#{obj.Cid}
    {{if unBlank .obj.UserID }},`userID` = #{obj.UserID}{{end}}
    {{if unBlank .obj.Name }},`name` = #{obj.Name}{{end}}
    {{if unBlank .obj.URL }},`URL` = #{obj.URL}{{end}}
    {{if unBlank .obj.CreatedAt }},`createdAt` = #{obj.CreatedAt}{{end}}
    {{if unBlank .obj.UpdatedAt }},`updatedAt` = #{obj.UpdatedAt}{{end}}
    {{if unBlank .obj.DeletedAt }},`deletedAt` = #{obj.DeletedAt}{{end}}

    WHERE `cid`=#{obj.Cid}
{{end}}
{{define "DeleteByID"}}
    delete FROM `albums` WHERE `cid`=#{Cid}
{{end}}
{{define "DeleteByIDs"}}
    delete FROM `albums` WHERE `cid` in{{tplfor .ids "(" ")" "," "ids"}}
{{end}}