// Package stock
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2026 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
// @AutoGenerate Version 2.18.6
package stock

import (
	"context"
	"hotgo/api/stock/momdata"
	"hotgo/internal/model/input/stockin"
	"hotgo/internal/service"
)

var (
	MomData = cMomData{}
)

type cMomData struct{}

// List 查看mom指标数据表列表
func (c *cMomData) List(ctx context.Context, req *momdata.ListReq) (res *momdata.ListRes, err error) {
	list, totalCount, err := service.StockMomData().List(ctx, &req.MomDataListInp)
	if err != nil {
		return
	}

	if list == nil {
		list = []*stockin.MomDataListModel{}
	}

	res = new(momdata.ListRes)
	res.List = list
	res.PageRes.Pack(req, totalCount)
	return
}

// Export 导出mom指标数据表列表
func (c *cMomData) Export(ctx context.Context, req *momdata.ExportReq) (res *momdata.ExportRes, err error) {
	err = service.StockMomData().Export(ctx, &req.MomDataListInp)
	return
}

// Edit 更新mom指标数据表
func (c *cMomData) Edit(ctx context.Context, req *momdata.EditReq) (res *momdata.EditRes, err error) {
	err = service.StockMomData().Edit(ctx, &req.MomDataEditInp)
	return
}

// View 获取指定mom指标数据表信息
func (c *cMomData) View(ctx context.Context, req *momdata.ViewReq) (res *momdata.ViewRes, err error) {
	data, err := service.StockMomData().View(ctx, &req.MomDataViewInp)
	if err != nil {
		return
	}

	res = new(momdata.ViewRes)
	res.MomDataViewModel = data
	return
}

// Delete 删除mom指标数据表
func (c *cMomData) Delete(ctx context.Context, req *momdata.DeleteReq) (res *momdata.DeleteRes, err error) {
	err = service.StockMomData().Delete(ctx, &req.MomDataDeleteInp)
	return
}

// GetMom 获取mom指标数据表
func (c *cSlowStochasticData) GetMom(ctx context.Context, req *momdata.GetMomReq) (res []*momdata.GetMomRes, err error) {
	_, err = service.StockMomData().GetMom(ctx, &req.GetMomDataInp)
	return
}
