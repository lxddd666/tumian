// Package stock
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2026 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
// @AutoGenerate Version 2.18.6
package stock

import (
	"context"
	"hotgo/api/stock/shareholderchange"
	"hotgo/internal/model/input/stockin"
	"hotgo/internal/service"
)

var (
	ShareholderChange = cShareholderChange{}
)

type cShareholderChange struct{}

// List 查看股东户数变化记录表 (记录相邻报告期的户数变化)列表
func (c *cShareholderChange) List(ctx context.Context, req *shareholderchange.ListReq) (res *shareholderchange.ListRes, err error) {
	list, totalCount, err := service.StockShareholderChange().List(ctx, &req.ShareholderChangeListInp)
	if err != nil {
		return
	}

	if list == nil {
		list = []*stockin.ShareholderChangeListModel{}
	}
	res = new(shareholderchange.ListRes)
	res.List = list
	res.PageRes.Pack(req, totalCount)
	return
}

// Export 导出股东户数变化记录表 (记录相邻报告期的户数变化)列表
func (c *cShareholderChange) Export(ctx context.Context, req *shareholderchange.ExportReq) (res *shareholderchange.ExportRes, err error) {
	err = service.StockShareholderChange().Export(ctx, &req.ShareholderChangeListInp)
	return
}

// Edit 更新股东户数变化记录表 (记录相邻报告期的户数变化)
func (c *cShareholderChange) Edit(ctx context.Context, req *shareholderchange.EditReq) (res *shareholderchange.EditRes, err error) {
	err = service.StockShareholderChange().Edit(ctx, &req.ShareholderChangeEditInp)
	return
}

// View 获取指定股东户数变化记录表 (记录相邻报告期的户数变化)信息
func (c *cShareholderChange) View(ctx context.Context, req *shareholderchange.ViewReq) (res *shareholderchange.ViewRes, err error) {
	data, err := service.StockShareholderChange().View(ctx, &req.ShareholderChangeViewInp)
	if err != nil {
		return
	}

	res = new(shareholderchange.ViewRes)
	res.ShareholderChangeViewModel = data
	return
}

// Delete 删除股东户数变化记录表 (记录相邻报告期的户数变化)
func (c *cShareholderChange) Delete(ctx context.Context, req *shareholderchange.DeleteReq) (res *shareholderchange.DeleteRes, err error) {
	err = service.StockShareholderChange().Delete(ctx, &req.ShareholderChangeDeleteInp)
	return
}

// GetShareholderChange 获取股东户数变化记录表数据
func (c *cShareholderChange) GetShareholderChange(ctx context.Context, req *shareholderchange.GetShareholderChangeReq) (res *shareholderchange.GetShareholderChangeRes, err error) {
	data, err := service.StockShareholderChange().GetShareholderChange(ctx, &req.ShareholderChangeGetShareholderChangeInp)
	if err != nil {
		return
	}

	res = new(shareholderchange.GetShareholderChangeRes)
	res.Data = data
	return
}
