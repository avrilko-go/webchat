package model

//性别定义
var (
	SexMan = 1 // 男
	SexWoMan = 2 // 女
	SexUnKnown =  3 // 未知
)

// 在线状态定义
var (
	Online = 1 // 在线
	UnOnline = 2 // 离线
)


type User struct {
	UserId int `xorm:"pk not null autoincr int(11) comment('主键id') " form:"user_id" json:"user_id"`
	Mobile string `xorm:"varchar(20) not null default '' comment('手机号码') " form:"mobile" json:"mobile"`
	Password string `xorm:"varchar(40) not null default '' comment('密码') " form:"password" json:"password"`
	Avatar string `xorm:"varchar(255) not null default '' comment('头像') " form:"avatar" json:"avatar"`
	Sex int8 `xorm:"tinyint(1) not null default 1 comment('性别 1为男 2为女 3为未知')" json:"sex" form:"sex"`
	NickName string `xorm:"varchar(20) not null default '' comment('昵称')" json:"nick_name" form:"nick_name"`
	Salt string `xorm:"varchar(10) not null default '' comment('密码加密字符串')" json:"salt" form:"salt"`
	Online int8 `xorm:"tinyint(1) not null default 1 comment('是否在线 1为在线 2为离线')" json:"online" form:"online"`
	Token string `xorm:"varchar(40) not null default '' comment('客户端鉴权唯一凭证')" json:"token" form:"token"`
	Memo string `xorm:"varchar(140) not null default '' comment('暂时不知道')" json:"memo" form:"memo"`
	CreateTime string `xorm:"datetime comment('创建时间')" json:"create_time" form:"create_time"`
}


