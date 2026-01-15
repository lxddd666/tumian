// Package stock
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2026 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
// @AutoGenerate Version 2.18.6
package stock

import (
	"context"
	"hotgo/api/stock/macddata"
	"hotgo/internal/model/input/stockin"
	"hotgo/internal/service"
)

var (
	MacdData = cMacdData{}
)

type cMacdData struct{}

// List 查看MACD指标数据表列表
func (c *cMacdData) List(ctx context.Context, req *macddata.ListReq) (res *macddata.ListRes, err error) {
	list, totalCount, err := service.StockMacdData().List(ctx, &req.MacdDataListInp)
	if err != nil {
		return
	}

	if list == nil {
		list = []*stockin.MacdDataListModel{}
	}

	res = new(macddata.ListRes)
	res.List = list
	res.PageRes.Pack(req, totalCount)
	return
}

// Export 导出MACD指标数据表列表
func (c *cMacdData) Export(ctx context.Context, req *macddata.ExportReq) (res *macddata.ExportRes, err error) {
	err = service.StockMacdData().Export(ctx, &req.MacdDataListInp)
	return
}

// Edit 更新MACD指标数据表
func (c *cMacdData) Edit(ctx context.Context, req *macddata.EditReq) (res *macddata.EditRes, err error) {
	err = service.StockMacdData().Edit(ctx, &req.MacdDataEditInp)
	return
}

// View 获取指定MACD指标数据表信息
func (c *cMacdData) View(ctx context.Context, req *macddata.ViewReq) (res *macddata.ViewRes, err error) {
	data, err := service.StockMacdData().View(ctx, &req.MacdDataViewInp)
	if err != nil {
		return
	}

	res = new(macddata.ViewRes)
	res.MacdDataViewModel = data
	return
}

// Delete 删除MACD指标数据表
func (c *cMacdData) Delete(ctx context.Context, req *macddata.DeleteReq) (res *macddata.DeleteRes, err error) {
	err = service.StockMacdData().Delete(ctx, &req.MacdDataDeleteInp)
	return
}
