// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// StockListDao is the data access object for the table hg_stock_list.
type StockListDao struct {
	table   string           // table is the underlying table name of the DAO.
	group   string           // group is the database configuration group name of the current DAO.
	columns StockListColumns // columns contains all the column names of Table for convenient usage.
}

// StockListColumns defines and stores column names for the table hg_stock_list.
type StockListColumns struct {
	Id           string // 自增主键
	Dm           string // 股票代码 (唯一业务标识，如: 000001)
	Mc           string // 股票名称 (如: 平安银行)
	Jys          string // 交易所代码 (如: sh, sz, bj)
	ExchangeName string // 交易所全称
	Symbol       string // 标准股票代码 (如: 000001.SZ)
	Status       string // 状态: 1-正常, 0-退市
	ListDate     string // 上市日期
	DataSource   string // 数据来源
	CreatedAt    string // 创建时间
	UpdatedAt    string // 更新时间
}

// stockListColumns holds the columns for the table hg_stock_list.
var stockListColumns = StockListColumns{
	Id:           "id",
	Dm:           "dm",
	Mc:           "mc",
	Jys:          "jys",
	ExchangeName: "exchange_name",
	Symbol:       "symbol",
	Status:       "status",
	ListDate:     "list_date",
	DataSource:   "data_source",
	CreatedAt:    "created_at",
	UpdatedAt:    "updated_at",
}

// NewStockListDao creates and returns a new DAO object for table data access.
func NewStockListDao() *StockListDao {
	return &StockListDao{
		group:   "default",
		table:   "hg_stock_list",
		columns: stockListColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *StockListDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *StockListDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *StockListDao) Columns() StockListColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *StockListDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *StockListDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rolls back the transaction and returns the error if function f returns a non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note: Do not commit or roll back the transaction in function f,
// as it is automatically handled by this function.
func (dao *StockListDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
