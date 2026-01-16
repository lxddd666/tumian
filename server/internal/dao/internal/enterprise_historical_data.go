// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// EnterpriseHistoricalDataDao is the data access object for the table hg_enterprise_historical_data.
type EnterpriseHistoricalDataDao struct {
	table    string                          // table is the underlying table name of the DAO.
	group    string                          // group is the database configuration group name of the current DAO.
	columns  EnterpriseHistoricalDataColumns // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler              // handlers for customized model modification.
}

// EnterpriseHistoricalDataColumns defines and stores column names for the table hg_enterprise_historical_data.
type EnterpriseHistoricalDataColumns struct {
	Id        string // 主键ID
	Symbol    string // 公司代码/股票代码 (例如: 000001.SZ, AAPL)
	T         string // 交易时间
	O         string // 开盘价
	H         string // 最高价
	L         string // 最低价
	C         string // 收盘价
	V         string // 成交量
	A         string // 成交额
	Pc        string // 前收盘价
	Sf        string // 停牌状态: 1停牌, 0不停牌
	CreatedAt string // 创建时间
	UpdatedAt string // 更新时间
}

// enterpriseHistoricalDataColumns holds the columns for the table hg_enterprise_historical_data.
var enterpriseHistoricalDataColumns = EnterpriseHistoricalDataColumns{
	Id:        "id",
	Symbol:    "symbol",
	T:         "t",
	O:         "o",
	H:         "h",
	L:         "l",
	C:         "c",
	V:         "v",
	A:         "a",
	Pc:        "pc",
	Sf:        "sf",
	CreatedAt: "created_at",
	UpdatedAt: "updated_at",
}

// NewEnterpriseHistoricalDataDao creates and returns a new DAO object for table data access.
func NewEnterpriseHistoricalDataDao(handlers ...gdb.ModelHandler) *EnterpriseHistoricalDataDao {
	return &EnterpriseHistoricalDataDao{
		group:    "default",
		table:    "hg_enterprise_historical_data",
		columns:  enterpriseHistoricalDataColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *EnterpriseHistoricalDataDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *EnterpriseHistoricalDataDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *EnterpriseHistoricalDataDao) Columns() EnterpriseHistoricalDataColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *EnterpriseHistoricalDataDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *EnterpriseHistoricalDataDao) Ctx(ctx context.Context) *gdb.Model {
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
func (dao *EnterpriseHistoricalDataDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
