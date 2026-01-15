// Package stock
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2026 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
// @AutoGenerate Version 2.18.6
package stock

import (
	"context"
	"hotgo/api/stock/stocklist"
	"hotgo/internal/model/input/stockin"
	"hotgo/internal/service"
)

var (
	StockList = cStockList{}
)

type cStockList struct{}

// List 查看股票列表核心表列表
func (c *cStockList) List(ctx context.Context, req *stocklist.ListReq) (res *stocklist.ListRes, err error) {
	list, totalCount, err := service.StockList().List(ctx, &req.StockListListInp)
	if err != nil {
		return
	}

	if list == nil {
		list = []*stockin.StockListListModel{}
	}

	res = new(stocklist.ListRes)
	res.List = list
	res.PageRes.Pack(req, totalCount)
	return
}

// Export 导出股票列表核心表列表
func (c *cStockList) Export(ctx context.Context, req *stocklist.ExportReq) (res *stocklist.ExportRes, err error) {
	err = service.StockList().Export(ctx, &req.StockListListInp)
	return
}

// Edit 更新股票列表核心表
func (c *cStockList) Edit(ctx context.Context, req *stocklist.EditReq) (res *stocklist.EditRes, err error) {
	err = service.StockList().Edit(ctx, &req.StockListEditInp)
	return
}

// View 获取指定股票列表核心表信息
func (c *cStockList) View(ctx context.Context, req *stocklist.ViewReq) (res *stocklist.ViewRes, err error) {
	data, err := service.StockList().View(ctx, &req.StockListViewInp)
	if err != nil {
		return
	}

	res = new(stocklist.ViewRes)
	res.StockListViewModel = data
	return
}

// Delete 删除股票列表核心表
func (c *cStockList) Delete(ctx context.Context, req *stocklist.DeleteReq) (res *stocklist.DeleteRes, err error) {
	err = service.StockList().Delete(ctx, &req.StockListDeleteInp)
	return
}

// Status 更新股票列表核心表状态
func (c *cStockList) Status(ctx context.Context, req *stocklist.StatusReq) (res *stocklist.StatusRes, err error) {
	err = service.StockList().Status(ctx, &req.StockListStatusInp)
	return
}

// GetStockList 获取股票列表核心表数据
func (c *cStockList) GetStockList(ctx context.Context, req *stocklist.GetStockListReq) (res *stocklist.GetStockListRes, err error) {
	data, err := service.StockList().GetStockList(ctx, &req.StockListGetStockListInp)
	if err != nil {
		return
	}

	res = new(stocklist.GetStockListRes)
	res.Data = data
	return
}
