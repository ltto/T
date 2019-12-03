{{define "base"}}
id,user_id,total_amount,total_profit,can_use_amount,freeze,pool_total_amount,pool_amount,pool_profit,pool_freeze,accumulation_profit,integral,create_time,update_time,status,delete_flag
{{end}}
{{define "Save"}}
INSERT INTO `biz_property`({{template "base"}})
VALUES(null,#{obj.UserId},#{obj.TotalAmount},#{obj.TotalProfit},#{obj.CanUseAmount},#{obj.Freeze},#{obj.PoolTotalAmount},#{obj.PoolAmount},#{obj.PoolProfit},#{obj.PoolFreeze},#{obj.AccumulationProfit},#{obj.Integral},#{obj.CreateTime},#{obj.UpdateTime},#{obj.Status},#{obj.DeleteFlag})
{{end}}
{{define "SelectByID"}}
	select {{template "base"}} from  biz_property where `id`=#{Id}
{{end}}
{{define "SelectLimit"}}
	SELECT {{template "base"}} FROM `biz_property` limit #{o},#{l}
{{end}}
{{define "SelectCount"}}
	SELECT count(1) FROM `biz_property`
{{end}}
{{define "UpdateByID"}}
	UPDATE `biz_property` SET `id`=#{obj.Id}
	{{if unBlank .obj.UserId }},`user_id` = #{obj.UserId}{{end}}
{{if unBlank .obj.TotalAmount }},`total_amount` = #{obj.TotalAmount}{{end}}
{{if unBlank .obj.TotalProfit }},`total_profit` = #{obj.TotalProfit}{{end}}
{{if unBlank .obj.CanUseAmount }},`can_use_amount` = #{obj.CanUseAmount}{{end}}
{{if unBlank .obj.Freeze }},`freeze` = #{obj.Freeze}{{end}}
{{if unBlank .obj.PoolTotalAmount }},`pool_total_amount` = #{obj.PoolTotalAmount}{{end}}
{{if unBlank .obj.PoolAmount }},`pool_amount` = #{obj.PoolAmount}{{end}}
{{if unBlank .obj.PoolProfit }},`pool_profit` = #{obj.PoolProfit}{{end}}
{{if unBlank .obj.PoolFreeze }},`pool_freeze` = #{obj.PoolFreeze}{{end}}
{{if unBlank .obj.AccumulationProfit }},`accumulation_profit` = #{obj.AccumulationProfit}{{end}}
{{if unBlank .obj.Integral }},`integral` = #{obj.Integral}{{end}}
{{if unBlank .obj.CreateTime }},`create_time` = #{obj.CreateTime}{{end}}
{{if unBlank .obj.UpdateTime }},`update_time` = #{obj.UpdateTime}{{end}}
{{if unBlank .obj.Status }},`status` = #{obj.Status}{{end}}
{{if unBlank .obj.DeleteFlag }},`delete_flag` = #{obj.DeleteFlag}{{end}}

	WHERE `id`=#{obj.Id}
{{end}}
{{define "DeleteByID"}}
	delete FROM `biz_property` WHERE `id`=#{Id}
{{end}}
{{define "DeleteByIDs"}}
	delete FROM `biz_property` WHERE `id` in{{tplfor .ids "(" ")" "," "ids"}}
{{end}}