package models

import (
	"fmt"

	"github.com/Jie-cv/go_admin/conf"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	// 初始化数据库
	dsn := conf.DB.DSN()
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true,
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("数据库连接成功....")
	// 自动建表
	err = db.AutoMigrate(&SysUser{})
	if err != nil {
		panic(err)
	}
	fmt.Println("建表成功....")
	DB = db
}
