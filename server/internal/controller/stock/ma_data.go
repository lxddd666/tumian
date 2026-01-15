// Package stock
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2026 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
// @AutoGenerate Version 2.18.6
package stock

import (
	"context"
	"hotgo/api/stock/madata"
	"hotgo/internal/model/input/stockin"
	"hotgo/internal/service"
)

var (
	MaData = cMaData{}
)

type cMaData struct{}

// List 查看移动平均线(MA)指标数据表列表
func (c *cMaData) List(ctx context.Context, req *madata.ListReq) (res *madata.ListRes, err error) {
	list, totalCount, err := service.StockMaData().List(ctx, &req.MaDataListInp)
	if err != nil {
		return
	}

	if list == nil {
		list = []*stockin.MaDataListModel{}
	}

	res = new(madata.ListRes)
	res.List = list
	res.PageRes.Pack(req, totalCount)
	return
}

// Export 导出移动平均线(MA)指标数据表列表
func (c *cMaData) Export(ctx context.Context, req *madata.ExportReq) (res *madata.ExportRes, err error) {
	err = service.StockMaData().Export(ctx, &req.MaDataListInp)
	return
}

// Edit 更新移动平均线(MA)指标数据表
func (c *cMaData) Edit(ctx context.Context, req *madata.EditReq) (res *madata.EditRes, err error) {
	err = service.StockMaData().Edit(ctx, &req.MaDataEditInp)
	return
}

// View 获取指定移动平均线(MA)指标数据表信息
func (c *cMaData) View(ctx context.Context, req *madata.ViewReq) (res *madata.ViewRes, err error) {
	data, err := service.StockMaData().View(ctx, &req.MaDataViewInp)
	if err != nil {
		return
	}

	res = new(madata.ViewRes)
	res.MaDataViewModel = data
	return
}

// Delete 删除移动平均线(MA)指标数据表
func (c *cMaData) Delete(ctx context.Context, req *madata.DeleteReq) (res *madata.DeleteRes, err error) {
	err = service.StockMaData().Delete(ctx, &req.MaDataDeleteInp)
	return
}

// GetMa 获取移动平均线(MA)指标数据
func (c *cMaData) GetMa(ctx context.Context, req *madata.GetMaReq) (res *madata.GetMaRes, err error) {
	data, err := service.StockMaData().GetMa(ctx, &req.MaDataGetMaInp)
	if err != nil {
		return
	}

	res = new(madata.GetMaRes)
	res.Data = data
	return
}
