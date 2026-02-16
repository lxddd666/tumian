package stock

import (
	"context"
	"fmt"
	"github.com/chromedp/chromedp"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/text/gstr"
	"github.com/gogf/gf/v2/util/gconv"
	"hotgo/internal/dao"
	"hotgo/internal/model/entity"
	"hotgo/internal/model/input/stockin"
	"hotgo/utility/simple"
	"sync"
	"time"
)

// BaiduFinanceCode 百度财经爬虫
func (s *sStockSelfCode) BaiduFinanceCode(ctx context.Context, in *stockin.SelfCodeIndicatorsApiInp) (err error) {

	// 获取全部股票
	var list []*entity.StockAllCode
	err = dao.StockAllCode.Ctx(ctx).Scan(&list)
	if err != nil {
		return
	}
	if len(list) == 0 {
		return
	}

	// 创建容量为3的协程池
	pool := make(chan struct{}, 100)

	// 创建WaitGroup等待所有任务完成
	var wg sync.WaitGroup

	// 错误收集
	var mu sync.Mutex
	var errs []error

	for _, stockCode := range list {
		stock := stockCode

		wg.Add(1)

		// 启动协程
		simple.SafeGo(ctx, func(ctx context.Context) {
			defer wg.Done()

			// 获取池中的令牌
			pool <- struct{}{}
			defer func() {
				// 释放令牌
				<-pool
			}()

			// 这里我们可以选择：
			// 1. 顺序执行三个任务（简单但可能不是最有效率）
			// 2. 或者创建子协程（但这样会超出池容量限制）

			// 方案1：顺序执行三个任务
			subErr := supportResistance(ctx, stock)
			if subErr != nil {
				mu.Lock()
				errs = append(errs, fmt.Errorf("supportResistance error for %s: %v", stock.Dm, subErr))
				mu.Unlock()
			}

			subErr = scoreMain(ctx, stock)
			if subErr != nil {
				mu.Lock()
				errs = append(errs, fmt.Errorf("scoreMain error for %s: %v", stock.Dm, subErr))
				mu.Unlock()
			}

			subErr = valuationIndicators(ctx, stock, "3Y")
			if subErr != nil {
				mu.Lock()
				errs = append(errs, fmt.Errorf("valuationIndicators error for %s: %v", stock.Dm, subErr))
				mu.Unlock()
			}
		})
	}

	// 等待所有任务完成
	wg.Wait()

	// 如果有错误，返回第一个错误或合并错误
	if len(errs) > 0 {
		// 可以选择返回第一个错误，或者合并所有错误
		err = fmt.Errorf("处理过程中发生 %d 个错误，第一个错误: %v", len(errs), errs[0])
	}

	return
}

func scoreMain(ctx context.Context, stock *entity.StockAllCode) (err error) {

	code := gstr.Split(stock.Dm, ".")[0]
	now := time.Now()

	var score *entity.StockScoreMain
	err = dao.StockScoreMain.Ctx(ctx).Where(dao.StockScoreMain.Columns().Symbol, stock.Dm).Scan(&score)
	if err != nil {
		return
	}
	if score == nil || gtime.Now().Sub(score.T).Hours() > 24 {
		var result map[string]interface{}
		_ = g.Client().GetVar(ctx, fmt.Sprintf("https://finance.pae.baidu.com/vapi/v1/analysis?code=%s&market=ab&isNew=1&finClientType=pc", code)).Scan(&result)
		newScore := new(entity.StockScoreMain)
		newScore.Symbol = stock.Dm
		newScore.T = gtime.New(time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location()))
		if result != nil {
			re := result["Result"].(map[string]interface{})
			if re["synthesisScore"] == nil {
				return
			}
			synthesisScore := re["synthesisScore"].(map[string]interface{})
			if stock.Industry == "" {
				stock.Industry = gconv.String(synthesisScore["secondIndustryName"])
				dao.StockAllCode.Ctx(ctx).Where(dao.StockAllCode.Columns().Dm, stock.Dm).Data(dao.StockAllCode.Columns().Industry, stock.Industry).Update()
			}
			newScore.ComprehensiveScore = gconv.Float64(synthesisScore["rating"])
			newScore.PriceScore = gconv.Float64(re["priceScore"].(map[string]interface{})["score"])
			newScore.IncomeScore = gconv.Float64(re["incomeScore"].(map[string]interface{})["score"])
			newScore.ValuationScore = gconv.Float64(re["valuationScore"].(map[string]interface{})["score"])
			newScore.FinancialScore = gconv.Float64(re["financeScore"].(map[string]interface{})["score"])
			newScore.RiskScore = gconv.Float64(re["riskScore"].(map[string]interface{})["score"])
			newScore.Mc = stock.Mc
			// 财报指标
			_ = FinancialIndicatorMtt(ctx, stock, re)
			dao.StockScoreMain.Ctx(ctx).Where(dao.StockScoreMain.Columns().Symbol, stock.Dm).Save(newScore)
		}
	}
	return
}

