package DB

import (
	"database/sql"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func init() {
	var err error
	DB, err = sql.Open("mysql", "root:12345678@tcp(localhost:3306)/Iris?charset=utf8")
	if err != nil {
		log.Fatal(err)
	}
	// defer DB.close()
	DB.SetConnMaxLifetime(time.Minute * 3)
	//设置数据库最大连接数
	DB.SetConnMaxLifetime(100)
	//设置上数据库最大闲置连接数
	DB.SetMaxIdleConns(10)
	//验证连接
	if err := DB.Ping(); err != nil {
		log.Fatal(err)
	}
	// defer Db.Close()
}

const (
	USERNAME = "root"
	PASSWORD = "12345678"
	NETWORK  = "tcp"
	SERVER   = "localhost"
	PORT     = 3306
	DATABASE = "Iris"
)

// func StartDB(DB *sql.DB) {

// 	// defer DB.close()
// 	DB.SetConnMaxLifetime(time.Minute * 3)
// 	//设置数据库最大连接数
// 	DB.SetConnMaxLifetime(100)
// 	//设置上数据库最大闲置连接数
// 	DB.SetMaxIdleConns(10)
// 	//验证连接
// 	if err := DB.Ping(); err != nil {
// 		log.Fatal(err)
// 	}
// 	// fmt.Println("connnect success")
// }

// 插入数据
func Insert(username string, password string, email string) bool {
	Data := "insert into User_information(username,password,email) values(?,?,?)"
	_, err := DB.Exec(Data, username, password, email)
	if err != nil {
		// log.Fatal(err)
		return false
	}
	// log.Fatal("插入成功")
	return true
}

// 查询username是否在数据库里
// func Select(DB *sql.DB, username string, password string, email string) bool {

// }
// 查询用户名是否存在
func Select_user(username string) bool {
	Data := "select username,password from User_information where username = " + "'" + username + "'"
	// info, err := DB.Exec(Data, username)
	var mation SelectInfo
	// err := DB.Select(&mation,Data,username)
	err := DB.QueryRow(Data).Scan(&mation.Username, &mation.Password)
	if err != nil {
		// log.Fatal(err)
		return false
	}
	return true

}

// 查询用户名和密码是否匹配
func Select_user_pwd(username string, password string) bool {
	Data := "select username,password from User_information where username = " + "'" + username + "'"
	// info, err := DB.Exec(Data, username)
	var mation SelectInfo
	// err := DB.Select(&mation,Data,username)
	err := DB.QueryRow(Data).Scan(&mation.Username, &mation.Password)
	if err != nil || mation.Password != password {
		// log.Fatal(err)
		return false
	}
	return true
}
