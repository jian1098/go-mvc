package api

import (
	"go-mvc/app/models"
	"go-mvc/app/utils/db"
	"go-mvc/app/utils/log"
	"go-mvc/app/utils/response"

	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/spf13/cast" // 使用cast包进行类型转换
)

// 继承BaseController
type IndexController struct {
	BaseController
}

// Demo方法
func (con IndexController) Demo(cxt *gin.Context) {
	response.Success(cxt, "hello gin!", "请求成功")
	return
}

// Get方法
func (con IndexController) Get(cxt *gin.Context) {
	name := cxt.Query("name")                //获取get参数值
	name = cxt.DefaultQuery("name", "Guest") //设置默认参数值

	response.Success(cxt, nil, "hello "+name)
	return
}

// Post Post方法
func (con IndexController) Post(cxt *gin.Context) {
	name := cxt.PostForm("name") //获取post参数值

	response.Success(cxt, nil, "hello "+name)
	return
}

// Bind方法
func (con IndexController) Bind(cxt *gin.Context) {
	var person models.Person
	err := cxt.ShouldBind(&person)

	if err.Error() != "" {
		response.Fail(cxt, err.Error())
	} else {
		response.Success(cxt, nil, "hello "+person.Name+" your age is "+cast.ToString(person.Age))
	}

	return
}

// 文件上传
func (con IndexController) Upload(cxt *gin.Context) {
	file, _ := cxt.FormFile("file") //表单的文件name="file"
	//文件上传路径和文件名
	err := cxt.SaveUploadedFile(file, "./uploads/images/"+file.Filename)
	//记录错误日志
	log.Instance().Error(err)
	if err != nil {
		response.Fail(cxt, err.Error())
	} else {
		response.Success(cxt, nil, "上传成功")
	}
	return
}

// 数据库操作
func (con IndexController) Db(cxt *gin.Context) {
	var user models.User
	err := db.Instance().Where("id = ?", 0).First(&user).Error
	if err != nil {
		log.Instance().Error(err)
		response.Fail(cxt, err.Error())
	} else {
		response.Success(cxt, user, "请求成功")
	}
	return
}
