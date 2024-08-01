package middlewares

import (
	"fmt"
	"go-mvc/app/common"
	"go-mvc/app/constants"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

// jwt登录验证
func JwtAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.Header.Get("Authorization")
		token = strings.Replace(token, "Bearer ", "", -1)
		// 解析获取用户载荷信息
		payLoad, ok := common.ParseJwt(token)
		fmt.Println("payLoad:", token, payLoad, ok)
		if !ok {
			var code = constants.API_FAIL_CODE
			c.JSON(http.StatusOK, gin.H{
				"code": code,
				"msg":  "请先登录",
				"data": nil,
			})
			c.Abort()
			return
		}

		// 在上下文设置载荷信息
		c.Set("UserId", payLoad.UserId) //用户id

		c.Next()
	}
}
