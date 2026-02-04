// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// FinancialIndicatorsMttDao is the data access object for the table hg_financial_indicators_mtt.
type FinancialIndicatorsMttDao struct {
	table    string                        // table is the underlying table name of the DAO.
	group    string                        // group is the database configuration group name of the current DAO.
	columns  FinancialIndicatorsMttColumns // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler            // handlers for customized model modification.
}

// FinancialIndicatorsMttColumns defines and stores column names for the table hg_financial_indicators_mtt.
type FinancialIndicatorsMttColumns struct {
	Id                        string // 主键ID
	Symbol                    string // 股票代码
	Mc                        string // 股票名称
	ReportDate                string // 报告日期（年-月-日）
	T                         string // 更新时间
	RevenueGrowthTtm          string // 营收增长率TTM (%) - 最近12个月营收同比增长率
	RevenueGrowthTtmYoy       string // 营收增长率TTM (%) - 最近12个月营收同比增长率 同比
	GrossProfitMarginTtm      string // 销售毛利率TTM (%) - （营业收入-营业成本）/营业收入 × 100%
	GrossProfitMarginTtmYoy   string // 销售毛利率TTM (%) - （营业收入-营业成本）/营业收入 × 100% 同比
	RoaTtm                    string // 资产收益率TTM (%) - 净利润/总资产 × 100%
	RoaTtmYoy                 string // 资产收益率TTM (%) - 净利润/总资产 × 100% 同比
	NetProfitMarginTtm        string // 销售净利率TTM (%) - 净利润/营业收入 × 100%
	NetProfitMarginTtmYoy     string // 销售净利率TTM (%) - 净利润/营业收入 × 100% 同比
	CurrentRatio              string // 流动比率 - 流动资产/流动负债
	CurrentRatioYoy           string // 流动比率 - 流动资产/流动负债 同比
	DebtToAssetRatio          string // 资产负债率 (%) - 总负债/总资产 × 100%
	DebtToAssetRatioYoy       string // 资产负债率 (%) - 总负债/总资产 × 100% 同比
	QuickRatio                string // 速动比率 - （流动资产-存货）/流动负债
	QuickRatioYoy             string // 速动比率 - （流动资产-存货）/流动负债 同比
	InterestCoverage          string // 利息保障倍数 - 息税前利润/利息费用
	InterestCoverageYoy       string // 利息保障倍数 - 息税前利润/利息费用 同比
	NetProfitCashContent      string // 净利润现金含量 - 经营活动现金流净额/净利润
	NetProfitCashContentYoy   string // 净利润现金含量 - 经营活动现金流净额/净利润 同比
	InventoryTurnoverDays     string // 存货周转天数(天/次) - 365/(营业成本/平均存货余额)
	InventoryTurnoverDaysYoy  string // 存货周转天数(天/次) - 365/(营业成本/平均存货余额) 同比
	ReceivableTurnoverDays    string // 应收周转天数(天/次) - 365/(营业收入/平均应收账款余额)
	ReceivableTurnoverDaysYoy string // 应收周转天数(天/次) - 365/(营业收入/平均应收账款余额) 同比
	DividendGrowthTtm         string // 股息增长率TTM (%) - 最近12个月股息同比增长率
	DividendGrowthTtmYoy      string // 股息增长率TTM (%) - 最近12个月股息同比增长率 同比
	DividendPayoutRatioTtm    string // 股息支付率TTM (%) - 每股股利/每股收益 × 100%
	DividendPayoutRatioTtmYoy string // 股息支付率TTM (%) - 每股股利/每股收益 × 100% 同比
	DividendYieldTtm          string // 股息率TTM (%) - 每股股利/股价 × 100%
	DividendYieldTtmYoy       string // 股息率TTM (%) - 每股股利/股价 × 100% 同比
	DataType                  string // Q1 Q2 Q3
}

// financialIndicatorsMttColumns holds the columns for the table hg_financial_indicators_mtt.
var financialIndicatorsMttColumns = FinancialIndicatorsMttColumns{
	Id:                        "id",
	Symbol:                    "symbol",
	Mc:                        "mc",
	ReportDate:                "report_date",
	T:                         "t",
	RevenueGrowthTtm:          "revenue_growth_ttm",
	RevenueGrowthTtmYoy:       "revenue_growth_ttm_yoy",
	GrossProfitMarginTtm:      "gross_profit_margin_ttm",
	GrossProfitMarginTtmYoy:   "gross_profit_margin_ttm_yoy",
	RoaTtm:                    "roa_ttm",
	RoaTtmYoy:                 "roa_ttm_yoy",
	NetProfitMarginTtm:        "net_profit_margin_ttm",
	NetProfitMarginTtmYoy:     "net_profit_margin_ttm_yoy",
	CurrentRatio:              "current_ratio",
	CurrentRatioYoy:           "current_ratio_yoy",
	DebtToAssetRatio:          "debt_to_asset_ratio",
	DebtToAssetRatioYoy:       "debt_to_asset_ratio_yoy",
	QuickRatio:                "quick_ratio",
	QuickRatioYoy:             "quick_ratio_yoy",
	InterestCoverage:          "interest_coverage",
	InterestCoverageYoy:       "interest_coverage_yoy",
	NetProfitCashContent:      "net_profit_cash_content",
	NetProfitCashContentYoy:   "net_profit_cash_content_yoy",
	InventoryTurnoverDays:     "inventory_turnover_days",
	InventoryTurnoverDaysYoy:  "inventory_turnover_days_yoy",
	ReceivableTurnoverDays:    "receivable_turnover_days",
	ReceivableTurnoverDaysYoy: "receivable_turnover_days_yoy",
	DividendGrowthTtm:         "dividend_growth_ttm",
	DividendGrowthTtmYoy:      "dividend_growth_ttm_yoy",
	DividendPayoutRatioTtm:    "dividend_payout_ratio_ttm",
	DividendPayoutRatioTtmYoy: "dividend_payout_ratio_ttm_yoy",
	DividendYieldTtm:          "dividend_yield_ttm",
	DividendYieldTtmYoy:       "dividend_yield_ttm_yoy",
	DataType:                  "data_type",
}

// NewFinancialIndicatorsMttDao creates and returns a new DAO object for table data access.
func NewFinancialIndicatorsMttDao(handlers ...gdb.ModelHandler) *FinancialIndicatorsMttDao {
	return &FinancialIndicatorsMttDao{
		group:    "default",
		table:    "hg_financial_indicators_mtt",
		columns:  financialIndicatorsMttColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *FinancialIndicatorsMttDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *FinancialIndicatorsMttDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *FinancialIndicatorsMttDao) Columns() FinancialIndicatorsMttColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *FinancialIndicatorsMttDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *FinancialIndicatorsMttDao) Ctx(ctx context.Context) *gdb.Model {
	model := dao.DB().Model(dao.table)
	for _, handler := range dao.handlers {
		model = handler(model)
	}
	return model.Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rolls back the transaction and returns the error if function f returns a non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note: Do not commit or roll back the transaction in function f,
// as it is automatically handled by this function.
func (dao *FinancialIndicatorsMttDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
