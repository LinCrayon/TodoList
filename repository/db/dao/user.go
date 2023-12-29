package dao

import (
	"context"
	"gorm.io/gorm"
	model "todo_list_v2.01/repository/db/model"
)

type userDao struct {
	db *gorm.DB
}

func NewUserDao(ctx context.Context) *userDao {
	if ctx == nil {
		/*
			背景上下文通常用于作为整个应用程序的根上下文，
			或者在没有更好的上下文可用时作为默认的上下文。
			由于背景上下文是空的，它不包含任何值，也没有截止时间。
			其他上下文可以通过派生（Derive）操作从背景上下文中创建，并在需要传递上下文信息的地方使用。
		*/
		ctx = context.Background()
	}
	return &userDao{NewDBClient(ctx)} //用传入的上下文对象初始化数据库客户端
}

// FindUserByUserName 根据用户名查询用户
func (s *userDao) FindUserByUserName(userName string) (user *model.UserModel, err error) {
	err = s.db.Model(&model.UserModel{}).
		Where("user_name = ?", userName).
		First(&user).Error
	return
}

// FindUserByUserId 根据Id查询用户
func (s *userDao) FindUserByUserId(userId int64) (user *model.UserModel, err error) {
	err = s.db.Model(&model.UserModel{}).
		Where("id = ?", userId).
		First(&user).Error
	return
}

// CreateUser 创建用户
func (s *userDao) CreateUser(in *model.UserModel) error {
	return s.db.Create(in).Error
}
