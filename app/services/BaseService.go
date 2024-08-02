package services

import (
	"go-mvc/app/utils"

	"go.uber.org/zap"
)

// 基类服务
type BaseService struct {
}

/**
 * api控制器公共变量
 */
var logger *zap.SugaredLogger //全局日志

// 构造函数
func init() {
	logger = utils.GetLogger()
}
