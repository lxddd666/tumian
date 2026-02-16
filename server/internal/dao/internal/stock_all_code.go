// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// StockAllCodeDao is the data access object for the table hg_stock_all_code.
type StockAllCodeDao struct {
	table    string              // table is the underlying table name of the DAO.
	group    string              // group is the database configuration group name of the current DAO.
	columns  StockAllCodeColumns // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler  // handlers for customized model modification.
}

// StockAllCodeColumns defines and stores column names for the table hg_stock_all_code.
type StockAllCodeColumns struct {
	Dm        string //
	Mc        string // 股票名称
	Jys       string // 交易所
	CreatedAt string // 创建时间
	Industry  string // 所属行业
}

// stockAllCodeColumns holds the columns for the table hg_stock_all_code.
var stockAllCodeColumns = StockAllCodeColumns{
	Dm:        "dm",
	Mc:        "mc",
	Jys:       "jys",
	CreatedAt: "created_at",
	Industry:  "industry",
}

// NewStockAllCodeDao creates and returns a new DAO object for table data access.
func NewStockAllCodeDao(handlers ...gdb.ModelHandler) *StockAllCodeDao {
	return &StockAllCodeDao{
		group:    "default",
		table:    "hg_stock_all_code",
		columns:  stockAllCodeColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *StockAllCodeDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *StockAllCodeDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *StockAllCodeDao) Columns() StockAllCodeColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *StockAllCodeDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *StockAllCodeDao) Ctx(ctx context.Context) *gdb.Model {
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
func (dao *StockAllCodeDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
