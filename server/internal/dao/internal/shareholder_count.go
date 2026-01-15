// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// ShareholderCountDao is the data access object for the table hg_shareholder_count.
type ShareholderCountDao struct {
	table   string                  // table is the underlying table name of the DAO.
	group   string                  // group is the database configuration group name of the current DAO.
	columns ShareholderCountColumns // columns contains all the column names of Table for convenient usage.
}

// ShareholderCountColumns defines and stores column names for the table hg_shareholder_count.
type ShareholderCountColumns struct {
	Id            string // 自增主键
	Symbol        string // 公司代码/股票代码 (例如: 000001.SZ, 600000.SS)
	Jzrq          string // 截止日期 (报告期结束日, 如2023-09-30)
	ReportYear    string // 报告年度
	ReportQuarter string // 报告季度 (1-4)
	ReportType    string // 报告类型: annual-年报, half_year-中报, quarter-季报
	Gdzs          string // 股东总数 (户)
	Agdhs         string // A股东户数 (户)
	Bgdhs         string // B股东户数 (户)
	Hgdhs         string // H股东户数 (户)
	Yltgdhs       string // 已流通股东户数 (户)
	Wltgdhs       string // 未流通股东户数 (户)
	AgRatio       string // A股股东占比(%)
	YltRatio      string // 已流通股东占比(%)
	DataSource    string // 数据来源
	IsLatest      string // 是否为该报告期最新数据
	CreatedAt     string // 创建时间
	UpdatedAt     string // 更新时间
}

// shareholderCountColumns holds the columns for the table hg_shareholder_count.
var shareholderCountColumns = ShareholderCountColumns{
	Id:            "id",
	Symbol:        "symbol",
	Jzrq:          "jzrq",
	ReportYear:    "report_year",
	ReportQuarter: "report_quarter",
	ReportType:    "report_type",
	Gdzs:          "gdzs",
	Agdhs:         "agdhs",
	Bgdhs:         "bgdhs",
	Hgdhs:         "hgdhs",
	Yltgdhs:       "yltgdhs",
	Wltgdhs:       "wltgdhs",
	AgRatio:       "ag_ratio",
	YltRatio:      "ylt_ratio",
	DataSource:    "data_source",
	IsLatest:      "is_latest",
	CreatedAt:     "created_at",
	UpdatedAt:     "updated_at",
}

// NewShareholderCountDao creates and returns a new DAO object for table data access.
func NewShareholderCountDao() *ShareholderCountDao {
	return &ShareholderCountDao{
		group:   "default",
		table:   "hg_shareholder_count",
		columns: shareholderCountColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *ShareholderCountDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *ShareholderCountDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *ShareholderCountDao) Columns() ShareholderCountColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *ShareholderCountDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *ShareholderCountDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rolls back the transaction and returns the error if function f returns a non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note: Do not commit or roll back the transaction in function f,
// as it is automatically handled by this function.
func (dao *ShareholderCountDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
