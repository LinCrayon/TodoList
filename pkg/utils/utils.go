package utils

import (
	"github.com/golang-jwt/jwt/v4"
	"time"
)

/*JWT三部分组成：头部（Header）、载荷（Payload）和签名（Signature）*/

var JWTsecret = []byte("ABAB") //JWT的签名密钥，用于签名和验证

type Claims struct {
	Id       uint   `json:"id"`
	UserName string `json:"user_name"`
	Password string `json:"password"`
	jwt.StandardClaims
}

// GenerateToken 签发token
func GenerateToken(id uint, username, password string) (string, error) {
	notTime := time.Now()
	expireTime := notTime.Add(24 * time.Hour)
	claims := Claims{
		Id:       id,
		UserName: username,
		Password: password,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer:    "todo_list",
		},
	}
	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims) //创建JWT，claims 通常包含了 JWT 的有效负载（Payload）
	token, err := tokenClaims.SignedString(JWTsecret)                //对创建的 JWT 进行签名
	return token, err
}

// ParseToken 验证token
func ParseToken(token string) (*Claims, error) {
	//空Claims 结构的实例，用于存储解析后的声明（payload）。解析密钥的逻辑函数，用JWTsecret解析
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return JWTsecret, nil
	})
	if tokenClaims != nil { //解析和验证过程没有发生错误。 //类型断言
		if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid { //tokenClaims.Valid检查JWT的有效性
			return claims, nil
		}
	}
	return nil, err
}
