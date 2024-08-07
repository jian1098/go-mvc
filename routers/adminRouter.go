package routers

import (
	admin "go-mvc/app/admin/controllers"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

type AdminRouter struct {
	BaseRouter
}

// admin后台路由
func initAdminRouter(router *gin.Engine) {
	//加载模板
	router.LoadHTMLGlob("templates/admin/**/*")
	//加载静态文件
	router.Static("/static", filepath.Join(os.Getenv("GOPATH"), "/static"))

	adminRouter := router.Group("/admin") //路由组前缀
	{
		//后台首页
		adminRouter.GET("/home/index", admin.HomeController{}.Index)
		//管理员列表
		adminRouter.GET("/admin/index", admin.AdminController{}.Index)
	}
}
