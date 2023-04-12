package model

import (
	"fmt"

	"log"
	"myapp/app/Index/utils"

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
	err = engine.Sync(new(Role_mag))
	if err != nil {
		fmt.Println("初始化角色管理错误")
		log.Fatal(err)
	}
	// err = engine.Sync(new(User_information))
	// if err != nil {
	// 	fmt.Println("创建错误")
	// 	log.Fatal(err)
	// }

}

// 获取用户密码
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

// 更新密码操作
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

// 查询用户邮箱
func Select_Email(username string) string {
	var User User_information
	has, _ := engine.Where("username = ?", username).Get(&User)
	// fmt.Println("查询到的email", User.Email)
	if has == false {
		fmt.Println("查询用户邮箱失败")
		return "nil"
	} else {
		return User.Email
	}
}

// 加入用户信息
func Insert_info(data *utils.UserUpadte) bool {
	affected, err := engine.Exec("insert into Person_information(realname,nickname,gender,mobile,email,address,intro) value(?,?,?,?,?,?,?)", data.Realname, data.Nickname, data.Gender, data.Mobile, data.Email, data.Address, data.Intro)
	if err != nil {
		fmt.Println("用户个人信息更新失败")
		return false
	}
	fmt.Println(affected)
	return true
}

// 判断用户是否已经更新过个人信息
func Select_info(username string) bool {
	has, err := engine.SQL("select * from Person_information where nickname = ?", username).Exist()
	if err != nil {
		fmt.Println(err)
	}
	return has
}

// 如果已经存在则直接更新
func Update_info(nickname string, data *utils.UserUpadte) bool {
	affected, err := engine.Exec("update Person_information set realname=?,nickname=?,gender=?,mobile=?,email=?,address=?,intro=? where nickname=?", data.Realname, data.Nickname, data.Gender, data.Mobile, data.Email, data.Address, data.Intro, nickname)
	if err != nil {
		fmt.Println("用户个人信息更新失败...")
		return false
	}
	fmt.Println(affected)
	return true

}

// 从Person_information中获取信息
func Select_Personinfo(nickname string) (info utils.UserUpadte) {
	var data Person_information
	has, err := engine.Where("nickname = ?", nickname).Get(&data)
	if !has {
		fmt.Println("从Person_information中获取信息失败", err)
	} else {
		fmt.Println("准备返回信息")

		info.Realname = data.Realname
		info.Nickname = data.Nickname
		info.Gender = data.Gender
		info.Mobile = data.Mobile
		info.Email = data.Email
		info.Address = data.Address
		info.Intro = data.Intro
		fmt.Println(info, "返回的信息")
	}
	return
}

// 更新User_information里的用户名字和email
func Update_userinfo(oldname string, newname string, email string) bool {
	affected, err := engine.Exec("update User_information set username=?,email=? where username = ?", newname, email, oldname)
	if err != nil {
		fmt.Println("昵称用户更新失败")
		return false
	}
	fmt.Println(affected)
	return true
}

// // 查询用户角色信息
// func Select_role()

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
