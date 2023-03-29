package model

import (
	"fmt"

	"log"

	_ "github.com/go-sql-driver/mysql"
	"xorm.io/xorm"
)

var engine *xorm.Engine

// var User User_information

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

func Select_userpwd(username string, pwd string) bool {
	var User User_information
	has, _ := engine.Where("username = ?", username).Get(&User)
	if has == false {
		fmt.Println("取用户密码失败")
		return false
	} else {
		if User.Password == pwd {
			return true
		}
		fmt.Println("密码不正确,重新输入!", User.Password)
		return false
	}
}

func Servise_pwd(username string, pwd string) bool {
	// has, _ := engine.Where("username = ? and password = ?", username, pwd).Get(&User)
	// if has == false {
	// 	return false
	// } else {
	affected, err := engine.Exec("update User_information set password = ? where username = ?", pwd, username)
	if err != nil {
		fmt.Println("密码更新失败")
		return false
	}
	fmt.Println(affected)
	return true
	// }

}

// func Select_user(username string) bool {
// 	var User User_information
// 	// has,err := engine.Exist(new())

// 	// 查询用户是否存在
// 	// var name string
// 	// has, err := engine.Table(&User).Where("username = ?",).
// 	// has, err := engine.SQL("select * from User_information where username = ?", username).Exist()
// 	has, err := engine.Where("username = ?", username).Get(&User)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	fmt.Println(User)
// 	return has
// }
