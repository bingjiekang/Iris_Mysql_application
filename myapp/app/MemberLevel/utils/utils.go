package utils

import (
	"time"
)

// 职称数据库结构体
type MemberLevel_mag struct {
	Id          int       `json:"id"`
	Name        string    `json:"name"`
	Sort        int       `json:"sort"`
	Create_time time.Time `json:"create_time"`
	Update_time time.Time `json:"update_time"`
}

// 职称添加结构体
type MemberLevelAddReq struct {
	Id   int    `form:"id"`
	Name string `form:"name" validate:"required"`
	Sort int    `json:"sort"`
}
