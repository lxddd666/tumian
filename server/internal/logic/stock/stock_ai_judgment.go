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
	"github.com/gogf/gf/v2/text/gstr"
	"github.com/gogf/gf/v2/util/gconv"
	"hotgo/internal/dao"
	"hotgo/internal/library/hgorm/handler"
	"hotgo/internal/model/entity"
	"hotgo/internal/model/input/stockin"
	"hotgo/internal/service"
	"hotgo/utility/convert"
	"hotgo/utility/excel"
	"hotgo/utility/simple"
	"sync"
	"time"
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

const SystemMessage = `你资深量化分析师。根据1d的K线及各种股票指标以及财报、公司估值、财报指标、历史股票支撑压力位、估值等分析行情等，输出JSON交易决策。`

// AiJudgmentComprehensiveData 综合
func (s *sStockAiJudgment) AiJudgmentComprehensiveData(ctx context.Context, in *stockin.StockAiJudgmentAiJudgmentInp) {
	// 获取自己观察的股票
	list, err := service.StockSelfCode().GetAllSelfCode(ctx, in.Symbol)
	if err != nil {
		return
	}

	var aiList []*entity.StockSelfAi
	err = service.StockSelfAi().Model(ctx).Scan(&aiList)
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
		scripts := make([]string, 0)
		scripts = append(scripts, SystemMessage)

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

		var rsi *entity.RsiData
		_ = service.StockRsiData().Model(ctx).Where(dao.RsiData.Columns().Symbol, stockCode.Dm).OrderDesc(dao.RsiData.Columns().T).Scan(&rsi)

		var cci *entity.CciData
		_ = service.StockCciData().Model(ctx).Where(dao.CciData.Columns().Symbol, stockCode.Dm).OrderDesc(dao.RsiData.Columns().T).Scan(&cci)

		var Stoch *entity.SlowStochasticData
		_ = service.StockSlowStochasticData().Model(ctx).Where(dao.SlowStochasticData.Columns().Symbol, stockCode.Dm).OrderDesc(dao.RsiData.Columns().T).Scan(&Stoch)

		var mom *entity.MomData
		_ = service.StockMomData().Model(ctx).Where(dao.SlowStochasticData.Columns().Symbol, stockCode.Dm).OrderDesc(dao.RsiData.Columns().T).Scan(&mom)

		var wmsr *entity.WilliamsData
		_ = service.StockWilliamsData().Model(ctx).Where(dao.SlowStochasticData.Columns().Symbol, stockCode.Dm).OrderDesc(dao.RsiData.Columns().T).Scan(&wmsr)

		var kst *entity.KstData
		_ = service.StockKstData().Model(ctx).Where(dao.SlowStochasticData.Columns().Symbol, stockCode.Dm).OrderDesc(dao.RsiData.Columns().T).Scan(&kst)

		var fastk *entity.FastkData
		_ = service.StockFastkData().Model(ctx).Where(dao.SlowStochasticData.Columns().Symbol, stockCode.Dm).OrderDesc(dao.RsiData.Columns().T).Scan(&fastk)
		// 查询所有指标
		indicatorStr := fmt.Sprintf(`
			股票代码[%s]股票名称[%s]
			[1D趋势]
			K线收盘价:%.3f,前收盘价:%.3f,今日成交量:%.3f,今日成交额:%.3f;
			MACD(12,26,9)中，diff:%.3f,dea:%.3f,macd:%.3f,ema12:%.3f,ema26:%.3f;
			MA中，m3:%.3f,m5:%.3f,m10:%.3f,m15:%.3f,m20:%.3f,m30:%.3f,m60:%.3f;
			BOLL中 上轨:%.3f,下轨:%.3f,中轨:%.3f;
			KDJ中: K值:%.3f,D值:%.3f,J值:%.3f,;
			Rsi(6,14,24)中: rsi-6值: %.3f, rsi-14值: %.3f, rsi-24值: %.3f;
			资金流向明细:大单动向:%.3f,大单差分:%.3f,主买特大单成交额:%.3f,主卖特大单成交额:%.3f,主买大单成交额:%.3f,主卖大单成交额:%.3f,被动买特大单成交额:%.3f,被动卖特大单成交额:%.3f,主买特大单成交额增量:%.3f,主买大单成交额增量:%.3f,涨跌动因:%.3f,主买单总单数:%d,主卖单总单数:%d,主买特大单成交量:%d,成交笔数增量:%d,
			`, stockCode.Dm, stockCode.Mc,
			kLine.C, kLine.Pc, kLine.V, kLine.A,
			macd.Diff, macd.Dea, macd.Macd, macd.Ema12, macd.Ema26,
			ma.Ma3, ma.Ma5, ma.Ma10, ma.Ma15, ma.Ma20, ma.Ma30, ma.Ma60,
			boll.U, boll.D, boll.M,
			kdj.K, kdj.D, kdj.J,
			rsi.Rsi6, rsi.Rsi14, rsi.Rsi24,
			flowOfFunds.Dddx, flowOfFunds.Ddcf, flowOfFunds.Zmbtdcje, flowOfFunds.Zmstdcje, flowOfFunds.Zmbddcje, flowOfFunds.Zmsddcje, flowOfFunds.Bdmbtdcje, flowOfFunds.Bdmstdcje, flowOfFunds.Zmbtdcjzl, flowOfFunds.Zmbddcjzl, flowOfFunds.Zddy, flowOfFunds.Zmbzds, flowOfFunds.Zmszds, flowOfFunds.Zmbtdcjl, flowOfFunds.Cjbszl,
		)

		if cci != nil {
			indicatorStr += fmt.Sprintf("\n CCI 20中，CCI:%.3f;", cci.Cci)
		}
		if Stoch != nil {
			indicatorStr += fmt.Sprintf("\n Slow Stochastic 14 3 3中， %%K线:%.3f, %%D线: %.3f;", Stoch.K, Stoch.D)
		}

		if mom != nil {
			indicatorStr += fmt.Sprintf("\n Mom 10 close中，Mom: %.3f;", mom.Mon)
		}

		if wmsr != nil {
			indicatorStr += fmt.Sprintf("\n Williams %%R 14中，%%R: %.3f;", wmsr.R)
		}

		if kst != nil {
			indicatorStr += fmt.Sprintf("\n 短期KST 10 15 20 30 10 10 10 159中，KST: %.3f, Signal: %.3f;", kst.Kst, kst.Signal)
		}

		if fastk != nil {
			indicatorStr += fmt.Sprintf("\n Fast Stochastic 14 1 3中, %%K:%.3f,%%D: %.3f;", fastk.K, fastk.D)
		}
		// 财务指标
		var financialIndicator *entity.FinancialIndicators
		_ = service.StockFinancialIndicators().Model(ctx).Where(dao.FinancialIndicators.Columns().Symbol, stockCode.Dm).OrderDesc(dao.FinancialIndicators.Columns().Date).Scan(&financialIndicator)
		if financialIndicator == nil {
			return
		}
		var inCome *entity.IncomeStatement
		_ = service.StockIncomeStatement().Model(ctx).Where(dao.IncomeStatement.Columns().Symbol, stockCode.Dm).OrderDesc(dao.IncomeStatement.Columns().Plrq).Scan(&inCome)
		if inCome == nil {
			return
		}
		var quarterlyProfit *entity.QuarterlyProfit
		_ = service.StockQuarterlyProfit().Model(ctx).Where(dao.QuarterlyProfit.Columns().Symbol, stockCode.Dm).OrderDesc(dao.QuarterlyProfit.Columns().Date).Scan(&quarterlyProfit)
		if quarterlyProfit == nil {
			return
		}
		var shareholderChange *entity.ShareholderChange
		_ = service.StockShareholderChange().Model(ctx).Where(dao.ShareholderChange.Columns().Symbol, stockCode.Dm).OrderDesc(dao.ShareholderChange.Columns().Jzrq).Scan(&shareholderChange)
		if shareholderChange == nil {
			return
		}
		// 查询所有指标
		financialStr := fmt.Sprintf(`
			股票代码[%s]股票名称[%s]
			[公司股票、财务、财报指标数据 时间:%s]
			摊薄每股收益(元):%.3f,加权每股收益(元):%.3f,净利润增长率:%.3f,流动比率:%.3f,速动比率:%.3f,应收账款周转率(次):%.3f,应收账款周转天数(天):%.3f,存货周转率(次):%.3f;
			主营业务收入增长率:%.3f,总资产周转率(次):%.3f,主营业务利润率:%.3f,经营现金净流量与净利润的比率:%.3f,扣除非经常性损益后的净利润(元):%.3f,扣除非经常性损益后的每股收益(元):%.3f;
			资产负债率:%.3f营业利润率:%.3f销售净利率:%.3f扣除非经常性损益后的净利润(元):%.3f加权净资产收益率:%.3f;

			[公司利润点 时间:%s-%s]
			营业收入:%.3f,营业总收入:%.3f,净利润:%.3f,归属于母公司所有者的净利润:%.3f,净利润(扣除非经常性损益后):%.3f,营业利润:%.3f,毛利率:%.3f,营业利润率:%.3f,投资收益:%.3f,公允价值变动收益:%.3f,资产减值损失:%.3f;
			
			[公司季度利润数据 时间:%s 会计时间:%d]
			基本每股收益(元/股): %.3f

			[公司股东变化数 截至时间: %s 公告时间:%s]
			股东户数:%d,比上期变化百分比:%.3f;
			`, stockCode.Dm, stockCode.Mc, financialIndicator.Date,
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

		// 财报评分
		var financialMttScript string

		var financialIndicatorsMtt *entity.FinancialIndicatorsMtt
		_ = dao.FinancialIndicatorsMtt.Ctx(ctx).Where(dao.FinancialIndicatorsMtt.Columns().Symbol, stock.Dm).Scan(&financialIndicatorsMtt)
		if financialIndicatorsMtt != nil {
			var profitMttScript string       // 盈利能力
			var solvencyScript string        // 偿债能力
			var profitQualityScript string   // 盈利质量
			var dividendReturnsScript string // 股息回报
			if financialIndicatorsMtt.RevenueGrowthTtm != 0 || financialIndicatorsMtt.GrossProfitMarginTtm != 0 || financialIndicatorsMtt.RoaTtm != 0 || financialIndicatorsMtt.NetProfitMarginTtm != 0 {
				profitMttScript = "[公司盈利能力] \n"
				if financialIndicatorsMtt.RevenueGrowthTtm != 0 {
					profitMttScript += fmt.Sprintf("营收增长率TTM(%%):%.3f,营收增长率TTM同比增长(%%):%.3f;", financialIndicatorsMtt.RevenueGrowthTtm, financialIndicatorsMtt.RevenueGrowthTtmYoy)
				}
				if financialIndicatorsMtt.GrossProfitMarginTtm != 0 {
					profitMttScript += fmt.Sprintf("销售毛利率TTM(%%):%.3f,销售毛利率TTM同比增长(%%):%.3f;", financialIndicatorsMtt.GrossProfitMarginTtm, financialIndicatorsMtt.GrossProfitMarginTtmYoy)
				}
				if financialIndicatorsMtt.RoaTtm != 0 {
					profitMttScript += fmt.Sprintf("净资产收益率TTM(%%):%.3f,净资产收益率TTM同比增长(%%):%.3f;", financialIndicatorsMtt.RoaTtm, financialIndicatorsMtt.RoaTtmYoy)
				}
				if financialIndicatorsMtt.NetProfitMarginTtm != 0 {
					profitMttScript += fmt.Sprintf("销售净利率TTM(%%):%.3f,销售净利率TTM同比增长(%%):%.3f;", financialIndicatorsMtt.NetProfitMarginTtm, financialIndicatorsMtt.NetProfitMarginTtmYoy)
				}
			}
			if financialIndicatorsMtt.CurrentRatio != 0 || financialIndicatorsMtt.DebtToAssetRatio != 0 || financialIndicatorsMtt.QuickRatio != 0 {
				solvencyScript = "[公司偿债能力]\n"
				if financialIndicatorsMtt.CurrentRatio != 0 {
					solvencyScript += fmt.Sprintf("流动比率(%%):%.3f,流动比率同比增长(%%):%.3f;", financialIndicatorsMtt.CurrentRatio, financialIndicatorsMtt.CurrentRatioYoy)
				}
				if financialIndicatorsMtt.DebtToAssetRatio != 0 {
					solvencyScript += fmt.Sprintf("资产负债率(%%):%.3f,资产负债率同比增长(%%):%.3f;", financialIndicatorsMtt.DebtToAssetRatio, financialIndicatorsMtt.DebtToAssetRatioYoy)
				}
				if financialIndicatorsMtt.QuickRatio != 0 {
					solvencyScript += fmt.Sprintf("速动比率(%%):%.3f,速动比率同比增长(%%):%.3f;", financialIndicatorsMtt.QuickRatio, financialIndicatorsMtt.QuickRatioYoy)
				}
				if financialIndicatorsMtt.InterestCoverage != 0 {
					solvencyScript += fmt.Sprintf("利息保障倍数:%.3f,利息保障倍数同比增长(%%):%.3f;", financialIndicatorsMtt.InterestCoverage, financialIndicatorsMtt.InterestCoverageYoy)
				}
			}
			if financialIndicatorsMtt.NetProfitCashContent != 0 || financialIndicatorsMtt.InventoryTurnoverDays != 0 || financialIndicatorsMtt.ReceivableTurnoverDays != 0 {
				profitQualityScript = "[公司盈利质量]\n"
				if financialIndicatorsMtt.NetProfitCashContent != 0 {
					profitQualityScript += fmt.Sprintf("净利润现金含量(%%):%.3f,净利润现金含量同比增长(%%):%.3f;", financialIndicatorsMtt.NetProfitCashContent, financialIndicatorsMtt.NetProfitCashContentYoy)
				}
				if financialIndicatorsMtt.InventoryTurnoverDays != 0 {
					profitQualityScript += fmt.Sprintf("存货周转天数(天/次):%.3f,存货周转天数同比增长(%%):%.3f;", financialIndicatorsMtt.InventoryTurnoverDays, financialIndicatorsMtt.InventoryTurnoverDaysYoy)
				}
				if financialIndicatorsMtt.ReceivableTurnoverDays != 0 {
					profitQualityScript += fmt.Sprintf("应收周转天数(天/次):%.3f,应收周转天数同比增长(%%):%.3f;", financialIndicatorsMtt.ReceivableTurnoverDays, financialIndicatorsMtt.ReceivableTurnoverDaysYoy)
				}
			}
			if financialIndicatorsMtt.DividendGrowthTtm != 0 || financialIndicatorsMtt.DividendPayoutRatioTtm != 0 || financialIndicatorsMtt.DividendYieldTtm != 0 {
				dividendReturnsScript = "[公司股息回报]\n"
				if financialIndicatorsMtt.DividendGrowthTtm != 0 {
					dividendReturnsScript += fmt.Sprintf("股息增长率TTM(%%):%.3f,股息增长率TTM同比增长(%%):%.3f;", financialIndicatorsMtt.DividendGrowthTtm, financialIndicatorsMtt.DividendGrowthTtmYoy)
				}
				if financialIndicatorsMtt.DividendPayoutRatioTtm != 0 {
					dividendReturnsScript += fmt.Sprintf("股息支付率TTM(%%):%.3f,股息支付率TTM同比增长(%%):%.3f;", financialIndicatorsMtt.DividendPayoutRatioTtm, financialIndicatorsMtt.DividendPayoutRatioTtmYoy)
				}
				if financialIndicatorsMtt.DividendYieldTtm != 0 {
					dividendReturnsScript += fmt.Sprintf("股息率TTM(%%):%.3f,股息率TTM同比增长(%%):%.3f;", financialIndicatorsMtt.DividendYieldTtm, financialIndicatorsMtt.DividendYieldTtmYoy)
				}
			}
			financialMttScript = profitMttScript + "\n" + solvencyScript + "\n" + profitQualityScript + "\n" + dividendReturnsScript
			if financialMttScript != "" {
				scripts = append(scripts, financialMttScript)
			}
		}

		// 支撑位 压力位
		var support *entity.StockSupportResistance
		_ = dao.StockSupportResistance.Ctx(ctx).Where(dao.StockSupportResistance.Columns().Symbol, stock.Dm).Scan(&support)

		var suppScript string
		if support != nil {

			if support.Zc != 0 || support.Yl != 0 {
				if support.Price != 0 {
					suppScript = fmt.Sprintf("股票如今价格:%.3f，", kLine.C)

					if support.Zc != 0 {
						suppScript += fmt.Sprintf("股票价格支撑位:%.3f，", support.Zc)
					}
					if support.Yl != 0 {
						suppScript += fmt.Sprintf("股票价格压力位:%.3f，", support.Yl)
					}
				}
			}
			scripts = append(scripts, suppScript)
		}

		// 估值
		var valuationMtt *entity.ValuationIndicators
		_ = dao.ValuationIndicators.Ctx(ctx).Where(dao.ValuationIndicators.Columns().Symbol, stock.Dm).Scan(&valuationMtt)
		var valuationScript string
		if valuationMtt != nil {
			valuationScript = "[股票相对估值]\n"
			if valuationMtt.PeTtm != 0 {
				valuationScript += fmt.Sprintf("现市盈率(TTM):%.3f, 近三年市盈率30分位置:%.3f, 近三年市盈率分70位置:%.3f \n ", valuationMtt.PeTtm, valuationMtt.PePercentile30, valuationMtt.PePercentile70)
			}
			if valuationMtt.Pb != 0 {
				valuationScript += fmt.Sprintf("现市净率(TTM):%.3f, 近三年市净率30分位置:%.3f, 近三年市净率分70位置:%.3f \n ", valuationMtt.Pb, valuationMtt.PbPercentile30, valuationMtt.PbPercentile70)
			}
			scripts = append(scripts, valuationScript) // 财报指标

		}

		// 输出格式
		outputFormat := fmt.Sprintf(`
			请按此JSON,仅仅输出{}内容,格式输出决策(无Markdown):
			{{
			 "indicatorsJudgment": "指标MACD、MA、BOLL、KDJ、Rsi、CCI、Slow Stochastic、MOM、WMSR、KST、FASTK,给出一个短中长期投资建议以及理由<100字"
			 "indicatorsFlag": "仅根据当日数据指标MACD、MA、BOLL、KDJ、Rsi、CCI、Slow Stochastic、MOM、WMSR、KST、FASTK判断是否买入 1买入 0不买(填1/0)",
			 "financialJudgment: "仅根据公司股票、财务、财报指标数据判断是否应该买入/卖出，给出一个短中长期投资建议以及理由<100字",
			 "financialFlag": "仅根据公司股票、财务、财报指标数据判断是否应该买入 1买入 0不买(填1/0)",
			 "financialRatingJudgment": "仅靠公司的盈利能力、偿债能力、盈利质量、股息回报，给出一个短中长期投资建议以及理由<100字",
			 "financialRatingFlag": "仅靠公司的盈利能力、偿债能力、盈利质量、股息回报，判断股票是否应该买入，1买入 0不买(填1/0)",
			 "marketAnalysisJudgment": "根据资金流向明细判断股票是否应该买入/卖出，给出一个投资简易以及理由<50字",
			 "marketAnalysisFlag": "根据资金流向明细判断股票是否应该买入/卖出，1买入 0不买(填1/0)",
			 "valueAssessmentJudgment: "根据股票价格支撑位、压力位以及股票相对估值，市盈率TTM、市净率TTM，给出一个投资建议以及理由<50字",
			 "valueAssessmentFlag: "根据股票价格支撑位、压力位以及股票相对估值，市盈率TTM、市净率TTM，1买入 0不买(填1/0)",
			 "comprehensiveJudgment": "综合指标，根据上面所诉内容，即当前股票数据指标和当前公司的公司股票、财务、财务评分（盈利能力、偿债能力、盈利质量、股息回报）、财报指标数据、股票压力位支撑位、股票相对估值、资金流动方向、来给出一个综合的短中投资建议理由<150字"
			 "comprehensiveFlag": "根据综合指标判断是否买入 1买入 0不买(填1/0)
			 "target": "根据综合指标，买入必须确定一个预计止盈价",
			 "stop": "根据综合指标，买入必须确定一个预计止损价"
			}}
			要求:
			1.仅输出JSON
			2.输出一个中短期的策略
			3.开仓必填止盈损
		`)

		scripts = append(scripts, indicatorStr) // 财报指标
		scripts = append(scripts, financialStr) // 财务指标
		scripts = append(scripts, outputFormat) // 输出格式

		aiModelList, _ := service.StockSelfAi().GetAllAi(ctx, "")
		wg := sync.WaitGroup{}
		for _, aimodel := range aiModelList {
			model := aimodel
			now := GetRecentWeekdayClear()
			//nt := now.Format("2006-01-02")
			flag, _ := service.StockAiJudgment().Model(ctx).Where(dao.StockAiJudgment.Columns().Symbol, stockCode.Dm).Where(dao.StockAiJudgment.Columns().T, now).Where(dao.StockAiJudgment.Columns().AiId, model.Id).Exist()
			if flag {
				return
			}
			wg.Add(1)
			simple.SafeGo(gctx.New(), func(ctx context.Context) {
				defer wg.Done()
				res, gerr := service.StockSelfAi().InvokeAi(ctx, model, scripts)
				if gerr != nil {
					return
				}

				res = gstr.Replace(res, "`", "")
				res = gstr.Replace(res, "json", "")

				var resMap map[string]interface{}
				err = gconv.Scan(res, &resMap)
				if err != nil {
					return
				} // 当前时间
				if resMap == nil {
					return
				}
				resMap["t"] = now // 固定格
				resMap["aiId"] = model.Id
				resMap["aiName"] = model.Name
				resMap["judgmentIndicatorsScript"] = indicatorStr
				resMap["judgmentFinancialScript"] = financialStr
				resMap["financialMttScript"] = financialMttScript
				resMap["supportScript"] = suppScript
				resMap["valuationMttScript"] = valuationScript
				resMap["symbol"] = stockCode.Dm
				resMap["mc"] = stockCode.Mc
				if resMap["comprehensiveFlag"] != nil && resMap["comprehensiveFlag"] == true {
					resMap["comprehensiveFlag"] = 1
				}
				if resMap["comprehensiveFlag"] != nil && resMap["comprehensiveFlag"] == false {
					resMap["comprehensiveFlag"] = 0
				}
				if resMap["indicatorsFlag"] != nil && resMap["indicatorsFlag"] == true {
					resMap["indicatorsFlag"] = 1
				}
				if resMap["indicatorsFlag"] != nil && resMap["indicatorsFlag"] == false {
					resMap["indicatorsFlag"] = 0
				}
				if resMap["financialFlag"] != nil && resMap["financialFlag"] == true {
					resMap["financialFlag"] = 1
				}
				if resMap["financialFlag"] != nil && resMap["financialFlag"] == false {
					resMap["financialFlag"] = 0
				}
				if resMap["target"] != nil && resMap["target"] == "" {
					resMap["target"] = 0
				}
				if resMap["stop"] != nil && resMap["stop"] == "" {
					resMap["stop"] = 0
				}
				_, err = service.StockAiJudgment().Model(ctx).Insert(resMap)
				if err != nil {
					fmt.Println(err)
				}
			})
		}
		wg.Wait()
		//})
	}
	//wg.Wait()
	return
}

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

		var rsi *entity.RsiData
		_ = service.StockRsiData().Model(ctx).Where(dao.RsiData.Columns().Symbol, stockCode.Dm).OrderDesc(dao.RsiData.Columns().T).Scan(&rsi)

		// 查询所有指标
		indicatorStr := fmt.Sprintf(`
			股票代码[%s]股票名称[%s]
			[1D趋势]
			K线收盘价:%.3f,前收盘价:%.3f,今日成交量:%.3f,今日成交额:%.3f;
			MACD中，diff:%.3f,dea:%.3f,macd:%.3f,ema12:%.3f,ema26:%.3f;
			MA中，m3:%.3f,m5:%.3f,m10:%.3f,m15:%.3f,m20:%.3f,m30:%.3f,m60:%.3f;
			BOLL中 上轨:%.3f,下轨:%.3f,中轨:%.3f;
			KDJ中: K值:%.3f,D值:%.3f,J值:%.3f,;
			资金流向明细:大单动向:%.3f,大单差分:%.3f,主买特大单成交额:%.3f,主卖特大单成交额:%.3f,主买大单成交额:%.3f,主卖大单成交额:%.3f,被动买特大单成交额:%.3f,被动卖特大单成交额:%.3f,主买特大单成交额增量:%.3f,主买大单成交额增量:%.3f,涨跌动因:%.3f,主买单总单数:%d,主卖单总单数:%d,主买特大单成交量:%d,成交笔数增量:%d,
			请按此JSON格式输出决策(无Markdown):
			{{
			 "indicatorsJudgment": "指标判断,给出一个短中长期投资建议以及理由<100字"
			 "indicatorsFlag": "仅根据当日数据指标判断是否买入 true/false",
			 "target": "根据综合指标，买入必须确定一个预计止盈价",
			 "stop": "根据综合指标，买入必须确定一个预计止损价"
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
		scripts = append(scripts, indicatorStr) // 财报指标

		aiModelList, _ := service.StockSelfAi().GetAllAi(ctx, "")
		for _, model := range aiModelList {
			res, gerr := service.StockSelfAi().InvokeAi(ctx, model, scripts)
			if gerr != nil {
				return
			}
			var resMap map[string]interface{}
			err = gconv.Scan(res, &resMap)
			if err != nil {
				return
			} // 当前时间
			now := time.Now()
			resMap["t"] = now.Format("2006-01-02") // 固定格
			resMap["aiId"] = model.Id
			resMap["aiName"] = model.Name
			resMap["judgmentIndicatorsScript"] = indicatorStr
			resMap["symbol"] = stockCode.Dm
			resMap["mc"] = stockCode.Mc
			service.StockAiJudgment().Model(ctx).Insert(resMap)
		}
		//})
	}
	//wg.Wait()
	return
}

// AiJudgmentFinancialData 财报数据ai鉴定
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

	for _, stockCode := range list {
		stock := stockCode
		// 财务指标
		var financialIndicator *entity.FinancialIndicators
		_ = service.StockFinancialIndicators().Model(ctx).Where(dao.FinancialIndicators.Columns().Symbol, stock.Dm).OrderDesc(dao.FinancialIndicators.Columns().Date).Scan(&financialIndicator)

		var inCome *entity.IncomeStatement
		_ = service.StockIncomeStatement().Model(ctx).Where(dao.IncomeStatement.Columns().Symbol, stock.Dm).OrderDesc(dao.IncomeStatement.Columns().Plrq).Scan(&inCome)

		var quarterlyProfit *entity.QuarterlyProfit
		_ = service.StockQuarterlyProfit().Model(ctx).Where(dao.QuarterlyProfit.Columns().Symbol, stock.Dm).OrderDesc(dao.QuarterlyProfit.Columns().Date).Scan(&quarterlyProfit)

		var shareholderChange *entity.ShareholderChange
		_ = service.StockShareholderChange().Model(ctx).Where(dao.ShareholderChange.Columns().Symbol, stock.Dm).OrderDesc(dao.ShareholderChange.Columns().Jzrq).Scan(&shareholderChange)

		// 查询所有指标
		indicatorStr := fmt.Sprintf(`
			股票代码[%s]股票名称[%s]
			[公司股票、财务、财报指标数据 时间:%s]
			摊薄每股收益(元):%.3f,加权每股收益(元):%.3f,净利润增长率:%.3f,流动比率:%.3f,速动比率:%.3f,应收账款周转率(次):%.3f,应收账款周转天数(天):%.3f,存货周转率(次):%.3f;
			主营业务收入增长率:%.3f,总资产周转率(次):%.3f,主营业务利润率:%.3f,经营现金净流量与净利润的比率:%.3f,扣除非经常性损益后的净利润(元):%.3f,扣除非经常性损益后的每股收益(元):%.3f;
			资产负债率:%.3f营业利润率:%.3f销售净利率:%.3f扣除非经常性损益后的净利润(元):%.3f加权净资产收益率:%.3f;

			[公司利润点 时间:%s-%s]
			营业收入:%.3f,营业总收入:%.3f,净利润:%.3f,归属于母公司所有者的净利润:%.3f,净利润(扣除非经常性损益后):%.3f,营业利润:%.3f,毛利率:%.3f,营业利润率:%.3f,投资收益:%.3f,公允价值变动收益:%.3f,资产减值损失:%.3f;
			
			[公司季度利润数据 时间:%s 会计时间:%d]
			基本每股收益(元/股): %.3f

			[公司股东变化数 截至时间: %s 公告时间:%s]
			股东户数:%d,比上期变化百分比:%.3f;
			请按此JSON格式输出决策(无Markdown):
			{{
			 "financialJudgment: "仅根据公司股票、财务、财报指标数据判断是否应该买入/卖出，给出一个短中长期投资建议以及理由<100字",
			 "financialFlag": "仅根据公司股票、财务、财报指标数据判断是否应该买入 true/false",
			 "target": "根据综合指标，买入必须确定一个预计止盈价",
			 "stop": "根据综合指标，买入必须确定一个预计止损价"
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
			res, gerr := service.StockSelfAi().InvokeAi(ctx, model, scripts)
			if gerr != nil {
				return
			}
			var resMap map[string]interface{}
			err = gconv.Scan(res, &resMap)
			if err != nil {
				return
			} // 当前时间
			now := time.Now()
			resMap["t"] = now.Format("2006-01-02") // 固定格
			resMap["aiId"] = model.Id
			resMap["aiName"] = model.Name
			resMap["judgmentIndicatorsScript"] = indicatorStr
			resMap["symbol"] = stockCode.Dm
			resMap["mc"] = stockCode.Mc
			service.StockAiJudgment().Model(ctx).Insert(resMap)
		}
	}
}

