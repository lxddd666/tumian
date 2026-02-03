// Package stock
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2026 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
// @AutoGenerate Version 2.18.6
package stock

import (
	"context"
	"hotgo/api/stock/stockselfcode"
	"hotgo/internal/service"
)

var (
	StockSelfCode = cStockSelfCode{}
)

type cStockSelfCode struct{}

// SelfCodeIndicatorsApi 技术指标获取api获取
func (c *cStockSelfCode) SelfCodeIndicatorsApi(ctx context.Context, req *stockselfcode.SelfCodeIndicatorsApiReq) (res *stockselfcode.SelfCodeIndicatorsApiRes, err error) {
	err = service.StockSelfCode().SelfCodeIndicatorsApi(ctx, &req.SelfCodeIndicatorsApiInp)
	return
}

// SelfStockWorkingCapitalInfoApi 运营资金情况（智兔api获取）
func (c *cStockSelfCode) SelfStockWorkingCapitalInfoApi(ctx context.Context, req *stockselfcode.SelfStockWorkingCapitalInfoApiReq) (res *stockselfcode.SelfStockWorkingCapitalInfoApiRes, err error) {
	err = service.StockSelfCode().SelfStockWorkingCapitalInfoApi(ctx, &req.SelfCodeIndicatorsApiInp)
	return
}

func (c *cStockSelfCode) BaiduFinanceCode(ctx context.Context, req *stockselfcode.BaiduFinanceCodeReq) (res *stockselfcode.BaiduFinanceCodeRes, err error) {
	err = service.StockSelfCode().BaiduFinanceCode(ctx, &req.SelfCodeIndicatorsApiInp)

	return
}
