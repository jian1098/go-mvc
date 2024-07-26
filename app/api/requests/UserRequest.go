package requests

import "github.com/gookit/validate"

func (f UserForm) ConfigValidation(v *validate.Validation) {
	//验证场景
	v.WithScenes(validate.SValues{
		"login": []string{"Mobile", "Password"},
	})
}

// 用户表单
type UserForm struct {
	Mobile   string `form:"mobile" label:"手机号码" validate:"required|cn_mobile"`
	Password string `form:"password" label:"登录密码" validate:"required|minLen:6"`
}
