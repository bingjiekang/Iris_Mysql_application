package midware

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
	fmt.Println("登录验证中间件")
	// 放行设置
	urlItem := []string{"/captcha", "/login"}
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

// 登录用户ID
func Uid(ctx iris.Context) int {
	fmt.Println("全局获取用户ID")
	//sessValues := sessions.Get(ctx).GetAll()
	//fmt.Println(len(sessValues))
	//for k, v := range sessValues {
	//	fmt.Println(k, v)
	//}
	userId := sessions.Get(ctx).GetIntDefault("userId", 0)
	return userId
}

// 判断用户登录状态
func IsLogin(ctx iris.Context) bool {
	// 初始化session对象
	fmt.Println("初始化SESSION")
	userId := Uid(ctx)
	return userId > 0
}
