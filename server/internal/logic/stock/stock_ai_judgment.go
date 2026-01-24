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

// InvokeIndicatorsJudgment 指标判断股票
func (s *sStockAiJudgment) InvokeIndicatorsJudgment(ctx context.Context, in *stockin.StockAiJudgmentAiJudgmentInp) {
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
		return
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

		var flowOfFunds *entity.FlowOfFunds
		_ = service.StockFlowOfFunds().Model(ctx).Where(dao.FlowOfFunds.Columns().Symbol, stockCode.Dm).OrderDesc(dao.KdjData.Columns().T).Scan(&flowOfFunds)
		// 查询所有指标
		indicatorStr := fmt.Sprintf(`
			股票代码[%s]股票名称[%s]
			[1D趋势]
			K线收盘价/前收盘价/今日成交量/今日成交额:%f/%f/%f/%f
			MACD中，diff/dea/macd/ema12/ema26:%f/%f/%f/%f/%f
			MA中，m3/m5/m10/m15/m20/m30/m60:%f/%f/%f/%f/%f/%f/%f
			BOLL中 上轨/下轨/中轨:%f/%f/%f
			KDJ中: K值/D值/J值:%f/%f/%f
			资金流向明细:大单动向/大单差分/主买特大单成交额/主卖特大单成交额/主买大单成交额/主卖大单成交额/被动买特大单成交额/被动卖特大单成交额/主买特大单成交额增量/主买大单成交额增量/涨跌动因/主买单总单数/主卖单总单数/主买特大单成交量/成交笔数增量:%f/%f/%f/%f/%f/%f/%f/%f/%f/%f/%f/%d/%d/%d/%d
			
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
			flowOfFunds.Dddx, flowOfFunds.Ddcf, flowOfFunds.Zmbtdcje, flowOfFunds.Zmstdcje, flowOfFunds.Zmbddcje, flowOfFunds.Zmsddcje, flowOfFunds.Bdmbtdcje, flowOfFunds.Bdmstdcje, flowOfFunds.Zmbtdcjzl, flowOfFunds.Zmbddcjzl, flowOfFunds.Zddy, flowOfFunds.Zmbzds, flowOfFunds.Zmszds, flowOfFunds.Zmbtdcjl, flowOfFunds.Cjbszl,
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

func (s *sStockAiJudgment) AiJudgmentFinancialData(ctx context.Context, in *stockin.StockAiJudgmentAiJudgmentInp) {
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
		return
	}

	for _, stock := range list {

		// 财务指标
		var financialIndicator *entity.FinancialIndicators
		_ = service.StockFinancialIndicators().Model(ctx).Where(dao.FinancialIndicators.Columns().Symbol, stock.Dm).OrderDesc(dao.FinancialIndicators.Columns().Date).Scan(&financialIndicator)

		var inCome *entity.IncomeStatement
		_ = service.StockIncomeStatement().Model(ctx).Where(dao.IncomeStatement.Columns().Symbol, stock.Dm).OrderDesc(dao.IncomeStatement.Columns().Plrq).Scan(&financialIndicator)

		var quarterlyProfit *entity.QuarterlyProfit
		_ = service.StockQuarterlyProfit().Model(ctx).Where(dao.QuarterlyProfit.Columns().Symbol, stock.Dm).OrderDesc(dao.QuarterlyProfit.Columns().Date).Scan(&financialIndicator)

		var shareholderChange *entity.ShareholderChange
		// 查询所有指标
		indicatorStr := fmt.Sprintf(`
			股票代码[%s]股票名称[%s]
			[公司财务指标数据 时间:%s]
			摊薄每股收益(元)/加权每股收益(元)/净利润增长率/流动比率/速动比率/应收账款周转率(次)/应收账款周转天数(天)/存货周转率(次):%f/%f/%f/%f/%f/%f/%f/%f
			主营业务收入增长率/总资产周转率(次)/主营业务利润率/经营现金净流量与净利润的比率/扣除非经常性损益后的净利润(元)/扣除非经常性损益后的每股收益(元):%f/%f/%f/%f/%f/%f
			资产负债率/营业利润率/销售净利率/扣除非经常性损益后的净利润(元)/加权净资产收益率:%f/%f/%f/%f/%f

			[公司利润点 时间:%s-%s]
			营业收入/营业总收入/净利润/归属于母公司所有者的净利润/净利润(扣除非经常性损益后)/营业利润/毛利率/营业利润率/投资收益/公允价值变动收益/资产减值损失:%f/%f/%f/%f/%f/%f/%f/%f/%f/%f/%f
			
			[公司季度利润数据 时间:%s 会计时间:%d]
			基本每股收益(元/股): %f

			[公司股东变化数 截至时间: %s 公告时间:%s]
			股东户数/比上期变化百分比: %d/%f
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
			`, stock.Dm, stock.Mc, financialIndicator.Date,
			financialIndicator.Tbmg, financialIndicator.Jqmg, financialIndicator.Jlzz, financialIndicator.Ldbl, financialIndicator.Sdbl, financialIndicator.Yszz, financialIndicator.Yszzt, financialIndicator.Chzzl,
			financialIndicator.Zysr, financialIndicator.Zzzzl, financialIndicator.Zylr, financialIndicator.Jylrb, financialIndicator.Kflr, financialIndicator.Kfmg,
			financialIndicator.Zcfzl, financialIndicator.Yylr, financialIndicator.Xsjl, financialIndicator.Kflr, financialIndicator.Jqjz,

			inCome.Jzrq, inCome.Plrq,
			inCome.Yysr, inCome.Yyzsr, inCome.Jlr, inCome.Gsmgsyzzdjlr, inCome.Jlrhfcjcx, inCome.Yylr, inCome.GrossMargin, inCome.OperatingMargin, inCome.Tzsy, inCome.Gyjzbdsy, inCome.Zcjzss,

			quarterlyProfit.Date, quarterlyProfit.ReportQuarter,
			quarterlyProfit.Basege,

			shareholderChange.Jzrq, shareholderChange.AnnDate,
			shareholderChange.Gdhs, shareholderChange.Bh,
		)

		scripts := make([]string, 0)
		scripts = append(scripts, SystemMessage)
		scripts = append(scripts, indicatorStr)

		aiModelList, _ := service.StockSelfAi().GetAllAi(ctx, "")
		for _, model := range aiModelList {
			service.StockSelfAi().InvokeAi(ctx, model, scripts)
		}
	}
}
