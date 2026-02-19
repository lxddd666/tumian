// Package stock
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2026 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
// @AutoGenerate Version 2.18.6
package stock

import (
	"context"
	"hotgo/api/stock/ccidata"
	"hotgo/internal/model/input/stockin"
	"hotgo/internal/service"
)

var (
	CciData = cCciData{}
)

type cCciData struct{}

// List 查看cci指标数据表列表
func (c *cCciData) List(ctx context.Context, req *ccidata.ListReq) (res *ccidata.ListRes, err error) {
	list, totalCount, err := service.StockCciData().List(ctx, &req.CciDataListInp)
	if err != nil {
		return
	}

	if list == nil {
		list = []*stockin.CciDataListModel{}
	}

	res = new(ccidata.ListRes)
	res.List = list
	res.PageRes.Pack(req, totalCount)
	return
}

// Export 导出cci指标数据表列表
func (c *cCciData) Export(ctx context.Context, req *ccidata.ExportReq) (res *ccidata.ExportRes, err error) {
	err = service.StockCciData().Export(ctx, &req.CciDataListInp)
	return
}

// Edit 更新cci指标数据表
func (c *cCciData) Edit(ctx context.Context, req *ccidata.EditReq) (res *ccidata.EditRes, err error) {
	err = service.StockCciData().Edit(ctx, &req.CciDataEditInp)
	return
}

// View 获取指定cci指标数据表信息
func (c *cCciData) View(ctx context.Context, req *ccidata.ViewReq) (res *ccidata.ViewRes, err error) {
	data, err := service.StockCciData().View(ctx, &req.CciDataViewInp)
	if err != nil {
		return
	}

	res = new(ccidata.ViewRes)
	res.CciDataViewModel = data
	return
}

// Delete 删除cci指标数据表
func (c *cCciData) Delete(ctx context.Context, req *ccidata.DeleteReq) (res *ccidata.DeleteRes, err error) {
	err = service.StockCciData().Delete(ctx, &req.CciDataDeleteInp)
	return
}

// GetCci 获取cci数据指标
func (c *cCciData) GetCci(ctx context.Context, req *ccidata.GetCciReq) (res *ccidata.GetCciRes, err error) {
	//_, err = service.StockCciData().GetCci(ctx, &req.GetCciDataInp)
	//service.StockRsiData().GetRsiData(ctx, &stockin.GetRsiDataInp{Code: req.Symbol})
	return
}
