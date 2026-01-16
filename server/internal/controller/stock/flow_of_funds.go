// Package stock
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2026 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
// @AutoGenerate Version 2.18.6
package stock

import (
	"context"
	"hotgo/api/stock/flowoffunds"
	"hotgo/internal/model/input/stockin"
	"hotgo/internal/service"
)

var (
	FlowOfFunds = cFlowOfFunds{}
)

type cFlowOfFunds struct{}

// List 查看资金流向明细表列表
func (c *cFlowOfFunds) List(ctx context.Context, req *flowoffunds.ListReq) (res *flowoffunds.ListRes, err error) {
	list, totalCount, err := service.StockFlowOfFunds().List(ctx, &req.FlowOfFundsListInp)
	if err != nil {
		return
	}

	if list == nil {
		list = []*stockin.FlowOfFundsListModel{}
	}

	res = new(flowoffunds.ListRes)
	res.List = list
	res.PageRes.Pack(req, totalCount)
	return
}

// Export 导出资金流向明细表列表
func (c *cFlowOfFunds) Export(ctx context.Context, req *flowoffunds.ExportReq) (res *flowoffunds.ExportRes, err error) {
	err = service.StockFlowOfFunds().Export(ctx, &req.FlowOfFundsListInp)
	return
}

// Edit 更新资金流向明细表
func (c *cFlowOfFunds) Edit(ctx context.Context, req *flowoffunds.EditReq) (res *flowoffunds.EditRes, err error) {
	err = service.StockFlowOfFunds().Edit(ctx, &req.FlowOfFundsEditInp)
	return
}

// View 获取指定资金流向明细表信息
func (c *cFlowOfFunds) View(ctx context.Context, req *flowoffunds.ViewReq) (res *flowoffunds.ViewRes, err error) {
	data, err := service.StockFlowOfFunds().View(ctx, &req.FlowOfFundsViewInp)
	if err != nil {
		return
	}

	res = new(flowoffunds.ViewRes)
	res.FlowOfFundsViewModel = data
	return
}

// Delete 删除资金流向明细表
func (c *cFlowOfFunds) Delete(ctx context.Context, req *flowoffunds.DeleteReq) (res *flowoffunds.DeleteRes, err error) {
	err = service.StockFlowOfFunds().Delete(ctx, &req.FlowOfFundsDeleteInp)
	return
}

// GetFlowOfFunds 获取资金流向明细表数据
func (c *cFlowOfFunds) GetFlowOfFunds(ctx context.Context, req *flowoffunds.GetFlowOfFundsReq) (res *flowoffunds.GetFlowOfFundsRes, err error) {
	data, err := service.StockFlowOfFunds().GetFlowOfFunds(ctx, &req.GetFlowOfFundsInp)
	if err != nil {
		return
	}

	res = new(flowoffunds.GetFlowOfFundsRes)
	res.Data = data
	return
}
