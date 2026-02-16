// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// CciDataDao is the data access object for the table hg_cci_data.
type CciDataDao struct {
	table    string             // table is the underlying table name of the DAO.
	group    string             // group is the database configuration group name of the current DAO.
	columns  CciDataColumns     // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler // handlers for customized model modification.
}

// CciDataColumns defines and stores column names for the table hg_cci_data.
type CciDataColumns struct {
	Id        string // 自增主键
	Symbol    string // 股票或标的代码 (例如: AAPL, 000001.SZ)
	T         string // 交易时间 (统一为datetime类型，日线数据时间部分设为00:00:00)
	Cci       string // cci
	CreatedAt string // 数据创建时间
	UpdatedAt string // 数据更新时间
}

// cciDataColumns holds the columns for the table hg_cci_data.
var cciDataColumns = CciDataColumns{
	Id:        "id",
	Symbol:    "symbol",
	T:         "t",
	Cci:       "cci",
	CreatedAt: "created_at",
	UpdatedAt: "updated_at",
}

// NewCciDataDao creates and returns a new DAO object for table data access.
func NewCciDataDao(handlers ...gdb.ModelHandler) *CciDataDao {
	return &CciDataDao{
		group:    "default",
		table:    "hg_cci_data",
		columns:  cciDataColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *CciDataDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *CciDataDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *CciDataDao) Columns() CciDataColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *CciDataDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *CciDataDao) Ctx(ctx context.Context) *gdb.Model {
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
func (dao *CciDataDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
