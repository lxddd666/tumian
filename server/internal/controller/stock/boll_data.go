// Package stock
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2026 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
// @AutoGenerate Version 2.18.6
package stock

import (
	"context"
	"hotgo/api/stock/bolldata"
	"hotgo/internal/model/input/stockin"
	"hotgo/internal/service"
)

var (
	BollData = cBollData{}
)

type cBollData struct{}

// List 查看布林带(BOLL)指标数据表列表
func (c *cBollData) List(ctx context.Context, req *bolldata.ListReq) (res *bolldata.ListRes, err error) {
	list, totalCount, err := service.StockBollData().List(ctx, &req.BollDataListInp)
	if err != nil {
		return
	}

	if list == nil {
		list = []*stockin.BollDataListModel{}
	}

	res = new(bolldata.ListRes)
	res.List = list
	res.PageRes.Pack(req, totalCount)
	return
}

// Export 导出布林带(BOLL)指标数据表列表
func (c *cBollData) Export(ctx context.Context, req *bolldata.ExportReq) (res *bolldata.ExportRes, err error) {
	err = service.StockBollData().Export(ctx, &req.BollDataListInp)
	return
}

// Edit 更新布林带(BOLL)指标数据表
func (c *cBollData) Edit(ctx context.Context, req *bolldata.EditReq) (res *bolldata.EditRes, err error) {
	err = service.StockBollData().Edit(ctx, &req.BollDataEditInp)
	return
}

// View 获取指定布林带(BOLL)指标数据表信息
func (c *cBollData) View(ctx context.Context, req *bolldata.ViewReq) (res *bolldata.ViewRes, err error) {
	data, err := service.StockBollData().View(ctx, &req.BollDataViewInp)
	if err != nil {
		return
	}

	res = new(bolldata.ViewRes)
	res.BollDataViewModel = data
	return
}

// Delete 删除布林带(BOLL)指标数据表
func (c *cBollData) Delete(ctx context.Context, req *bolldata.DeleteReq) (res *bolldata.DeleteRes, err error) {
	err = service.StockBollData().Delete(ctx, &req.BollDataDeleteInp)
	return
}
