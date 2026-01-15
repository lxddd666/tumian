// Package stock
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2026 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
// @AutoGenerate Version 2.18.6
package stock

import (
	"context"
	"hotgo/api/stock/toptencirculatingholders"
	"hotgo/internal/model/input/stockin"
	"hotgo/internal/service"
)

var (
	TopTenCirculatingHolders = cTopTenCirculatingHolders{}
)

type cTopTenCirculatingHolders struct{}

// List 查看公司十大流通股东表 (数据来源于定期报告)[citation:4]列表
func (c *cTopTenCirculatingHolders) List(ctx context.Context, req *toptencirculatingholders.ListReq) (res *toptencirculatingholders.ListRes, err error) {
	list, totalCount, err := service.StockTopTenCirculatingHolders().List(ctx, &req.TopTenCirculatingHoldersListInp)
	if err != nil {
		return
	}

	if list == nil {
		list = []*stockin.TopTenCirculatingHoldersListModel{}
	}

	res = new(toptencirculatingholders.ListRes)
	res.List = list
	res.PageRes.Pack(req, totalCount)
	return
}

// Export 导出公司十大流通股东表 (数据来源于定期报告)[citation:4]列表
func (c *cTopTenCirculatingHolders) Export(ctx context.Context, req *toptencirculatingholders.ExportReq) (res *toptencirculatingholders.ExportRes, err error) {
	err = service.StockTopTenCirculatingHolders().Export(ctx, &req.TopTenCirculatingHoldersListInp)
	return
}

// Edit 更新公司十大流通股东表 (数据来源于定期报告)[citation:4]
func (c *cTopTenCirculatingHolders) Edit(ctx context.Context, req *toptencirculatingholders.EditReq) (res *toptencirculatingholders.EditRes, err error) {
	err = service.StockTopTenCirculatingHolders().Edit(ctx, &req.TopTenCirculatingHoldersEditInp)
	return
}

// View 获取指定公司十大流通股东表 (数据来源于定期报告)[citation:4]信息
func (c *cTopTenCirculatingHolders) View(ctx context.Context, req *toptencirculatingholders.ViewReq) (res *toptencirculatingholders.ViewRes, err error) {
	data, err := service.StockTopTenCirculatingHolders().View(ctx, &req.TopTenCirculatingHoldersViewInp)
	if err != nil {
		return
	}

	res = new(toptencirculatingholders.ViewRes)
	res.TopTenCirculatingHoldersViewModel = data
	return
}

// Delete 删除公司十大流通股东表 (数据来源于定期报告)[citation:4]
func (c *cTopTenCirculatingHolders) Delete(ctx context.Context, req *toptencirculatingholders.DeleteReq) (res *toptencirculatingholders.DeleteRes, err error) {
	err = service.StockTopTenCirculatingHolders().Delete(ctx, &req.TopTenCirculatingHoldersDeleteInp)
	return
}
