package admin

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// 继承BaseController
type HomeController struct {
	BaseController
}

// Home.Index方法
func (con HomeController) Index(c *gin.Context) {
	c.HTML(http.StatusOK, "admin/home/index.html", gin.H{
		"title": "admin homepage",
	})
}
