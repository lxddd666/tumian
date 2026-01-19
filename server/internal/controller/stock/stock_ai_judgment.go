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
	"hotgo/internal/model/input/stockin"
	"hotgo/internal/service"
)

var (
	StockAiJudgment = cStockAiJudgment{}
)

type cStockAiJudgment struct{}

// List 查看ai 选股判断列表
func (c *cStockAiJudgment) List(ctx context.Context, req *stockaijudgment.ListReq) (res *stockaijudgment.ListRes, err error) {
	list, totalCount, err := service.StockAiJudgment().List(ctx, &req.StockAiJudgmentListInp)
	if err != nil {
		return
	}

	if list == nil {
		list = []*stockin.StockAiJudgmentListModel{}
	}

	res = new(stockaijudgment.ListRes)
	res.List = list
	res.PageRes.Pack(req, totalCount)
	return
}

// Export 导出ai 选股判断列表
func (c *cStockAiJudgment) Export(ctx context.Context, req *stockaijudgment.ExportReq) (res *stockaijudgment.ExportRes, err error) {
	err = service.StockAiJudgment().Export(ctx, &req.StockAiJudgmentListInp)
	return
}

// Edit 更新ai 选股判断
func (c *cStockAiJudgment) Edit(ctx context.Context, req *stockaijudgment.EditReq) (res *stockaijudgment.EditRes, err error) {
	err = service.StockAiJudgment().Edit(ctx, &req.StockAiJudgmentEditInp)
	return
}

// View 获取指定ai 选股判断信息
func (c *cStockAiJudgment) View(ctx context.Context, req *stockaijudgment.ViewReq) (res *stockaijudgment.ViewRes, err error) {
	data, err := service.StockAiJudgment().View(ctx, &req.StockAiJudgmentViewInp)
	if err != nil {
		return
	}

	res = new(stockaijudgment.ViewRes)
	res.StockAiJudgmentViewModel = data
	return
}

// Delete 删除ai 选股判断
func (c *cStockAiJudgment) Delete(ctx context.Context, req *stockaijudgment.DeleteReq) (res *stockaijudgment.DeleteRes, err error) {
	err = service.StockAiJudgment().Delete(ctx, &req.StockAiJudgmentDeleteInp)
	return
}
