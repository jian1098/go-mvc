package services

import (
	"go-mvc/app/common"

	"github.com/jinzhu/gorm"
	"go.uber.org/zap"
)

// 基类服务
type BaseService struct {
}

/**
 * api控制器公共变量
 */
var logger *zap.SugaredLogger //全局日志
var db *gorm.DB               //数据库

// 构造函数
func init() {
	logger = common.GetLogger()
	db = common.GetDB()
}
