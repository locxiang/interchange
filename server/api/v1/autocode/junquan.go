package autocode

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/autocode"
	autocodeReq "github.com/flipped-aurora/gin-vue-admin/server/model/autocode/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/service"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
)

type JunquanApi struct {
}

var JqService = service.ServiceGroupApp.AutoCodeServiceGroup.JunquanService

// CreateJunquan 创建Junquan
// @Tags Junquan
// @Summary 创建Junquan
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body autocode.Junquan true "创建Junquan"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /Jq/createJunquan [post]
func (JqApi *JunquanApi) CreateJunquan(c *gin.Context) {
	var Jq autocode.Junquan
	_ = c.ShouldBindJSON(&Jq)
	if err := JqService.CreateJunquan(Jq); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteJunquan 删除Junquan
// @Tags Junquan
// @Summary 删除Junquan
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body autocode.Junquan true "删除Junquan"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /Jq/deleteJunquan [delete]
func (JqApi *JunquanApi) DeleteJunquan(c *gin.Context) {
	var Jq autocode.Junquan
	_ = c.ShouldBindJSON(&Jq)
	if err := JqService.DeleteJunquan(Jq); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteJunquanByIds 批量删除Junquan
// @Tags Junquan
// @Summary 批量删除Junquan
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除Junquan"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /Jq/deleteJunquanByIds [delete]
func (JqApi *JunquanApi) DeleteJunquanByIds(c *gin.Context) {
	var IDS request.IdsReq
	_ = c.ShouldBindJSON(&IDS)
	if err := JqService.DeleteJunquanByIds(IDS); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateJunquan 更新Junquan
// @Tags Junquan
// @Summary 更新Junquan
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body autocode.Junquan true "更新Junquan"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /Jq/updateJunquan [put]
func (JqApi *JunquanApi) UpdateJunquan(c *gin.Context) {
	var Jq autocode.Junquan
	_ = c.ShouldBindJSON(&Jq)
	if err := JqService.UpdateJunquan(Jq); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindJunquan 用id查询Junquan
// @Tags Junquan
// @Summary 用id查询Junquan
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query autocode.Junquan true "用id查询Junquan"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /Jq/findJunquan [get]
func (JqApi *JunquanApi) FindJunquan(c *gin.Context) {
	var Jq autocode.Junquan
	_ = c.ShouldBindQuery(&Jq)
	if err, reJq := JqService.GetJunquan(Jq.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"reJq": reJq}, c)
	}
}

// GetJunquanList 分页获取Junquan列表
// @Tags Junquan
// @Summary 分页获取Junquan列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query autocodeReq.JunquanSearch true "分页获取Junquan列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /Jq/getJunquanList [get]
func (JqApi *JunquanApi) GetJunquanList(c *gin.Context) {
	var pageInfo autocodeReq.JunquanSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if err, list, total := JqService.GetJunquanInfoList(pageInfo); err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithDetailed(response.PageResult{
			List:     list,
			Total:    total,
			Page:     pageInfo.Page,
			PageSize: pageInfo.PageSize,
		}, "获取成功", c)
	}
}

// GetJunquanJump 获取军犬跳转地址
// @Tags Junquan
// @Summary 获取军犬跳转地址
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /Jq/getJunquanJump [get]
func (JqApi *JunquanApi) GetJunquanJump(c *gin.Context) {

	userId := utils.GetUserID(c)
	err, data := JqService.Jump(userId)
	if err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
	}

	c.Redirect(http.StatusFound, data)

}
