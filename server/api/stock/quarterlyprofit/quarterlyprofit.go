// Package quarterlyprofit
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2026 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
// @AutoGenerate Version 2.18.6
package quarterlyprofit

import (
	"hotgo/internal/model/entity"
	"hotgo/internal/model/input/form"
	"hotgo/internal/model/input/stockin"

	"github.com/gogf/gf/v2/frame/g"
)

// ListReq 查询季度利润数据表 (近一年各季度)列表
type ListReq struct {
	g.Meta `path:"/quarterlyProfit/list" method:"get" tags:"季度利润数据表 (近一年各季度)" summary:"获取季度利润数据表 (近一年各季度)列表"`
	stockin.QuarterlyProfitListInp
}

type ListRes struct {
	form.PageRes
	List []*stockin.QuarterlyProfitListModel `json:"list"   dc:"数据列表"`
}

// ExportReq 导出季度利润数据表 (近一年各季度)列表
type ExportReq struct {
	g.Meta `path:"/quarterlyProfit/export" method:"get" tags:"季度利润数据表 (近一年各季度)" summary:"导出季度利润数据表 (近一年各季度)列表"`
	stockin.QuarterlyProfitListInp
}

type ExportRes struct{}

// ViewReq 获取季度利润数据表 (近一年各季度)指定信息
type ViewReq struct {
	g.Meta `path:"/quarterlyProfit/view" method:"get" tags:"季度利润数据表 (近一年各季度)" summary:"获取季度利润数据表 (近一年各季度)指定信息"`
	stockin.QuarterlyProfitViewInp
}

type ViewRes struct {
	*stockin.QuarterlyProfitViewModel
}

// EditReq 修改/新增季度利润数据表 (近一年各季度)
type EditReq struct {
	g.Meta `path:"/quarterlyProfit/edit" method:"post" tags:"季度利润数据表 (近一年各季度)" summary:"修改/新增季度利润数据表 (近一年各季度)"`
	stockin.QuarterlyProfitEditInp
}

type EditRes struct{}

// DeleteReq 删除季度利润数据表 (近一年各季度)
type DeleteReq struct {
	g.Meta `path:"/quarterlyProfit/delete" method:"post" tags:"季度利润数据表 (近一年各季度)" summary:"删除季度利润数据表 (近一年各季度)"`
	stockin.QuarterlyProfitDeleteInp
}

type DeleteRes struct{}

// GetQuarterlyProfitReq 获取季度利润数据表数据
type GetQuarterlyProfitReq struct {
	g.Meta `path:"/quarterlyProfit/getQuarterlyProfit" method:"get" tags:"季度利润数据表 (近一年各季度)" summary:"获取季度利润数据表数据"`
	stockin.QuarterlyProfitGetQuarterlyProfitInp
}

type GetQuarterlyProfitRes struct {
	Data []*entity.QuarterlyProfit `json:"data" dc:"返回数据"`
}
