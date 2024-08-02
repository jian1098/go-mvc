package response

import (
	"go-mvc/app/constants"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Success 请求成功响应方法
func Success(c *gin.Context, data interface{}, msg ...string) {
	var code = constants.API_SUCCESS_CODE
	var msgStr = ""
	for _, res := range msg {
		msgStr += res
	}
	if msgStr == "" {
		msgStr = constants.API_SUCCESS_MSG
	}

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  msgStr,
		"data": data,
	})

	c.Abort() // 阻止执行后续代码
}

// Fail 请求失败响应方法
func Fail(c *gin.Context, msg string) {
	if msg == "" {
		msg = constants.API_FAIL_MSG
	}

	var code = constants.API_FAIL_CODE
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  msg,
		"data": nil,
	})

	c.Abort() // 阻止执行后续代码
}
