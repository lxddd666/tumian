// Package enterprisehistoricaldata
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2026 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
// @AutoGenerate Version 2.18.6
package enterprisehistoricaldata

import (
	"hotgo/internal/model/input/form"
	"hotgo/internal/model/input/stockin"

	"github.com/gogf/gf/v2/frame/g"
)

// ListReq 查询企业级历史行情数据表 (K线数据)列表
type ListReq struct {
	g.Meta `path:"/enterpriseHistoricalData/list" method:"get" tags:"企业级历史行情数据表 (K线数据)" summary:"获取企业级历史行情数据表 (K线数据)列表"`
	stockin.EnterpriseHistoricalDataListInp
}

type ListRes struct {
	form.PageRes
	List []*stockin.EnterpriseHistoricalDataListModel `json:"list"   dc:"数据列表"`
}

// ExportReq 导出企业级历史行情数据表 (K线数据)列表
type ExportReq struct {
	g.Meta `path:"/enterpriseHistoricalData/export" method:"get" tags:"企业级历史行情数据表 (K线数据)" summary:"导出企业级历史行情数据表 (K线数据)列表"`
	stockin.EnterpriseHistoricalDataListInp
}

type ExportRes struct{}

// ViewReq 获取企业级历史行情数据表 (K线数据)指定信息
type ViewReq struct {
	g.Meta `path:"/enterpriseHistoricalData/view" method:"get" tags:"企业级历史行情数据表 (K线数据)" summary:"获取企业级历史行情数据表 (K线数据)指定信息"`
	stockin.EnterpriseHistoricalDataViewInp
}

type ViewRes struct {
	*stockin.EnterpriseHistoricalDataViewModel
}

// EditReq 修改/新增企业级历史行情数据表 (K线数据)
type EditReq struct {
	g.Meta `path:"/enterpriseHistoricalData/edit" method:"post" tags:"企业级历史行情数据表 (K线数据)" summary:"修改/新增企业级历史行情数据表 (K线数据)"`
	stockin.EnterpriseHistoricalDataEditInp
}

type EditRes struct{}

// DeleteReq 删除企业级历史行情数据表 (K线数据)
type DeleteReq struct {
	g.Meta `path:"/enterpriseHistoricalData/delete" method:"post" tags:"企业级历史行情数据表 (K线数据)" summary:"删除企业级历史行情数据表 (K线数据)"`
	stockin.EnterpriseHistoricalDataDeleteInp
}

type DeleteRes struct{}
