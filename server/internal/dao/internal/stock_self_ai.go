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
	table   string             // table is the underlying table name of the DAO.
	group   string             // group is the database configuration group name of the current DAO.
	columns StockSelfAiColumns // columns contains all the column names of Table for convenient usage.
}

// StockSelfAiColumns defines and stores column names for the table hg_stock_self_ai.
type StockSelfAiColumns struct {
	Id        string // 自增主键
	Name      string // ai模型全称名称 例如deepseek-plus 千问
	Model     string // ai model
	BaseUrl   string // ai base url
	ApiKey    string // ai api key
	CreatedAt string // 创建时间
	AiModel   string // 语言模型 qianwen deepseek
	Status    string // 0正常 -1不正常
}

// stockSelfAiColumns holds the columns for the table hg_stock_self_ai.
var stockSelfAiColumns = StockSelfAiColumns{
	Id:        "id",
	Name:      "name",
	Model:     "model",
	BaseUrl:   "base_url",
	ApiKey:    "api_key",
	CreatedAt: "created_at",
	AiModel:   "ai_model",
	Status:    "status",
}

// NewStockSelfAiDao creates and returns a new DAO object for table data access.
func NewStockSelfAiDao() *StockSelfAiDao {
	return &StockSelfAiDao{
		group:   "default",
		table:   "hg_stock_self_ai",
		columns: stockSelfAiColumns,
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
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
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
