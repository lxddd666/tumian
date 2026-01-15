// Package stock
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2026 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
// @AutoGenerate Version 2.18.6
package stock

import (
	"context"
	"hotgo/api/stock/stockbasicinfo"
	"hotgo/internal/model/input/stockin"
	"hotgo/internal/service"
)

var (
	StockBasicInfo = cStockBasicInfo{}
)

type cStockBasicInfo struct{}

// List 查看股票基础信息表列表
func (c *cStockBasicInfo) List(ctx context.Context, req *stockbasicinfo.ListReq) (res *stockbasicinfo.ListRes, err error) {
	list, totalCount, err := service.StockBasicInfo().List(ctx, &req.StockBasicInfoListInp)
	if err != nil {
		return
	}

	if list == nil {
		list = []*stockin.StockBasicInfoListModel{}
	}

	res = new(stockbasicinfo.ListRes)
	res.List = list
	res.PageRes.Pack(req, totalCount)
	return
}

// Export 导出股票基础信息表列表
func (c *cStockBasicInfo) Export(ctx context.Context, req *stockbasicinfo.ExportReq) (res *stockbasicinfo.ExportRes, err error) {
	err = service.StockBasicInfo().Export(ctx, &req.StockBasicInfoListInp)
	return
}

// Edit 更新股票基础信息表
func (c *cStockBasicInfo) Edit(ctx context.Context, req *stockbasicinfo.EditReq) (res *stockbasicinfo.EditRes, err error) {
	err = service.StockBasicInfo().Edit(ctx, &req.StockBasicInfoEditInp)
	return
}

// View 获取指定股票基础信息表信息
func (c *cStockBasicInfo) View(ctx context.Context, req *stockbasicinfo.ViewReq) (res *stockbasicinfo.ViewRes, err error) {
	data, err := service.StockBasicInfo().View(ctx, &req.StockBasicInfoViewInp)
	if err != nil {
		return
	}

	res = new(stockbasicinfo.ViewRes)
	res.StockBasicInfoViewModel = data
	return
}

// Delete 删除股票基础信息表
func (c *cStockBasicInfo) Delete(ctx context.Context, req *stockbasicinfo.DeleteReq) (res *stockbasicinfo.DeleteRes, err error) {
	err = service.StockBasicInfo().Delete(ctx, &req.StockBasicInfoDeleteInp)
	return
}
