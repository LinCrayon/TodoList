package api

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
	"todo_list/pkg/utils"
	"todo_list/service"
)

func CreateTask(c *gin.Context) {
	var createTask service.CreateTaskService
	claim, _ := utils.ParseToken(c.GetHeader("Authorization"))
	if err := c.ShouldBind(&createTask); err == nil {
		res := createTask.Create(claim.Id)
		c.JSON(http.StatusOK, res)
	} else {
		logrus.Error(err)
		c.JSON(http.StatusBadRequest, ErrorResponse(err))
	}
}

func ShowTask(c *gin.Context) {
	var showTask service.ShowTaskService
	if err := c.ShouldBind(&showTask); err == nil {
		res := showTask.Show(c.Param("id"))
		c.JSON(http.StatusOK, res)
	} else {
		logrus.Error(err)
		c.JSON(http.StatusBadRequest, ErrorResponse(err))
	}
}

func ListTask(c *gin.Context) {
	var listTask service.ListTaskService
	claim, _ := utils.ParseToken(c.GetHeader("Authorization"))
	if err := c.ShouldBind(&listTask); err == nil {
		res := listTask.List(claim.Id)
		c.JSON(http.StatusOK, res)
	} else {
		logrus.Error(err)
		c.JSON(http.StatusBadRequest, ErrorResponse(err))
	}
}
func UpdateTask(c *gin.Context) {
	var updateTask service.UpdateTaskService
	if err := c.ShouldBind(&updateTask); err == nil {
		res := updateTask.Update(c.Param("id"))
		c.JSON(http.StatusOK, res)
	} else {
		logrus.Error(err)
		c.JSON(http.StatusBadRequest, ErrorResponse(err))
	}
}

func SearchTask(c *gin.Context) {
	var searchTask service.SearchTaskService
	claim, _ := utils.ParseToken(c.GetHeader("Authorization"))
	if err := c.ShouldBind(&searchTask); err == nil {
		res := searchTask.Search(claim.Id)
		c.JSON(http.StatusOK, res)
	} else {
		logrus.Error(err)
		c.JSON(http.StatusBadRequest, ErrorResponse(err))
	}
}
func DeleteTask(c *gin.Context) {
	var deleteTask service.DeleteTaskService
	if err := c.ShouldBind(&deleteTask); err == nil {
		res := deleteTask.Delete(c.Param("id"))
		c.JSON(http.StatusOK, res)
	} else {
		logrus.Error(err)
		c.JSON(http.StatusBadRequest, ErrorResponse(err))
	}
}
