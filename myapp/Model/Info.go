package model

import (
	"time"
)

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

// 角色信息数据库
type Role_mag struct {
	Id       int    `xorm:"pk autoincr"`
	Username string `xorm:"varchar(200)"`
	Status   bool
	Sort     int
	Created  time.Time `xorm:"created"`
	Updated  time.Time `xorm:"updated"`
}

// 职级信息数据库
type Level struct {
	Id          int    `xorm:"pk autoincr"`
	Name        string `xorm:"varchar(200)"`
	Status      bool
	Sort        int
	Create_time time.Time `xorm:"created"`
	Update_time time.Time `xorm:"updated"`
}

// 岗位信息数据库
type Positions struct {
	Id          int    `xorm:"pk autoincr"`
	Name        string `xorm:"varchar(200)"`
	Status      bool
	Sort        int
	Create_time time.Time `xorm:"created"`
	Update_time time.Time `xorm:"updated"`
}

// 会员等级数据库
type MemberLevel struct {
	Id          int    `xorm:"pk autoincr"`
	Name        string `xorm:"varchar(200)"`
	Sort        int
	Create_time time.Time `xorm:"created"`
	Update_time time.Time `xorm:"updated"`
}

// 用户数据库结构体
type Users struct {
	Id           int       `xorm:"pk autoincr"`
	Realname     string    `xorm:"varchar(200)"` //真实姓名
	Gender       int       //性别
	Nickname     string    `xorm:"varchar(200)"` // 昵称
	Password     string    `xorm:"varchar(200)"`
	Status       bool      `xorm:"varchar(200)"`
	LevelName    string    `xorm:"varchar(200)"` //职称名
	PositionName string    `xorm:"varchar(200)"` //岗位名
	RoleName     string    `xorm:"varchar(200)"` //角色名
	Mobile       string    `xorm:"varchar(200)"`
	Email        string    `xorm:"varchar(200)"`
	Address      string    `xorm:"varchar(200)"`
	Sort         int       //序号
	Note         string    `xorm:"varchar(400)"` //备注
	Create_time  time.Time `json:"created"`
	Update_time  time.Time `json:"updated"`
}
