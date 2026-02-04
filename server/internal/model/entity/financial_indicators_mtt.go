// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// FinancialIndicatorsMtt is the golang structure for table financial_indicators_mtt.
type FinancialIndicatorsMtt struct {
	Id                        uint        `json:"id"                        orm:"id"                            description:"主键ID"`
	Symbol                    string      `json:"symbol"                    orm:"symbol"                        description:"股票代码"`
	Mc                        string      `json:"mc"                        orm:"mc"                            description:"股票名称"`
	ReportDate                string      `json:"reportDate"                orm:"report_date"                   description:"报告日期（年-月-日）"`
	T                         *gtime.Time `json:"t"                         orm:"t"                             description:"更新时间"`
	RevenueGrowthTtm          float64     `json:"revenueGrowthTtm"          orm:"revenue_growth_ttm"            description:"营收增长率TTM (%) - 最近12个月营收同比增长率"`
	RevenueGrowthTtmYoy       float64     `json:"revenueGrowthTtmYoy"       orm:"revenue_growth_ttm_yoy"        description:"营收增长率TTM (%) - 最近12个月营收同比增长率 同比"`
	GrossProfitMarginTtm      float64     `json:"grossProfitMarginTtm"      orm:"gross_profit_margin_ttm"       description:"销售毛利率TTM (%) - （营业收入-营业成本）/营业收入 × 100%"`
	GrossProfitMarginTtmYoy   float64     `json:"grossProfitMarginTtmYoy"   orm:"gross_profit_margin_ttm_yoy"   description:"销售毛利率TTM (%) - （营业收入-营业成本）/营业收入 × 100% 同比"`
	RoaTtm                    float64     `json:"roaTtm"                    orm:"roa_ttm"                       description:"资产收益率TTM (%) - 净利润/总资产 × 100%"`
	RoaTtmYoy                 float64     `json:"roaTtmYoy"                 orm:"roa_ttm_yoy"                   description:"资产收益率TTM (%) - 净利润/总资产 × 100% 同比"`
	NetProfitMarginTtm        float64     `json:"netProfitMarginTtm"        orm:"net_profit_margin_ttm"         description:"销售净利率TTM (%) - 净利润/营业收入 × 100%"`
	NetProfitMarginTtmYoy     float64     `json:"netProfitMarginTtmYoy"     orm:"net_profit_margin_ttm_yoy"     description:"销售净利率TTM (%) - 净利润/营业收入 × 100% 同比"`
	CurrentRatio              float64     `json:"currentRatio"              orm:"current_ratio"                 description:"流动比率 - 流动资产/流动负债"`
	CurrentRatioYoy           float64     `json:"currentRatioYoy"           orm:"current_ratio_yoy"             description:"流动比率 - 流动资产/流动负债 同比"`
	DebtToAssetRatio          float64     `json:"debtToAssetRatio"          orm:"debt_to_asset_ratio"           description:"资产负债率 (%) - 总负债/总资产 × 100%"`
	DebtToAssetRatioYoy       float64     `json:"debtToAssetRatioYoy"       orm:"debt_to_asset_ratio_yoy"       description:"资产负债率 (%) - 总负债/总资产 × 100% 同比"`
	QuickRatio                float64     `json:"quickRatio"                orm:"quick_ratio"                   description:"速动比率 - （流动资产-存货）/流动负债"`
	QuickRatioYoy             float64     `json:"quickRatioYoy"             orm:"quick_ratio_yoy"               description:"速动比率 - （流动资产-存货）/流动负债 同比"`
	InterestCoverage          float64     `json:"interestCoverage"          orm:"interest_coverage"             description:"利息保障倍数 - 息税前利润/利息费用"`
	InterestCoverageYoy       float64     `json:"interestCoverageYoy"       orm:"interest_coverage_yoy"         description:"利息保障倍数 - 息税前利润/利息费用 同比"`
	NetProfitCashContent      float64     `json:"netProfitCashContent"      orm:"net_profit_cash_content"       description:"净利润现金含量 - 经营活动现金流净额/净利润"`
	NetProfitCashContentYoy   float64     `json:"netProfitCashContentYoy"   orm:"net_profit_cash_content_yoy"   description:"净利润现金含量 - 经营活动现金流净额/净利润 同比"`
	InventoryTurnoverDays     float64     `json:"inventoryTurnoverDays"     orm:"inventory_turnover_days"       description:"存货周转天数(天/次) - 365/(营业成本/平均存货余额)"`
	InventoryTurnoverDaysYoy  float64     `json:"inventoryTurnoverDaysYoy"  orm:"inventory_turnover_days_yoy"   description:"存货周转天数(天/次) - 365/(营业成本/平均存货余额) 同比"`
	ReceivableTurnoverDays    float64     `json:"receivableTurnoverDays"    orm:"receivable_turnover_days"      description:"应收周转天数(天/次) - 365/(营业收入/平均应收账款余额)"`
	ReceivableTurnoverDaysYoy float64     `json:"receivableTurnoverDaysYoy" orm:"receivable_turnover_days_yoy"  description:"应收周转天数(天/次) - 365/(营业收入/平均应收账款余额) 同比"`
	DividendGrowthTtm         float64     `json:"dividendGrowthTtm"         orm:"dividend_growth_ttm"           description:"股息增长率TTM (%) - 最近12个月股息同比增长率"`
	DividendGrowthTtmYoy      float64     `json:"dividendGrowthTtmYoy"      orm:"dividend_growth_ttm_yoy"       description:"股息增长率TTM (%) - 最近12个月股息同比增长率 同比"`
	DividendPayoutRatioTtm    float64     `json:"dividendPayoutRatioTtm"    orm:"dividend_payout_ratio_ttm"     description:"股息支付率TTM (%) - 每股股利/每股收益 × 100%"`
	DividendPayoutRatioTtmYoy float64     `json:"dividendPayoutRatioTtmYoy" orm:"dividend_payout_ratio_ttm_yoy" description:"股息支付率TTM (%) - 每股股利/每股收益 × 100% 同比"`
	DividendYieldTtm          float64     `json:"dividendYieldTtm"          orm:"dividend_yield_ttm"            description:"股息率TTM (%) - 每股股利/股价 × 100%"`
	DividendYieldTtmYoy       float64     `json:"dividendYieldTtmYoy"       orm:"dividend_yield_ttm_yoy"        description:"股息率TTM (%) - 每股股利/股价 × 100% 同比"`
	DataType                  string      `json:"dataType"                  orm:"data_type"                     description:"Q1 Q2 Q3"`
}
