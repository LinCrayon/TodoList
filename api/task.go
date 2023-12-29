package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"todo_list_v2.01/consts"
	"todo_list_v2.01/pkg/utils"
	"todo_list_v2.01/service"
	"todo_list_v2.01/types"
)

func CreateTaskHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req types.CreateTaskReq
		if err := ctx.ShouldBind(&req); err == nil { //请求数据绑定结构体 。根据请求的Content-Type自动选择合适的绑定器（Binder）进行数据绑定。
			// 参数校验
			l := service.GetTaskSrv() //单例对象的安全初始化
			resp, err := l.CreateTask(ctx.Request.Context(), &req)
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

// ListTaskHandler 分页查询当前用户的备忘录
func ListTaskHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req types.ListTaskReq
		if err := ctx.ShouldBind(&req); err == nil {
			// 参数校验
			if req.Limit == 0 {
				req.Limit = consts.BasePageSize
			}
			l := service.GetTaskSrv()
			resp, err := l.ListTask(ctx.Request.Context(), &req)
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

func ShowTaskHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req types.ShowTaskReq
		if err := ctx.ShouldBind(&req); err == nil {
			// 参数校验
			l := service.GetTaskSrv()
			resp, err := l.ShowTask(ctx.Request.Context(), &req)
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

// UpdateTaskHandler 更新备忘录
func UpdateTaskHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req types.UpdateTaskReq
		if err := ctx.ShouldBind(&req); err == nil {
			// 参数校验
			l := service.GetTaskSrv()
			resp, err := l.UpdateTask(ctx.Request.Context(), &req)
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

// SearchTaskHandler 搜索备忘录
func SearchTaskHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req types.SearchTaskReq
		if err := ctx.ShouldBind(&req); err == nil {
			// 参数校验
			l := service.GetTaskSrv()
			resp, err := l.SearchTask(ctx.Request.Context(), &req)
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
func DeleteTaskHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req types.DeleteTaskReq
		if err := ctx.ShouldBind(&req); err == nil {
			// 参数校验
			l := service.GetTaskSrv()
			resp, err := l.DeleteTask(ctx.Request.Context(), &req)
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
