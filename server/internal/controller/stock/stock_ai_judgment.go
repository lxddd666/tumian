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

// AiJudgment 删除ai 选股判断
func (c *cStockAiJudgment) AiJudgment(ctx context.Context, req *stockaijudgment.AiJudgmentReq) (res *stockaijudgment.AiJudgmentRes, err error) {
	service.StockAiJudgment().InvokeJudgment(ctx, &req.StockAiJudgmentAiJudgmentInp)
	return
}
