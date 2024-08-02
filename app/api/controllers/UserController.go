package api

import (
	"go-mvc/app/api/requests"
	"go-mvc/app/models"
	"go-mvc/app/services"
	"go-mvc/app/utils"
	"go-mvc/app/utils/response"

	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// 继承BaseController
type UserController struct {
	BaseController
}

// 用户登录
func (con UserController) Login(cxt *gin.Context) {
	var form requests.LoginForm
	if err := cxt.ShouldBind(&form); err != nil && err.Error() != "" {
		response.Fail(cxt, err.Error())
		return
	}

	//查找用户
	var user models.User
	err := db.Where("mobile = ?", form.Mobile).First(&user).Error
	if err != nil {
		response.Fail(cxt, "手机号码未注册")
		return
	}

	// 验证密码
	if user.Password != utils.MD5(form.Password) {
		response.Fail(cxt, "密码错误")
		return
	}

	//生成token
	res, err := utils.CreateJwt(user.Id)
	if err != nil {
		response.Fail(cxt, err.Error())
		return
	}

	//响应数据
	responseData := map[string]interface{}{
		"token":      res.Token,
		"expires_at": res.ExpiresAt,
		"user_id":    user.Id,
		"user_name":  user.Name,
	}

	response.Success(cxt, responseData, "登录成功")
	return
}

// 用户信息
func (con UserController) Info(cxt *gin.Context) {
	userId := GetUserId(cxt)

	//获取用户信息
	userInfo, err := services.UserService{}.GetUserInfo(userId)
	if err != nil {
		response.Fail(cxt, err.Error())
		return
	}

	response.Success(cxt, userInfo)
	return
}
