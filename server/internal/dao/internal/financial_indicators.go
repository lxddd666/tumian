// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// FinancialIndicatorsDao is the data access object for the table hg_financial_indicators.
type FinancialIndicatorsDao struct {
	table   string                     // table is the underlying table name of the DAO.
	group   string                     // group is the database configuration group name of the current DAO.
	columns FinancialIndicatorsColumns // columns contains all the column names of Table for convenient usage.
}

// FinancialIndicatorsColumns defines and stores column names for the table hg_financial_indicators.
type FinancialIndicatorsColumns struct {
	Id             string // 自增主键
	Symbol         string // 公司代码/股票代码 (例如: 000001.SZ, AAPL)
	Jzrq           string // 截止日期 (报告期结束日)
	Plrq           string // 披露日期
	ReportYear     string // 报告年度
	ReportQuarter  string // 报告季度 (1-4)
	ReportType     string // 报告类型: annual-年报, quarter-季报
	FiscalPeriod   string // 会计期间 (衍生字段，如2023Q1)
	Mgzbgjj        string // 每股资本公积金
	Mgjyhdxjl      string // 每股经营活动现金流量
	Mgjzc          string // 每股净资产
	Jbmgsy         string // 基本每股收益
	Xsmgsy         string // 稀释每股收益
	Mgwfplr        string // 每股未分配利润
	Kfmgsy         string // 扣非每股收益
	Jzcsyl         string // 净资产收益率(%)
	Jqjzcsyl       string // 加权净资产收益率(%)
	Tbjzcsyl       string // 摊薄净资产收益率(%)
	Tbzzcsyl       string // 摊薄总资产收益率(%)
	Xsmlv          string // 销售毛利率(%)
	Mlv            string // 毛利率(%)
	Jlv            string // 净利率(%)
	Sjslv          string // 实际税率(%)
	Zyyrsrzz       string // 主营收入同比增长(%)
	Jlrzz          string // 净利润同比增长(%)
	Gsmgsyzzdjlrzz string // 归属于母公司所有者的净利润同比增长(%)
	Kfjlrzz        string // 扣非净利润同比增长(%)
	Yyzsrgdhbzz    string // 营业总收入滚动环比增长(%)
	Sljlrjqhbzz    string // 归属净利润滚动环比增长(%)
	Kfjlrgdhbzz    string // 扣非净利润滚动环比增长(%)
	Yskyysr        string // 预收款/营业收入
	Xsxjlyysr      string // 销售现金流/营业收入
	Zcfzl          string // 资产负债比率(%)
	Chzzl          string // 存货周转率(次)
	DataSource     string // 数据来源
	Currency       string // 货币单位
	Unit           string // 单位: yuan-元
	IsCalculated   string // 是否为计算指标: 0-原始数据, 1-计算得出
	CalcVersion    string // 计算版本
	CreatedAt      string // 创建时间
	UpdatedAt      string // 更新时间
}

// financialIndicatorsColumns holds the columns for the table hg_financial_indicators.
var financialIndicatorsColumns = FinancialIndicatorsColumns{
	Id:             "id",
	Symbol:         "symbol",
	Jzrq:           "jzrq",
	Plrq:           "plrq",
	ReportYear:     "report_year",
	ReportQuarter:  "report_quarter",
	ReportType:     "report_type",
	FiscalPeriod:   "fiscal_period",
	Mgzbgjj:        "mgzbgjj",
	Mgjyhdxjl:      "mgjyhdxjl",
	Mgjzc:          "mgjzc",
	Jbmgsy:         "jbmgsy",
	Xsmgsy:         "xsmgsy",
	Mgwfplr:        "mgwfplr",
	Kfmgsy:         "kfmgsy",
	Jzcsyl:         "jzcsyl",
	Jqjzcsyl:       "jqjzcsyl",
	Tbjzcsyl:       "tbjzcsyl",
	Tbzzcsyl:       "tbzzcsyl",
	Xsmlv:          "xsmlv",
	Mlv:            "mlv",
	Jlv:            "jlv",
	Sjslv:          "sjslv",
	Zyyrsrzz:       "zyyrsrzz",
	Jlrzz:          "jlrzz",
	Gsmgsyzzdjlrzz: "gsmgsyzzdjlrzz",
	Kfjlrzz:        "kfjlrzz",
	Yyzsrgdhbzz:    "yyzsrgdhbzz",
	Sljlrjqhbzz:    "sljlrjqhbzz",
	Kfjlrgdhbzz:    "kfjlrgdhbzz",
	Yskyysr:        "yskyysr",
	Xsxjlyysr:      "xsxjlyysr",
	Zcfzl:          "zcfzl",
	Chzzl:          "chzzl",
	DataSource:     "data_source",
	Currency:       "currency",
	Unit:           "unit",
	IsCalculated:   "is_calculated",
	CalcVersion:    "calc_version",
	CreatedAt:      "created_at",
	UpdatedAt:      "updated_at",
}

// NewFinancialIndicatorsDao creates and returns a new DAO object for table data access.
func NewFinancialIndicatorsDao() *FinancialIndicatorsDao {
	return &FinancialIndicatorsDao{
		group:   "default",
		table:   "hg_financial_indicators",
		columns: financialIndicatorsColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *FinancialIndicatorsDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *FinancialIndicatorsDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *FinancialIndicatorsDao) Columns() FinancialIndicatorsColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *FinancialIndicatorsDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *FinancialIndicatorsDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rolls back the transaction and returns the error if function f returns a non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note: Do not commit or roll back the transaction in function f,
// as it is automatically handled by this function.
func (dao *FinancialIndicatorsDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
