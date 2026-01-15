// Package stock
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2026 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
// @AutoGenerate Version 2.18.6
package stock

import (
	"context"
	"hotgo/api/stock/kdjdata"
	"hotgo/internal/model/input/stockin"
	"hotgo/internal/service"
)

var (
	KdjData = cKdjData{}
)

type cKdjData struct{}

// List 查看KDJ随机指标数据表列表
func (c *cKdjData) List(ctx context.Context, req *kdjdata.ListReq) (res *kdjdata.ListRes, err error) {
	list, totalCount, err := service.StockKdjData().List(ctx, &req.KdjDataListInp)
	if err != nil {
		return
	}

	if list == nil {
		list = []*stockin.KdjDataListModel{}
	}

	res = new(kdjdata.ListRes)
	res.List = list
	res.PageRes.Pack(req, totalCount)
	return
}

// Export 导出KDJ随机指标数据表列表
func (c *cKdjData) Export(ctx context.Context, req *kdjdata.ExportReq) (res *kdjdata.ExportRes, err error) {
	err = service.StockKdjData().Export(ctx, &req.KdjDataListInp)
	return
}

// Edit 更新KDJ随机指标数据表
func (c *cKdjData) Edit(ctx context.Context, req *kdjdata.EditReq) (res *kdjdata.EditRes, err error) {
	err = service.StockKdjData().Edit(ctx, &req.KdjDataEditInp)
	return
}

// View 获取指定KDJ随机指标数据表信息
func (c *cKdjData) View(ctx context.Context, req *kdjdata.ViewReq) (res *kdjdata.ViewRes, err error) {
	data, err := service.StockKdjData().View(ctx, &req.KdjDataViewInp)
	if err != nil {
		return
	}

	res = new(kdjdata.ViewRes)
	res.KdjDataViewModel = data
	return
}

// Delete 删除KDJ随机指标数据表
func (c *cKdjData) Delete(ctx context.Context, req *kdjdata.DeleteReq) (res *kdjdata.DeleteRes, err error) {
	err = service.StockKdjData().Delete(ctx, &req.KdjDataDeleteInp)
	return
}

// GetKdj 获取KDJ随机指标数据
func (c *cKdjData) GetKdj(ctx context.Context, req *kdjdata.GetKdjReq) (res *kdjdata.GetKdjRes, err error) {
	data, err := service.StockKdjData().GetKdj(ctx, &req.KdjDataGetKdjInp)
	if err != nil {
		return
	}

	res = new(kdjdata.GetKdjRes)
	res.Data = data
	return
}
