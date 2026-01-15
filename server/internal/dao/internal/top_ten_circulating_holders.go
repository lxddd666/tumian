// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// TopTenCirculatingHoldersDao is the data access object for the table hg_top_ten_circulating_holders.
type TopTenCirculatingHoldersDao struct {
	table   string                          // table is the underlying table name of the DAO.
	group   string                          // group is the database configuration group name of the current DAO.
	columns TopTenCirculatingHoldersColumns // columns contains all the column names of Table for convenient usage.
}

// TopTenCirculatingHoldersColumns defines and stores column names for the table hg_top_ten_circulating_holders.
type TopTenCirculatingHoldersColumns struct {
	Id            string // 自增主键
	Symbol        string // 公司代码/股票代码 (例如: 000001.SZ)
	Jzrq          string // 截止日期 (报告期结束日, 如2023-09-30)[citation:4]
	Ggrq          string // 公告日期 (信息发布日期)
	ReportYear    string // 报告年度
	ReportQuarter string // 报告季度 (1-4)
	ReportType    string // 报告类型: annual-年报, half_year-中报, quarter-季报[citation:4]
	Gdmc          string // 股东名称
	Gdlx          string // 股东类型 (如: 基金、社保、个人等)
	Gfxz          string // 股份性质 (如: 流通A股、限售A股等)
	Cgsl          string // 持股数量 (股)
	Cgbl          string // 持股比例 (%)
	Cgpm          string // 持股排名 (1-10)
	Bdyy          string // 变动原因
	BdType        string // 变动类型: increase-增持, decrease-减持, new-新进, unchanged-不变, other-其他
	DataSource    string // 数据来源 (如: 交易所公告)[citation:4]
	IsLatest      string // 是否为该报告期最新数据: 0-历史快照, 1-最新
	Version       string // 数据版本
	CreatedAt     string // 创建时间
	UpdatedAt     string // 更新时间
}

// topTenCirculatingHoldersColumns holds the columns for the table hg_top_ten_circulating_holders.
var topTenCirculatingHoldersColumns = TopTenCirculatingHoldersColumns{
	Id:            "id",
	Symbol:        "symbol",
	Jzrq:          "jzrq",
	Ggrq:          "ggrq",
	ReportYear:    "report_year",
	ReportQuarter: "report_quarter",
	ReportType:    "report_type",
	Gdmc:          "gdmc",
	Gdlx:          "gdlx",
	Gfxz:          "gfxz",
	Cgsl:          "cgsl",
	Cgbl:          "cgbl",
	Cgpm:          "cgpm",
	Bdyy:          "bdyy",
	BdType:        "bd_type",
	DataSource:    "data_source",
	IsLatest:      "is_latest",
	Version:       "version",
	CreatedAt:     "created_at",
	UpdatedAt:     "updated_at",
}

// NewTopTenCirculatingHoldersDao creates and returns a new DAO object for table data access.
func NewTopTenCirculatingHoldersDao() *TopTenCirculatingHoldersDao {
	return &TopTenCirculatingHoldersDao{
		group:   "default",
		table:   "hg_top_ten_circulating_holders",
		columns: topTenCirculatingHoldersColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *TopTenCirculatingHoldersDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *TopTenCirculatingHoldersDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *TopTenCirculatingHoldersDao) Columns() TopTenCirculatingHoldersColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *TopTenCirculatingHoldersDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *TopTenCirculatingHoldersDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rolls back the transaction and returns the error if function f returns a non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note: Do not commit or roll back the transaction in function f,
// as it is automatically handled by this function.
func (dao *TopTenCirculatingHoldersDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
