package middlewares

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

// 访问ip黑名单
func IPLimit() gin.HandlerFunc {
	return func(c *gin.Context) {
		ip := c.ClientIP()
		blackList := ""

		if strings.Contains(blackList, ip) {
			c.Abort()
			c.JSON(http.StatusUnauthorized, gin.H{
				"code": 0,
				"msg":  "IP " + ip + " 没有访问权限",
				"data": nil,
			})
			return // return也是可以省略的，执行了abort操作，会内置在中间件defer前，return，写出来也只是解答为什么Abort()之后，还能执行返回JSON数据
		}
	}
}
