package service

import (
	"context"
	"sync"
	"time"
	"todo_list_v2.01/pkg/ctl"
	"todo_list_v2.01/pkg/utils"
	"todo_list_v2.01/repository/db/dao"
	model "todo_list_v2.01/repository/db/model"
	"todo_list_v2.01/types"
)

var TaskSrvIns *TaskSrv

var TaskSrvOnce sync.Once

type TaskSrv struct {
}

// GetTaskSrv 单例对象的安全初始化
func GetTaskSrv() *TaskSrv {
	TaskSrvOnce.Do(func() { //接收一个函数作为参数，这个函数将被确保只执行一次
		TaskSrvIns = &TaskSrv{}
	})
	return TaskSrvIns
}

// CreateTask 创建备忘录
func (s *TaskSrv) CreateTask(ctx context.Context, req *types.CreateTaskReq) (resp any, err error) {
	u, err := ctl.GetUserInfo(ctx) //获取用户信息根据content中的id
	if err != nil {
		utils.LogrusObj.Info(err)
		return
	}
	user, err := dao.NewUserDao(ctx).FindUserByUserId(u.Id) //content获取的用户Id查
	if err != nil {
		utils.LogrusObj.Info(err)
		return
	}
	task := &model.TaskModel{
		Uid:       user.Id,
		Title:     req.Title,
		Content:   req.Content,
		Status:    req.Status,
		StartTime: time.Now().Unix(),
	}
	err = dao.NewTaskDao(ctx).CreateTask(task) //获取用户连接数据库的content，创建用户
	if err != nil {
		utils.LogrusObj.Info(err)
		return
	}
	return ctl.RespSuccess(), nil
}

// ListTask 分页查询当前用户的备忘录（带总条数）
func (s *TaskSrv) ListTask(ctx context.Context, req *types.ListTaskReq) (resp any, err error) {
	u, err := ctl.GetUserInfo(ctx) //获取用户
	if err != nil {
		utils.LogrusObj.Info(err)
		return
	}
	tasks, total, err := dao.NewTaskDao(ctx).ListTask(req.Start, req.Limit, u.Id) //用户data,条数
	tRespList := make([]*types.ListTaskResp, 0)                                   //创建切片
	for _, v := range tasks {
		tRespList = append(tRespList, &types.ListTaskResp{ //切片存储用户data
			Id:        v.Id,
			Title:     v.Title,
			Content:   v.Content,
			View:      v.View(),
			Status:    v.Status,
			CreatedAt: v.CreatedAt.Unix(),
			StartTime: v.StartTime,
			EndTime:   v.EndTime,
		})
	}
	return ctl.RespList(tRespList, total), nil
}

// ShowTask 根据id和uid查询当前用户的单条task
func (s *TaskSrv) ShowTask(ctx context.Context, req *types.ShowTaskReq) (resp interface{}, err error) {
	u, err := ctl.GetUserInfo(ctx)
	if err != nil {
		utils.LogrusObj.Info(err)
		return
	}
	task, err := dao.NewTaskDao(ctx).FindTaskByIdAndUserId(req.Id, u.Id)
	if err != nil {
		utils.LogrusObj.Info(err)
		return
	}
	task.AddView() //增加点击次数
	tResp := &types.ListTaskResp{
		Id:        task.Id,
		Title:     task.Title,
		Content:   task.Content,
		View:      task.View(), //task_show点击数
		Status:    task.Status,
		CreatedAt: task.CreatedAt.Unix(),
		StartTime: task.StartTime,
		EndTime:   task.EndTime,
	}
	return ctl.RespSuccessWithData(tResp), nil //带data成功返回
}

// UpdateTask 更新备忘录
func (s *TaskSrv) UpdateTask(ctx context.Context, req *types.UpdateTaskReq) (resp any, err error) {
	u, err := ctl.GetUserInfo(ctx)
	if err != nil {
		utils.LogrusObj.Info(err)
		return
	}
	err = dao.NewTaskDao(ctx).UpdateTask(u.Id, req)
	if err != nil {
		utils.LogrusObj.Info(err)
	}
	return ctl.RespSuccess(), nil
}

// SearchTask 搜索备忘录
func (s *TaskSrv) SearchTask(ctx context.Context, req *types.SearchTaskReq) (resp interface{}, err error) {
	u, err := ctl.GetUserInfo(ctx)
	if err != nil {
		utils.LogrusObj.Info(err)
		return
	}

	tasks, err := dao.NewTaskDao(ctx).SearchTask(u.Id, req.Info)
	if err != nil {
		utils.LogrusObj.Info(err)
	}
	tRespList := make([]*types.ListTaskResp, 0)
	for _, v := range tasks {
		tRespList = append(tRespList, &types.ListTaskResp{
			Id:        v.Id,
			Title:     v.Title,
			Content:   v.Content,
			View:      v.View(),
			Status:    v.Status,
			CreatedAt: v.CreatedAt.Unix(),
			StartTime: v.StartTime,
			EndTime:   v.EndTime,
		})
	}
	return ctl.RespSuccessWithData(tRespList), nil
}

// DeleteTask 删除用户
func (s *TaskSrv) DeleteTask(ctx context.Context, req *types.DeleteTaskReq) (resp interface{}, err error) {
	u, err := ctl.GetUserInfo(ctx)
	if err != nil {
		utils.LogrusObj.Info(err)
		return
	}
	err = dao.NewTaskDao(ctx).DeleteTaskByIdAndUserId(req.Id, u.Id)
	if err != nil {
		utils.LogrusObj.Info(err)
		return
	}
	return ctl.RespSuccess(), nil
}
