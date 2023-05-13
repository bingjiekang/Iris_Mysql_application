package utils

import (
	Analysis "myapp/app/Analysis/view"
	City "myapp/app/City/view"
	Index "myapp/app/Index/view"
	Level "myapp/app/Level/view"
	Memberlevel "myapp/app/MemberLevel/view"
	Position "myapp/app/Position/view"
	User "myapp/app/User/view"

	_ "myapp/app/Login/middleware"
	Login "myapp/app/Login/view"
	Regist "myapp/app/Regist/view"
	Role "myapp/app/Role/view"
	Time "myapp/app/Time/view"

	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/sessions"
	_ "github.com/mojocn/base64Captcha"
)

func Register(app *iris.Application) {
	// 注册SESSION中间件
	// session := sessions.New(sessions.Config{
	// 	Cookie: sessions.DefaultCookieName,
	// })
	SessionId := "SESSION"
	Sess := sessions.New(sessions.Config{
		Cookie: SessionId,
	})

	// SESSION中间件
	app.Use(Sess.Handler())
	// app.Use(Sess.Start)
	// 登录验证中间件
	// app.Use(middleware.CheckLogin)

	// tmpl注册html页面,并重载所有方法
	tmpl := iris.HTML("./template", ".html")
	// //
	tmpl.Reload(true)

	// app注册tmpl

	app.RegisterView(tmpl)

	// 访问静态文件
	app.HandleDir("/static", "./static")
	app.HandleDir("/assets", "./static/assets")

	// 登陆主页处理
	index := app.Party("/")
	{
		index.Any("/", Login.Login)
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
		index.Get("/userInfo", Index.Userinfo)     // 用户个人信息
		index.Post("/userInfo", Index.Userinfo)    // 用户个人信息
	}

	// 角色管理界面
	role := app.Party("role")
	{
		role.Get("/index", Role.Index) // 角色界面
		role.Post("/index", Role.Index)
		role.Post("/list", Role.List)                     // 角色信息界面
		role.Get("/edit/{id:int}", Role.Role_append)      // 显示信息及更新信息界面
		role.Post("/add", Role.Add)                       // 添加角色
		role.Post("/update", Role.Update)                 // 更新角色信息
		role.Post("/delete/{id:string}", Role.Delete_mag) // 删除角色信息
		role.Post("/setStatus", Index.Setstatus)          //状态栏返回信息
	}

	// 职称管理界面
	level := app.Party("/level")
	{
		level.Get("/index", Level.Index)                    // 职称显示界面
		level.Post("/list", Level.List)                     // 职称信息界面
		level.Get("/edit/{id:int}", Level.Level_append)     // 职级显示及更新界面
		level.Post("/add", Level.Add)                       //添加职级
		level.Post("/update", Level.Update)                 //更新职称信息
		level.Post("/delete/{id:string}", Level.Delete_mag) //删除职称信息
		level.Post("/setStatus", Index.Setstatus)           //状态栏返回信息
	}

	// 岗位管理界面
	position := app.Party("/position")
	{
		position.Get("/index", Position.Index)                    // 岗位显示界面
		position.Post("/list", Position.List)                     // 岗位信息界面
		position.Get("/edit/{id:int}", Position.Position_append)  // 岗位显示及更新界面
		position.Post("/add", Position.Add)                       //添加岗位
		position.Post("/update", Position.Update)                 //更新岗位信息
		position.Post("/delete/{id:string}", Position.Delete_mag) //删除岗位信息
		position.Post("/setStatus", Index.Setstatus)              //状态栏返回信息
	}

	// 会员等级管理界面
	memberlevel := app.Party("/memberlevel")
	{
		memberlevel.Get("/index", Memberlevel.Index)                      // 等级显示界面
		memberlevel.Post("/list", Memberlevel.List)                       // 等级信息界面
		memberlevel.Get("/edit/{id:int}", Memberlevel.Memberlevel_append) // 等级显示及更新界面
		memberlevel.Post("/add", Memberlevel.Add)                         //添加等级
		memberlevel.Post("/update", Memberlevel.Update)                   //更新等级信息
		memberlevel.Post("/delete/{id:string}", Memberlevel.Delete_mag)   //删除等级信息
	}

	// 用户管理界面
	user := app.Party("/user")
	{
		user.Get("/index", User.Index)                    // 用户显示界面
		user.Post("/list", User.List)                     // 显示信息
		user.Get("/edit/{id:int}", User.User_append)      // 显示及更新
		user.Post("/add", User.Add)                       //添加用户
		user.Post("/update", User.Update)                 //更新用户信息
		user.Post("/delete/{id:string}", User.Delete_mag) //删除用户信息
		user.Post("/setStatus", Index.Setstatus)          //状态栏返回信息
	}

	// 实践管理界面
	time := app.Party("time")
	{
		time.Get("/index", Time.Time)          // 时间管理界面
		time.Get("/countdown", Time.Countdown) // 时间倒计时
	}

	// 数据分析界面
	analysis := app.Party("analysis")
	{
		analysis.Any("/index", Analysis.Analysis) // 数据分析管理界面
	}

	// 城市地区
	city := app.Party("city")
	{
		city.Any("/index", City.City) // 地域地区显示
	}

}
