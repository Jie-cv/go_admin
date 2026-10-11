package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Cors() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		method := ctx.Request.Method
		ctx.Header("Access-Control-Allow-Origin", "*")
		ctx.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, PATCH")
		ctx.Header("Access-Control-Allow-Headers", "Content-Type, Authorization, Origin,X-Requested-With,Accept,AccessToken,RefreshToken")
		ctx.Header("Access-Control-Expose-Headers", "Authorization,Content-Length,AccessToken,RefreshToken,X-Request-Id,Response-Time,X-Response-Status")
		ctx.Header("Access-Control-Allow-Credentials", "true")
		if method == "OPTIONS" {
			ctx.AbortWithStatus(http.StatusNoContent)
		}
		ctx.Next()
	}
}
