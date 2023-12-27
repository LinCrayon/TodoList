package service

import (
	"github.com/jinzhu/gorm"
	"todo_list/model"
	"todo_list/pkg/e"
	"todo_list/pkg/utils"
	"todo_list/serializer"
)

type UserService struct {
	UserName string `form:"user_name" json:"user_name" binding:"required,min=3,max=15"`
	Password string `form:"password" json:"password" binding:"required,min=5,max=16"`
}

func (service *UserService) Register() serializer.Response {
	var user model.User
	var count int
	//&model.User{}用于指定要操作的数据库表以及该表对应的模型
	model.DB.Model(&model.User{}).Where("user_name=?", service.UserName).
		First(&user).Count(&count)
	//First(&user)从数据库中获取第一条匹配的记录，并将其存储到user变量中
	//Count(&count)计算满足前一个查询条件的记录数
	if count == 1 {
		return serializer.Response{
			Status: e.ErrorExistUser,
			Msg:    "用户已存在，无需注册",
		}
	}
	user.UserName = service.UserName
	//对密码加密
	if err := user.SetPassword(service.Password); err != nil {
		return serializer.Response{
			Status: e.ErrorAuthCheckTokenFail,
			Msg:    err.Error(),
		}
	}

	//创建用户
	if err := model.DB.Create(&user).Error; err != nil {
		return serializer.Response{
			Status: e.ErrorDatabase,
			Msg:    "数据库操作错误",
		}
	}
	return serializer.Response{
		Status: e.SUCCESS,
		Msg:    "用户注册成功",
	}

}

func (service *UserService) Login() serializer.Response {
	var user model.User
	//先去db看看有没有这个人
	//Error: 这是GORM提供的链式调用，用于获取执行查询时的错误。如果查询成功，err 将为 nil；
	if err := model.DB.Where("user_name=?", service.UserName).First(&user).Error; err != nil {
		//gorm.IsRecordNotFoundError()用于检查错误是否是由于查询未找到记录而引起的。
		if gorm.IsRecordNotFoundError(err) { //如果命令不存在
			return serializer.Response{
				Status: e.ErrorNotExistUser,
				Msg:    "用户不存在，先登录",
			}
		}
		//如果不是用户不存在，是其他不可抗因素导致错误
		return serializer.Response{
			Status: e.ErrorDatabase,
			Msg:    "数据库错误",
		}
	}
	if user.CheckPassword(service.Password) == false {
		return serializer.Response{
			Status: e.ErrorAuthCheckTokenFail,
			Msg:    "密码错误",
		}
	}
	//发送一个token，为了其他功能需要身份验证给前端存储的。
	//创建备忘录需要token
	token, err := utils.GenerateToken(user.ID, service.UserName, service.Password)
	if err != nil {
		return serializer.Response{
			Status: e.ErrorAuthToken,
			Msg:    "Token签发错误",
		}
	}
	return serializer.Response{
		Status: e.SUCCESS,
		Data:   serializer.TokenData{User: serializer.BuildUser(user), Token: token},
		Msg:    "登录成功",
	}
}
