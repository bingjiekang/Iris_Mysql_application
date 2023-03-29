package main

import (
	"myapp/utils"

	// "github.com/kataras/iris/v12/sessions"
	"github.com/kataras/iris/v12"
	// "github.com/kataras/iris/v12/sessions"
	_ "github.com/mojocn/base64Captcha"
)

func main() {
	// 创建一个应用
	app := iris.New()

	// 路由注册
	utils.Register(app)

	// session配置
	// SessionId := "SESSION"
	// Sess = sessions.New(sessions.Config{
	// 	Cookie: SessionId,
	// })

	// 初始化配置
	// config := iris.WithConfiguration(iris.YAML("./conf/config.yaml"))

	// 监听8080并运行
	app.Run(iris.Addr(":8080"))
}
