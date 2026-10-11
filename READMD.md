# go_admin

基于 Go + Gin + GORM 的后台管理系统。

## 环境要求

- Go 1.18+（当前 go.mod 指定 1.27.2）
- MySQL

## 安装依赖

```bash
go mod tidy
```

## 运行

```bash
go run main.go
```

服务默认监听 `:3382` 端口。

## 依赖命令

```bash
# Gin Web 框架
go get github.com/gin-gonic/gin@v1.12.0

# GORM ORM 框架
go get gorm.io/gorm@v1.31.2

# GORM MySQL 驱动
go get gorm.io/driver/mysql@v1.6.0

# JWT Token 生成与解析
go get github.com/dgrijalva/jwt-go@v3.2.0+incompatible
```