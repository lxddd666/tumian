// Package stock
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2026 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
// @AutoGenerate Version 2.18.6
package stock

import (
	"context"
	"hotgo/api/stock/stockscoremain"
	"hotgo/internal/model/input/stockin"
	"hotgo/internal/service"
)

var (
	StockScoreMain = cStockScoreMain{}
)

type cStockScoreMain struct{}

// List 查看股票评分主表列表
func (c *cStockScoreMain) List(ctx context.Context, req *stockscoremain.ListReq) (res *stockscoremain.ListRes, err error) {
	list, totalCount, err := service.StockScoreMain().List(ctx, &req.StockScoreMainListInp)
	if err != nil {
		return
	}

	if list == nil {
		list = []*stockin.StockScoreMainListModel{}
	}

	res = new(stockscoremain.ListRes)
	res.List = list
	res.PageRes.Pack(req, totalCount)
	return
}

// Export 导出股票评分主表列表
func (c *cStockScoreMain) Export(ctx context.Context, req *stockscoremain.ExportReq) (res *stockscoremain.ExportRes, err error) {
	err = service.StockScoreMain().Export(ctx, &req.StockScoreMainListInp)
	return
}

// Edit 更新股票评分主表
func (c *cStockScoreMain) Edit(ctx context.Context, req *stockscoremain.EditReq) (res *stockscoremain.EditRes, err error) {
	err = service.StockScoreMain().Edit(ctx, &req.StockScoreMainEditInp)
	return
}

// View 获取指定股票评分主表信息
func (c *cStockScoreMain) View(ctx context.Context, req *stockscoremain.ViewReq) (res *stockscoremain.ViewRes, err error) {
	data, err := service.StockScoreMain().View(ctx, &req.StockScoreMainViewInp)
	if err != nil {
		return
	}

	res = new(stockscoremain.ViewRes)
	res.StockScoreMainViewModel = data
	return
}

// Delete 删除股票评分主表
func (c *cStockScoreMain) Delete(ctx context.Context, req *stockscoremain.DeleteReq) (res *stockscoremain.DeleteRes, err error) {
	err = service.StockScoreMain().Delete(ctx, &req.StockScoreMainDeleteInp)
	return
}
