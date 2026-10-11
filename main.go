package main

import (
	"fmt"

	"github.com/Jie-cv/go_admin/model"
	"github.com/Jie-cv/go_admin/router"
)

func main() {
	model.InitDB()
	r := router.AppRouter()
	fmt.Println("路由初始化成功....")
	r.Run(":3382")
	fmt.Println("服务启动成功....")
}
