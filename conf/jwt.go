package conf

import (
	"time"

	"github.com/dgrijalva/jwt-go"
)

type JwtConfig struct {
	JwtKey             string
	AccessTokenExpire  int64
	RefreshTokenExpire int64
}

var Jwt = JwtConfig{
	JwtKey:             "338286",
	AccessTokenExpire:  time.Now().Add(time.Second * 3600 * 1).Unix(),      // 请求token
	RefreshTokenExpire: time.Now().Add(time.Second * 3600 * 24 * 7).Unix(), // 登录token
}

// 用户权限
type UserClaim struct {
	Id      uint
	Name    string
	IsAdmin bool
	jwt.StandardClaims
}
