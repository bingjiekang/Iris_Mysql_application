package utils

// 系统注册
type RegistReq struct {
	UserName string `form:"username" validate:"required"`
	Password string `form:"password" validate:"required"`
	Email    string `form:"email"`
	Captcha  string `form:"captcha" validate:"required"`
	IdKey    string `form:"idKey" validate:"required"`
}

// 返回信息
type JsonResult struct {
	Code  int         `json:"code"`  // 响应编码：0成功 401请登录 403无权限 500错误
	Msg   string      `json:"msg"`   // 消息提示语
	Data  interface{} `json:"data"`  // 数据对象
	Count int64       `json:"count"` // 记录总数
}

// 邮箱登录前验证
type SendEmail struct {
	Email   string `form:"email"`
	Captcha string `form:"captcha" validate:"required"`
	IdKey   string `form:"idkey" validate:"required"`
}

// 邮箱登陆验证码
type LaEmail struct {
	Email  string `form:"email"`
	Verify string `form:"verify"`
}

// 注册结构体
type RegistUser struct {
	Username   string `form:"username"`
	Password   string `form:"password"`
	Repassword string `form:"rePassword"`
	Email      string `form:"email"`
	Captcha    string `form:"captcha"`
}
