package admin

import (
	"go-mvc/app/common"
	"go.uber.org/zap"
)

// 基类控制器
type BaseController struct {
}

/**
 * api控制器公共变量
 */
var logger *zap.SugaredLogger //全局日志
var moduleName = "admin"      //模块名称

// 构造函数
func init() {
	logger = common.GetLogger()
}

// 获取当前模块名称
func (con BaseController) GetModuleName() string {
	return moduleName
}