func (s *sStockAiJudgment) Export(ctx context.Context) (err error) {
	now := GetRecentWeekdayClear()
	//type totalScoreModel struct {
	//	totalScore int
	//	symbol     string
	//}
	//totalScoreList := make([]totalScoreModel, 0)
	//err = dao.StockAiJudgment.Ctx(ctx).Fields("sum(comprehensive_flag) totalScore, symbol").Where(dao.StockAiJudgment.Columns().T, now).Group(dao.StockAiJudgment.Columns().Symbol).OrderDesc("totalScore").Scan(&totalScoreList)
	//if err != nil {
	//	return
	//}
	//scoreMap := make(map[string]int)
	//for _, l := range totalScoreList {
	//	scoreMap[l.symbol] = l.totalScore
	//}

	list := make([]stockin.StockAiJudgmentExportModel, 0)

	sqlinnerJoin := fmt.Sprintf(" (SELECT symbol, SUM(comprehensive_flag) as totalScore    FROM hg_stock_ai_judgment    WHERE t = '%s'    GROUP BY symbol) as score", now)
	err = dao.StockAiJudgment.Ctx(ctx).As("aj").Fields("aj.* ,score.totalScore").InnerJoin(sqlinnerJoin, "aj.symbol = score.symbol").Where("aj.t", now).OrderDesc("score.totalScore,aj.symbol").Scan(&list)
	if err != nil {
		return
	}

	tags, err := convert.GetEntityDescTags(stockin.StockAiJudgmentExportModel{})
	if err != nil {
		return
	}
	var (
		fileName  = "导出今天ai选股-" + GetRecentWeekdayClear()
		sheetName = fmt.Sprintf("导出%d条数据", len(list))
	)
	err = excel.ExportByStructs(ctx, tags, list, fileName, sheetName)
	return
}
