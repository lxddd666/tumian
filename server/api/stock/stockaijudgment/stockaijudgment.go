// Package stockaijudgment
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2026 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
// @AutoGenerate Version 2.18.6
package stockaijudgment

import (
	"hotgo/internal/model/input/form"
	"hotgo/internal/model/input/stockin"

	"github.com/gogf/gf/v2/frame/g"
)

// ListReq 查询ai 选股判断列表
type ListReq struct {
	g.Meta `path:"/stockAiJudgment/list" method:"get" tags:"ai 选股判断" summary:"获取ai 选股判断列表"`
	stockin.StockAiJudgmentListInp
}

type ListRes struct {
	form.PageRes
	List []*stockin.StockAiJudgmentListModel `json:"list"   dc:"数据列表"`
}

// ExportReq 导出ai 选股判断列表
type ExportReq struct {
	g.Meta `path:"/stockAiJudgment/export" method:"get" tags:"ai 选股判断" summary:"导出ai 选股判断列表"`
	stockin.StockAiJudgmentListInp
}

type ExportRes struct{}

// ViewReq 获取ai 选股判断指定信息
type ViewReq struct {
	g.Meta `path:"/stockAiJudgment/view" method:"get" tags:"ai 选股判断" summary:"获取ai 选股判断指定信息"`
	stockin.StockAiJudgmentViewInp
}

type ViewRes struct {
	*stockin.StockAiJudgmentViewModel
}

// EditReq 修改/新增ai 选股判断
type EditReq struct {
	g.Meta `path:"/stockAiJudgment/edit" method:"post" tags:"ai 选股判断" summary:"修改/新增ai 选股判断"`
	stockin.StockAiJudgmentEditInp
}

type EditRes struct{}

// DeleteReq 删除ai 选股判断
type DeleteReq struct {
	g.Meta `path:"/stockAiJudgment/delete" method:"post" tags:"ai 选股判断" summary:"删除ai 选股判断"`
	stockin.StockAiJudgmentDeleteInp
}

type DeleteRes struct{}

type AiJudgmentReq struct {
	g.Meta `path:"/stockAiJudgment/aiJudgment" method:"post" tags:"ai 选股判断" summary:"ai鉴股"`
	stockin.StockAiJudgmentAiJudgmentInp
}

type AiJudgmentRes struct{}

type AiJudgmentFinancialDataReq struct {
	g.Meta `path:"/stockAiJudgment/aiJudgmentFinancialDataReq" method:"post" tags:"ai 选股判断" summary:"ai财务数据鉴定股票鉴股"`
	stockin.StockAiJudgmentAiJudgmentInp
}

type AiJudgmentFinancialDataRes struct{}

type AiJudgmentComprehensiveDataReq struct {
	g.Meta `path:"/stockAiJudgment/AiJudgmentComprehensiveDataReq" method:"post" tags:"ai 选股判断" summary:"ai综合指标数据鉴定股票鉴股"`
	stockin.StockAiJudgmentAiJudgmentInp
}

type AiJudgmentComprehensiveDataRes struct{}
