package main

import (
	"myapp/utils"

	"github.com/kataras/iris/v12"
	_ "github.com/mojocn/base64Captcha"
)

func main() {
	// 创建一个应用
	app := iris.New()

	// 路由注册
	utils.Register(app)

	// 监听8080并运行
	app.Run(iris.Addr(":8080"))
}
