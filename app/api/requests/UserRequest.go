package requests

// 用户登录表单
type LoginForm struct {
	Mobile   string `form:"mobile" label:"手机号码" validate:"required|cn_mobile"`
	Password string `form:"password" label:"登录密码" validate:"required|minLen:6"`
}
