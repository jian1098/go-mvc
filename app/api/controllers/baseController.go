package api

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	common2 "go-mvc/app/common"
	"go-mvc/app/constants"
	"go.uber.org/zap"
	"net/http"
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
	logger = common2.GetLogger()
	db = common2.GetDB()
}

// 获取当前模块名称
func (con BaseController) GetModuleName() string {
	return moduleName
}

// api请求成功响应方法
func ApiRequsetSuccess(c *gin.Context, data interface{}, msg string) {
	if msg == "" {
		msg = constants.API_SUCCESS_MSG
	}
	c.JSON(http.StatusOK, gin.H{
		"code": constants.API_SUCCESS_CODE,
		"msg":  msg,
		"data": data,
	})
	return
}

// api请求失败响应方法
func ApiRequsetFail(c *gin.Context, msg string) {
	if msg == "" {
		msg = constants.API_FAIL_MSG
	}
	c.JSON(http.StatusOK, gin.H{
		"code": constants.API_FAIL_CODE,
		"msg":  msg,
		"data": nil,
	})
	return
}
