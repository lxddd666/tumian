// Package stock
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2026 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
// @AutoGenerate Version 2.18.6
package stock

import (
	"context"
	"hotgo/api/stock/stocksupportresistance"
	"hotgo/internal/model/input/stockin"
	"hotgo/internal/service"
)

var (
	StockSupportResistance = cStockSupportResistance{}
)

type cStockSupportResistance struct{}

// List 查看股票支撑阻力表列表
func (c *cStockSupportResistance) List(ctx context.Context, req *stocksupportresistance.ListReq) (res *stocksupportresistance.ListRes, err error) {
	list, totalCount, err := service.StockSupportResistance().List(ctx, &req.StockSupportResistanceListInp)
	if err != nil {
		return
	}

	if list == nil {
		list = []*stockin.StockSupportResistanceListModel{}
	}

	res = new(stocksupportresistance.ListRes)
	res.List = list
	res.PageRes.Pack(req, totalCount)
	return
}

// Export 导出股票支撑阻力表列表
func (c *cStockSupportResistance) Export(ctx context.Context, req *stocksupportresistance.ExportReq) (res *stocksupportresistance.ExportRes, err error) {
	err = service.StockSupportResistance().Export(ctx, &req.StockSupportResistanceListInp)
	return
}

// Edit 更新股票支撑阻力表
func (c *cStockSupportResistance) Edit(ctx context.Context, req *stocksupportresistance.EditReq) (res *stocksupportresistance.EditRes, err error) {
	err = service.StockSupportResistance().Edit(ctx, &req.StockSupportResistanceEditInp)
	return
}

// View 获取指定股票支撑阻力表信息
func (c *cStockSupportResistance) View(ctx context.Context, req *stocksupportresistance.ViewReq) (res *stocksupportresistance.ViewRes, err error) {
	data, err := service.StockSupportResistance().View(ctx, &req.StockSupportResistanceViewInp)
	if err != nil {
		return
	}

	res = new(stocksupportresistance.ViewRes)
	res.StockSupportResistanceViewModel = data
	return
}

// Delete 删除股票支撑阻力表
func (c *cStockSupportResistance) Delete(ctx context.Context, req *stocksupportresistance.DeleteReq) (res *stocksupportresistance.DeleteRes, err error) {
	err = service.StockSupportResistance().Delete(ctx, &req.StockSupportResistanceDeleteInp)
	return
}
