package api

import (
	"go-mvc/app/api/requests"

	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// 继承BaseController
type UserController struct {
	BaseController
}

// 登录方法
func (con UserController) Login(c *gin.Context) {
	var form requests.UserForm
	if err := c.ShouldBind(&form); err != nil && err.Error() != "" {
		ResponseError(c, err.Error())
		return
	}

	ResponseSuccess(c, form, "登录成功")
}
