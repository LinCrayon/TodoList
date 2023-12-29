package dao

import (
	"github.com/spf13/cast"
	"gorm.io/gorm"
	"todo_list_v2.01/repository/cache"
)

type TaskModel struct {
	gorm.Model
	Id        int64  `gorm:"column:id;primary_key;"`
	Uid       int64  `gorm:"column:uid;"`
	Title     string `gorm:"column:title;"`
	Content   string `gorm:"column:content;"`
	StartTime int64  `gorm:"column:start_time;"`
	EndTime   int64  `gorm:"column:end_time;"`
	Status    int    `form:"status" json:"status"` // 0 待办   1已完成
}

func (*TaskModel) TableName() string {
	return "task"
}

func (Task *TaskModel) View() int64 {
	// 增加点击数
	countStr, _ := cache.RedisClient.Get(cache.TaskViewKey(Task.Id)).Result()
	return cast.ToInt64(countStr)
}

// AddView 增加点击
func (Task *TaskModel) AddView() {
	cache.RedisClient.Incr(cache.TaskViewKey(Task.Id)) // 增加视频点击数
}
