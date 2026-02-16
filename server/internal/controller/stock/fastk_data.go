// Package stock
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2026 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
// @AutoGenerate Version 2.18.6
package stock

import (
	"context"
	"hotgo/api/stock/fastkdata"
	"hotgo/internal/model/input/stockin"
	"hotgo/internal/service"
)

var (
	FastkData = cFastkData{}
)

type cFastkData struct{}

// List 查看fastk指标数据表列表
func (c *cFastkData) List(ctx context.Context, req *fastkdata.ListReq) (res *fastkdata.ListRes, err error) {
	list, totalCount, err := service.StockFastkData().List(ctx, &req.FastkDataListInp)
	if err != nil {
		return
	}

	if list == nil {
		list = []*stockin.FastkDataListModel{}
	}

	res = new(fastkdata.ListRes)
	res.List = list
	res.PageRes.Pack(req, totalCount)
	return
}

// Export 导出fastk指标数据表列表
func (c *cFastkData) Export(ctx context.Context, req *fastkdata.ExportReq) (res *fastkdata.ExportRes, err error) {
	err = service.StockFastkData().Export(ctx, &req.FastkDataListInp)
	return
}

// Edit 更新fastk指标数据表
func (c *cFastkData) Edit(ctx context.Context, req *fastkdata.EditReq) (res *fastkdata.EditRes, err error) {
	err = service.StockFastkData().Edit(ctx, &req.FastkDataEditInp)
	return
}

// View 获取指定fastk指标数据表信息
func (c *cFastkData) View(ctx context.Context, req *fastkdata.ViewReq) (res *fastkdata.ViewRes, err error) {
	data, err := service.StockFastkData().View(ctx, &req.FastkDataViewInp)
	if err != nil {
		return
	}

	res = new(fastkdata.ViewRes)
	res.FastkDataViewModel = data
	return
}

// Delete 删除fastk指标数据表
func (c *cFastkData) Delete(ctx context.Context, req *fastkdata.DeleteReq) (res *fastkdata.DeleteRes, err error) {
	err = service.StockFastkData().Delete(ctx, &req.FastkDataDeleteInp)
	return
}
