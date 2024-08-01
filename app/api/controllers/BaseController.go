package api

import (
	"go-mvc/app/constants"
	"go-mvc/app/utils"
	"net/http"

	"github.com/spf13/cast"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"go.uber.org/zap"
)

// 基类控制器
type BaseController struct {
}

/**
 * api控制器公共变量
 */
var logger *zap.SugaredLogger //全局日志
var moduleName = "api"        //模块名称
var db *gorm.DB               //数据库

// 构造函数
func init() {
	logger = utils.GetLogger()
	db = utils.GetDB()
}

// 获取当前模块名称
func (con BaseController) GetModuleName() string {
	return moduleName
}

// api请求成功响应方法
func ResponseSuccess(c *gin.Context, data interface{}, msg string, statusCode ...int) {
	if msg == "" {
		msg = constants.API_SUCCESS_MSG
	}

	var code int = constants.API_SUCCESS_CODE
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  msg,
		"data": data,
	})
	return
}

// api请求失败响应方法
func ResponseError(c *gin.Context, msg string, statusCode ...int) {
	if msg == "" {
		msg = constants.API_FAIL_MSG
	}

	var code int = constants.API_FAIL_CODE
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  msg,
		"data": nil,
	})
	return
}

// 获取当前登录用户id
func GetUserId(c *gin.Context) int {
	userId, _ := c.Get("UserId")
	return cast.ToInt(userId)
}
