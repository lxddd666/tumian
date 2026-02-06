// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// StockSupportResistanceDao is the data access object for the table hg_stock_support_resistance.
type StockSupportResistanceDao struct {
	table   string                        // table is the underlying table name of the DAO.
	group   string                        // group is the database configuration group name of the current DAO.
	columns StockSupportResistanceColumns // columns contains all the column names of Table for convenient usage.
}

// StockSupportResistanceColumns defines and stores column names for the table hg_stock_support_resistance.
type StockSupportResistanceColumns struct {
	Id        string // 主键ID
	T         string // 时间
	Symbol    string // 股票或标的代码 (例如: AAPL, 000001.SZ)
	Price     string // 当前股票价格
	Yl        string // 压力位
	Zc        string // 支撑位
	CreatedAt string // 创建时间
	Mc        string // mc
}

// stockSupportResistanceColumns holds the columns for the table hg_stock_support_resistance.
var stockSupportResistanceColumns = StockSupportResistanceColumns{
	Id:        "id",
	T:         "t",
	Symbol:    "symbol",
	Price:     "price",
	Yl:        "yl",
	Zc:        "zc",
	CreatedAt: "created_at",
	Mc:        "mc",
}

// NewStockSupportResistanceDao creates and returns a new DAO object for table data access.
func NewStockSupportResistanceDao() *StockSupportResistanceDao {
	return &StockSupportResistanceDao{
		group:   "default",
		table:   "hg_stock_support_resistance",
		columns: stockSupportResistanceColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *StockSupportResistanceDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *StockSupportResistanceDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *StockSupportResistanceDao) Columns() StockSupportResistanceColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *StockSupportResistanceDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *StockSupportResistanceDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rolls back the transaction and returns the error if function f returns a non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note: Do not commit or roll back the transaction in function f,
// as it is automatically handled by this function.
func (dao *StockSupportResistanceDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
