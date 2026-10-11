package conf

import "fmt"

// DBConfig 数据库配置
type DBConfig struct {
	UserName string
	Password string
	Host     string
	Port     string
	DBName   string
	Charset  string
}

// DB 全局数据库配置
var DB = DBConfig{
	UserName: "root",
	Password: "338286",
	Host:     "127.0.0.1",
	Port:     "3306",
	DBName:   "go_admin",
	Charset:  "utf8mb4",
}

// DSN 生成数据库连接字符串
func (c DBConfig) DSN() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=True&loc=Local",
		c.UserName, c.Password, c.Host, c.Port, c.DBName, c.Charset)
}
