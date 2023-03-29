package Index

import (
	// "myapp/Login/middleware"

	"fmt"
	"log"
	model "myapp/Model"
	Index_utils "myapp/app/Index/utils"
	_ "myapp/app/Login/middleware"
	"myapp/app/Login/utils"
	"regexp"

	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/sessions"
)

func Index(ctx iris.Context) {
	var req utils.Status
	// fmt.Println("判断用户是否登录和那个用户在登录!!!")
	// sessValues := sessions.Get(ctx).GetAll()
	// fmt.Println(len(sessValues))
	// for k, v := range sessValues {
	// 	fmt.Println(k, v)
	// }
	req.Trickname = sessions.Get(ctx).GetString("trickname")
	// req.Status, = sessions.Get(ctx).GetBoolean("status")
	ctx.ViewData("user", req)
	// if trickname != "" {
	// 	fmt.Println(trickname, "hhhhhh")
	// }
	// if middleware.IsLogin(ctx) {
	// 	fmt.Println("index登录成功")
	// } else {
	// 	fmt.Println("index登录失败")
	// }
	ctx.View("index.html")
}

// 密码修改
func Update_Pwd(ctx iris.Context) {
	if ctx.Method() == "POST" {
		var req Index_utils.UpdatePwd
		err := ctx.ReadForm(&req)
		if err != nil {
			fmt.Println("密码修改获取错误!")
		}
		trickname := sessions.Get(ctx).GetString("trickname")
		if !model.Select_userpwd(trickname, req.Oldpassword) {
			ctx.JSON(utils.JsonResult{
				Code: -1,
				Msg:  "原始密码不正确,请重新输入",
			})
			fmt.Println(trickname)
			fmt.Println(req.Oldpassword)
			return
		}
		// 密码校验:是否合理
		verify_password := "^[a-zA-Z0-9]{4,11}$"
		Pbol, err := regexp.MatchString(verify_password, req.Newpassword)
		if err != nil {
			log.Fatal(err)
		}
		if !Pbol {
			ctx.JSON(utils.JsonResult{
				Code: -1,
				Msg:  "密码需满足5-12字符,且必须为字母或数字",
			})
			return
		}
		// 密码校验:两次密码是否一致
		// fmt.Println(req.Password, req.Repassword)
		if req.Newpassword != req.Repassword { // 是否一致
			ctx.JSON(utils.JsonResult{
				Code: -1,
				Msg:  "两次密码不同,请重新输入",
			})
			return
		}

		if !model.Servise_pwd(trickname, req.Newpassword) {
			ctx.JSON(utils.JsonResult{
				Code: 0,
				Msg:  "更改密码失败,请查看mysql数据库信息",
			})
			return
		}

		ctx.JSON(utils.JsonResult{
			Code: 0,
			Msg:  "更改密码成功!",
		})
		return
	}
	ctx.View("index.html")

}

func Userinfo(ctx iris.Context) {
	ctx.View("user_info/index.html")
}

// 退出
func Logout(ctx iris.Context) {
	sessions.Get(ctx).Delete("trickname")
	sessions.Get(ctx).Delete("status")
	ctx.View("login.html")
}

// 默认界面
func Default(ctx iris.Context) {
	ctx.View("welcome.html")
}
