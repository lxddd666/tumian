// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// FastkDataDao is the data access object for the table hg_fastk_data.
type FastkDataDao struct {
	table    string             // table is the underlying table name of the DAO.
	group    string             // group is the database configuration group name of the current DAO.
	columns  FastkDataColumns   // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler // handlers for customized model modification.
}

// FastkDataColumns defines and stores column names for the table hg_fastk_data.
type FastkDataColumns struct {
	Id        string // 自增主键
	Symbol    string // 股票或标的代码 (例如: AAPL, 000001.SZ)
	T         string // 交易时间 (统一为datetime类型，日线数据时间部分设为00:00:00)
	K         string // kst
	D         string // signal
	CreatedAt string // 数据创建时间
	UpdatedAt string // 数据更新时间
}

// fastkDataColumns holds the columns for the table hg_fastk_data.
var fastkDataColumns = FastkDataColumns{
	Id:        "id",
	Symbol:    "symbol",
	T:         "t",
	K:         "k",
	D:         "d",
	CreatedAt: "created_at",
	UpdatedAt: "updated_at",
}

// NewFastkDataDao creates and returns a new DAO object for table data access.
func NewFastkDataDao(handlers ...gdb.ModelHandler) *FastkDataDao {
	return &FastkDataDao{
		group:    "default",
		table:    "hg_fastk_data",
		columns:  fastkDataColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *FastkDataDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *FastkDataDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *FastkDataDao) Columns() FastkDataColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *FastkDataDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *FastkDataDao) Ctx(ctx context.Context) *gdb.Model {
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
func (dao *FastkDataDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
