package dao

import (
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"todo_list_v2.01/consts"
)

type UserModel struct {
	gorm.Model
	Id             int64  `gorm:"primary"`
	UserName       string `gorm:"column:user_name;unique"`
	PasswordDigest string `gorm:"column:password_digest;"`
}

func (*UserModel) TableName() string {
	return "user"
}

// SetPassword 设置密码
func (user *UserModel) SetPassword(password string) error {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), consts.PasswordCost)
	if err != nil {
		return err
	}
	user.PasswordDigest = string(bytes)
	return nil
}

// CheckPassword 校验密码
func (user *UserModel) CheckPassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(user.PasswordDigest), []byte(password))
	return err == nil
}
