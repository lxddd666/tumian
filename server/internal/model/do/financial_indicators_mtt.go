// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// FinancialIndicatorsMtt is the golang structure of table hg_financial_indicators_mtt for DAO operations like Where/Data.
type FinancialIndicatorsMtt struct {
	g.Meta                    `orm:"table:hg_financial_indicators_mtt, do:true"`
	Id                        any         // 主键ID
	Symbol                    any         // 股票代码
	Mc                        any         // 股票名称
	ReportDate                any         // 报告日期（年-月-日）
	T                         *gtime.Time // 更新时间
	RevenueGrowthTtm          any         // 营收增长率TTM (%) - 最近12个月营收同比增长率
	RevenueGrowthTtmYoy       any         // 营收增长率TTM (%) - 最近12个月营收同比增长率 同比
	GrossProfitMarginTtm      any         // 销售毛利率TTM (%) - （营业收入-营业成本）/营业收入 × 100%
	GrossProfitMarginTtmYoy   any         // 销售毛利率TTM (%) - （营业收入-营业成本）/营业收入 × 100% 同比
	RoaTtm                    any         // 资产收益率TTM (%) - 净利润/总资产 × 100%
	RoaTtmYoy                 any         // 资产收益率TTM (%) - 净利润/总资产 × 100% 同比
	NetProfitMarginTtm        any         // 销售净利率TTM (%) - 净利润/营业收入 × 100%
	NetProfitMarginTtmYoy     any         // 销售净利率TTM (%) - 净利润/营业收入 × 100% 同比
	CurrentRatio              any         // 流动比率 - 流动资产/流动负债
	CurrentRatioYoy           any         // 流动比率 - 流动资产/流动负债 同比
	DebtToAssetRatio          any         // 资产负债率 (%) - 总负债/总资产 × 100%
	DebtToAssetRatioYoy       any         // 资产负债率 (%) - 总负债/总资产 × 100% 同比
	QuickRatio                any         // 速动比率 - （流动资产-存货）/流动负债
	QuickRatioYoy             any         // 速动比率 - （流动资产-存货）/流动负债 同比
	InterestCoverage          any         // 利息保障倍数 - 息税前利润/利息费用
	InterestCoverageYoy       any         // 利息保障倍数 - 息税前利润/利息费用 同比
	NetProfitCashContent      any         // 净利润现金含量 - 经营活动现金流净额/净利润
	NetProfitCashContentYoy   any         // 净利润现金含量 - 经营活动现金流净额/净利润 同比
	InventoryTurnoverDays     any         // 存货周转天数(天/次) - 365/(营业成本/平均存货余额)
	InventoryTurnoverDaysYoy  any         // 存货周转天数(天/次) - 365/(营业成本/平均存货余额) 同比
	ReceivableTurnoverDays    any         // 应收周转天数(天/次) - 365/(营业收入/平均应收账款余额)
	ReceivableTurnoverDaysYoy any         // 应收周转天数(天/次) - 365/(营业收入/平均应收账款余额) 同比
	DividendGrowthTtm         any         // 股息增长率TTM (%) - 最近12个月股息同比增长率
	DividendGrowthTtmYoy      any         // 股息增长率TTM (%) - 最近12个月股息同比增长率 同比
	DividendPayoutRatioTtm    any         // 股息支付率TTM (%) - 每股股利/每股收益 × 100%
	DividendPayoutRatioTtmYoy any         // 股息支付率TTM (%) - 每股股利/每股收益 × 100% 同比
	DividendYieldTtm          any         // 股息率TTM (%) - 每股股利/股价 × 100%
	DividendYieldTtmYoy       any         // 股息率TTM (%) - 每股股利/股价 × 100% 同比
	DataType                  any         // Q1 Q2 Q3
}
