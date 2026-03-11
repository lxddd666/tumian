// Package stock
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2026 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
// @AutoGenerate Version 2.18.6
package stock

import (
	"context"
	"hotgo/api/stock/stockaijudgment"
	"hotgo/internal/service"
)

var (
	StockAiJudgment = cStockAiJudgment{}
)

type cStockAiJudgment struct{}

// AiJudgment ai 选股判断
func (c *cStockAiJudgment) AiJudgment(ctx context.Context, req *stockaijudgment.AiJudgmentReq) (res *stockaijudgment.AiJudgmentRes, err error) {
	service.StockAiJudgment().InvokeIndicatorsJudgment(ctx, &req.StockAiJudgmentAiJudgmentInp)
	return
}

// AiJudgmentFinancialData ai财务数据鉴定股票鉴股
func (c *cStockAiJudgment) AiJudgmentFinancialData(ctx context.Context, req *stockaijudgment.AiJudgmentFinancialDataReq) (res *stockaijudgment.AiJudgmentFinancialDataRes, err error) {
	service.StockAiJudgment().AiJudgmentFinancialData(ctx, &req.StockAiJudgmentAiJudgmentInp)
	return
}

// AiJudgmentComprehensiveData ai综合指标数据鉴定股票鉴股
func (c *cStockAiJudgment) AiJudgmentComprehensiveData(ctx context.Context, req *stockaijudgment.AiJudgmentComprehensiveDataReq) (res *stockaijudgment.AiJudgmentComprehensiveDataRes, err error) {
	service.StockAiJudgment().AiJudgmentComprehensiveData(ctx, &req.StockAiJudgmentAiJudgmentInp)
	return
}

// AiJudgmentComprehensiveData ai综合指标数据鉴定股票鉴股
func (c *cStockAiJudgment) Export(ctx context.Context, req *stockaijudgment.ExportReq) (res *stockaijudgment.ExportRes, err error) {
	err = service.StockAiJudgment().Export(ctx)
	return
}
