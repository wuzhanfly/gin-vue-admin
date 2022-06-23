package IndexPage

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type IndexPageRouter struct {
}

// InitIndexPageRouter 初始化 IndexPage 路由信息
func (s *IndexPageRouter) InitIndexPageRouter(Router *gin.RouterGroup) {
	InPageRouter := Router.Group("InPage").Use(middleware.OperationRecord())
	InPageRouterWithoutRecord := Router.Group("InPage")
	var InPageApi = v1.ApiGroupApp.IndexpageApiGroup.IndexPageApi
	{
		InPageRouter.POST("createIndexPage", InPageApi.CreateIndexPage)             // 新建IndexPage
		InPageRouter.DELETE("deleteIndexPage", InPageApi.DeleteIndexPage)           // 删除IndexPage
		InPageRouter.DELETE("deleteIndexPageByIds", InPageApi.DeleteIndexPageByIds) // 批量删除IndexPage
		InPageRouter.PUT("updateIndexPage", InPageApi.UpdateIndexPage)              // 更新IndexPage
	}
	{
		InPageRouterWithoutRecord.GET("findIndexPage", InPageApi.FindIndexPage)       // 根据ID获取IndexPage
		InPageRouterWithoutRecord.GET("getIndexPageList", InPageApi.GetIndexPageList) // 获取IndexPage列表
	}
}
