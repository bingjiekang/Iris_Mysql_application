package model

import (
	"fmt"
	"time"

	"log"
	"myapp/app/Index/utils"
	level_utils "myapp/app/Level/utils"
	member_level "myapp/app/MemberLevel/utils"
	position_utils "myapp/app/Position/utils"
	role_utils "myapp/app/Role/utils"
	user_utils "myapp/app/User/utils"

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
	// 初始化role_mag数据库
	err = engine.Sync(new(Role_mag))
	if err != nil {
		fmt.Println("初始化角色管理数据库错误")
		log.Fatal(err)
	}

	// 初始化level数据库
	err = engine.Sync(new(Level))
	if err != nil {
		fmt.Println("初始化职级管理数据库错误")
		log.Fatal(err)
	}

	// 初始化position数据库
	err = engine.Sync(new(Positions))
	if err != nil {
		fmt.Println("初始化岗位管理数据库错误")
		log.Fatal(err)
	}

	// 初始化memberlevel数据库
	err = engine.Sync(new(MemberLevel))
	if err != nil {
		fmt.Println("初始化会员等级数据库错误")
		log.Fatal(err)
	}

	// 初始化users数据库
	err = engine.Sync(new(Users))
	if err != nil {
		fmt.Println("初始化users数据库错误")
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

// 查询用户角色信息
func Select_role() (role_mag []Role_mag) {
	err := engine.Find(&role_mag)
	if err != nil {
		fmt.Print("查询role message失败!")
	}
	return
}

// 根据角色名称返回限定的信息
func Select_role_limit(name string) (role_mag []Role_mag) {
	err := engine.Where("username like ?", "%"+name+"%").Find(&role_mag)
	if err != nil {
		fmt.Print("查询限制的roel_mag失败!")
	}
	return
}

// 查询角色是否已存在
func Select_role_exit(username string) bool {
	has, err := engine.SQL("select * from role_mag where username  = ?", username).Exist()
	if err != nil {
		fmt.Println(err)
	}
	return has
}

// 根据角色id查询角色信息
func Select_id(id int) Role_mag {
	var role_mag Role_mag
	_, err := engine.Where("id = ?", id).Get(&role_mag)
	if err != nil {
		fmt.Println(err)
	}
	return role_mag

}

// 向role_mag数据库加入信息
func Insert_role_mag(data *role_utils.RoleAddReq) bool {
	affected, err := engine.Exec("insert into role_mag(username,status,sort,created,updated) value(?,?,?,?,?)", data.Name, data.Status, data.Sort, time.Now().Format("2006-01-02 15:04:05"), time.Now().Format("2006-01-02 15:04:05"))
	if err != nil {
		fmt.Println("角色信息添加失败")
		return false
	}
	fmt.Println(affected)
	return true
}

// 更新role_mag信息
func Update_role_mag(data *role_utils.RoleAddReq) bool {
	affected, err := engine.Exec("update role_mag set username=?,status=?,sort=?,updated=? where id = ?", data.Name, data.Status, data.Sort, time.Now().Format("2006-01-02 15:04:05"), data.Id)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(affected)
	return true

}

// 删除role_mag指定id的用户信息
func Delete_role_mag(id int) bool {
	affected, err := engine.Table("role_mag").Where("id = ?", id).Delete()
	if err != nil {
		fmt.Println(err)
		return false
	}
	fmt.Println(affected)
	return true
}

// 查询职称角色信息
func Select_level() (level []Level) {
	err := engine.Find(&level)
	if err != nil {
		fmt.Print("查询level职称信息失败!")
	}
	return
}

// 根据职称名称返回限定的信息
func Select_level_limit(name string) (level []Level) {
	err := engine.Where("name like ?", "%"+name+"%").Find(&level)
	if err != nil {
		fmt.Print("查询限制的level失败!")
	}
	return
}

// 查询职称是否已存在
func Select_level_exit(name string) bool {
	has, err := engine.SQL("select * from level where name  = ?", name).Exist()
	if err != nil {
		fmt.Println(err)
	}
	return has
}

// 根据职称id查询角色信息
func Select_level_id(id int) Level {
	var level Level
	_, err := engine.Where("id = ?", id).Get(&level)
	if err != nil {
		fmt.Println(err)
	}
	return level

}

// 向level数据库加入信息
func Insert_level_mag(data *level_utils.LevelAddReq) bool {
	affected, err := engine.Exec("insert into level(name,status,sort,create_time,update_time) value(?,?,?,?,?)", data.Name, data.Status, data.Sort, time.Now().Format("2006-01-02 15:04:05"), time.Now().Format("2006-01-02 15:04:05"))
	if err != nil {
		fmt.Println("职称信息添加失败")
		return false
	}
	fmt.Println(affected)
	return true
}

// 更新level信息
func Update_level_mag(data *level_utils.LevelAddReq) bool {
	affected, err := engine.Exec("update level set name=?,status=?,sort=?,update_time=? where id = ?", data.Name, data.Status, data.Sort, time.Now().Format("2006-01-02 15:04:05"), data.Id)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(affected)
	return true

}

// 删除level指定id的用户信息
func Delete_level_mag(id int) bool {
	affected, err := engine.Table("level").Where("id = ?", id).Delete()
	if err != nil {
		fmt.Println(err)
		return false
	}
	fmt.Println(affected)
	return true
}

// 岗位position
// 查询岗位信息
func Select_position() (position []Positions) {
	err := engine.Find(&position)
	if err != nil {
		fmt.Print("查询position岗位信息失败!")
	}
	return
}

// 根据岗位名称返回限定的信息
func Select_position_limit(name string) (position []Positions) {
	err := engine.Where("name like ?", "%"+name+"%").Find(&position)
	if err != nil {
		fmt.Print("查询限制的岗位信息失败!")
	}
	return
}

// 查询岗位是否已存在
func Select_position_exit(name string) bool {
	has, err := engine.SQL("select * from positions where name  = ?", name).Exist()
	if err != nil {
		fmt.Println(err)
	}
	return has
}

// 根据岗位id查询岗位信息
func Select_position_id(id int) Positions {
	var position Positions
	_, err := engine.Where("id = ?", id).Get(&position)
	if err != nil {
		fmt.Println(err)
	}
	return position

}

// 向position数据库加入信息
func Insert_position_mag(data *position_utils.PositionAddReq) bool {
	affected, err := engine.Exec("insert into positions(name,status,sort,create_time,update_time) value(?,?,?,?,?)", data.Name, data.Status, data.Sort, time.Now().Format("2006-01-02 15:04:05"), time.Now().Format("2006-01-02 15:04:05"))
	if err != nil {
		fmt.Println("岗位信息添加失败", err)
		return false
	}
	fmt.Println(affected)
	return true
}

// 更新position信息
func Update_position_mag(data *position_utils.PositionAddReq) bool {
	affected, err := engine.Exec("update positions set name=?,status=?,sort=?,update_time=? where id = ?", data.Name, data.Status, data.Sort, time.Now().Format("2006-01-02 15:04:05"), data.Id)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(affected)
	return true

}

// 删除position指定id的用户信息
func Delete_position_mag(id int) bool {
	affected, err := engine.Table("positions").Where("id = ?", id).Delete()
	if err != nil {
		fmt.Println(err)
		return false
	}
	fmt.Println(affected)
	return true
}

// 查询会员等级
func Select_memberlevel() (memberlevel []MemberLevel) {
	err := engine.Find(&memberlevel)
	if err != nil {
		fmt.Print("查询memberlevel信息失败!")
	}
	return
}

// 查询会员等级是否已存在
func Select_memberlevel_exit(name string) bool {
	has, err := engine.SQL("select * from member_level where name  = ?", name).Exist()
	if err != nil {
		fmt.Println(err)
	}
	return has
}

// 根据会员id查询等级信息
func Select_memberlevel_id(id int) MemberLevel {
	var memberlevel MemberLevel
	_, err := engine.Where("id = ?", id).Get(&memberlevel)
	if err != nil {
		fmt.Println(err)
	}
	return memberlevel

}

// 向memberlevel数据库加入信息
func Insert_memberlevel_mag(data *member_level.MemberLevelAddReq) bool {
	affected, err := engine.Exec("insert into member_level(name,sort,create_time,update_time) value(?,?,?,?)", data.Name, data.Sort, time.Now().Format("2006-01-02 15:04:05"), time.Now().Format("2006-01-02 15:04:05"))
	if err != nil {
		fmt.Println("会员等级信息添加失败", err)
		return false
	}
	fmt.Println(affected)
	return true
}

// 更新会员等级信息
func Update_memberlevel_mag(data *member_level.MemberLevelAddReq) bool {
	affected, err := engine.Exec("update member_level set name=?,sort=?,update_time=? where id = ?", data.Name, data.Sort, time.Now().Format("2006-01-02 15:04:05"), data.Id)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(affected)
	return true

}

// 删除memberlevel指定id的用户信息
func Delete_memberlevel_mag(id int) bool {
	affected, err := engine.Table("member_level").Where("id = ?", id).Delete()
	if err != nil {
		fmt.Println(err)
		return false
	}
	fmt.Println(affected)
	return true
}

// 查询所有角色name
func Select_all_role() (role []string) {
	engine.SQL("select username from role_mag").Find(&role)
	return
}

// 查询所有level的name
func Select_all_level() (level []string) {
	engine.SQL("select name from level").Find(&level)
	return
}

// 查询所有position的name
func Select_all_position() (position []string) {
	engine.SQL("select name from positions").Find(&position)
	return
}

// 查询user信息并返回
func Select_user() (users []Users) {
	err := engine.Find(&users)
	if err != nil {
		fmt.Print("查询users信息失败!")
	}
	return
}

// 查询user的限制信息
func Select_user_limit(name string, gender int) (users []Users) {
	err := engine.Where("realname like ?", "%"+name+"%").And("gender = ?", gender).Find(&users)
	if err != nil {
		fmt.Print("查询users限定的信息失败!")
	}
	return

}

// 查询user的昵称nickname是否已存在
func Select_user_exit(name string) bool {
	has, err := engine.SQL("select * from users where nickname  = ?", name).Exist()
	if err != nil {
		fmt.Println(err)
	}
	return has
}

// 根据id查询用户信息
func Select_user_id(id int) Users {
	var users Users
	_, err := engine.Where("id = ?", id).Get(&users)
	if err != nil {
		fmt.Println(err)
	}
	return users

}

// 向users数据库加入信息
func Insert_user_mag(data *user_utils.UserAddReq) bool {
	affected, err := engine.Exec("insert into users(realname,gender,nickname,password,status,level_name,position_name,role_name,mobile,email,address,sort,note,create_time,update_time) value(?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)", data.Realname, data.Gender, data.Nickname, data.Password, data.Status, data.LevelName, data.PositionName, data.RoleName, data.Mobile, data.Email, data.Address, data.Sort, data.Note, time.Now().Format("2006-01-02 15:04:05"), time.Now().Format("2006-01-02 15:04:05"))
	if err != nil {
		fmt.Println("用户信息添加失败", err)
		return false
	}
	fmt.Println(affected)
	return true
}

// 更新users信息
func Update_user_mag(data *user_utils.UserAddReq) bool {
	affected, err := engine.Exec("update users set realname=?,gender=?,nickname=?,password=?,status=?,level_name=?,position_name=?,role_name=?,mobile=?,email=?,address=?,sort=?,note=?,update_time=? where id = ?", data.Realname, data.Gender, data.Nickname, data.Password, data.Status, data.LevelName, data.PositionName, data.RoleName, data.Mobile, data.Email, data.Address, data.Sort, data.Note, time.Now().Format("2006-01-02 15:04:05"), data.Id)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(affected)
	return true

}

// 删除users指定id的用户信息
func Delete_user_mag(id int) bool {
	affected, err := engine.Table("users").Where("id = ?", id).Delete()
	if err != nil {
		fmt.Println(err)
		return false
	}
	fmt.Println(affected)
	return true
}
