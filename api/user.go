package api

import (
	"net/http"
	"todo_list_v2.01/pkg/utils"
	"todo_list_v2.01/service"
	"todo_list_v2.01/types"

	"github.com/gin-gonic/gin"
)

func UserRegisterHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req types.UserRegisterReq
		if err := ctx.ShouldBind(&req); err == nil {
			// 参数校验
			l := service.GetUserSrv()
			resp, err := l.UserRegister(ctx.Request.Context(), &req)
			if err != nil {
				ctx.JSON(http.StatusInternalServerError, ErrorResponse(err))
				return
			}
			ctx.JSON(http.StatusOK, resp)
		} else {
			utils.LogrusObj.Infoln(err)
			ctx.JSON(http.StatusBadRequest, ErrorResponse(err))
		}
	}
}

func UserLoginHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req types.UserLoginReq
		if err := ctx.ShouldBind(&req); err == nil {
			// 参数校验
			l := service.GetUserSrv()
			resp, err := l.UserLogin(ctx.Request.Context(), &req)
			if err != nil {
				ctx.JSON(http.StatusInternalServerError, ErrorResponse(err))
				return
			}
			ctx.JSON(http.StatusOK, resp)
		} else {
			utils.LogrusObj.Infoln(err)
			ctx.JSON(http.StatusBadRequest, ErrorResponse(err))
		}
	}
}
