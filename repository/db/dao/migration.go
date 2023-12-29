package dao

import (
	model "todo_list_v2.01/repository/db/model"
)

// 执行数据迁移
func migration() {
	// 设置了表的选项、自动迁移模式
	err := _db.Set("gorm:table_options", "charset=utf8mb4").
		AutoMigrate(&model.TaskModel{}, &model.UserModel{})
	if err != nil {
		return
	}
	// DB.Model(&Task{}).AddForeignKey("uid","User(id)","CASCADE","CASCADE")
}
