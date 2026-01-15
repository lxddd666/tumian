// Package financialindicators
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2026 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
// @AutoGenerate Version 2.18.6
package financialindicators

import (
	"hotgo/internal/model/entity"
	"hotgo/internal/model/input/form"
	"hotgo/internal/model/input/stockin"

	"github.com/gogf/gf/v2/frame/g"
)

// ListReq 查询财务指标分析表列表
type ListReq struct {
	g.Meta `path:"/financialIndicators/list" method:"get" tags:"财务指标分析表" summary:"获取财务指标分析表列表"`
	stockin.FinancialIndicatorsListInp
}

type ListRes struct {
	form.PageRes
	List []*stockin.FinancialIndicatorsListModel `json:"list"   dc:"数据列表"`
}

// ExportReq 导出财务指标分析表列表
type ExportReq struct {
	g.Meta `path:"/financialIndicators/export" method:"get" tags:"财务指标分析表" summary:"导出财务指标分析表列表"`
	stockin.FinancialIndicatorsListInp
}

type ExportRes struct{}

// ViewReq 获取财务指标分析表指定信息
type ViewReq struct {
	g.Meta `path:"/financialIndicators/view" method:"get" tags:"财务指标分析表" summary:"获取财务指标分析表指定信息"`
	stockin.FinancialIndicatorsViewInp
}

type ViewRes struct {
	*stockin.FinancialIndicatorsViewModel
}

// EditReq 修改/新增财务指标分析表
type EditReq struct {
	g.Meta `path:"/financialIndicators/edit" method:"post" tags:"财务指标分析表" summary:"修改/新增财务指标分析表"`
	stockin.FinancialIndicatorsEditInp
}

type EditRes struct{}

// DeleteReq 删除财务指标分析表
type DeleteReq struct {
	g.Meta `path:"/financialIndicators/delete" method:"post" tags:"财务指标分析表" summary:"删除财务指标分析表"`
	stockin.FinancialIndicatorsDeleteInp
}

type DeleteRes struct{}

// GetFinancialIndicatorsReq 获取财务指标分析表数据
type GetFinancialIndicatorsReq struct {
	g.Meta `path:"/financialIndicators/getFinancialIndicators" method:"get" tags:"财务指标分析表" summary:"获取财务指标分析表数据"`
	stockin.FinancialIndicatorsGetFinancialIndicatorsInp
}

type GetFinancialIndicatorsRes struct {
	Data *entity.FinancialIndicators `json:"data" dc:"返回数据"`
}