func supportResistance(ctx context.Context, stock *entity.StockAllCode) (err error) {
	code := gstr.Split(stock.Dm, ".")[0]

	now := time.Now()
	var support *entity.StockSupportResistance
	err = dao.StockSupportResistance.Ctx(ctx).Where(dao.StockSupportResistance.Columns().Symbol, stock.Dm).Scan(&support)
	if err != nil {
		return
	}
	if support == nil || gtime.Now().Sub(support.T).Hours() > 24 {
		var result map[string]interface{}
		_ = g.Client().GetVar(ctx, fmt.Sprintf("https://finance.pae.baidu.com/sapi/v1/get_analysis_quotation?all=1&newFormat=1&ktype=day&group=quotation_analysis_kline&code=%s&market_type=ab&finClientType=pc", code)).Scan(&result)
		if result != nil {
			newSupport := new(entity.StockSupportResistance)
			newSupport.Symbol = stock.Dm
			newSupport.T = gtime.New(time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location()))
			if result["Result"] != nil {
				res := result["Result"].(map[string]interface{})
				analysisData := res["analysisData"].(map[string]interface{})
				newSupport.Price = gconv.Float64(analysisData["price"])
				newSupport.Yl = gconv.Float64(analysisData["yl"])
				newSupport.Zc = gconv.Float64(analysisData["zc"])
				newSupport.Mc = stock.Mc
				dao.StockSupportResistance.Ctx(ctx).Where(dao.StockSupportResistance.Columns().Symbol, stock.Dm).Save(newSupport)
			}

		}
	}
	return
}

