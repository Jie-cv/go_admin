package service

import (
	"errors"
	"net/http"

	"github.com/Jie-cv/go_admin/conf"
	"github.com/Jie-cv/go_admin/contant/code"
	"github.com/Jie-cv/go_admin/helper"
	"github.com/Jie-cv/go_admin/model"
	"github.com/Jie-cv/go_admin/utils"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Login(ctx *gin.Context) {
	in := new(LoginRequest)
	// in := &LoginRequest{}
	err := ctx.ShouldBindJSON(in)
	if err != nil {
		ctx.JSON(http.StatusOK, utils.FormateError(code.Fail, err.Error()))
		return
	}
	sysUser, err := model.GetUserByUsernamePwd(in.UserName, in.Password)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			ctx.JSON(http.StatusOK, utils.FormateError(code.Fail, "用户名或密码错误"))
			return
		}
		ctx.JSON(http.StatusOK, utils.FormateError(code.Fail, err.Error()))
		return
	}
	accessToken, err := helper.GenerateToken(sysUser.ID, sysUser.UserName, conf.Jwt.AccessTokenExpire)
	if err != nil {
		ctx.JSON(http.StatusOK, utils.FormateError(code.Fail, err.Error()))
		return
	}
	refreshToken, err := helper.GenerateToken(sysUser.ID, sysUser.UserName, conf.Jwt.RefreshTokenExpire)
	if err != nil {
		ctx.JSON(http.StatusOK, utils.FormateError(code.Fail, err.Error()))
		return
	}
	ctx.JSON(http.StatusOK, utils.FormateSuccess(code.Success, "登录成功", &LoginResonse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}))
}
