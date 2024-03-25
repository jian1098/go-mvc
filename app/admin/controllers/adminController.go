package admin

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// 继承BaseController
type AdminController struct {
	BaseController
}

// Admin.Index方法
func (con AdminController) Index(c *gin.Context) {
	c.HTML(http.StatusOK, "admin/admin/index.html", gin.H{
		"title": "admin homepage",
	})
}
