package types

// UserRegisterReq 用户注册请求
type UserRegisterReq struct {
	UserName string `form:"user_name" json:"user_name" binding:"required,min=3,max=15" example:"FanOne"`
	Password string `form:"password" json:"password" binding:"required,min=5,max=16" example:"FanOne666"`
}

// UserLoginReq 用户登录请求
type UserLoginReq struct {
	UserName string `form:"user_name" json:"user_name" binding:"required,min=3,max=15" example:"FanOne"`
	Password string `form:"password" json:"password" binding:"required,min=5,max=16" example:"FanOne666"`
}

// UserLoginResp 用户登录信息返回
type UserLoginResp struct {
	Id       int64  `json:"id"`
	UserName string `json:"user_name"`
	CreateAt int64  `json:"create_at"`
}

// TokenData 用户登录专属的，带用户信息和token的返回resp
type TokenData struct {
	User        any    `json:"user"`
	AccessToken string `json:"access_token"`
	//RefreshToken string `json:"refresh_token"`
}
