// Package stock
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2026 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
// @AutoGenerate Version 2.18.6
package stock

import (
	"context"
	"hotgo/api/stock/slowstochasticdata"
	"hotgo/internal/model/input/stockin"
	"hotgo/internal/service"
)

var (
	SlowStochasticData = cSlowStochasticData{}
)

type cSlowStochasticData struct{}

// List 查看stoch指标数据表列表
func (c *cSlowStochasticData) List(ctx context.Context, req *slowstochasticdata.ListReq) (res *slowstochasticdata.ListRes, err error) {
	list, totalCount, err := service.StockSlowStochasticData().List(ctx, &req.SlowStochasticDataListInp)
	if err != nil {
		return
	}

	if list == nil {
		list = []*stockin.SlowStochasticDataListModel{}
	}

	res = new(slowstochasticdata.ListRes)
	res.List = list
	res.PageRes.Pack(req, totalCount)
	return
}

// Export 导出stoch指标数据表列表
func (c *cSlowStochasticData) Export(ctx context.Context, req *slowstochasticdata.ExportReq) (res *slowstochasticdata.ExportRes, err error) {
	err = service.StockSlowStochasticData().Export(ctx, &req.SlowStochasticDataListInp)
	return
}

// Edit 更新stoch指标数据表
func (c *cSlowStochasticData) Edit(ctx context.Context, req *slowstochasticdata.EditReq) (res *slowstochasticdata.EditRes, err error) {
	err = service.StockSlowStochasticData().Edit(ctx, &req.SlowStochasticDataEditInp)
	return
}

// View 获取指定stoch指标数据表信息
func (c *cSlowStochasticData) View(ctx context.Context, req *slowstochasticdata.ViewReq) (res *slowstochasticdata.ViewRes, err error) {
	data, err := service.StockSlowStochasticData().View(ctx, &req.SlowStochasticDataViewInp)
	if err != nil {
		return
	}

	res = new(slowstochasticdata.ViewRes)
	res.SlowStochasticDataViewModel = data
	return
}

// Delete 删除stoch指标数据表
func (c *cSlowStochasticData) Delete(ctx context.Context, req *slowstochasticdata.DeleteReq) (res *slowstochasticdata.DeleteRes, err error) {
	err = service.StockSlowStochasticData().Delete(ctx, &req.SlowStochasticDataDeleteInp)
	return
}
