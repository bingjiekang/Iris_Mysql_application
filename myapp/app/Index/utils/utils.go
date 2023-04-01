package utils

// 更新密码结构体
type UpdatePwd struct {
	Oldpassword string `form:"oldPassword"`
	Newpassword string `form:"newPassword"`
	Repassword  string `form:"rePassword"`
}

// 用户没有更新个人信息时获取内容
type UserUpadte struct {
	Realname string `xorm:"char(100)" form:"realname" validate:"required"` // 真实姓名
	Nickname string `form:"nickname"`                                      // 昵称
	Gender   int    `form:"gender"`                                        // 性别:1男 2女 3保密
	Mobile   string `xorm:"char(100)" form:"mobile"`                       // 手机号码
	Email    string `xorm:"char(100)" form:"email" validate:"required"`    // 邮箱地址
	Address  string `xorm:"varchar(200)" form:"address"`                   // 详细地址
	Intro    string `xorm:"varchar(300)" form:"intro"`                     // 个人简介
}
