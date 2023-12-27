package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"todo_list/service"
)

// UserRegister 用户注册
func UserRegister(c *gin.Context) {
	var userRegister service.UserService
	//ShouldBind用于将请求中的数据绑定到指定的结构体变量上,会根据请求的Content-Type自动选择适当的绑定器（如 form绑定、JSON绑定等）
	if err := c.ShouldBind(&userRegister); err != nil {
		res := userRegister.Register()
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusBadRequest, ErrorResponse(err))
	}
}

// UserLogin 用户登录
func UserLogin(c *gin.Context) {
	var userLogin service.UserService
	if err := c.ShouldBind(&userLogin); err != nil {
		res := userLogin.Login()
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusBadRequest, ErrorResponse(err))
	}
}
