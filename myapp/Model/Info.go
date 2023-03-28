package model

type User_information struct {
	Id       int    `xorm:not null`
	Username string `xorm:"char(100)"`
	Password string `xorm:"char(100)"`
	Email    string `xorm:"char(100)"`
}