func valuationIndicators(ctx context.Context, stock *entity.StockAllCode, CalcWindow string) (err error) {
	code := gstr.Split(stock.Dm, ".")[0]

	var ValuationIndicators *entity.ValuationIndicators
	err = dao.ValuationIndicators.Ctx(ctx).Where(dao.ValuationIndicators.Columns().Symbol, stock.Dm).Scan(&ValuationIndicators)
	if err != nil {
		return
	}

	if ValuationIndicators != nil && gtime.New(GetRecentWeekday()).Sub(ValuationIndicators.T).Hours() < 24 {
		return
	}
	mtt := new(entity.ValuationIndicators)
	mtt.Symbol = stock.Dm
	mtt.Mc = stock.Mc
	mtt.CalcWindow = CalcWindow
	var peMap map[string]interface{}
	_ = g.Client().GetVar(ctx, fmt.Sprintf("https://gushitong.baidu.com/opendata?openapi=1&dspName=iphone&tn=tangram&client=app&query=市盈率(TTM)&code=%s&word=&resource_id=51171&srcid=51171&market=ab&tag=市盈率(TTM)&skip_industry=1&chart_select=&from=rating&isNew=1&finClientType=pc", code)).Scan(&peMap)
	if peMap != nil {
		list := peMap["Result"].([]interface{})
		if len(list) > 0 {
			displayData := list[0].(map[string]interface{})["DisplayData"]
			resultData := displayData.(map[string]interface{})["resultData"].(map[string]interface{})
			chartInfo := resultData["tplData"].(map[string]interface{})["result"].(map[string]interface{})["chartInfo"].([]interface{})
			if len(chartInfo) > 0 {
				chart := chartInfo[0].(map[string]interface{})
				mtt.PeTtm = gconv.Float64(chart["curTagValue"])
				mtt.PePercentile30 = gconv.Float64(chart["thirtyTantile"])
				mtt.PePercentile70 = gconv.Float64(chart["seventyTantile"])
				// 当前市盈率
				bodyList := chart["body"].([]interface{})
				if len(bodyList) > 0 {
					body := bodyList[len(bodyList)-1].([]interface{})
					mtt.T = gtime.New(gconv.String(body[0]))
					mtt.PePercentileCurrent = gconv.Float64(body[1])
				}
			}
		}
	}

	var pbMap map[string]interface{}
	_ = g.Client().GetVar(ctx, fmt.Sprintf("https://gushitong.baidu.com/opendata?openapi=1&dspName=iphone&tn=tangram&client=app&query=市净率&code=%s&word=&resource_id=51171&srcid=51171&market=ab&tag=市净率&skip_industry=1&chart_select=&from=rating&isNew=1&finClientType=pc", code)).Scan(&pbMap)
	if pbMap != nil {
		list := pbMap["Result"].([]interface{})
		if len(list) > 0 {
			displayData := list[0].(map[string]interface{})["DisplayData"]
			resultData := displayData.(map[string]interface{})["resultData"].(map[string]interface{})
			chartInfo := resultData["tplData"].(map[string]interface{})["result"].(map[string]interface{})["chartInfo"].([]interface{})
			if len(chartInfo) > 0 {
				chart := chartInfo[0].(map[string]interface{})
				mtt.Pb = gconv.Float64(chart["curTagValue"])
				mtt.PbPercentile30 = gconv.Float64(chart["thirtyTantile"])
				mtt.PbPercentile70 = gconv.Float64(chart["seventyTantile"])
				// 当前市盈率
				bodyList := chart["body"].([]interface{})
				if len(bodyList) > 0 {
					body := bodyList[len(bodyList)-1].([]interface{})
					mtt.T = gtime.New(gconv.String(body[0]))
					mtt.PbPercentileCurrent = gconv.Float64(body[1])
				}
			}
		}
	}
	if mtt.T == nil {
		now := time.Now()
		mtt.T = gtime.New(time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location()))

	}
	_, err = dao.ValuationIndicators.Ctx(ctx).Save(mtt)
	return
}

