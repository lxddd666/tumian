// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// ValuationIndicatorsDao is the data access object for the table hg_valuation_indicators.
type ValuationIndicatorsDao struct {
	table    string                     // table is the underlying table name of the DAO.
	group    string                     // group is the database configuration group name of the current DAO.
	columns  ValuationIndicatorsColumns // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler         // handlers for customized model modification.
}

// ValuationIndicatorsColumns defines and stores column names for the table hg_valuation_indicators.
type ValuationIndicatorsColumns struct {
	Id                  string // 主键ID
	Symbol              string // 股票代码
	Mc                  string // 股票名称
	T                   string // 数据日期
	CalcWindow          string // 分位值计算窗口: 1Y, 3Y, 5Y, 10Y
	PeTtm               string // 市盈率(TTM) - 股价/每股收益(TTM)
	PePercentile30      string // 市盈率30分位值
	PePercentile70      string // 市盈率70分位值
	PePercentileCurrent string // 当前市盈率历史百分位
	Pb                  string // 市净率 - 股价/每股净资产
	PbPercentile30      string // 市净率30分位值
	PbPercentile70      string // 市净率70分位值
	PbPercentileCurrent string // 当前市净率历史百分位
}

// valuationIndicatorsColumns holds the columns for the table hg_valuation_indicators.
var valuationIndicatorsColumns = ValuationIndicatorsColumns{
	Id:                  "id",
	Symbol:              "symbol",
	Mc:                  "mc",
	T:                   "t",
	CalcWindow:          "calc_window",
	PeTtm:               "pe_ttm",
	PePercentile30:      "pe_percentile_30",
	PePercentile70:      "pe_percentile_70",
	PePercentileCurrent: "pe_percentile_current",
	Pb:                  "pb",
	PbPercentile30:      "pb_percentile_30",
	PbPercentile70:      "pb_percentile_70",
	PbPercentileCurrent: "pb_percentile_current",
}

// NewValuationIndicatorsDao creates and returns a new DAO object for table data access.
func NewValuationIndicatorsDao(handlers ...gdb.ModelHandler) *ValuationIndicatorsDao {
	return &ValuationIndicatorsDao{
		group:    "default",
		table:    "hg_valuation_indicators",
		columns:  valuationIndicatorsColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *ValuationIndicatorsDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *ValuationIndicatorsDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *ValuationIndicatorsDao) Columns() ValuationIndicatorsColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *ValuationIndicatorsDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *ValuationIndicatorsDao) Ctx(ctx context.Context) *gdb.Model {
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
func (dao *ValuationIndicatorsDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
