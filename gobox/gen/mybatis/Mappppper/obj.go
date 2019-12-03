package Mappppper

import (
	"github.com/guregu/null"
)

type Albums struct {
	Cid       null.Int    `json:"cid"`
	UserID    null.Int    `json:"userID"`
	Name      null.String `json:"name"`
	URL       null.String `json:"URL"`
	CreatedAt null.Time   `json:"createdAt"`
	UpdatedAt null.Time   `json:"updatedAt"`
	DeletedAt null.Time   `json:"deletedAt"`
}
type BizProperty struct {
	Id                 null.String `json:"id"`
	UserId             null.String `json:"user_id"`
	TotalAmount        null.Float  `json:"total_amount"`
	TotalProfit        null.Float  `json:"total_profit"`
	CanUseAmount       null.Float  `json:"can_use_amount"`
	Freeze             null.Float  `json:"freeze"`
	PoolTotalAmount    null.Float  `json:"pool_total_amount"`
	PoolAmount         null.Float  `json:"pool_amount"`
	PoolProfit         null.Float  `json:"pool_profit"`
	PoolFreeze         null.Float  `json:"pool_freeze"`
	AccumulationProfit null.Float  `json:"accumulation_profit"`
	Integral           null.Float  `json:"integral"`
	CreateTime         null.Time   `json:"create_time"`
	UpdateTime         null.Time   `json:"update_time"`
	Status             null.Int    `json:"status"`
	DeleteFlag         null.Int    `json:"delete_flag"`
}
type Photos struct {
	Id        null.Int    `json:"id"`
	AlbumID   null.Int    `json:"albumID"`
	URL       null.String `json:"URL"`
	CreatedAt null.Time   `json:"createdAt"`
	UpdatedAt null.Time   `json:"updatedAt"`
	DeletedAt null.Time   `json:"deletedAt"`
}
type BizActivity struct {
	Id         null.String `json:"id"`
	Name       null.String `json:"name"`
	Uuid       null.String `json:"uuid"`
	PcLink     null.String `json:"pc_link"`
	H5Link     null.String `json:"h5_link"`
	Remark     null.String `json:"remark"`
	Version    null.Int    `json:"version"`
	CreateTime null.Time   `json:"create_time"`
	DeleteFlag null.Int    `json:"delete_flag"`
}
type Users struct {
	Id        null.Int    `json:"id"`
	Nick      null.String `json:"nick"`
	OpenID    null.String `json:"openID"`
	CreatedAt null.Time   `json:"createdAt"`
	UpdatedAt null.Time   `json:"updatedAt"`
	DeletedAt null.Time   `json:"deletedAt"`
}
