// Package stock
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2026 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
// @AutoGenerate Version 2.18.6
package stock

import (
	"context"
	"hotgo/api/stock/financialindicators"
	"hotgo/internal/model/input/stockin"
	"hotgo/internal/service"
)

var (
	FinancialIndicators = cFinancialIndicators{}
)

type cFinancialIndicators struct{}

// List 查看财务指标分析表列表
func (c *cFinancialIndicators) List(ctx context.Context, req *financialindicators.ListReq) (res *financialindicators.ListRes, err error) {
	list, totalCount, err := service.StockFinancialIndicators().List(ctx, &req.FinancialIndicatorsListInp)
	if err != nil {
		return
	}

	if list == nil {
		list = []*stockin.FinancialIndicatorsListModel{}
	}

	res = new(financialindicators.ListRes)
	res.List = list
	res.PageRes.Pack(req, totalCount)
	return
}

// Export 导出财务指标分析表列表
func (c *cFinancialIndicators) Export(ctx context.Context, req *financialindicators.ExportReq) (res *financialindicators.ExportRes, err error) {
	err = service.StockFinancialIndicators().Export(ctx, &req.FinancialIndicatorsListInp)
	return
}

// Edit 更新财务指标分析表
func (c *cFinancialIndicators) Edit(ctx context.Context, req *financialindicators.EditReq) (res *financialindicators.EditRes, err error) {
	err = service.StockFinancialIndicators().Edit(ctx, &req.FinancialIndicatorsEditInp)
	return
}

// View 获取指定财务指标分析表信息
func (c *cFinancialIndicators) View(ctx context.Context, req *financialindicators.ViewReq) (res *financialindicators.ViewRes, err error) {
	data, err := service.StockFinancialIndicators().View(ctx, &req.FinancialIndicatorsViewInp)
	if err != nil {
		return
	}

	res = new(financialindicators.ViewRes)
	res.FinancialIndicatorsViewModel = data
	return
}

// Delete 删除财务指标分析表
func (c *cFinancialIndicators) Delete(ctx context.Context, req *financialindicators.DeleteReq) (res *financialindicators.DeleteRes, err error) {
	err = service.StockFinancialIndicators().Delete(ctx, &req.FinancialIndicatorsDeleteInp)
	return
}

// GetFinancialIndicators 获取财务指标分析表数据
func (c *cFinancialIndicators) GetFinancialIndicators(ctx context.Context, req *financialindicators.GetFinancialIndicatorsReq) (res *financialindicators.GetFinancialIndicatorsRes, err error) {
	data, err := service.StockFinancialIndicators().GetFinancialIndicators(ctx, &req.FinancialIndicatorsGetFinancialIndicatorsInp)
	if err != nil {
		return
	}

	res = new(financialindicators.GetFinancialIndicatorsRes)
	res.Data = data
	return
}
