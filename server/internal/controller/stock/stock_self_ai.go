// Package stock
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2026 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
// @AutoGenerate Version 2.18.6
package stock

import (
	"context"
	"hotgo/api/stock/stockselfai"
	"hotgo/internal/model/input/stockin"
	"hotgo/internal/service"
)

var (
	StockSelfAi = cStockSelfAi{}
)

type cStockSelfAi struct{}

// List 查看ai 基本信息列表
func (c *cStockSelfAi) List(ctx context.Context, req *stockselfai.ListReq) (res *stockselfai.ListRes, err error) {
	list, totalCount, err := service.StockSelfAi().List(ctx, &req.StockSelfAiListInp)
	if err != nil {
		return
	}

	if list == nil {
		list = []*stockin.StockSelfAiListModel{}
	}

	res = new(stockselfai.ListRes)
	res.List = list
	res.PageRes.Pack(req, totalCount)
	return
}

// Export 导出ai 基本信息列表
func (c *cStockSelfAi) Export(ctx context.Context, req *stockselfai.ExportReq) (res *stockselfai.ExportRes, err error) {
	err = service.StockSelfAi().Export(ctx, &req.StockSelfAiListInp)
	return
}

// Edit 更新ai 基本信息
func (c *cStockSelfAi) Edit(ctx context.Context, req *stockselfai.EditReq) (res *stockselfai.EditRes, err error) {
	err = service.StockSelfAi().Edit(ctx, &req.StockSelfAiEditInp)
	return
}

// View 获取指定ai 基本信息信息
func (c *cStockSelfAi) View(ctx context.Context, req *stockselfai.ViewReq) (res *stockselfai.ViewRes, err error) {
	data, err := service.StockSelfAi().View(ctx, &req.StockSelfAiViewInp)
	if err != nil {
		return
	}

	res = new(stockselfai.ViewRes)
	res.StockSelfAiViewModel = data
	return
}

// Delete 删除ai 基本信息
func (c *cStockSelfAi) Delete(ctx context.Context, req *stockselfai.DeleteReq) (res *stockselfai.DeleteRes, err error) {
	err = service.StockSelfAi().Delete(ctx, &req.StockSelfAiDeleteInp)
	return
}
