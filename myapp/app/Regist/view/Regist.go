package Regist

import (
	"fmt"
	"log"
	"math/rand"
	"myapp/DB"
	"myapp/app/Regist/utils"
	"net/smtp"
	"regexp"
	"strconv"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/kataras/iris/v12"
	"github.com/mojocn/base64Captcha"
)

var verify_code utils.LaEmail

// 注册界面
func Regist(ctx iris.Context) {
	if ctx.Method() == "POST" {
		// 定义req为注册信息的结构体变量
		var req utils.RegistUser
		// 接收传过来的信息
		ctx.ReadJSON(&req)
		// 数据校验
		// 用户名校验
		verify_username := "^[a-zA-Z][a-zA-Z0-9_]{4,15}$" // 是否合理
		Bol, err := regexp.MatchString(verify_username, req.Username)
		if err != nil {
			log.Fatal(err)
		}
		if !Bol {
			ctx.JSON(utils.JsonResult{
				Code: -1,
				Msg:  "用户名需满足5-16字符,且首字母必须为字母",
			})
			return
		}
		// 用户名是否已存在
		// 链接数据库
		// Db, err := sql.Open("mysql", "root:12345678@tcp(localhost:3306)/Iris?charset=utf8")
		// if err != nil {
		// 	log.Fatal(err)
		// }
		// defer Db.Close()
		// DB.StartDB(Db)
		if DB.Select_user(req.Username) {
			ctx.JSON(utils.JsonResult{
				Code: -1,
				Msg:  "用户名已存在,请更改后注册",
			})
			return
		}

		// 密码校验:是否合理
		verify_password := "^[a-zA-Z0-9]{4,11}$"
		Pbol, err := regexp.MatchString(verify_password, req.Password)
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
		if req.Password != req.Repassword { // 是否一致
			ctx.JSON(utils.JsonResult{
				Code: -1,
				Msg:  "两次密码不同,请重新输入",
			})
			return
		}
		// 邮箱及验证码是否正确
		// 上线的话 需要验证邮箱是否重复注册
		// var emailReg = "^([a-zA-Z]|[0-9])(\w|\-)+@[a-zA-Z0-9]+\.([a-zA-Z]{2,4})$"
		// 邮箱验证码:不正确或超时
		Verify_email_code := DB.Get_Redis(req.Email)
		if Verify_email_code == "nil" {
			ctx.JSON(utils.JsonResult{
				Code: -1,
				Msg:  "验证码不正确或失效",
			})
			return
		}
		if req.Captcha != Verify_email_code {
			ctx.JSON(utils.JsonResult{
				Code: -1,
				Msg:  "邮箱验证码错误...请重新输入",
			})
			return
		}
		ctx.JSON(utils.JsonResult{
			Code: 0,
			Msg:  "登录成功!",
		})

		// 向数据库存入用户数据
		// if err != nil {
		// 	log.Fatal(err)
		// }
		// defer Db.Close()
		// DB.StartDB(Db)
		if !DB.Insert(req.Username, req.Password, req.Email) {
			ctx.JSON(utils.JsonResult{
				Code: -1,
				Msg:  "注册失败,请联系管理员",
			})
		}

		// 将redis数据库的验证信息删除
		if DB.Del_Redis(req.Email) == "OK" {
			fmt.Println("删除成功")
		} else {
			fmt.Println("删除失败")
		}
		return

	}
	ctx.View("regist.html")
}

// 邮箱验证
func SendEmail(ctx iris.Context) {
	if ctx.Method() == "POST" {
		// 定义req为接受验证码的结构体变量
		var req utils.SendEmail
		// req接受传过来的信息
		// ctx.ReadForm(&req)
		ctx.ReadJSON(&req)
		// fmt.Println(req.Email, req.Captcha, req.IdKey)

		verifyRes := base64Captcha.VerifyCaptcha(req.IdKey, req.Captcha)
		// 校验验证码是否正确
		if !verifyRes {
			ctx.JSON(utils.JsonResult{
				Code: -1,
				Msg:  "验证码错误...",
			})
			return
		}
		// 正确 返回code 0
		ctx.JSON(utils.JsonResult{
			Code: 0,
			Msg:  "验证码正确!",
		})

		// 设置随机验证码 1000   10000
		rand.Seed(time.Now().Unix())
		randm := rand.Intn(9000) + 1000
		// 记录对应email和验证码
		verify_code.Email = req.Email
		// 将randm转为字符串
		verify_code.Verify = strconv.FormatInt(int64(randm), 10)
		// 将验证码和对应用户邮箱存储到redis里
		DB.Line_Redis()
		DB.Set_Redis(verify_code.Email, verify_code.Verify)
		// 转为字符串
		// rm := strconv.FormatInt(int64(randm), 10)
		// 发送邮箱
		Send_Email_func(req.Email, verify_code.Verify)
		// fmt.Println(Send_Email_func(req.Email, rm), randm, rm)
		// fmt.Println(verify_code.Email, verify_code.Verify)

		return
	}
}

func Send_Email_func(email string, randm string) bool {
	// 发件人及收件人信息确认
	// sender := "kangbingjie2023@outlook.com"
	// receiver := email
	// pwd := "wmznnqmklquelquy"
	// mail_type := "smtp.office365.com"
	// randm = randm(int64)
	mail_type := "smtp.qq.com"
	mail_user := "942844582@qq.com" // 邮箱授权者信息
	mail_otac := "cljmrzxqgdvqbedb" // 授权码
	receiver := email

	auth := smtp.PlainAuth("", mail_user, mail_otac, mail_type)
	to := []string{receiver}
	msg := []byte("To:" + receiver + "\r\n" +
		"Subject:Isis框架登录验证\r\n\r\n" +
		"【Iris+Layui开发框架】您的验证码为" + randm +
		".验证码3分钟后自动过期,若非本人操作,请忽略此条信息.\r\n")

	// smtpobj.ehlo() // 验证
	// smtpobj.starttls()
	// smtpobj.login(sender, pwdcore)
	err := smtp.SendMail(mail_type+":25", auth, mail_user, to, msg)
	if err != nil {
		log.Fatal(err)
		return false
	}
	return true
}
