// Package stock
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2026 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
// @AutoGenerate Version 2.18.6
package stock

import (
	"context"
	"hotgo/api/stock/williamsdata"
	"hotgo/internal/model/input/stockin"
	"hotgo/internal/service"
)

var (
	WilliamsData = cWilliamsData{}
)

type cWilliamsData struct{}

// List 查看wmsr指标数据表列表
func (c *cWilliamsData) List(ctx context.Context, req *williamsdata.ListReq) (res *williamsdata.ListRes, err error) {
	list, totalCount, err := service.StockWilliamsData().List(ctx, &req.WilliamsDataListInp)
	if err != nil {
		return
	}

	if list == nil {
		list = []*stockin.WilliamsDataListModel{}
	}

	res = new(williamsdata.ListRes)
	res.List = list
	res.PageRes.Pack(req, totalCount)
	return
}

// Export 导出wmsr指标数据表列表
func (c *cWilliamsData) Export(ctx context.Context, req *williamsdata.ExportReq) (res *williamsdata.ExportRes, err error) {
	err = service.StockWilliamsData().Export(ctx, &req.WilliamsDataListInp)
	return
}

// Edit 更新wmsr指标数据表
func (c *cWilliamsData) Edit(ctx context.Context, req *williamsdata.EditReq) (res *williamsdata.EditRes, err error) {
	err = service.StockWilliamsData().Edit(ctx, &req.WilliamsDataEditInp)
	return
}

// View 获取指定wmsr指标数据表信息
func (c *cWilliamsData) View(ctx context.Context, req *williamsdata.ViewReq) (res *williamsdata.ViewRes, err error) {
	data, err := service.StockWilliamsData().View(ctx, &req.WilliamsDataViewInp)
	if err != nil {
		return
	}

	res = new(williamsdata.ViewRes)
	res.WilliamsDataViewModel = data
	return
}

// Delete 删除wmsr指标数据表
func (c *cWilliamsData) Delete(ctx context.Context, req *williamsdata.DeleteReq) (res *williamsdata.DeleteRes, err error) {
	err = service.StockWilliamsData().Delete(ctx, &req.WilliamsDataDeleteInp)
	return
}
