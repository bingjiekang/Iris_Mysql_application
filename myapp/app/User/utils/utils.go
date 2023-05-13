package utils

import "time"

// 用户数据库结构体
type User_mag struct {
	Id           int         `json:"id"`
	Realname     string      `json:"realname"` //真实姓名
	Gender       int         `json:"gender"`   //性别
	Nickname     string      `json:"nickname"` // 昵称
	Password     string      `json:"password"`
	Status       bool        `json:"status"`
	LevelName    interface{} `json:"levelName"`    //职称名
	PositionName interface{} `json:"positionName"` //岗位名
	RoleName     interface{} `json:"RoleName"`     //角色名
	Mobile       string      `json:"mobile"`
	Email        string      `json:"email"`
	Address      string      `json:"address"`
	Sort         int         `json:"sort"` //序号
	Note         string      `json:"note"` //备注
	Create_time  time.Time   `json:"create_time"`
	Update_time  time.Time   `json:"update_time"`
}

// 添加结构体
type UserAddReq struct {
	Id           int    `form:"id"`
	Realname     string `form:"realname"` //真实姓名
	Gender       int    `form:"gender"`   //性别
	Nickname     string `form:"nickname"` // 昵称
	Password     string `form:"password"`
	Status       bool   `form:"status"`
	LevelName    string `form:"levelname"`    //职称名
	PositionName string `form:"positionname"` //岗位名
	RoleName     string `form:"rolename"`     //角色名
	Mobile       string `form:"mobile"`
	Email        string `form:"email"`
	Address      string `form:"address"`
	Sort         int    `form:"sort"` //序号
	Note         string `form:"note"` //备注
}

// 用户分页查询条件
type UserPageReq struct {
	Name   string `form:"name"`   // 用户名
	Gender int    `form:"gender"` // 性别
	Page   int    `form:"page"`   // 页码
	Limit  int    `form:"limit"`  // 每页数
}
