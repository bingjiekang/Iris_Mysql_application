package utils

import "time"

// 角色数据库结构体
type Role_mag struct {
	Id       int       `json:"id"`
	Username string    `json:"username"`
	Status   bool      `json:"status"`
	Sort     int       `json:"sort"`
	Created  time.Time `json:"created"`
	Updated  time.Time `json:"updated"`
}

// 添加结构体
type RoleAddReq struct {
	Id     int    `form:"id"`
	Name   string `form:"name" validate:"required"`
	Status int    `form:"status" validate:"int"`
	Sort   int    `form:"sort" validate:"int"`
}
