package routers

import (
	"github.com/gin-gonic/gin"
	"go-mvc/app/middlewares"
)

type BaseRouter struct {
}

// 初始化路由
func (br BaseRouter) InitRouter() *gin.Engine {
	router := gin.Default()
	//设置一个信任ip,不然会报警告：[GIN-debug] [WARNING] You trusted all proxies, this is NOT safe. We recommend you to set a value.
	err := router.SetTrustedProxies([]string{"127.0.0.1"})
	if err != nil {
		return nil
	}

	//使用中间件
	router.Use(middlewares.IPLimit())

	//api路由
	initApiRouter(router)

	//admin路由
	initAdminRouter(router)

	return router
}
