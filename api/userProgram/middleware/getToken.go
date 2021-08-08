package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

//获取Token
//如果存在就带入下个页面
//不存在则不处理
func GetToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.Header.Get("token")

		if token != "" {
			j := NewJWT()
			// parseToken 解析token包含的信息
			claims, err := j.ParseToken(token)
			if err != nil {
				if err == TokenExpired {
					c.JSON(http.StatusOK, gin.H{
						"code":   401,
						"status": -1,
						"msg":    "授权已过期",
					})
					c.Abort()
					return
				} else {
					fmt.Println(err.Error())
				}
				c.JSON(http.StatusOK, gin.H{
					"code":   401,
					"status": -1,
					"msg":    err.Error(),
				})
				c.Abort()
				return
			}
			c.Set("claims", claims)
		}
		//继续交由下一个路由处理
	}
}
