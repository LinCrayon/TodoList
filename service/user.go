package service

import (
	"context"
	"errors"
	"gorm.io/gorm"
	"sync"
	"todo_list_v2.01/pkg/ctl"
	"todo_list_v2.01/pkg/utils"
	"todo_list_v2.01/repository/db/dao"
	model "todo_list_v2.01/repository/db/model"
	"todo_list_v2.01/types"
)

var UserSrvIns *UserSrv

// UserSrvOnce 确保某个操作只执行一次的机制
var UserSrvOnce sync.Once //todo sync.Once来实现对UserSrv单例对象的安全初始化

type UserSrv struct {
}

// GetUserSrv 获取 UserSrv 的单例对象,只会在第一次调用时被初始化
func GetUserSrv() *UserSrv {
	UserSrvOnce.Do(func() { //todo Do()方法接收一个函数作为参数，这个函数将被确保只执行一次
		UserSrvIns = &UserSrv{}
	})
	return UserSrvIns
	//完成单例对象的初始化后。在后续的调用中，由于sync.Once的特性，不在初始化，而是直接返回已经初始化好的 UserSrvIns。
}

// UserRegister 用户注册
func (s *UserSrv) UserRegister(ctx context.Context, req *types.UserRegisterReq) (resp interface{}, err error) {
	userDao := dao.NewUserDao(ctx)
	u, err := userDao.FindUserByUserName(req.UserName)
	switch err {
	case gorm.ErrRecordNotFound: //数据库查询操作中没有找到符合条件的记录
		u = &model.UserModel{
			UserName: req.UserName,
		}
		// 密码加密存储
		if err = u.SetPassword(req.Password); err != nil {
			utils.LogrusObj.Info(err)
			return
		}
		//创建用户
		if err = userDao.CreateUser(u); err != nil {
			utils.LogrusObj.Info(err)
			return
		}
		return ctl.RespSuccess(), nil
	case nil:
		err = errors.New("用户已存在")
		return
	default:
		return
	}
}

// UserLogin 用户登录
func (s *UserSrv) UserLogin(ctx context.Context, req *types.UserLoginReq) (resp any, err error) {
	userDao := dao.NewUserDao(ctx)                     //查看上下文是否存在，不存在则创建context.Background()
	u, err := userDao.FindUserByUserName(req.UserName) //根据用户名查询用户是否存在
	if err != nil {
		utils.LogrusObj.Info(err)
		return
	}
	//校验密码
	if !u.CheckPassword(req.Password) {
		err = errors.New("账号/密码错误")
		return
	}
	//todo 分发token 用户信息放在token里发送给前端，让前端保存，请求带上token验证
	token, err := utils.GenerateToken(u.Id, u.UserName) //创建token
	if err != nil {
		utils.LogrusObj.Info(err)
		return
	}
	//UserInfo token
	userResp := types.TokenData{
		User: types.UserLoginResp{
			Id:       u.Id,
			UserName: u.UserName,
			CreateAt: u.CreatedAt.Unix(),
		},
		AccessToken: token,
	}
	return ctl.RespSuccessWithData(userResp), nil //带data成功返回

}
