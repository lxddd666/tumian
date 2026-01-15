// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// FundStockHoldingDao is the data access object for the table hg_fund_stock_holding.
type FundStockHoldingDao struct {
	table    string                  // table is the underlying table name of the DAO.
	group    string                  // group is the database configuration group name of the current DAO.
	columns  FundStockHoldingColumns // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler      // handlers for customized model modification.
}

// FundStockHoldingColumns defines and stores column names for the table hg_fund_stock_holding.
type FundStockHoldingColumns struct {
	Id         string // 自增主键
	Jzrq       string // 截止日期 (报告期，如2025-12-31)
	T          string // 交易时间 (衍生自jzrq，兼容时间序列查询)
	Jjmc       string // 基金名称
	Jjdm       string // 基金代码
	Symbol     string // 股票代码 (如: 000001.SZ)
	Ccsl       string // 持仓数量(股)
	Ltbl       string // 占流通股比例(%)
	Cgsz       string // 持股市值（元）
	Jzbl       string // 占净值比例（%）
	AvgCost    string // 估算持仓成本 (元/股)
	DataSource string // 数据来源 (如: 基金季报)
	ReportType string // 报告类型: 季报, 中报, 年报
	IsLatest   string // 是否为该基金对该股票的最新持仓: 0-否, 1-是
	CreatedAt  string // 创建时间
	UpdatedAt  string // 更新时间
}

// fundStockHoldingColumns holds the columns for the table hg_fund_stock_holding.
var fundStockHoldingColumns = FundStockHoldingColumns{
	Id:         "id",
	Jzrq:       "jzrq",
	T:          "t",
	Jjmc:       "jjmc",
	Jjdm:       "jjdm",
	Symbol:     "symbol",
	Ccsl:       "ccsl",
	Ltbl:       "ltbl",
	Cgsz:       "cgsz",
	Jzbl:       "jzbl",
	AvgCost:    "avg_cost",
	DataSource: "data_source",
	ReportType: "report_type",
	IsLatest:   "is_latest",
	CreatedAt:  "created_at",
	UpdatedAt:  "updated_at",
}

// NewFundStockHoldingDao creates and returns a new DAO object for table data access.
func NewFundStockHoldingDao(handlers ...gdb.ModelHandler) *FundStockHoldingDao {
	return &FundStockHoldingDao{
		group:    "default",
		table:    "hg_fund_stock_holding",
		columns:  fundStockHoldingColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *FundStockHoldingDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *FundStockHoldingDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *FundStockHoldingDao) Columns() FundStockHoldingColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *FundStockHoldingDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *FundStockHoldingDao) Ctx(ctx context.Context) *gdb.Model {
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
func (dao *FundStockHoldingDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
