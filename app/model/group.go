package model

const (
	GROUP_TYPE_NORMAL = 1
)

type Group struct {
	GroupId int `xorm:"pk autoincr int(11) not null comment('主键id')" form:"group_id" json:"group_id"`
	Name string `xorm:"varchar(30) not null comment('群名称')" form:"name" json:"name"`
	UserId int `xorm:"int(11) not null comment('群主id 关联用户表')" form:"user_id" json:"user_id"`
	Icon string `xorm:"varchar(255) not null default '' comment('群头像')" form:"icon" json:"icon"`
	Type int8 `xorm:"tinyint(1) not null default 1 comment('群类型 1为普通群')" form:"type" json:"type"`
	Memo string `xorm:"varchar(120) not null default '' comment('描述')" form:"memo" json:"memo"`
	CreateTime string `xorm:"datetime comment('创建时间')" form:"create_time" json:"create_time"`
}
