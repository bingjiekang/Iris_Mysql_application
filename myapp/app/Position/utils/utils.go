package utils

import (
	"time"
)

// 职称数据库结构体
type Position_mag struct {
	Id          int       `json:"id"`
	Name        string    `json:"name"`
	Status      bool      `json:"status"`
	Sort        int       `json:"sort"`
	Create_time time.Time `json:"create_time"`
	Update_time time.Time `json:"update_time"`
}

// 职称添加结构体
type PositionAddReq struct {
	Id     int    `form:"id"`
	Name   string `form:"name" validate:"required"`
	Status int    `form:"status" validate:"int"`
	Sort   int    `json:"sort"`
}

// 分页查询条件
type PositionPageReq struct {
	Name  string `form:"name"`  // 角色名称
	Page  int    `form:"page"`  // 页码
	Limit int    `form:"limit"` // 每页数
}
