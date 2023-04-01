package model

// 登录信息
type User_information struct {
	Id       int    `xorm:not null`
	Username string `xorm:"char(100)"`
	Password string `xorm:"char(100)"`
	Email    string `xorm:"char(100)"`
}

// 个人信息
type Person_information struct {
	Id       int
	Avatar   string `xorm:"varchar(200)"` // 头像
	Realname string `xorm:"char(100)"`    // 真实姓名
	Nickname string `xorm:"char(100)"`    // 昵称
	Gender   int    // 性别:1男 2女 3保密
	Mobile   string `xorm:"char(100)"`    // 手机号码
	Email    string `xorm:"char(100)"`    // 邮箱地址
	Address  string `xorm:"varchar(200)"` // 详细地址
	Intro    string `xorm:"varchar(300)"` // 个人简介
}
