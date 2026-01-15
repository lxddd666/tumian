// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// StockBasicInfoDao is the data access object for the table hg_stock_basic_info.
type StockBasicInfoDao struct {
	table   string                // table is the underlying table name of the DAO.
	group   string                // group is the database configuration group name of the current DAO.
	columns StockBasicInfoColumns // columns contains all the column names of Table for convenient usage.
}

// StockBasicInfoColumns defines and stores column names for the table hg_stock_basic_info.
type StockBasicInfoColumns struct {
	Id             string // 自增主键
	Symbol         string // 股票代码
	Ii             string //
	Ei             string //
	Exchange       string // 交易所名称
	Name           string //
	ShortName      string //
	EnName         string //
	Od             string // 上市日期
	DataUpdateDate string //
	Pc             string //
	Up             string //
	Dp             string //
	Pk             string //
	Fv             string //
	Tv             string //
	FloatRatio     string // 流通股比例 (%)
	Is             string //
	TradingStatus  string // 交易状态描述
	Industry       string //
	Sector         string //
	MarketType     string //
	DataSource     string //
	IsActive       string //
	Version        string //
	CreatedAt      string //
	UpdatedAt      string //
}

// stockBasicInfoColumns holds the columns for the table hg_stock_basic_info.
var stockBasicInfoColumns = StockBasicInfoColumns{
	Id:             "id",
	Symbol:         "symbol",
	Ii:             "ii",
	Ei:             "ei",
	Exchange:       "exchange",
	Name:           "name",
	ShortName:      "short_name",
	EnName:         "en_name",
	Od:             "od",
	DataUpdateDate: "data_update_date",
	Pc:             "pc",
	Up:             "up",
	Dp:             "dp",
	Pk:             "pk",
	Fv:             "fv",
	Tv:             "tv",
	FloatRatio:     "float_ratio",
	Is:             "is",
	TradingStatus:  "trading_status",
	Industry:       "industry",
	Sector:         "sector",
	MarketType:     "market_type",
	DataSource:     "data_source",
	IsActive:       "is_active",
	Version:        "version",
	CreatedAt:      "created_at",
	UpdatedAt:      "updated_at",
}

// NewStockBasicInfoDao creates and returns a new DAO object for table data access.
func NewStockBasicInfoDao() *StockBasicInfoDao {
	return &StockBasicInfoDao{
		group:   "default",
		table:   "hg_stock_basic_info",
		columns: stockBasicInfoColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *StockBasicInfoDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *StockBasicInfoDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *StockBasicInfoDao) Columns() StockBasicInfoColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *StockBasicInfoDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *StockBasicInfoDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rolls back the transaction and returns the error if function f returns a non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note: Do not commit or roll back the transaction in function f,
// as it is automatically handled by this function.
func (dao *StockBasicInfoDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
