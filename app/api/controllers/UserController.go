package api

import (
	"go-mvc/app/api/requests"
	"go-mvc/app/models"
	"go-mvc/app/services"
	"go-mvc/app/utils"

	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// 继承BaseController
type UserController struct {
	BaseController
}

// 用户登录
func (con UserController) Login(c *gin.Context) {
	var form requests.LoginForm
	if err := c.ShouldBind(&form); err != nil && err.Error() != "" {
		ResponseError(c, err.Error())
		return
	}

	//查找用户
	var user models.User
	err := db.Where("mobile = ?", form.Mobile).First(&user).Error
	if err != nil {
		ResponseError(c, err.Error())
		return
	}

	// 验证密码
	if user.Password != utils.MD5(form.Password) {
		ResponseError(c, "密码错误")
		return
	}

	//生成token
	res, err := utils.CreateJwt(user.Id)
	if err != nil {
		ResponseError(c, err.Error())
		return
	}

	//响应数据
	responseData := map[string]interface{}{
		"token":      res.Token,
		"expires_at": res.ExpiresAt,
		"user_id":    user.Id,
		"user_name":  user.Name,
	}

	ResponseSuccess(c, responseData, "登录成功")
	return
}

// 用户信息
func (con UserController) Info(c *gin.Context) {
	userId := GetUserId(c)

	//获取用户信息
	userInfo, err := services.UserService{}.GetUserInfo(userId)
	if err != nil {
		ResponseError(c, err.Error())
		return
	}

	ResponseSuccess(c, userInfo, "请求成功")
}
