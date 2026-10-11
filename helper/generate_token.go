package helper

import (
	"github.com/Jie-cv/go_admin/conf"
	"github.com/dgrijalva/jwt-go"
)

func GenerateToken(id uint, name string, expireAt int64) (string, error) {
	uc := conf.UserClaim{
		Id:        id,
		Name:      name,
		ExpiresAt: expireAt,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, uc)
	tokenString, err := token.SignedString([]byte(conf.Jwt.JwtKey))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}
