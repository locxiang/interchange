import service from '@/utils/request'

// @Tags Junquan
// @Summary 创建Junquan
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Junquan true "创建Junquan"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /Jq/createJunquan [post]
export const createJunquan = (data) => {
  return service({
    url: '/Jq/createJunquan',
    method: 'post',
    data
  })
}

// @Tags Junquan
// @Summary 删除Junquan
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Junquan true "删除Junquan"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /Jq/deleteJunquan [delete]
export const deleteJunquan = (data) => {
  return service({
    url: '/Jq/deleteJunquan',
    method: 'delete',
    data
  })
}

// @Tags Junquan
// @Summary 删除Junquan
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除Junquan"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /Jq/deleteJunquan [delete]
export const deleteJunquanByIds = (data) => {
  return service({
    url: '/Jq/deleteJunquanByIds',
    method: 'delete',
    data
  })
}

// @Tags Junquan
// @Summary 更新Junquan
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Junquan true "更新Junquan"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /Jq/updateJunquan [put]
export const updateJunquan = (data) => {
  return service({
    url: '/Jq/updateJunquan',
    method: 'put',
    data
  })
}

// @Tags Junquan
// @Summary 用id查询Junquan
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.Junquan true "用id查询Junquan"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /Jq/findJunquan [get]
export const findJunquan = (params) => {
  return service({
    url: '/Jq/findJunquan',
    method: 'get',
    params
  })
}

// @Tags Junquan
// @Summary 分页获取Junquan列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取Junquan列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /Jq/getJunquanList [get]
export const getJunquanList = (params) => {
  return service({
    url: '/Jq/getJunquanList',
    method: 'get',
    params
  })
}
