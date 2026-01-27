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

const SystemMessage = `你资深量化分析师。根据1d的K线及MACD/MA/BOLL/KDJ指标以及财报和公司分析行情，输出JSON交易决策。`

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
	wg := sync.WaitGroup{}
	for _, stock := range list {
		stockCode := stock
		wg.Add(1)
		simple.SafeGo(gctx.New(), func(ctx context.Context) {
			wg.Done()

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
			K线收盘价:%f,前收盘价:%f,今日成交量:%f,今日成交额:%f;
			MACD中，diff:%f,dea:%f,macd:%f,ema12:%f,ema26:%f;
			MA中，m3:%f,m5:%f,m10:%f,m15:%f,m20:%f,m30:%f,m60:%f;
			BOLL中 上轨:%f,下轨:%f,中轨:%f;
			KDJ中: K值:%f,D值:%f,J值:%f,;
			资金流向明细:大单动向:%f,大单差分:%f,主买特大单成交额:%f,主卖特大单成交额:%f,主买大单成交额:%f,主卖大单成交额:%f,被动买特大单成交额:%f,被动卖特大单成交额:%f,主买特大单成交额增量:%f,主买大单成交额增量:%f,涨跌动因:%f,主买单总单数:%d,主卖单总单数:%d,主买特大单成交量:%d,成交笔数增量:%d,
			
			请按此JSON格式输出决策(无Markdown):
			{{
			 "indicatorsJudgment": "指标判断,给出一个短中长期投资建议以及理由<100字"
			 "indicatorsFlag": "仅根据当日数据指标判断是否买入 true/false",
			 "financialJudgment: "仅根据公司股票、财务、财报指标数据判断是否应该买入/卖出，给出一个短中长期投资建议以及理由<100字",
			 "financialFlag": "仅根据公司股票、财务、财报指标数据判断是否应该买入 true/false",
			 "comprehensiveJudgment": "综合指标，根据当前数据指标和当前公司的公司股票、财务、财报指标数据来给出一个综合的短中长投资建议理由<100字"
			 "comprehensiveFlag": "根据综合指标判断是否买入 true/false"
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

			// 财务指标
			var financialIndicator *entity.FinancialIndicators
			_ = service.StockFinancialIndicators().Model(ctx).Where(dao.FinancialIndicators.Columns().Symbol, stock.Dm).OrderDesc(dao.FinancialIndicators.Columns().Date).Scan(&financialIndicator)
			if financialIndicator == nil {
				return
			}
			var inCome *entity.IncomeStatement
			_ = service.StockIncomeStatement().Model(ctx).Where(dao.IncomeStatement.Columns().Symbol, stock.Dm).OrderDesc(dao.IncomeStatement.Columns().Plrq).Scan(&inCome)
			if inCome == nil {
				return
			}
			var quarterlyProfit *entity.QuarterlyProfit
			_ = service.StockQuarterlyProfit().Model(ctx).Where(dao.QuarterlyProfit.Columns().Symbol, stock.Dm).OrderDesc(dao.QuarterlyProfit.Columns().Date).Scan(&quarterlyProfit)
			if quarterlyProfit == nil {
				return
			}
			var shareholderChange *entity.ShareholderChange
			_ = service.StockShareholderChange().Model(ctx).Where(dao.ShareholderChange.Columns().Symbol, stock.Dm).OrderDesc(dao.ShareholderChange.Columns().Jzrq).Scan(&shareholderChange)
			if shareholderChange == nil {
				return
			}
			// 查询所有指标
			financialStr := fmt.Sprintf(`
			股票代码[%s]股票名称[%s]
			[公司股票、财务、财报指标数据 时间:%s]
			摊薄每股收益(元):%f,加权每股收益(元):%f,净利润增长率:%f,流动比率:%f,速动比率:%f,应收账款周转率(次):%f,应收账款周转天数(天):%f,存货周转率(次):%f;
			主营业务收入增长率:%f,总资产周转率(次):%f,主营业务利润率:%f,经营现金净流量与净利润的比率:%f,扣除非经常性损益后的净利润(元):%f,扣除非经常性损益后的每股收益(元):%f;
			资产负债率:%f营业利润率:%f销售净利率:%f扣除非经常性损益后的净利润(元):%f加权净资产收益率:%f;

			[公司利润点 时间:%s-%s]
			营业收入:%f,营业总收入:%f,净利润:%f,归属于母公司所有者的净利润:%f,净利润(扣除非经常性损益后):%f,营业利润:%f,毛利率:%f,营业利润率:%f,投资收益:%f,公允价值变动收益:%f,资产减值损失:%f;
			
			[公司季度利润数据 时间:%s 会计时间:%d]
			基本每股收益(元/股): %f

			[公司股东变化数 截至时间: %s 公告时间:%s]
			股东户数:%d,比上期变化百分比:%f;
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

			// 输出格式
			outputFormat := fmt.Sprintf(`
			请按此JSON,仅仅输出{}内容,格式输出决策(无Markdown):
			{{
			 "indicatorsJudgment": "指标判断,给出一个短中长期投资建议以及理由<100字"
			 "indicatorsFlag": "仅根据当日数据指标判断是否买入 true/false",
			 "financialJudgment: "仅根据公司股票、财务、财报指标数据判断是否应该买入/卖出，给出一个短中长期投资建议以及理由<100字",
			 "financialFlag": "仅根据公司股票、财务、财报指标数据判断是否应该买入 true/false",
			 "comprehensiveJudgment": "综合指标，根据当前数据指标和当前公司的公司股票、财务、财报指标数据来给出一个综合的短中长投资建议理由<100字"
			 "comprehensiveFlag": "根据综合指标判断是否买入 true/false"
			 "target": "根据综合指标，买入必须确定一个预计止盈价",
			 "stop": "根据综合指标，买入必须确定一个预计止损价"
			}}
			要求:
			1.仅输出JSON
			2.输出一个中短期的策略
			3.开仓必填止盈损
		`)

			scripts := make([]string, 0)
			scripts = append(scripts, SystemMessage)
			scripts = append(scripts, indicatorStr) // 财报指标
			scripts = append(scripts, financialStr) // 财务指标
			scripts = append(scripts, outputFormat) // 输出格式

			aiModelList, _ := service.StockSelfAi().GetAllAi(ctx, "")
			for _, model := range aiModelList {
				now := time.Now()
				nt := now.Format("2006-01-02")
				flag, _ := service.StockAiJudgment().Model(ctx).Where(dao.StockAiJudgment.Columns().Symbol, stock.Dm).Where(dao.StockAiJudgment.Columns().T, nt).Where(dao.StockAiJudgment.Columns().AiId, model.Id).Exist()
				if flag {
					continue
				}
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
					continue
				}
				resMap["t"] = now.Format("2006-01-02") // 固定格
				resMap["aiId"] = model.Id
				resMap["aiName"] = model.Name
				resMap["judgmentIndicatorsScript"] = indicatorStr
				resMap["judgmentFinancialScript"] = financialStr
				resMap["symbol"] = stockCode.Dm
				resMap["mc"] = stockCode.Mc
				service.StockAiJudgment().Model(ctx).Insert(resMap)
			}
		})
	}
	wg.Wait()
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
		// 查询所有指标
		indicatorStr := fmt.Sprintf(`
			股票代码[%s]股票名称[%s]
			[1D趋势]
			K线收盘价:%f,前收盘价:%f,今日成交量:%f,今日成交额:%f;
			MACD中，diff:%f,dea:%f,macd:%f,ema12:%f,ema26:%f;
			MA中，m3:%f,m5:%f,m10:%f,m15:%f,m20:%f,m30:%f,m60:%f;
			BOLL中 上轨:%f,下轨:%f,中轨:%f;
			KDJ中: K值:%f,D值:%f,J值:%f,;
			资金流向明细:大单动向:%f,大单差分:%f,主买特大单成交额:%f,主卖特大单成交额:%f,主买大单成交额:%f,主卖大单成交额:%f,被动买特大单成交额:%f,被动卖特大单成交额:%f,主买特大单成交额增量:%f,主买大单成交额增量:%f,涨跌动因:%f,主买单总单数:%d,主卖单总单数:%d,主买特大单成交量:%d,成交笔数增量:%d,
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
			摊薄每股收益(元):%f,加权每股收益(元):%f,净利润增长率:%f,流动比率:%f,速动比率:%f,应收账款周转率(次):%f,应收账款周转天数(天):%f,存货周转率(次):%f;
			主营业务收入增长率:%f,总资产周转率(次):%f,主营业务利润率:%f,经营现金净流量与净利润的比率:%f,扣除非经常性损益后的净利润(元):%f,扣除非经常性损益后的每股收益(元):%f;
			资产负债率:%f营业利润率:%f销售净利率:%f扣除非经常性损益后的净利润(元):%f加权净资产收益率:%f;

			[公司利润点 时间:%s-%s]
			营业收入:%f,营业总收入:%f,净利润:%f,归属于母公司所有者的净利润:%f,净利润(扣除非经常性损益后):%f,营业利润:%f,毛利率:%f,营业利润率:%f,投资收益:%f,公允价值变动收益:%f,资产减值损失:%f;
			
			[公司季度利润数据 时间:%s 会计时间:%d]
			基本每股收益(元/股): %f

			[公司股东变化数 截至时间: %s 公告时间:%s]
			股东户数:%d,比上期变化百分比:%f;
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
