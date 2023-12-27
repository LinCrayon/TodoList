package model

func migrate() {
	DB.Set("gorm:table_options", "charset=utf8mb4").
		AutoMigrate(&User{}).
		AutoMigrate(&Task{})
	DB.Model(&Task{}).AddForeignKey("uid", "User(id)", "CASCADE", "CASCADE")
}

//添加了一个外键约束，将 Task 表中的 uid 字段与 User 表中的 id 字段关联
//&Task{}：表示要在 Task 模型上执行操作。
//"CASCADE"：表示在主键更新或删除时，相关的外键约束也会被更新或删除。
