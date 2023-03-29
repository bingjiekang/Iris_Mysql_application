package utils

import (
	Index "myapp/app/Index/view"
	"myapp/app/Login/middleware"
	Login "myapp/app/Login/view"
	Regist "myapp/app/Regist/view"

	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/sessions"
	_ "github.com/mojocn/base64Captcha"
)

func Register(app *iris.Application) {
	// 注册SESSION中间件
	session := sessions.New(sessions.Config{
		Cookie: sessions.DefaultCookieName,
	})

	// SESSION中间件
	app.Use(session.Handler())
	// // 登录验证中间件
	app.Use(middleware.CheckLogin)

	// tmpl注册html页面,并重载所有方法
	tmpl := iris.HTML("./template", ".html").Reload(true)
	// //
	// tmpl.Reload(true)

	// app注册tmpl
	app.RegisterView(tmpl)

	// 访问静态文件
	app.HandleDir("/static", "./static")
	app.HandleDir("/assets", "./static/assets")

	// 登陆主页处理
	index := app.Party("/")
	{
		index.Any("/", Index.Index)
		index.Get("/login", Login.Login)           // 登陆界面
		index.Post("/login", Login.Login)          // 提交到login的信息
		index.Get("/captcha", Login.Captcha)       // 验证码展示
		index.Get("/regist", Regist.Regist)        // 注册界面
		index.Post("/regist", Regist.Regist)       // 提交注册信息
		index.Post("/sendEmail", Regist.SendEmail) // 提交登陆验证码信息
		index.Get("/index", Index.Index)           // 详情界面信息
		index.Get("/default", Index.Default)       // 默认页面信息
		index.Get("/logout", Index.Logout)         // 退出
		index.Post("/updatePwd", Index.Update_Pwd) // 更新密码
		index.Get("/userinfo", Index.Userinfo)     // 用户个人信息
		index.Post("/userinfo", Index.Userinfo)    // 用户个人信息
	}

}
