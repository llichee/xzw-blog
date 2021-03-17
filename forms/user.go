package forms

type UserInfo struct {
	Username string `form:"username"`
	Password string `form:"password"`
}
