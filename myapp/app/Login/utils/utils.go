package utils

// 登陆图形验证码
type CaptchaRes struct {
	Code  int         `json:"code"`  //响应编码 0 成功 500 错误 403 无权限
	Msg   string      `json:"msg"`   //消息
	Data  interface{} `json:"data"`  //数据内容
	IdKey string      `json:"idkey"` //验证码ID
}

// 系统登录
type LoginReq struct {
	UserName string `form:"username" validate:"required"`
	Password string `form:"password" validate:"required"`
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
