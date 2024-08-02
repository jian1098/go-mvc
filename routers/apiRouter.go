package routers

import (
	api "go-mvc/app/api/controllers"
	"go-mvc/app/middlewares"

	"github.com/gin-gonic/gin"
)

// API接口路由
type ApiRouter struct {
	BaseRouter
}

// api路由
func initApiRouter(router *gin.Engine) {
	//需要登录的路由组
	loginGroup := router.Group("/api") //路由组前缀
	//使用中间件
	loginGroup.Use(middlewares.JwtAuth()) //jwt验证中间件
	loginGroup.Use(middlewares.Cors())    //跨域中间件
	{
		//用户路由组
		userRouter := loginGroup.Group("/user")
		userRouter.GET("/info", api.UserController{}.Info) //用户信息

		//其他路由组

	}

	//不需要登录的路由组
	noLoginGroup := router.Group("/api") //路由组前缀
	noLoginGroup.Use(middlewares.Cors()) //跨域中间件
	{
		//单独路由示例模块
		noLoginGroup.GET("/index/demo", api.IndexController{}.Demo)      //demo
		noLoginGroup.GET("/index/get", api.IndexController{}.Get)        //get
		noLoginGroup.POST("/index/post", api.IndexController{}.Post)     //post
		noLoginGroup.POST("/index/bind", api.IndexController{}.Bind)     //参数绑定到结构体
		noLoginGroup.POST("/index/upload", api.IndexController{}.Upload) //文件上传
		noLoginGroup.GET("/index/db", api.IndexController{}.Db)          //数据库操作

		//用户路由组
		noLoginUserRouter := noLoginGroup.Group("/user")
		noLoginUserRouter.POST("/login", api.UserController{}.Login) //用户登录
	}
}
