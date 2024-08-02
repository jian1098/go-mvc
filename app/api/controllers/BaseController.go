package api

import (
	"go-mvc/app/utils"

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

// 获取当前登录用户id
func GetUserId(c *gin.Context) int {
	userId, _ := c.Get("UserId")
	return cast.ToInt(userId)
}
