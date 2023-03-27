package DB

import (
	"fmt"
	_ "log"
	"time"

	"github.com/go-redis/redis"
)

var Client = redis.NewClient(&redis.Options{
	Addr:     "localhost:6379", // redis地址
	Password: "",               // redis密码，没有则留空
	DB:       0,                // 默认数据库，默认是0
})

// 链接redis
func Line_Redis() {
	_, err := Client.Ping().Result() // 连接验证
	if err != nil {
		fmt.Println("连接失败")
	}
}

// 存入数据
func Set_Redis(keys string, values string) bool {
	err := Client.Set(keys, values, 180*time.Second)
	if err != nil {
		fmt.Println(err)
		return false
	}
	return true
}

// 取数据
func Get_Redis(keys string) string {
	val, err := Client.Get(keys).Result()
	if err != nil {
		return "nil"
	} else {
		return val
	}
}

// 已成功注册,删除验证数据
func Del_Redis(keys string) string {
	_, err := Client.Del(keys).Result()
	if err != nil {
		return "NO"
	} else {
		return "OK"
	}
}