func FinancialIndicatorMtt(ctx context.Context, stock *entity.StockAllCode, re map[string]interface{}) (err error) {
	now := time.Now()
	financialMtt := new(entity.FinancialIndicatorsMtt)
	financialMtt.Symbol = stock.Dm
	financialMtt.Mc = stock.Mc
	financialMtt.T = gtime.New(time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location()))
	financial := re["financeScore"].(map[string]interface{})
	var dataType string
	if financial != nil {
		ratingContent := financial["ratingContent"].(map[string]interface{})
		if ratingContent["ratingChart"] != nil {
			ratingList := ratingContent["ratingChart"].([]interface{})
			for _, l := range ratingList {
				rList := l.(map[string]interface{})["chartList"].([]interface{})
				for _, rating := range rList {
					raMtt := rating.(map[string]interface{})
					switch raMtt["name"].(string) {
					case "营收增长率TTM(%)":
						financialMtt.RevenueGrowthTtm = gconv.Float64(raMtt["value"])
						financialMtt.RevenueGrowthTtmYoy = gconv.Float64(gstr.Replace(gconv.String(raMtt["yoy"]), "%", ""))
						if dataType == "" {
							// 什么类型财报/时间
							if raMtt["source"] != nil {
								sourceList := raMtt["source"].([]interface{})
								if len(sourceList) > 0 {
									sList := sourceList[len(sourceList)-1].([]interface{})
									if len(sList) > 0 {
										dataType = gconv.String(sList[0])
									}
								}
							}
						}
					case "销售毛利率TTM(%)":
						financialMtt.GrossProfitMarginTtm = gconv.Float64(raMtt["value"])
						financialMtt.GrossProfitMarginTtmYoy = gconv.Float64(gstr.Replace(gconv.String(raMtt["yoy"]), "%", ""))
					case "净资产收益率TTM(%)":
						financialMtt.RoaTtm = gconv.Float64(raMtt["value"])
						financialMtt.RoaTtmYoy = gconv.Float64(gstr.Replace(gconv.String(raMtt["yoy"]), "%", ""))
					case "销售净利率TTM(%)":
						financialMtt.NetProfitMarginTtm = gconv.Float64(raMtt["value"])
						financialMtt.NetProfitMarginTtmYoy = gconv.Float64(gstr.Replace(gconv.String(raMtt["yoy"]), "%", ""))
					case "流动比率":
						financialMtt.CurrentRatio = gconv.Float64(raMtt["value"])
						financialMtt.CurrentRatioYoy = gconv.Float64(gstr.Replace(gconv.String(raMtt["yoy"]), "%", ""))
					case "资产负债率(%)":
						financialMtt.DebtToAssetRatio = gconv.Float64(raMtt["value"])
						financialMtt.DebtToAssetRatioYoy = gconv.Float64(gstr.Replace(gconv.String(raMtt["yoy"]), "%", ""))
					case "速动比率":
						financialMtt.QuickRatio = gconv.Float64(raMtt["value"])
						financialMtt.QuickRatioYoy = gconv.Float64(gstr.Replace(gconv.String(raMtt["yoy"]), "%", ""))
					case "利息保障倍数":
						financialMtt.InterestCoverage = gconv.Float64(raMtt["value"])
						financialMtt.InterestCoverageYoy = gconv.Float64(gstr.Replace(gconv.String(raMtt["yoy"]), "%", ""))
					case "净利润现金含量(%)":
						financialMtt.NetProfitCashContent = gconv.Float64(raMtt["value"])
						financialMtt.NetProfitCashContentYoy = gconv.Float64(gstr.Replace(gconv.String(raMtt["yoy"]), "%", ""))
					case "存货周转天数(天/次)":
						financialMtt.InventoryTurnoverDays = gconv.Float64(raMtt["value"])
						financialMtt.InventoryTurnoverDaysYoy = gconv.Float64(gstr.Replace(gconv.String(raMtt["yoy"]), "%", ""))
					case "应收周转天数(天/次)":
						financialMtt.ReceivableTurnoverDays = gconv.Float64(raMtt["value"])
						financialMtt.ReceivableTurnoverDaysYoy = gconv.Float64(gstr.Replace(gconv.String(raMtt["yoy"]), "%", ""))
					case "股息增长率TTM(%)":
						financialMtt.DividendGrowthTtm = gconv.Float64(raMtt["value"])
						financialMtt.DividendGrowthTtmYoy = gconv.Float64(gstr.Replace(gconv.String(raMtt["yoy"]), "%", ""))
					case "股息支付率TTM(%)":
						financialMtt.DividendPayoutRatioTtm = gconv.Float64(raMtt["value"])
						financialMtt.DividendPayoutRatioTtmYoy = gconv.Float64(gstr.Replace(gconv.String(raMtt["yoy"]), "%", ""))
					case "股息率TTM(%)":
						financialMtt.DividendYieldTtm = gconv.Float64(raMtt["value"])
						financialMtt.DividendYieldTtmYoy = gconv.Float64(gstr.Replace(gconv.String(raMtt["yoy"]), "%", ""))
					}
				}
			}
		}
	}
	financialMtt.DataType = dataType
	_, err = dao.FinancialIndicatorsMtt.Ctx(ctx).Insert(financialMtt)
	return
}

func SetupBrowserOptions() []chromedp.ExecAllocatorOption {

	opts := append(chromedp.DefaultExecAllocatorOptions[:],
		chromedp.UserAgent("Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/131.0.0.0 Safari/537.36"),
		chromedp.Flag("headless", true),
		chromedp.Flag("no-sandbox", true),
		chromedp.Flag("ignore-certificate-errors", true),
		chromedp.Flag("disable-web-security", true),
		chromedp.Flag("proxy-bypass-list", "<-loopback>"),
		// 添加更多连接相关的标志
		chromedp.Flag("disable-background-networking", false),
		chromedp.Flag("enable-features", "NetworkService,NetworkServiceInProcess"),
		chromedp.Flag("disable-background-timer-throttling", true),
		chromedp.Flag("disable-client-side-phishing-detection", true),
		chromedp.Flag("disable-default-apps", true),
		chromedp.Flag("disable-extensions", true),
		chromedp.Flag("disable-popup-blocking", true),
		chromedp.Flag("disable-prompt-on-repost", true),
		chromedp.Flag("disable-sync", true),
	)
	return opts
}
