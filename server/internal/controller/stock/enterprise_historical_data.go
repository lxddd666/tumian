// Package stock
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2026 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
// @AutoGenerate Version 2.18.6
package stock

import (
	"context"
	"hotgo/api/stock/enterprisehistoricaldata"
	"hotgo/internal/model/input/stockin"
	"hotgo/internal/service"
)

var (
	EnterpriseHistoricalData = cEnterpriseHistoricalData{}
)

type cEnterpriseHistoricalData struct{}

// List 查看企业级历史行情数据表 (K线数据)列表
func (c *cEnterpriseHistoricalData) List(ctx context.Context, req *enterprisehistoricaldata.ListReq) (res *enterprisehistoricaldata.ListRes, err error) {
	list, totalCount, err := service.StockEnterpriseHistoricalData().List(ctx, &req.EnterpriseHistoricalDataListInp)
	if err != nil {
		return
	}

	if list == nil {
		list = []*stockin.EnterpriseHistoricalDataListModel{}
	}

	res = new(enterprisehistoricaldata.ListRes)
	res.List = list
	res.PageRes.Pack(req, totalCount)
	return
}

// Export 导出企业级历史行情数据表 (K线数据)列表
func (c *cEnterpriseHistoricalData) Export(ctx context.Context, req *enterprisehistoricaldata.ExportReq) (res *enterprisehistoricaldata.ExportRes, err error) {
	err = service.StockEnterpriseHistoricalData().Export(ctx, &req.EnterpriseHistoricalDataListInp)
	return
}

// Edit 更新企业级历史行情数据表 (K线数据)
func (c *cEnterpriseHistoricalData) Edit(ctx context.Context, req *enterprisehistoricaldata.EditReq) (res *enterprisehistoricaldata.EditRes, err error) {
	err = service.StockEnterpriseHistoricalData().Edit(ctx, &req.EnterpriseHistoricalDataEditInp)
	return
}

// View 获取指定企业级历史行情数据表 (K线数据)信息
func (c *cEnterpriseHistoricalData) View(ctx context.Context, req *enterprisehistoricaldata.ViewReq) (res *enterprisehistoricaldata.ViewRes, err error) {
	data, err := service.StockEnterpriseHistoricalData().View(ctx, &req.EnterpriseHistoricalDataViewInp)
	if err != nil {
		return
	}

	res = new(enterprisehistoricaldata.ViewRes)
	res.EnterpriseHistoricalDataViewModel = data
	return
}

// Delete 删除企业级历史行情数据表 (K线数据)
func (c *cEnterpriseHistoricalData) Delete(ctx context.Context, req *enterprisehistoricaldata.DeleteReq) (res *enterprisehistoricaldata.DeleteRes, err error) {
	err = service.StockEnterpriseHistoricalData().Delete(ctx, &req.EnterpriseHistoricalDataDeleteInp)
	return
}
