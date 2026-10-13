package router

import (
	"github.com/Jie-cv/go_admin/service"
	"github.com/gin-gonic/gin"
)

func InitUserRouter(r *gin.Engine) {
	userGroup := r.Group("/user")
	{
		userGroup.POST("/login", service.Login)
		userGroup.GET("/list", service.GetUserList)
		userGroup.POST("/add", service.AddUser)
		userGroup.POST("/update", service.UpdateUser)
		userGroup.POST("/delete", service.DeleteUser)
	}
}
