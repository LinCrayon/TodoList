package middleware

import (
	"net/http"
	"time"
	"todo_list_v2.01/pkg/ctl"
	"todo_list_v2.01/pkg/e"
	"todo_list_v2.01/pkg/utils"

	"github.com/gin-gonic/gin"
)

// JWT token验证中间件
func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		var code int
		code = e.SUCCESS
		token := c.GetHeader("Authorization") //从上下文中获取HTTP请求头部中名为"Authorization"的值。
		if token == "" {
			code = http.StatusNotFound
			c.JSON(e.InvalidParams, gin.H{
				"status": code,
				"msg":    e.GetMsg(code),
				"data":   "缺少Token",
			})
			c.Abort()
			return
		}

		claims, err := utils.ParseToken(token)
		if err != nil {
			code = e.ErrorAuthCheckTokenFail
		} else if time.Now().Unix() > claims.ExpiresAt {
			code = e.ErrorAuthCheckTokenTimeout
		}

		if code != e.SUCCESS {
			c.JSON(e.InvalidParams, gin.H{
				"status": code,
				"msg":    e.GetMsg(code),
				"data":   "可能是身份过期了，请重新登录",
			})
			c.Abort()
			return
		}
		//c.Request.WithContext()将原始请求中的上下文替换为新创建的上下文。
		//NewContext()创建一个新的上下文,接受两个参数，第一个是旧的上下文，第二个是要添加到上下文中的键值对。
		c.Request = c.Request.WithContext(ctl.NewContext(c.Request.Context(), &ctl.UserInfo{Id: claims.Id}))
		c.Next()
	}
}
