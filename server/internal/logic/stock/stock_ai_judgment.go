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
	"hotgo/internal/dao"
	"hotgo/internal/library/hgorm/handler"
	"hotgo/internal/model/entity"
	"hotgo/internal/model/input/stockin"
	"hotgo/internal/service"
)

type sStockAiJudgment struct{}

func NewStockAiJudgment() *sStockAiJudgment {
	return &sStockAiJudgment{}
}

func init() {
	service.RegisterStockAiJudgment(NewStockAiJudgment())
}

// Model ai 选股判断ORM模型
func (s *sStockAiJudgment) Model(ctx context.Context, option ...*handler.Option) *gdb.Model {
	return handler.Model(dao.StockAiJudgment.Ctx(ctx), option...)
}

// AITemp ai 提示模板
var AITemp = `交易对:%s 现价:%s
[1D趋势]
K线:%s
MACD中，diff/dea/macd/ema12/ema26:%d/%d/%d/%d%d
MA中，m3/m5/m10/m15/m20/m30/m60:%d/%d/%d/%d/%d/%d/%d
BOLL中 上轨/下轨/中轨:%d/%d/%d
KDJ中: K值/D值/J值:%d/%d/%d

请按此JSON格式输出决策(无Markdown):
{{
 "symbol": "BTC_USDT|ETH_USDT|null",
 "action": "Open(Short|Long)|Close(Short|Long)|HOLD",
 "confidence": 0.0-1.0,
 "rationale": "理由<50字",
 "target": "止盈价/0",
 "stop": "止损价/0"
}}
要求:
1.仅输出JSON
2.输出一个中短期的策略
3.开仓必填止盈损
`

const SystemMessage = `你资深量化分析师。根据1d的K线及MACD/MA/BOLL/KDJ指标以及财报和公司分析行情，输出JSON交易决策。`

// InvokeJudgment ai荐股
func (s *sStockAiJudgment) InvokeJudgment(ctx context.Context, in *stockin.StockAiJudgmentAiJudgmentInp) {
	// 获取自己观察的股票
	list, err := service.StockSelfCode().GetAllSelfCode(ctx, in.Symbol)
	if err != nil {
		return
	}

	var aiList []*entity.StockSelfAi
	err = service.StockSelfAi().Model(ctx).WhereNot(dao.StockSelfAi.Columns().Status, -1).Scan(&aiList)
	if err != nil {
		return
	}
	if len(aiList) == 0 {
		err = gerror.New("获取可用ai为空")
	}
	//wg := sync.WaitGroup{}
	for _, stock := range list {
		stockCode := stock
		//wg.Add(1)
		//simple.SafeGo(gctx.New(), func(ctx context.Context) {
		//	wg.Done()

		var kLine *entity.EnterpriseHistoricalData
		_ = service.StockEnterpriseHistoricalData().Model(ctx).Where(dao.EnterpriseHistoricalData.Columns().Symbol, stockCode.Dm).OrderDesc(dao.EnterpriseHistoricalData.Columns().T).Scan(&kLine)

		var macd *entity.MacdData
		_ = service.StockMacdData().Model(ctx).Where(dao.MacdData.Columns().Symbol, stockCode.Dm).OrderDesc(dao.MacdData.Columns().T).Scan(&macd)

		var ma *entity.MaData
		_ = service.StockMaData().Model(ctx).Where(dao.MaData.Columns().Symbol, stockCode.Dm).OrderDesc(dao.MaData.Columns().T).Scan(&ma)

		var boll *entity.BollData
		_ = service.StockBollData().Model(ctx).Where(dao.BollData.Columns().Symbol, stockCode.Dm).OrderDesc(dao.BollData.Columns().T).Scan(&boll)

		var kdj *entity.KdjData
		_ = service.StockKdjData().Model(ctx).Where(dao.KdjData.Columns().Symbol, stockCode.Dm).OrderDesc(dao.KdjData.Columns().T).Scan(&kdj)

		// 查询所有指标
		indicatorStr := fmt.Sprintf(`
			股票代码[%s]股票名称[%s]
			[1D趋势]
			K线收盘价/前收盘价/今日成交量/今日成交额:%f/%f/%f/%f
			MACD中，diff/dea/macd/ema12/ema26:%f/%f/%f/%f/%f
			MA中，m3/m5/m10/m15/m20/m30/m60:%f/%f/%f/%f/%f/%f/%f
			BOLL中 上轨/下轨/中轨:%f/%f/%f
			KDJ中: K值/D值/J值:%f/%f/%f
			
			请按此JSON格式输出决策(无Markdown):
			{{
			 "symbol": "BTC_USDT|ETH_USDT|null",
			 "action": "Open(Short|Long)|Close(Short|Long)|HOLD",
			 "rationale": "给出一个短中期投资建议以及理由<100字",
			 "financialJudgment: "是否应该买入/卖出，true/false",
			 "target": "预计止盈价",
			 "stop": "预计止损价"
			}}
			要求:
			1.仅输出JSON
			2.输出一个中短期的策略
			3.开仓必填止盈损
			`, stockCode.Dm, stockCode.Mc,
			kLine.C, kLine.Pc, kLine.V, kLine.A,
			macd.Diff, macd.Dea, macd.Macd, macd.Ema12, macd.Ema26,
			ma.Ma3, ma.Ma5, ma.Ma10, ma.Ma15, ma.Ma20, ma.Ma30, ma.Ma60,
			boll.U, boll.D, boll.M,
			kdj.K, kdj.D, kdj.J,
		)

		scripts := make([]string, 0)
		scripts = append(scripts, SystemMessage)
		scripts = append(scripts, indicatorStr)

		aiModelList, _ := service.StockSelfAi().GetAllAi(ctx, "")
		for _, model := range aiModelList {
			service.StockSelfAi().InvokeAi(ctx, model, scripts)
		}
		//})
	}
	//wg.Wait()
	return
}

func (s *sStockAiJudgment) indicatorsJudgment(ctx context.Context, kLine *entity.EnterpriseHistoricalData, macd *entity.MacdData, ma *entity.MaData, kdj *entity.KdjData, boll *entity.MacdData) {

}
