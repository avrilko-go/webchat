package model

const (
	CONCAT_TYPE_USER = 1  //用户
	CONCAT_TYPE_GROUP = 2 //群组
)

type Contact struct {
	ContactId int `xorm:"pk autoincr int(11) comment('主键id')" form:"contact_id" json:"contact_id"`
	UserId int `xorm:"int(11) not null comment('好友id 关联用户表')" form:"user_id" json:"user_id"`
	AddId int `xorm:"int(11) not null comment('加入的id 可以为好友或者群id')" form:"add_id" json:"add_id"`
	Type int8 `xorm:"tinyint(1) not null default 1 comment('类型 1为好友 2为群')" form:"type" json:"type"`
	Memo string `xorm:"varchar(120) not null default '' comment('描述')" form:"memo" json:"memo"`
	CreateTime string `xorm:"datetime comment('创建时间')" form:"create_time" json:"create_time"`
} 