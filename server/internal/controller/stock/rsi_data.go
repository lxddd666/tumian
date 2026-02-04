// Package stock
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2026 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
// @AutoGenerate Version 2.18.6
package stock

import (
	"context"
	"hotgo/api/stock/rsidata"
	"hotgo/internal/model/input/stockin"
	"hotgo/internal/service"
)

var (
	RsiData = cRsiData{}
)

type cRsiData struct{}

// List 查看rsi指标数据表列表
func (c *cRsiData) List(ctx context.Context, req *rsidata.ListReq) (res *rsidata.ListRes, err error) {
	list, totalCount, err := service.StockRsiData().List(ctx, &req.RsiDataListInp)
	if err != nil {
		return
	}

	if list == nil {
		list = []*stockin.RsiDataListModel{}
	}

	res = new(rsidata.ListRes)
	res.List = list
	res.PageRes.Pack(req, totalCount)
	return
}

// Export 导出rsi指标数据表列表
func (c *cRsiData) Export(ctx context.Context, req *rsidata.ExportReq) (res *rsidata.ExportRes, err error) {
	err = service.StockRsiData().Export(ctx, &req.RsiDataListInp)
	return
}

// Edit 更新rsi指标数据表
func (c *cRsiData) Edit(ctx context.Context, req *rsidata.EditReq) (res *rsidata.EditRes, err error) {
	err = service.StockRsiData().Edit(ctx, &req.RsiDataEditInp)
	return
}

// View 获取指定rsi指标数据表信息
func (c *cRsiData) View(ctx context.Context, req *rsidata.ViewReq) (res *rsidata.ViewRes, err error) {
	data, err := service.StockRsiData().View(ctx, &req.RsiDataViewInp)
	if err != nil {
		return
	}

	res = new(rsidata.ViewRes)
	res.RsiDataViewModel = data
	return
}

// Delete 删除rsi指标数据表
func (c *cRsiData) Delete(ctx context.Context, req *rsidata.DeleteReq) (res *rsidata.DeleteRes, err error) {
	err = service.StockRsiData().Delete(ctx, &req.RsiDataDeleteInp)
	return
}
