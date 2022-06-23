import service from '@/utils/request'

// @Tags IndexPage
// @Summary 创建IndexPage
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.IndexPage true "创建IndexPage"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /InPage/createIndexPage [post]
export const createIndexPage = (data) => {
  return service({
    url: '/InPage/createIndexPage',
    method: 'post',
    data
  })
}

// @Tags IndexPage
// @Summary 删除IndexPage
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.IndexPage true "删除IndexPage"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /InPage/deleteIndexPage [delete]
export const deleteIndexPage = (data) => {
  return service({
    url: '/InPage/deleteIndexPage',
    method: 'delete',
    data
  })
}

// @Tags IndexPage
// @Summary 删除IndexPage
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除IndexPage"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /InPage/deleteIndexPage [delete]
export const deleteIndexPageByIds = (data) => {
  return service({
    url: '/InPage/deleteIndexPageByIds',
    method: 'delete',
    data
  })
}

// @Tags IndexPage
// @Summary 更新IndexPage
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.IndexPage true "更新IndexPage"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /InPage/updateIndexPage [put]
export const updateIndexPage = (data) => {
  return service({
    url: '/InPage/updateIndexPage',
    method: 'put',
    data
  })
}

// @Tags IndexPage
// @Summary 用id查询IndexPage
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.IndexPage true "用id查询IndexPage"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /InPage/findIndexPage [get]
export const findIndexPage = (params) => {
  return service({
    url: '/InPage/findIndexPage',
    method: 'get',
    params
  })
}

// @Tags IndexPage
// @Summary 分页获取IndexPage列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取IndexPage列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /InPage/getIndexPageList [get]
export const getIndexPageList = (params) => {
  return service({
    url: '/InPage/getIndexPageList',
    method: 'get',
    params
  })
}
