package middleware

import (
	"github.com/gin-gonic/gin"
	"time"
	"todo_list/pkg/utils"
)

func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		code := 200
		//var data interface{}
		token := c.GetHeader("Authorization")
		if token == "" {
			code = 404
		} else {
			claim, err := utils.ParseToken(token)
			if err != nil {
				code = 403 //无权限，token是无权限的，是假的
			} else if time.Now().Unix() > claim.ExpiresAt {
				code = 401 //Token无效了
			}
		}
		if code != 200 {
			c.JSON(200, gin.H{
				"status": code,
				"msg":    "Token解析错误",
			})
			c.Abort() //中止请求处理链
			return    //使用 return 来确保函数的即时退出
		}
		c.Next() //显式调用请求处理链中的下一个处理程序,制权传递给链中的下一个处理程序
	}
}