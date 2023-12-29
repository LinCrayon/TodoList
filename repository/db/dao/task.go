package dao

import (
	"context"
	"gorm.io/gorm"
	model "todo_list_v2.01/repository/db/model"
	"todo_list_v2.01/types"
)

type TaskDao struct {
	db *gorm.DB
}

func NewTaskDao(ctx context.Context) *TaskDao {
	if ctx == nil {
		ctx = context.Background()
	}
	return &TaskDao{NewDBClient(ctx)}
}

// ListTask 分页查询当前用户的备忘录
func (this *TaskDao) ListTask(start, limit int, uid int64) (taskMode []*model.TaskModel, total int64, err error) {
	err = this.db.Model(&model.TaskModel{}).Where("uid=?", uid).Count(&total).Error //统计符合指定条件的记录数
	if err != nil {
		return
	}
	err = this.db.Model(&model.TaskModel{}). //Model(&model.TaskModel{})指定查询的模型,查询与TaskModel模型相关联的表
							Where("uid=?", uid).Limit(limit).
							Offset((start - 1) * limit).Find(&taskMode).Error //Find()存储查询结果
	//(start - 1) 表示实际上是获取上一页的最后一条记录之后的数据，因为页码是从1开始的，而偏移量是从0开始的。
	//.Error：检查执行查询时是否发生错误，如果有错误，会将错误信息返回。
	return
}

// FindTaskByIdAndUserId 根据id和uid查询用户的单条task
func (this *TaskDao) FindTaskByIdAndUserId(id, userId int64) (taskModel *model.TaskModel, err error) {
	err = this.db.Model(&model.TaskModel{}).
		Where("id=? AND uid=?", id, userId).
		Find(&taskModel).Error
	return
}

// CreateTask 创建备忘录
func (s *TaskDao) CreateTask(in *model.TaskModel) error {
	return s.db.Create(in).Error
}

// UpdateTask 更新备忘录
func (this *TaskDao) UpdateTask(uid int64, req *types.UpdateTaskReq) error {
	//判断task是否存在
	var taskModel model.TaskModel
	err := this.db.Model(&model.TaskModel{}).
		Where("id=? AND uid=?", req.Id, uid).
		First(&taskModel).Error //根据id和uid查询是否存在该task
	if err != nil {
		return err
	}
	if req.Status != 0 { // 0 待办   1已完成
		taskModel.Status = req.Status
	}

	if req.Title != "" {
		taskModel.Title = req.Title
	}

	if req.Content != "" {
		taskModel.Content = req.Content
	}

	return this.db.Save(&taskModel).Error
}

// SearchTask 搜索Task（模糊查询） 查询标题、内容中包含xxx的task
func (this *TaskDao) SearchTask(uid int64, info string) (taskModel []*model.TaskModel, err error) {
	//todo Preload()预加载关联的数据,将预加载与任务相关联的用户数据，避免了在后续使用任务时再次查询用户信息。
	//err = this.db.Where("uid=?", uid).Preload("user").First(&taskModel).Error
	//if err != nil {
	//	return
	//}

	err = this.db.Model(&model.TaskModel{}).Where("title LIKE ? OR content LIKE ?", //检索标题或内容中包含指定字符串info的记录
		"%"+info+"%", "%"+info+"%").Find(&taskModel).Error //将查询结果存储在 taskModel 切片中
	return //将当前函数的返回值直接返回
}

// DeleteTaskByIdAndUserId  通过id、uid删除备忘录
func (this *TaskDao) DeleteTaskByIdAndUserId(id, uid int64) (err error) {
	err = this.db.Model(&model.TaskModel{}).
		Where("id=? AND uid=?", id, uid).
		Delete(&model.TaskModel{}).Error
	return
}
