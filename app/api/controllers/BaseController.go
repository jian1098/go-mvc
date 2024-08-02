package api

import (
	"github.com/spf13/cast"

	"github.com/gin-gonic/gin"
)

// 基类控制器
type BaseController struct {
}

/**
 * api控制器公共变量
 */
var moduleName = "api" //模块名称

// 构造函数
func init() {

}

// 获取当前模块名称
func (con BaseController) GetModuleName() string {
	return moduleName
}

// 获取当前登录用户id
func GetUserId(c *gin.Context) int {
	userId, _ := c.Get("UserId")
	return cast.ToInt(userId)
}
