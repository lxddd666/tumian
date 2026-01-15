// Package stock
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2026 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
// @AutoGenerate Version 2.18.6
package stock

import (
	"context"
	"hotgo/api/stock/incomestatement"
	"hotgo/internal/model/input/stockin"
	"hotgo/internal/service"
)

var (
	IncomeStatement = cIncomeStatement{}
)

type cIncomeStatement struct{}

// List 查看利润表 (Income Statement)列表
func (c *cIncomeStatement) List(ctx context.Context, req *incomestatement.ListReq) (res *incomestatement.ListRes, err error) {
	list, totalCount, err := service.StockIncomeStatement().List(ctx, &req.IncomeStatementListInp)
	if err != nil {
		return
	}

	if list == nil {
		list = []*stockin.IncomeStatementListModel{}
	}

	res = new(incomestatement.ListRes)
	res.List = list
	res.PageRes.Pack(req, totalCount)
	return
}

// Export 导出利润表 (Income Statement)列表
func (c *cIncomeStatement) Export(ctx context.Context, req *incomestatement.ExportReq) (res *incomestatement.ExportRes, err error) {
	err = service.StockIncomeStatement().Export(ctx, &req.IncomeStatementListInp)
	return
}

// Edit 更新利润表 (Income Statement)
func (c *cIncomeStatement) Edit(ctx context.Context, req *incomestatement.EditReq) (res *incomestatement.EditRes, err error) {
	err = service.StockIncomeStatement().Edit(ctx, &req.IncomeStatementEditInp)
	return
}

// View 获取指定利润表 (Income Statement)信息
func (c *cIncomeStatement) View(ctx context.Context, req *incomestatement.ViewReq) (res *incomestatement.ViewRes, err error) {
	data, err := service.StockIncomeStatement().View(ctx, &req.IncomeStatementViewInp)
	if err != nil {
		return
	}

	res = new(incomestatement.ViewRes)
	res.IncomeStatementViewModel = data
	return
}

// Delete 删除利润表 (Income Statement)
func (c *cIncomeStatement) Delete(ctx context.Context, req *incomestatement.DeleteReq) (res *incomestatement.DeleteRes, err error) {
	err = service.StockIncomeStatement().Delete(ctx, &req.IncomeStatementDeleteInp)
	return
}

// GetIncomeStatement 获取利润表数据
func (c *cIncomeStatement) GetIncomeStatement(ctx context.Context, req *incomestatement.GetIncomeStatementReq) (res *incomestatement.GetIncomeStatementRes, err error) {
	data, err := service.StockIncomeStatement().GetIncomeStatement(ctx, &req.IncomeStatementGetIncomeStatementInp)
	if err != nil {
		return
	}

	res = new(incomestatement.GetIncomeStatementRes)
	res.Data = data
	return
}
