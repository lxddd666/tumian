// Package stock
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2026 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
// @AutoGenerate Version 2.18.6
package stock

import (
	"context"
	"hotgo/api/stock/kstdata"
	"hotgo/internal/model/input/stockin"
	"hotgo/internal/service"
)

var (
	KstData = cKstData{}
)

type cKstData struct{}

// List 查看kst指标数据表列表
func (c *cKstData) List(ctx context.Context, req *kstdata.ListReq) (res *kstdata.ListRes, err error) {
	list, totalCount, err := service.StockKstData().List(ctx, &req.KstDataListInp)
	if err != nil {
		return
	}

	if list == nil {
		list = []*stockin.KstDataListModel{}
	}

	res = new(kstdata.ListRes)
	res.List = list
	res.PageRes.Pack(req, totalCount)
	return
}

// Export 导出kst指标数据表列表
func (c *cKstData) Export(ctx context.Context, req *kstdata.ExportReq) (res *kstdata.ExportRes, err error) {
	err = service.StockKstData().Export(ctx, &req.KstDataListInp)
	return
}

// Edit 更新kst指标数据表
func (c *cKstData) Edit(ctx context.Context, req *kstdata.EditReq) (res *kstdata.EditRes, err error) {
	err = service.StockKstData().Edit(ctx, &req.KstDataEditInp)
	return
}

// View 获取指定kst指标数据表信息
func (c *cKstData) View(ctx context.Context, req *kstdata.ViewReq) (res *kstdata.ViewRes, err error) {
	data, err := service.StockKstData().View(ctx, &req.KstDataViewInp)
	if err != nil {
		return
	}

	res = new(kstdata.ViewRes)
	res.KstDataViewModel = data
	return
}

// Delete 删除kst指标数据表
func (c *cKstData) Delete(ctx context.Context, req *kstdata.DeleteReq) (res *kstdata.DeleteRes, err error) {
	err = service.StockKstData().Delete(ctx, &req.KstDataDeleteInp)
	return
}
