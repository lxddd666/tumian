// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// StockSelfAiDao is the data access object for the table hg_stock_self_ai.
type StockSelfAiDao struct {
	table    string             // table is the underlying table name of the DAO.
	group    string             // group is the database configuration group name of the current DAO.
	columns  StockSelfAiColumns // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler // handlers for customized model modification.
}

// StockSelfAiColumns defines and stores column names for the table hg_stock_self_ai.
type StockSelfAiColumns struct {
	Id        string // 自增主键
	Name      string // ai名称 例如deepseek 千问
	Model     string // ai model
	BaseUrl   string // ai base url
	ApiKey    string // ai api key
	CreatedAt string // 创建时间
}

// stockSelfAiColumns holds the columns for the table hg_stock_self_ai.
var stockSelfAiColumns = StockSelfAiColumns{
	Id:        "id",
	Name:      "name",
	Model:     "model",
	BaseUrl:   "base_url",
	ApiKey:    "api_key",
	CreatedAt: "created_at",
}

// NewStockSelfAiDao creates and returns a new DAO object for table data access.
func NewStockSelfAiDao(handlers ...gdb.ModelHandler) *StockSelfAiDao {
	return &StockSelfAiDao{
		group:    "default",
		table:    "hg_stock_self_ai",
		columns:  stockSelfAiColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *StockSelfAiDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *StockSelfAiDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *StockSelfAiDao) Columns() StockSelfAiColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *StockSelfAiDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *StockSelfAiDao) Ctx(ctx context.Context) *gdb.Model {
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
func (dao *StockSelfAiDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
