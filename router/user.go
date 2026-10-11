package router

import (
	"github.com/Jie-cv/go_admin/services"
	"github.com/gin-gonic/gin"
)

func InitUserRouter(r *gin.Engine) {
	userGroup := r.Group("/user")
	{
		userGroup.POST("/login", services.Login)
	}
}
