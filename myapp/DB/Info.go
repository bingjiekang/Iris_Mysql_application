package DB

// 登陆验证信息
type SelectInfo struct {
	Username string `db:"username"`
	Password string `db:"password"`
}
