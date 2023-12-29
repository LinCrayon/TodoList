package router

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"todo_list_v2.01/api"
	"todo_list_v2.01/middleware"
)

func NewRouter() *gin.Engine {
	ginRouter := gin.Default()
	ginRouter.Use(middleware.Cors())
	//store作为用于加密和验证cookie的密钥。这个密钥用于对会话数据进行签名，以确保数据在传输过程中不被篡改。
	store := cookie.NewStore([]byte("something-very-secret")) //创建了一个Cookie存储（store）
	ginRouter.Use(sessions.Sessions("mysession", store))      //创建一个用于管理会话的中间件,两个参数(会话的名称 , 会话的存储介质),中间件添加到你的应用程序中，以便在处理HTTP请求时启用会话管理
	v1 := ginRouter.Group("/api/v1")
	{
		v1.GET("ping", func(context *gin.Context) {
			context.JSON(200, "success")
		})
		v1.POST("user/register", api.UserRegisterHandler())
		v1.POST("user/login", api.UserLoginHandler())

		authed := v1.Group("/")      //需要登录验证
		authed.Use(middleware.JWT()) //登录JWT校验中间件
		{
			authed.POST("task_create", api.CreateTaskHandler()) //创建备忘录
			authed.GET("task_list", api.ListTaskHandler())      //Get分页查询当前用户的全部备忘录（带总条数）
			authed.GET("task_show", api.ShowTaskHandler())      //根据id和uid查询当前用户的单条task
			authed.POST("task_update", api.UpdateTaskHandler()) //根据id和uid更新备忘录
			authed.POST("task_search", api.SearchTaskHandler()) //查询标题、内容中包含xxx的task
			authed.POST("task_delete", api.DeleteTaskHandler()) //删除task根据id和uid
		}
	}
	return ginRouter
}

/*
"Context"（上下文）是一个用于在请求处理函数之间传递数据的对象，而不是整个服务共享的数据。
上下文对象是短暂的,仅在请求处理的生命周期内存在。
每个 HTTP 请求都会创建一个新的上下文对象，该对象在整个请求处理周期内存在，用于存储请求相关的信息。
当请求处理完成后，这个上下文对象会被销毁，其中存储的数据也会随之释放。
*/
