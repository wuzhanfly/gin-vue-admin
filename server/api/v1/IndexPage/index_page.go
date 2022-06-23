package IndexPage

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/IndexPage"
	IndexPageReq "github.com/flipped-aurora/gin-vue-admin/server/model/IndexPage/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/service"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type IndexPageApi struct {
}

var InPageService = service.ServiceGroupApp.IndexpageServiceGroup.IndexPageService

// CreateIndexPage 创建IndexPage
// @Tags IndexPage
// @Summary 创建IndexPage
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body IndexPage.IndexPage true "创建IndexPage"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /InPage/createIndexPage [post]
func (InPageApi *IndexPageApi) CreateIndexPage(c *gin.Context) {
	var InPage IndexPage.IndexPage
	_ = c.ShouldBindJSON(&InPage)
	if err := InPageService.CreateIndexPage(InPage); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteIndexPage 删除IndexPage
// @Tags IndexPage
// @Summary 删除IndexPage
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body IndexPage.IndexPage true "删除IndexPage"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /InPage/deleteIndexPage [delete]
func (InPageApi *IndexPageApi) DeleteIndexPage(c *gin.Context) {
	var InPage IndexPage.IndexPage
	_ = c.ShouldBindJSON(&InPage)
	if err := InPageService.DeleteIndexPage(InPage); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteIndexPageByIds 批量删除IndexPage
// @Tags IndexPage
// @Summary 批量删除IndexPage
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除IndexPage"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /InPage/deleteIndexPageByIds [delete]
func (InPageApi *IndexPageApi) DeleteIndexPageByIds(c *gin.Context) {
	var IDS request.IdsReq
	_ = c.ShouldBindJSON(&IDS)
	if err := InPageService.DeleteIndexPageByIds(IDS); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateIndexPage 更新IndexPage
// @Tags IndexPage
// @Summary 更新IndexPage
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body IndexPage.IndexPage true "更新IndexPage"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /InPage/updateIndexPage [put]
func (InPageApi *IndexPageApi) UpdateIndexPage(c *gin.Context) {
	var InPage IndexPage.IndexPage
	_ = c.ShouldBindJSON(&InPage)
	if err := InPageService.UpdateIndexPage(InPage); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindIndexPage 用id查询IndexPage
// @Tags IndexPage
// @Summary 用id查询IndexPage
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query IndexPage.IndexPage true "用id查询IndexPage"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /InPage/findIndexPage [get]
func (InPageApi *IndexPageApi) FindIndexPage(c *gin.Context) {
	var InPage IndexPage.IndexPage
	_ = c.ShouldBindQuery(&InPage)
	if reInPage, err := InPageService.GetIndexPage(InPage.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"reInPage": reInPage}, c)
	}
}

// GetIndexPageList 分页获取IndexPage列表
// @Tags IndexPage
// @Summary 分页获取IndexPage列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query IndexPageReq.IndexPageSearch true "分页获取IndexPage列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /InPage/getIndexPageList [get]
func (InPageApi *IndexPageApi) GetIndexPageList(c *gin.Context) {
	var pageInfo IndexPageReq.IndexPageSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if list, total, err := InPageService.GetIndexPageInfoList(pageInfo); err != nil {
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
