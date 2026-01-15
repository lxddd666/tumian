// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// ShareholderChangeDao is the data access object for the table hg_shareholder_change.
type ShareholderChangeDao struct {
	table   string                   // table is the underlying table name of the DAO.
	group   string                   // group is the database configuration group name of the current DAO.
	columns ShareholderChangeColumns // columns contains all the column names of Table for convenient usage.
}

// ShareholderChangeColumns defines and stores column names for the table hg_shareholder_change.
type ShareholderChangeColumns struct {
	Id              string // 自增主键
	Symbol          string // 股票代码 (如: 000001.SZ)
	Jzrq            string // 截止日期 (统计截止日，如2025-12-31)[citation:9]
	Gdhs            string // 股东户数 (统计截止日的总户数)[citation:3][citation:6]
	Bh              string // 比上期变化情况 (通常为百分比，如 -6.23 表示减少6.23%)[citation:9]
	ChangeDirection string // 变化方向 (衍生字段)
	DataSource      string // 数据来源 (如: 交易所公告[citation:6]、互动平台[citation:3])
	AnnDate         string // 公告日期 (信息发布日期)[citation:1]
	CreatedAt       string // 创建时间
	UpdatedAt       string // 更新时间
}

// shareholderChangeColumns holds the columns for the table hg_shareholder_change.
var shareholderChangeColumns = ShareholderChangeColumns{
	Id:              "id",
	Symbol:          "symbol",
	Jzrq:            "jzrq",
	Gdhs:            "gdhs",
	Bh:              "bh",
	ChangeDirection: "change_direction",
	DataSource:      "data_source",
	AnnDate:         "ann_date",
	CreatedAt:       "created_at",
	UpdatedAt:       "updated_at",
}

// NewShareholderChangeDao creates and returns a new DAO object for table data access.
func NewShareholderChangeDao() *ShareholderChangeDao {
	return &ShareholderChangeDao{
		group:   "default",
		table:   "hg_shareholder_change",
		columns: shareholderChangeColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *ShareholderChangeDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *ShareholderChangeDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *ShareholderChangeDao) Columns() ShareholderChangeColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *ShareholderChangeDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *ShareholderChangeDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rolls back the transaction and returns the error if function f returns a non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note: Do not commit or roll back the transaction in function f,
// as it is automatically handled by this function.
func (dao *ShareholderChangeDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
