package router

import (
	"github.com/Jie-cv/go_admin/service"
	"github.com/gin-gonic/gin"
)

func InitRoleRoutor(r *gin.Engine) {
	RoleGroup := r.Group("/role")
	RoleGroup.GET("/list", service.GetRoleList)
	RoleGroup.POST("/add", service.AddRole)
	RoleGroup.POST("/update", service.UpdateRole)
	RoleGroup.POST("/delete", service.DeleteRole)
}
