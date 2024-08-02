package middlewares

import (
	"fmt"
	"go-mvc/app/utils"
	"go-mvc/app/utils/response"
	"strings"

	"github.com/gin-gonic/gin"
)

// jwt登录验证
func JwtAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.Header.Get("Authorization")
		token = strings.Replace(token, "Bearer ", "", -1)
		// 解析获取用户载荷信息
		payLoad, ok := utils.ParseJwt(token)
		fmt.Println("payLoad:", token, payLoad, ok)
		if !ok {
			response.Fail(c, "请先登录")
			return
		}

		// 在上下文设置载荷信息
		c.Set("UserId", payLoad.UserId) //用户id

		c.Next()
	}
}
