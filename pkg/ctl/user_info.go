package ctl

import (
	"context"
	"errors"
)

type key int

var userKey key //todo 上下文中存储用户信息的键。

type UserInfo struct {
	Id       int64  `json:"id"`
	UserName string `json:"user_name"`
}

// GetUserInfo 用于从上下文中获取用户信息
func GetUserInfo(ctx context.Context) (*UserInfo, error) {
	user, ok := FromContext(ctx)
	if !ok {
		return nil, errors.New("获取用户信息错误")
	}
	return user, nil
}

// NewContext 创建一个新的 context,存储用户信息
func NewContext(ctx context.Context, u *UserInfo) context.Context {
	return context.WithValue(ctx, userKey, u) //创建带有键值对数据的context
}

// FromContext 获取用户信息
func FromContext(ctx context.Context) (*UserInfo, bool) {
	u, ok := ctx.Value(userKey).(*UserInfo) //根据UserInfo的id获取用户
	return u, ok
}
