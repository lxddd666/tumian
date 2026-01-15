// Package stock
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2026 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
// @AutoGenerate Version 2.18.6
package stock

import (
	"context"
	"hotgo/api/stock/fundstockholding"
	"hotgo/internal/model/input/stockin"
	"hotgo/internal/service"
)

var (
	FundStockHolding = cFundStockHolding{}
)

type cFundStockHolding struct{}

// List 查看基金持股明细表 (来源于基金定期报告)列表
func (c *cFundStockHolding) List(ctx context.Context, req *fundstockholding.ListReq) (res *fundstockholding.ListRes, err error) {
	list, totalCount, err := service.StockFundStockHolding().List(ctx, &req.FundStockHoldingListInp)
	if err != nil {
		return
	}

	if list == nil {
		list = []*stockin.FundStockHoldingListModel{}
	}

	res = new(fundstockholding.ListRes)
	res.List = list
	res.PageRes.Pack(req, totalCount)
	return
}

// Export 导出基金持股明细表 (来源于基金定期报告)列表
func (c *cFundStockHolding) Export(ctx context.Context, req *fundstockholding.ExportReq) (res *fundstockholding.ExportRes, err error) {
	err = service.StockFundStockHolding().Export(ctx, &req.FundStockHoldingListInp)
	return
}

// Edit 更新基金持股明细表 (来源于基金定期报告)
func (c *cFundStockHolding) Edit(ctx context.Context, req *fundstockholding.EditReq) (res *fundstockholding.EditRes, err error) {
	err = service.StockFundStockHolding().Edit(ctx, &req.FundStockHoldingEditInp)
	return
}

// View 获取指定基金持股明细表 (来源于基金定期报告)信息
func (c *cFundStockHolding) View(ctx context.Context, req *fundstockholding.ViewReq) (res *fundstockholding.ViewRes, err error) {
	data, err := service.StockFundStockHolding().View(ctx, &req.FundStockHoldingViewInp)
	if err != nil {
		return
	}

	res = new(fundstockholding.ViewRes)
	res.FundStockHoldingViewModel = data
	return
}

// Delete 删除基金持股明细表 (来源于基金定期报告)
func (c *cFundStockHolding) Delete(ctx context.Context, req *fundstockholding.DeleteReq) (res *fundstockholding.DeleteRes, err error) {
	err = service.StockFundStockHolding().Delete(ctx, &req.FundStockHoldingDeleteInp)
	return
}

// GetFundStockHolding 获取基金持股明细表数据
func (c *cFundStockHolding) GetFundStockHolding(ctx context.Context, req *fundstockholding.GetFundStockHoldingReq) (res *fundstockholding.GetFundStockHoldingRes, err error) {
	data, err := service.StockFundStockHolding().GetFundStockHolding(ctx, &req.FundStockHoldingGetFundStockHoldingInp)
	if err != nil {
		return
	}

	res = new(fundstockholding.GetFundStockHoldingRes)
	res.Data = data
	return
}
