package service

import (
	"net/http"

	"github.com/Jie-cv/go_admin/contant/code"
	"github.com/Jie-cv/go_admin/model"
	"github.com/Jie-cv/go_admin/utils"
	"github.com/gin-gonic/gin"
)

func GetUserList(ctx *gin.Context) {
	in := &GetUserListRequest{NewQueryRequest()}
	err := ctx.ShouldBind(in)
	if err != nil {
		ctx.JSON(http.StatusOK, utils.FormateError(code.Fail, "参数异常"))
		return
	}
	var (
		cnt       int64
		user_list = make([]*GetUserListRsponse, 0)
	)
	err = model.GetUserList(in.Keyword).Count(&cnt).Offset((in.Page - 1) * in.Size).Limit(in.Size).Find(&user_list).Error
	if err != nil {
		ctx.JSON(http.StatusOK, utils.FormateError(code.ServerError, "数据库异常"))
		return
	}
	ctx.JSON(http.StatusOK, utils.FormateSuccess(code.Success, "数据获取成功", gin.H{
		"user_list": user_list,
		"count":     cnt,
	}))
}
