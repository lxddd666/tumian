// Package stock
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2026 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
// @AutoGenerate Version 2.18.6
package stock

import (
	"context"
	"hotgo/api/stock/atrdata"
	"hotgo/internal/model/input/stockin"
	"hotgo/internal/service"
)

var (
	AtrData = cAtrData{}
)

type cAtrData struct{}

// List 查看atr指标数据表列表
func (c *cAtrData) List(ctx context.Context, req *atrdata.ListReq) (res *atrdata.ListRes, err error) {
	list, totalCount, err := service.StockAtrData().List(ctx, &req.AtrDataListInp)
	if err != nil {
		return
	}

	if list == nil {
		list = []*stockin.AtrDataListModel{}
	}

	res = new(atrdata.ListRes)
	res.List = list
	res.PageRes.Pack(req, totalCount)
	return
}

// Export 导出atr指标数据表列表
func (c *cAtrData) Export(ctx context.Context, req *atrdata.ExportReq) (res *atrdata.ExportRes, err error) {
	err = service.StockAtrData().Export(ctx, &req.AtrDataListInp)
	return
}

// Edit 更新atr指标数据表
func (c *cAtrData) Edit(ctx context.Context, req *atrdata.EditReq) (res *atrdata.EditRes, err error) {
	err = service.StockAtrData().Edit(ctx, &req.AtrDataEditInp)
	return
}

// View 获取指定atr指标数据表信息
func (c *cAtrData) View(ctx context.Context, req *atrdata.ViewReq) (res *atrdata.ViewRes, err error) {
	data, err := service.StockAtrData().View(ctx, &req.AtrDataViewInp)
	if err != nil {
		return
	}

	res = new(atrdata.ViewRes)
	res.AtrDataViewModel = data
	return
}

// Delete 删除atr指标数据表
func (c *cAtrData) Delete(ctx context.Context, req *atrdata.DeleteReq) (res *atrdata.DeleteRes, err error) {
	err = service.StockAtrData().Delete(ctx, &req.AtrDataDeleteInp)
	return
}
