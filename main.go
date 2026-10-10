package main

import (
	"github.com/Jie-cv/go_admin/router"
)

func main() {
	r := router.AppRouter()
	r.Run(":3382")
}
