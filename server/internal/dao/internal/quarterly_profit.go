// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// QuarterlyProfitDao is the data access object for the table hg_quarterly_profit.
type QuarterlyProfitDao struct {
	table   string                 // table is the underlying table name of the DAO.
	group   string                 // group is the database configuration group name of the current DAO.
	columns QuarterlyProfitColumns // columns contains all the column names of Table for convenient usage.
}

// QuarterlyProfitColumns defines and stores column names for the table hg_quarterly_profit.
type QuarterlyProfitColumns struct {
	Id                   string // 自增主键
	Symbol               string // 股票代码 (如: 000001.SZ)
	Date                 string // 截止日期 (报告期截止日，如2025-03-31)
	ReportYear           string // 报告年度
	ReportQuarter        string // 报告季度 (1-4)
	FiscalPeriod         string // 会计期间 (如2025Q1)
	Income               string // 营业收入（万元）
	Expend               string // 营业支出（万元）
	Profit               string // 营业利润（万元）
	Totalp               string // 利润总额（万元）
	Reprofit             string // 净利润（万元）
	Basege               string // 基本每股收益(元/股)
	Ettege               string // 稀释每股收益(元/股)
	Otherp               string // 其他综合收益（万元）
	Totalcp              string // 综合收益总额（万元）
	GrossProfitMargin    string // 毛利率(%)
	NetProfitMargin      string // 净利率(%)
	OperatingProfitRatio string // 营业利润率(%)
	ReportType           string // 报告类型: 一季报, 中报, 三季报, 年报
	DataSource           string // 数据来源 (如: 交易所财报)
	Currency             string // 货币单位
	IsAudited            string // 是否审计: 0-未审计, 1-已审计
	CreatedAt            string // 创建时间
	UpdatedAt            string // 更新时间
}

// quarterlyProfitColumns holds the columns for the table hg_quarterly_profit.
var quarterlyProfitColumns = QuarterlyProfitColumns{
	Id:                   "id",
	Symbol:               "symbol",
	Date:                 "date",
	ReportYear:           "report_year",
	ReportQuarter:        "report_quarter",
	FiscalPeriod:         "fiscal_period",
	Income:               "income",
	Expend:               "expend",
	Profit:               "profit",
	Totalp:               "totalp",
	Reprofit:             "reprofit",
	Basege:               "basege",
	Ettege:               "ettege",
	Otherp:               "otherp",
	Totalcp:              "totalcp",
	GrossProfitMargin:    "gross_profit_margin",
	NetProfitMargin:      "net_profit_margin",
	OperatingProfitRatio: "operating_profit_ratio",
	ReportType:           "report_type",
	DataSource:           "data_source",
	Currency:             "currency",
	IsAudited:            "is_audited",
	CreatedAt:            "created_at",
	UpdatedAt:            "updated_at",
}

// NewQuarterlyProfitDao creates and returns a new DAO object for table data access.
func NewQuarterlyProfitDao() *QuarterlyProfitDao {
	return &QuarterlyProfitDao{
		group:   "default",
		table:   "hg_quarterly_profit",
		columns: quarterlyProfitColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *QuarterlyProfitDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *QuarterlyProfitDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *QuarterlyProfitDao) Columns() QuarterlyProfitColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *QuarterlyProfitDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *QuarterlyProfitDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rolls back the transaction and returns the error if function f returns a non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note: Do not commit or roll back the transaction in function f,
// as it is automatically handled by this function.
func (dao *QuarterlyProfitDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
