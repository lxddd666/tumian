// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// MaDataDao is the data access object for the table hg_ma_data.
type MaDataDao struct {
	table   string        // table is the underlying table name of the DAO.
	group   string        // group is the database configuration group name of the current DAO.
	columns MaDataColumns // columns contains all the column names of Table for convenient usage.
}

// MaDataColumns defines and stores column names for the table hg_ma_data.
type MaDataColumns struct {
	Id           string // 自增主键
	Symbol       string // 股票或标的代码 (例如: AAPL, 000001.SZ)
	T            string // 交易时间 (统一为datetime类型，日线数据时间部分设为00:00:00)
	IntervalType string // 数据间隔: minute-短分时, day-日线
	Ma3          string // MA3值
	Ma5          string // MA5值
	Ma10         string // MA10值
	Ma15         string // MA15值
	Ma20         string // MA20值
	Ma30         string // MA30值
	Ma60         string // MA60值
	Ma120        string // MA120值
	Ma200        string // MA200值
	Ma250        string // MA250值
	CreatedAt    string // 数据创建时间
	UpdatedAt    string // 数据更新时间
}

// maDataColumns holds the columns for the table hg_ma_data.
var maDataColumns = MaDataColumns{
	Id:           "id",
	Symbol:       "symbol",
	T:            "t",
	IntervalType: "interval_type",
	Ma3:          "ma3",
	Ma5:          "ma5",
	Ma10:         "ma10",
	Ma15:         "ma15",
	Ma20:         "ma20",
	Ma30:         "ma30",
	Ma60:         "ma60",
	Ma120:        "ma120",
	Ma200:        "ma200",
	Ma250:        "ma250",
	CreatedAt:    "created_at",
	UpdatedAt:    "updated_at",
}

// NewMaDataDao creates and returns a new DAO object for table data access.
func NewMaDataDao() *MaDataDao {
	return &MaDataDao{
		group:   "default",
		table:   "hg_ma_data",
		columns: maDataColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *MaDataDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *MaDataDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *MaDataDao) Columns() MaDataColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *MaDataDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *MaDataDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rolls back the transaction and returns the error if function f returns a non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note: Do not commit or roll back the transaction in function f,
// as it is automatically handled by this function.
func (dao *MaDataDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
