package Index

import (
	// "myapp/Login/middleware"

	"fmt"
	"log"
	"myapp/DB"
	model "myapp/Model"
	Index_utils "myapp/app/Index/utils"
	"myapp/app/Login/middleware"
	"myapp/app/Login/utils"
	"regexp"

	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/sessions"
)

// 设置状态栏
func Setstatus(ctx iris.Context) {
	ctx.JSON(utils.JsonResult{
		Code: -1,
		Msg:  "请在编辑栏里操作!",
	})
	return
}

func Index(ctx iris.Context) {
	if !middleware.IsLogin(ctx) {
		ctx.Redirect("/login")
	}
	var req utils.Status
	// 获得登录用户的名字
	req.Trickname = sessions.Get(ctx).GetString("trickname")
	// req.Status, = sessions.Get(ctx).GetBoolean("status")
	// 向前端传送信息
	ctx.ViewData("user", req)
	// 渲染界面
	ctx.View("index.html")
}

// 密码修改
func Update_Pwd(ctx iris.Context) {
	if !middleware.IsLogin(ctx) {
		ctx.Redirect("/login")
	}
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
	if !middleware.IsLogin(ctx) {
		ctx.Redirect("/login")
	}
	// fmt.Println("Readly show view")
	var req Index_utils.UserUpadte
	// 获取用户名
	req.Nickname = sessions.Get(ctx).GetString("trickname")
	// 存储接收到的用户名
	evernickname := req.Nickname
	if ctx.Method() == "GET" {
		// 如果用户没有第一次更新个人信息
		if !model.Select_info(req.Nickname) {
			data := model.Select_Email(req.Nickname)
			if data == "nil" {
				ctx.JSON(utils.JsonResult{
					Code: -1,
					Msg:  "获取用户邮箱失败...",
				})
				return
			} else {
				req.Email = data
			}
			fmt.Println(data, "hhhhhh")
			ctx.ViewData("userInfo", req)
		} else { // 已经更新过第一次信息,从Person_information取内容
			// var tempdata model.Pserson_infomation
			tmreq := model.Select_Personinfo(req.Nickname)
			fmt.Println(tmreq, "HHHHHH")
			ctx.ViewData("userInfo", tmreq)
		}
		// ctx.View("user_info/index.html")
		// return
	}
	if ctx.Method() == "POST" {
		ctx.ReadForm(&req)
		fmt.Println("information:", req)
		// 用户没有第一次更新个人信息
		if !model.Select_info(evernickname) {
			// 查询nickname是否改变,若改变则要看是否已经被其他人注册
			if evernickname != req.Nickname && DB.Select_user(req.Nickname) {
				ctx.JSON(utils.JsonResult{
					Code: -1,
					Msg:  "昵称已存在,请修改后提交",
				})
				return
			}
			// 第一次更新用户信息
			if !model.Insert_info(&req) {
				ctx.JSON(utils.JsonResult{
					Code: -1,
					Msg:  "更新信息失败",
				})
				return
			}
			// ctx.View("user_info/index.html")
		} else { // 用户信息已存在
			// 查询nickname是否改变,若改变则要看是否已经被其他人注册
			if evernickname != req.Nickname && DB.Select_user(req.Nickname) {
				ctx.JSON(utils.JsonResult{
					Code: -1,
					Msg:  "昵称已存在,请修改后提交",
				})
				return
			}
			if !model.Update_info(evernickname, &req) {
				ctx.JSON(utils.JsonResult{
					Code: -1,
					Msg:  "更新已存在信息失败",
				})
				return
			}
		}
		// 如果用户更新了昵称 则同时要更新User_information里的username和email
		if req.Nickname != evernickname {
			if !model.Update_userinfo(evernickname, req.Nickname, req.Email) {
				ctx.JSON(utils.JsonResult{
					Code: -1,
					Msg:  "更新初始用户信息失败",
				})
				return
			}
		}
		ctx.JSON(utils.JsonResult{
			Code: 0,
			Msg:  "更新信息成功",
		})
		return
	}
	ctx.View("user_info/index.html")
	// ctx.View("user_info/index.html")

}

// 退出
func Logout(ctx iris.Context) {
	sessions.Get(ctx).Delete("trickname")
	sessions.Get(ctx).Delete("status")
	ctx.View("login.html")
}

// 默认界面
func Default(ctx iris.Context) {
	if !middleware.IsLogin(ctx) {
		ctx.Redirect("/login")
	}
	ctx.View("welcome.html")
}
