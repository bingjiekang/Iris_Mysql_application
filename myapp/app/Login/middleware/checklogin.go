package middleware

import (
	// _ "easygoadmin/utils"
	"fmt"
	"net/http"
	"strings"

	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/sessions"
)

// 登录验证中间件
func CheckLogin(ctx iris.Context) {
	// fmt.Println("登录验证中间件")
	// 放行设置
	urlItem := []string{"/captcha", "/login", "/regist"}
	if !InStringArray(ctx.Path(), urlItem) && !strings.Contains(ctx.Path(), "static") {
		if !IsLogin(ctx) {
			// 跳转登录页,方式：301(永久移动),308(永久重定向),307(临时重定向)
			ctx.Redirect("/login", http.StatusTemporaryRedirect)
			return
		}
	}
	// 前置中间件
	ctx.Application().Logger().Infof("Runs before %s", ctx.Path())
	ctx.Next()
}

func InStringArray(value string, array []string) bool {
	for _, v := range array {
		if v == value {
			return true
		}
	}
	return false
}

// 判断用户是否登录
func IsLogin(ctx iris.Context) bool {
	// fmt.Println("判断用户是否登录!!!")
	// sessValues := sessions.Get(ctx).GetAll()
	// fmt.Println(len(sessValues))
	// for k, v := range sessValues {
	// 	fmt.Println(k, v)
	// }
	// sessions.Get(ctx.Get)
	islogin, err := sessions.Get(ctx).GetBoolean("status")
	// fmt.Println(ctx, islogin)
	if err != nil {
		fmt.Println("未登录!请先登录")
	}
	return islogin
}

// // 判断用户登录状态
// func IsLogin(ctx iris.Context) bool {
// 	// 初始化session对象
// 	fmt.Println("初始化SESSION")
// 	userId := Uid(ctx)
// 	return userId > 0
// }
