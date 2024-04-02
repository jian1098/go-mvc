package api

import (
	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/spf13/cast" // 使用cast包进行类型转换
	"go-mvc/app/models"
	"net/http"
)

// 继承BaseController
type IndexController struct {
	BaseController
}

// Demo方法
func (con IndexController) Demo(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"module": con.GetModuleName(),
		"msg":    "hello gin!",
	})
}

// Get方法
func (con IndexController) Get(c *gin.Context) {
	name := c.Query("name")                //获取get参数值
	name = c.DefaultQuery("name", "Guest") //设置默认参数值

	c.JSON(http.StatusOK, gin.H{
		"module": con.GetModuleName(),
		"msg":    "hello " + name,
	})
}

// Post方法
func (con IndexController) Post(c *gin.Context) {
	name := c.PostForm("name") //获取post参数值
	c.JSON(http.StatusOK, gin.H{
		"module": con.GetModuleName(),
		"msg":    "hello " + name,
	})
}

// Bind方法
func (con IndexController) Bind(c *gin.Context) {
	var person models.Person
	err := c.ShouldBind(&person)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"module": con.GetModuleName(),
			"msg":    err.Error(),
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"module": con.GetModuleName(),
			"msg":    "hello " + person.Name + " your age is " + cast.ToString(person.Age),
		})
	}
}

// 文件上传
func (con IndexController) Upload(c *gin.Context) {
	file, _ := c.FormFile("file") //表单的文件name="file"
	//文件上传路径和文件名
	err := c.SaveUploadedFile(file, "./uploads/images/"+file.Filename)
	//记录错误日志
	logger.Error(err)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"module": con.GetModuleName(),
			"msg":    err.Error(),
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"module": con.GetModuleName(),
			"msg":    "upload success",
		})
	}
}

// 数据库操作
func (con IndexController) Db(c *gin.Context) {
	var user models.User
	err := db.Where("id = ?", 0).First(&user).Error
	if err != nil {
		logger.Error(err)
		ApiRequsetFail(c, err.Error())
	} else {
		ApiRequsetSuccess(c, user, "请求成功")
	}
}
