package Login

import (
	"database/sql"
	"log"
	"myapp/DB"
	"myapp/app/Login/utils"

	_ "github.com/go-sql-driver/mysql"
	"github.com/kataras/iris/v12"
	"github.com/mojocn/base64Captcha"
)

// 登陆验证
func Login(ctx iris.Context) {

	// 如果是post请求
	if ctx.Method() == "POST" {
		var req utils.LoginReq
		ctx.ReadForm(&req)

		// 判断用户是否存在
		Db, err := sql.Open("mysql", "root:12345678@tcp(localhost:3306)/Iris?charset=utf8")
		if err != nil {
			log.Fatal(err)
		}
		defer Db.Close()
		DB.StartDB(Db)
		if !DB.Select_user(Db, req.UserName) {
			ctx.JSON(utils.JsonResult{
				Code: -1,
				Msg:  "用户名不存在",
			})
			return
		}

		// 校验验证码是否正确
		verifyRes := base64Captcha.VerifyCaptcha(req.IdKey, req.Captcha)
		if !verifyRes {
			ctx.JSON(utils.JsonResult{
				Code: -1,
				Msg:  "验证码错误",
			})
			return
		}

		// 判断密码是否合理
		if !DB.Select_user_pwd(Db, req.UserName, req.Password) {
			ctx.JSON(utils.JsonResult{
				Code: -1,
				Msg:  "密码不正确",
			})
			return
		}

		// 登陆成功
		ctx.JSON(utils.JsonResult{
			Code: 0,
			Msg:  "登陆成功!",
		})
		// ctx.View("index.html")
		return
	}

	// get请求或其他请求 渲染登陆界面
	ctx.View("login.html")

}

// 验证码
func Captcha(ctx iris.Context) {
	// 验证码参数配置：字符,公式,验证码配置
	var configC = base64Captcha.ConfigCharacter{
		Height: 60,
		Width:  240,
		//const CaptchaModeNumber:数字,CaptchaModeAlphabet:字母,CaptchaModeArithmetic:算术,CaptchaModeNumberAlphabet:数字字母混合.
		Mode:               base64Captcha.CaptchaModeNumberAlphabet,
		ComplexOfNoiseText: base64Captcha.CaptchaComplexLower,
		ComplexOfNoiseDot:  base64Captcha.CaptchaComplexLower,
		IsShowHollowLine:   true,
		IsShowNoiseDot:     false,
		IsShowNoiseText:    false,
		IsShowSlimeLine:    false,
		IsShowSineLine:     false,
		CaptchaLen:         6,
		UseCJKFonts:        true,
	}
	///create a characters captcha.
	idKeyC, capC := base64Captcha.GenerateCaptcha("", configC)
	//以base64编码
	base64stringC := base64Captcha.CaptchaWriteToBase64Encoding(capC)

	// 返回结果集
	ctx.JSON(utils.CaptchaRes{
		Code:  0,
		IdKey: idKeyC,
		Data:  base64stringC,
		Msg:   "操作成功",
	})
}
