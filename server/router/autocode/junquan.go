package autocode

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type JunquanRouter struct {
}

// InitJunquanRouter 初始化 Junquan 路由信息
func (s *JunquanRouter) InitJunquanRouter(Router *gin.RouterGroup) {
	JqRouter := Router.Group("Jq").Use(middleware.OperationRecord())
	JqRouterWithoutRecord := Router.Group("Jq")
	var JqApi = v1.ApiGroupApp.AutoCodeApiGroup.JunquanApi
	{
		JqRouter.POST("createJunquan", JqApi.CreateJunquan)             // 新建Junquan
		JqRouter.DELETE("deleteJunquan", JqApi.DeleteJunquan)           // 删除Junquan
		JqRouter.DELETE("deleteJunquanByIds", JqApi.DeleteJunquanByIds) // 批量删除Junquan
		JqRouter.PUT("updateJunquan", JqApi.UpdateJunquan)              // 更新Junquan
	}
	{
		JqRouterWithoutRecord.GET("findJunquan", JqApi.FindJunquan)       // 根据ID获取Junquan
		JqRouterWithoutRecord.GET("getJunquanList", JqApi.GetJunquanList) // 获取Junquan列表
		JqRouterWithoutRecord.GET("getJunquanJump", JqApi.GetJunquanJump)              //获取跳转地址

	}
}
