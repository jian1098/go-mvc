package routers

import (
	"github.com/gin-gonic/gin"
	"go-mvc/app/api/controllers"
)

type ApiRouter struct {
	BaseRouter
}

// api路由
func initApiRouter(router *gin.Engine) {
	apiRouter := router.Group("/api") //路由组前缀
	{
		//demo
		apiRouter.GET("/index/demo", api.IndexController{}.Demo)
		//get
		apiRouter.GET("/index/get", api.IndexController{}.Get)
		//post
		apiRouter.POST("/index/post", api.IndexController{}.Post)
		//参数绑定到结构体
		apiRouter.POST("/index/bind", api.IndexController{}.Bind)
		//文件上传
		apiRouter.POST("/index/upload", api.IndexController{}.Upload)
		//数据库操作
		apiRouter.GET("/index/db", api.IndexController{}.Db)
	}
}
