package router

import (
	"github.com/gin-gonic/gin"
)

func AppRouter() *gin.Engine {
	r := gin.Default()
	return r
}
