// Package stock
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2026 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
// @AutoGenerate Version 2.18.6
package stock

import (
	"context"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/errors/gerror"
	"hotgo/internal/dao"
	"hotgo/internal/library/hgorm/handler"
	"hotgo/internal/model/entity"
	"hotgo/internal/model/input/stockin"
	"hotgo/internal/service"
)

type sStockSelfCode struct{}

func NewStockSelfCode() *sStockSelfCode {
	return &sStockSelfCode{}
}

func init() {
	service.RegisterStockSelfCode(NewStockSelfCode())
}

// Model 自选股票ORM模型
func (s *sStockSelfCode) Model(ctx context.Context, option ...*handler.Option) *gdb.Model {
	return handler.Model(dao.StockSelfCode.Ctx(ctx), option...)
}

// SelfCodeIndicatorsApi 技术指标获取api获取（智兔api获取）
func (s *sStockSelfCode) SelfCodeIndicatorsApi(ctx context.Context, in *stockin.SelfCodeIndicatorsApiInp) (err error) {
	// 获取自己观察的股票
	list, err := s.GetAllSelfCode(ctx)
	if err != nil {
		return
	}
	for _, stockCode := range list {

		// k线
		_, _ = service.StockEnterpriseHistoricalData().GetEnterpriseHistoricalData(ctx, &stockin.EnterpriseHistoricalDataGetEnterpriseHistoricalDataInp{Symbol: stockCode.Dm})

		// ma
		_, _ = service.StockMaData().GetMa(ctx, &stockin.MaDataGetMaInp{Symbol: stockCode.Dm})
		//
		// macd
		_, _ = service.StockMacdData().GetMacd(ctx, &stockin.MacdDataGetMacdInp{Symbol: stockCode.Dm})
		//
		// kdj
		_, _ = service.StockKdjData().GetKdj(ctx, &stockin.KdjDataGetKdjInp{Symbol: stockCode.Dm})
		//
		// boll
		_, _ = service.StockBollData().GetBoll(ctx, &stockin.BollDataGetBollInp{Symbol: stockCode.Dm})

		// 资金流向
		_, _ = service.StockFlowOfFunds().GetFlowOfFunds(ctx, &stockin.GetFlowOfFundsInp{Symbol: stockCode.Dm})
	}
	return
}

// SelfStockWorkingCapitalInfoApi 运营资金情况（智兔api获取）
func (s *sStockSelfCode) SelfStockWorkingCapitalInfoApi(ctx context.Context, in *stockin.SelfCodeIndicatorsApiInp) (err error) {
	// 获取自己观察的股票
	list, err := s.GetAllSelfCode(ctx)
	if err != nil {
		return
	}
	for _, stockCode := range list {
		// 获取财务指标分析表数据
		_, _ = service.StockFinancialIndicators().GetFinancialIndicators(ctx, &stockin.FinancialIndicatorsGetFinancialIndicatorsInp{Symbol: stockCode.Dm})
		//
		//// 基金持股明细表
		_, _ = service.StockFundStockHolding().GetFundStockHolding(ctx, &stockin.FundStockHoldingGetFundStockHoldingInp{Symbol: stockCode.Dm})

		// 企业利润表
		_, _ = service.StockIncomeStatement().GetIncomeStatement(ctx, &stockin.IncomeStatementGetIncomeStatementInp{Symbol: stockCode.Dm})

		// 季度利润
		_, _ = service.StockQuarterlyProfit().GetQuarterlyProfit(ctx, &stockin.QuarterlyProfitGetQuarterlyProfitInp{Symbol: stockCode.Dm})

		// 股东户数变化记录表
		_, _ = service.StockShareholderChange().GetShareholderChange(ctx, &stockin.ShareholderChangeGetShareholderChangeInp{Symbol: stockCode.Dm})

		// 公司股东户数统计表数据 (有问题)
		_, _ = service.StockShareholderCount().GetShareholderCount(ctx, &stockin.ShareholderCountGetShareholderCountInp{Symbol: stockCode.Dm})

		// 股票基础信息表数据
		// //_, _ = service.StockBasicInfo().GetStockBasicInfo(ctx, &stockin.StockBasicInfoGetStockBasicInfoInp{Symbol: stockCode.Dm})

		// 公司十大流通股东
		// //_, _ = service.StockTopTenCirculatingHolders().GetTopTenCirculatingHolders(ctx, &stockin.TopTenCirculatingHoldersGetTopTenCirculatingHoldersInp{Symbol: stockCode.Dm})

	}
	return
}

// GetAllSelfCode 获取所有code
func (s *sStockSelfCode) GetAllSelfCode(ctx context.Context) (list []*entity.StockSelfCode, err error) {
	err = s.Model(ctx).Scan(&list)
	if err != nil {
		return
	}
	if len(list) == 0 {
		err = gerror.New("获取自选股票为空")
		return
	}
	return
}
