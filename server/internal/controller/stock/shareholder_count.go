// Package stock
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2026 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
// @AutoGenerate Version 2.18.6
package stock

import (
	"context"
	"hotgo/api/stock/shareholdercount"
	"hotgo/internal/model/input/stockin"
	"hotgo/internal/service"
)

var (
	ShareholderCount = cShareholderCount{}
)

type cShareholderCount struct{}

// List 查看公司股东户数统计表 (按报告期统计)列表
func (c *cShareholderCount) List(ctx context.Context, req *shareholdercount.ListReq) (res *shareholdercount.ListRes, err error) {
	list, totalCount, err := service.StockShareholderCount().List(ctx, &req.ShareholderCountListInp)
	if err != nil {
		return
	}

	if list == nil {
		list = []*stockin.ShareholderCountListModel{}
	}

	res = new(shareholdercount.ListRes)
	res.List = list
	res.PageRes.Pack(req, totalCount)
	return
}

// Export 导出公司股东户数统计表 (按报告期统计)列表
func (c *cShareholderCount) Export(ctx context.Context, req *shareholdercount.ExportReq) (res *shareholdercount.ExportRes, err error) {
	err = service.StockShareholderCount().Export(ctx, &req.ShareholderCountListInp)
	return
}

// Edit 更新公司股东户数统计表 (按报告期统计)
func (c *cShareholderCount) Edit(ctx context.Context, req *shareholdercount.EditReq) (res *shareholdercount.EditRes, err error) {
	err = service.StockShareholderCount().Edit(ctx, &req.ShareholderCountEditInp)
	return
}

// View 获取指定公司股东户数统计表 (按报告期统计)信息
func (c *cShareholderCount) View(ctx context.Context, req *shareholdercount.ViewReq) (res *shareholdercount.ViewRes, err error) {
	data, err := service.StockShareholderCount().View(ctx, &req.ShareholderCountViewInp)
	if err != nil {
		return
	}

	res = new(shareholdercount.ViewRes)
	res.ShareholderCountViewModel = data
	return
}

// Delete 删除公司股东户数统计表 (按报告期统计)
func (c *cShareholderCount) Delete(ctx context.Context, req *shareholdercount.DeleteReq) (res *shareholdercount.DeleteRes, err error) {
	err = service.StockShareholderCount().Delete(ctx, &req.ShareholderCountDeleteInp)
	return
}

// GetShareholderCount 获取公司股东户数统计表数据
func (c *cShareholderCount) GetShareholderCount(ctx context.Context, req *shareholdercount.GetShareholderCountReq) (res *shareholdercount.GetShareholderCountRes, err error) {
	data, err := service.StockShareholderCount().GetShareholderCount(ctx, &req.ShareholderCountGetShareholderCountInp)
	if err != nil {
		return
	}

	res = new(shareholdercount.GetShareholderCountRes)
	res.Data = data
	return
}
