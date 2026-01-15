// Package stock
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2026 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
// @AutoGenerate Version 2.18.6
package stock

import (
	"context"
	"hotgo/api/stock/quarterlyprofit"
	"hotgo/internal/model/input/stockin"
	"hotgo/internal/service"
)

var (
	QuarterlyProfit = cQuarterlyProfit{}
)

type cQuarterlyProfit struct{}

// List 查看季度利润数据表 (近一年各季度)列表
func (c *cQuarterlyProfit) List(ctx context.Context, req *quarterlyprofit.ListReq) (res *quarterlyprofit.ListRes, err error) {
	list, totalCount, err := service.StockQuarterlyProfit().List(ctx, &req.QuarterlyProfitListInp)
	if err != nil {
		return
	}

	if list == nil {
		list = []*stockin.QuarterlyProfitListModel{}
	}

	res = new(quarterlyprofit.ListRes)
	res.List = list
	res.PageRes.Pack(req, totalCount)
	return
}

// Export 导出季度利润数据表 (近一年各季度)列表
func (c *cQuarterlyProfit) Export(ctx context.Context, req *quarterlyprofit.ExportReq) (res *quarterlyprofit.ExportRes, err error) {
	err = service.StockQuarterlyProfit().Export(ctx, &req.QuarterlyProfitListInp)
	return
}

// Edit 更新季度利润数据表 (近一年各季度)
func (c *cQuarterlyProfit) Edit(ctx context.Context, req *quarterlyprofit.EditReq) (res *quarterlyprofit.EditRes, err error) {
	err = service.StockQuarterlyProfit().Edit(ctx, &req.QuarterlyProfitEditInp)
	return
}

// View 获取指定季度利润数据表 (近一年各季度)信息
func (c *cQuarterlyProfit) View(ctx context.Context, req *quarterlyprofit.ViewReq) (res *quarterlyprofit.ViewRes, err error) {
	data, err := service.StockQuarterlyProfit().View(ctx, &req.QuarterlyProfitViewInp)
	if err != nil {
		return
	}

	res = new(quarterlyprofit.ViewRes)
	res.QuarterlyProfitViewModel = data
	return
}

// Delete 删除季度利润数据表 (近一年各季度)
func (c *cQuarterlyProfit) Delete(ctx context.Context, req *quarterlyprofit.DeleteReq) (res *quarterlyprofit.DeleteRes, err error) {
	err = service.StockQuarterlyProfit().Delete(ctx, &req.QuarterlyProfitDeleteInp)
	return
}
