package model

// 登录信息
type User_information struct {
	Id       int    `xorm:not null`
	Username string `xorm:"char(100)"`
	Password string `xorm:"char(100)"`
	Email    string `xorm:"char(100)"`
}

// 个人信息
type Pserson_infomation struct {
	Id       int
	Avatar   string `xorm:"varchar(200)" form:"avatar"`                    // 头像
	Realname string `xorm:"char(100)" form:"realname" validate:"required"` // 真实姓名
	Nickname string `xorm:"char(100)" form:"nickname" validate:"required"` // 昵称
	Gender   int    `form:"gender"`                                        // 性别:1男 2女 3保密
	Mobile   string `xorm:"char(100)" form:"mobile"`                       // 手机号码
	Email    string `xorm:"char(100)" form:"email" validate:"required"`    // 邮箱地址
	Address  string `xorm:"varchar(200)" form:"address"`                   // 详细地址
	Intro    string `xorm:"varchar(300)" form:"intro"`                     // 个人简介
}
