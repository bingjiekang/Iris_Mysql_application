package model

import (
	"fmt"

	"log"

	_ "github.com/go-sql-driver/mysql"
	"xorm.io/xorm"
)

var engine *xorm.Engine

func init() {
	var err error
	engine, err = xorm.NewEngine("mysql", "root:12345678@tcp(localhost:3306)/Iris?charset=utf8")
	if err != nil {
		fmt.Println("连接数据库错误")
		log.Fatal(err)
	}

	// err = engine.Sync(new(User_information))
	// if err != nil {
	// 	fmt.Println("创建错误")
	// 	log.Fatal(err)
	// }

}

func Select_user(username string) bool {
	var User User_information
	// has,err := engine.Exist(new())

	// 查询用户是否存在
	// var name string
	// has, err := engine.Table(&User).Where("username = ?",).
	// has, err := engine.SQL("select * from User_information where username = ?", username).Exist()
	has, err := engine.Where("username = ?", username).Get(&User)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(User)
	return has
}
