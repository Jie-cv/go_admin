package service

import (
	"net/http"
	"time"

	"github.com/Jie-cv/go_admin/contant/code"
	"github.com/Jie-cv/go_admin/model"
	"github.com/Jie-cv/go_admin/utils"
	"github.com/gin-gonic/gin"
)

func GetRoleList(ctx *gin.Context) {
	in := &RoleListRequest{NewQueryRequest()}
	err := ctx.ShouldBind(in)
	if err != nil {
		ctx.JSON(http.StatusOK, utils.FormateError(code.Fail, "参数异常"))
		return
	}
	var (
		cnt       int64 = 0
		role_list       = make([]*RoleListResponse, 0)
	)
	err = model.GetRoleList(in.Keyword).Count(&cnt).Offset((in.Page - 1) * in.Size).Limit(in.Size).Find(&role_list).Error
	if err != nil {
		ctx.JSON(http.StatusOK, utils.FormateError(code.ServerError, err.Error()))
		return
	}
	ctx.JSON(http.StatusOK, utils.FormateSuccess(code.Success, "查询成功", gin.H{
		"cnt":       cnt,
		"role_list": role_list,
	}))
}

func AddRole(ctx *gin.Context) {
	in := new(AddRoleRequest)
	err := ctx.ShouldBindJSON(in)
	if err != nil {
		ctx.JSON(http.StatusOK, utils.FormateError(code.Fail, "参数异常"))
		return
	}
	var cnt int64
	err = model.DB.Model(new(model.SysRole)).Where("name = ?", in.Name).Count(&cnt).Error
	if err != nil {
		ctx.JSON(http.StatusOK, utils.FormateError(code.ServerError, err.Error()))
		return
	}
	if cnt > 0 {
		ctx.JSON(http.StatusOK, utils.FormateError(code.Fail, "角色名称已存在"))
		return
	}
	err = model.DB.Create(&model.SysRole{
		IsAdmin: in.Is_Admin,
		Name:    in.Name,
		Sort:    in.Sort,
		Remarks: in.Remarks,
	}).Error
	if err != nil {
		ctx.JSON(http.StatusOK, utils.FormateError(code.ServerError, err.Error()))
		return
	}
	ctx.JSON(http.StatusOK, utils.FormateSuccess(code.Success, "添加角色成功", gin.H{}))
}

func UpdateRole(ctx *gin.Context) {
	in := new(UpdateRoleRequest)
	err := ctx.ShouldBindJSON(in)
	if err != nil {
		ctx.JSON(http.StatusOK, utils.FormateError(code.Fail, "参数异常"))
		return
	}
	var cnt int64
	err = model.DB.Model(new(model.SysRole)).Where("id != ? AND name = ?", in.ID, in.Name).Count(&cnt).Error
	if err != nil {
		ctx.JSON(http.StatusOK, utils.FormateError(code.ServerError, err.Error()))
		return
	}
	if cnt > 0 {
		ctx.JSON(http.StatusOK, utils.FormateError(code.Fail, "角色名称已存在"))
		return
	}
	err = model.DB.Model(new(model.SysRole)).Where("id = ?", in.ID).Updates(map[string]any{
		"name":       in.Name,
		"is_admin":   in.Is_Admin,
		"remarks":    in.Remarks,
		"sort":       in.Sort,
		"updated_at": time.Now(),
	}).Error
	if err != nil {
		ctx.JSON(http.StatusOK, utils.FormateError(code.ServerError, err.Error()))
		return
	}
	ctx.JSON(http.StatusOK, utils.FormateSuccess(code.Success, "更新角色成功", gin.H{}))
}

func DeleteRole(ctx *gin.Context) {
	in := new(DeleteRoleRequest)
	err := ctx.ShouldBindJSON(in)
	if err != nil {
		ctx.JSON(http.StatusOK, utils.FormateError(code.Fail, "参数异常"))
		return
	}
	err = model.DB.Model(new(model.SysRole)).Where("id = ?", in.ID).Delete(&model.SysRole{}).Error
	if err != nil {
		ctx.JSON(http.StatusOK, utils.FormateError(code.ServerError, err.Error()))
		return
	}
	ctx.JSON(http.StatusOK, utils.FormateSuccess(code.Success, "删除角色成功", gin.H{}))
}
