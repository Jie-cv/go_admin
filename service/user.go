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
		ctx.JSON(http.StatusOK, utils.FormateError(code.ServerError, err.Error()))
		return
	}
	ctx.JSON(http.StatusOK, utils.FormateSuccess(code.Success, "数据获取成功", gin.H{
		"user_list": user_list,
		"count":     cnt,
	}))
}

func AddUser(ctx *gin.Context) {
	in := new(AddUserRequest)
	err := ctx.ShouldBindJSON(in)
	if err != nil {
		ctx.JSON(http.StatusOK, utils.FormateError(code.Fail, "参数异常"))
		return
	}
	var cnt int64
	err = model.DB.Model(new(model.SysUser)).Where("username = ?", in.Username).Count(&cnt).Error
	if cnt > 0 {
		ctx.JSON(http.StatusOK, utils.FormateError(code.Fail, "用户已存在"))
		return
	}
	err = model.DB.Create(&model.SysUser{
		UserName: in.Username,
		Password: in.Password,
		Phone:    in.Phone,
		Email:    in.Email,
		Remarks:  in.Remarks,
	}).Error

	if err != nil {
		ctx.JSON(http.StatusOK, utils.FormateError(code.ServerError, "添加失败"))
		return
	}
	ctx.JSON(http.StatusOK, utils.FormateSuccess(code.Success, "添加成功", gin.H{}))
}
func UpdateUser(ctx *gin.Context) {
	in := new(UpdateUserRequest)
	err := ctx.ShouldBindJSON(in)
	if err != nil {
		ctx.JSON(http.StatusOK, utils.FormateError(code.Fail, "参数异常"))
		return
	}
	var cnt int64
	// 用户名不存在
	err = model.DB.Model(new(model.SysUser)).Where("id != ? AND username = ?", in.ID, in.Username).Count(&cnt).Error
	if err != nil {
		ctx.JSON(http.StatusOK, utils.FormateError(code.ServerError, err.Error()))
		return
	}
	if cnt > 0 {
		ctx.JSON(http.StatusOK, utils.FormateError(code.Fail, "用户名存在"))
		return
	}
	err = model.DB.Model(new(model.SysUser)).Where("id = ?", in.ID).Updates(map[string]interface{}{
		"password": in.Password,
		"phone":    in.Phone,
		"email":    in.Email,
		"remarks":  in.Remarks,
		"username": in.Username,
	}).Error
	if err != nil {
		ctx.JSON(http.StatusOK, utils.FormateError(code.ServerError, err.Error()))
		return
	}
	ctx.JSON(http.StatusOK, utils.FormateSuccess(code.Success, "更新成功", gin.H{}))
}
