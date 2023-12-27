package routes

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"todo_list/api"
	"todo_list/middleware"
)

func NewRouter() *gin.Engine {
	r := gin.Default()
	store := cookie.NewStore([]byte("something-very-secret")) //创建了cookie 的存储，用于会话管理。NewStore 函数的参数是一个用于签名和验证会话 cookie 完整性的秘钥。
	r.Use(sessions.Sessions("mysession", store))              //会话信息将存储在客户端的 cookie中，
	v1 := r.Group("api/v1")
	{
		//用户操作
		v1.POST("user/register", api.UserRegister)
		v1.POST("user/login", api.UserLogin)

		authed := v1.Group("/")
		authed.Use(middleware.JWT()) // 全局中间件
		{
			authed.POST("task", api.CreateTask)
			authed.GET("task/:id", api.ShowTask)
			authed.GET("tasks", api.ListTask)
			//更新备忘录
			authed.PUT("tasks/:id", api.UpdateTask)
			//查询备忘录
			authed.POST("search", api.SearchTask)
			//删除备忘录
			authed.DELETE("task/:id", api.DeleteTask)
		}

	}

	return r
}
