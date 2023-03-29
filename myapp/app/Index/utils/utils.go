package utils

type UpdatePwd struct {
	Oldpassword string `form:"oldPassword"`
	Newpassword string `form:"newPassword"`
	Repassword  string `form:"rePassword"`
}
