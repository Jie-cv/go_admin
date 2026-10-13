package router

import (
	"net/http"

	"github.com/Jie-cv/go_admin/middleware"
	"github.com/gin-gonic/gin"
)

func AppRouter() *gin.Engine {
	r := gin.Default()
	r.Use(middleware.Cors())
	r.LoadHTMLGlob("templates/*") // 加载模板文件
	r.GET("/", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "index.html", gin.H{})
	})
	InitUserRouter(r) // 用户相关路由
	InitRoleRoutor(r) // 角色相关路由
	return r
}
