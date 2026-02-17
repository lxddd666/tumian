// Package stock
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2026 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
// @AutoGenerate Version 2.18.6
package stock

import (
	"context"
	"fmt"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gogf/gf/v2/util/gconv"
	"hotgo/internal/dao"
	"hotgo/internal/library/hgorm/handler"
	"hotgo/internal/model/entity"
	"hotgo/internal/model/input/stockin"
	"hotgo/internal/service"
	"hotgo/utility/simple"
	"sync"
	"time"
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
	list, err := s.GetAllSelfCode(ctx, in.Code)
	if err != nil {
		return
	}

	wg := sync.WaitGroup{}

	for i, stock := range list {
		stockCode := stock
		wg.Add(1)
		fmt.Println(i)
		time.Sleep(gconv.Duration(i*100) * time.Millisecond)

		simple.SafeGo(gctx.New(), func(ctx context.Context) {
			defer wg.Done()
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

			// rsi
			_, _ = service.StockRsiData().GetRsiData(ctx, &stockin.GetRsiDataInp{Code: stockCode.Dm})

			// atr
			//_, _ = service.StockAtrData().GetAtrData(ctx, &stockin.GetAtrDataInp{Symbol: stockCode.Dm})

			// cci
			_, _ = service.StockCciData().GetCci(ctx, &stockin.GetCciDataInp{Symbol: stockCode.Dm})

			// stoch
			_, _ = service.StockSlowStochasticData().GetStoch(ctx, &stockin.GetStochDataInp{Symbol: stockCode.Dm})

			// mom
			_, _ = service.StockMomData().GetMom(ctx, &stockin.GetMomDataInp{Symbol: stockCode.Dm})

			//wmsr
			_, _ = service.StockWilliamsData().GetWmsr(ctx, &stockin.GetWilliamsDataInp{Symbol: stockCode.Dm})

			// kst
			_, _ = service.StockKstData().GetKst(ctx, &stockin.GetKstDataInp{Symbol: stockCode.Dm})

			// FASTK
			_, _ = service.StockFastkData().GetFastk(ctx, &stockin.GetFastkDataInp{Symbol: stockCode.Dm})

		})
	}
	wg.Done()
	return
}

// SelfStockWorkingCapitalInfoApi 运营资金情况（智兔api获取）
func (s *sStockSelfCode) SelfStockWorkingCapitalInfoApi(ctx context.Context, in *stockin.SelfCodeIndicatorsApiInp) (err error) {
	// 获取自己观察的股票
	list, err := s.GetAllSelfCode(ctx, in.Code)
	if err != nil {
		return
	}

	wg := sync.WaitGroup{}

	for i, stockCode := range list {
		wg.Add(1)
		time.Sleep(gconv.Duration(i*100) * time.Millisecond)

		simple.SafeGo(gctx.New(), func(ctx context.Context) {
			defer wg.Done()
			// 获取财务指标分析表数据
			_, _ = service.StockFinancialIndicators().GetFinancialIndicators(ctx, &stockin.FinancialIndicatorsGetFinancialIndicatorsInp{Symbol: stockCode.Dm})
			//
			//// 基金持股明细表
			//_, _ = service.StockFundStockHolding().GetFundStockHolding(ctx, &stockin.FundStockHoldingGetFundStockHoldingInp{Symbol: stockCode.Dm})

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
		})
	}
	wg.Done()

	return
}

// GetAllSelfCode 获取所有code
func (s *sStockSelfCode) GetAllSelfCode(ctx context.Context, code string) (list []*entity.StockSelfCode, err error) {
	mod := s.Model(ctx)
	if code != "" {
		mod.Where(dao.StockSelfCode.Columns().Dm, code)
	}
	err = mod.Scan(&list)
	if err != nil {
		return
	}
	if len(list) == 0 {
		err = gerror.New("获取自选股票为空")
		return
	}
	return
}

func GetSelfCodeDataForAi(ctx context.Context) {

}
