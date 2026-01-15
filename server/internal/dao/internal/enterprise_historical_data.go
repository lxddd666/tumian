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
	Id            string // 自增主键
	Symbol        string // 证券代码 (如: 000001.SZ, AAPL)
	T             string // 交易时间 (精确到分钟或日)
	Date          string // 交易日期 (衍生字段)
	Year          string // 交易年份
	Month         string // 交易月份
	Weekday       string // 星期几 (1=周日,7=周六)
	O             string // 开盘价
	H             string // 最高价
	L             string // 最低价
	C             string // 收盘价
	Pc            string // 前收盘价
	V             string // 成交量 (股/手)
	A             string // 成交额 (元)
	Change        string // 涨跌额
	ChangePct     string // 涨跌幅 (%)
	Amplitude     string // 振幅 (%)
	Sf            string // 停牌标志: 0-正常, 1-停牌
	TradingStatus string // 交易状态描述
	Period        string // 数据周期: 1min, 5min, 15min, 30min, 60min, day, week, month
	IsAdjusted    string // 是否复权: 0-不复权, 1-前复权, 2-后复权
	DataQuality   string // 数据质量: 0-异常, 1-正常, 2-补全
	IsVerified    string // 是否已验证: 0-未验证, 1-已验证
	DataSource    string // 数据来源
	Version       string // 数据版本
	CreatedAt     string // 创建时间
	UpdatedAt     string // 更新时间
}

// enterpriseHistoricalDataColumns holds the columns for the table hg_enterprise_historical_data.
var enterpriseHistoricalDataColumns = EnterpriseHistoricalDataColumns{
	Id:            "id",
	Symbol:        "symbol",
	T:             "t",
	Date:          "date",
	Year:          "year",
	Month:         "month",
	Weekday:       "weekday",
	O:             "o",
	H:             "h",
	L:             "l",
	C:             "c",
	Pc:            "pc",
	V:             "v",
	A:             "a",
	Change:        "change",
	ChangePct:     "change_pct",
	Amplitude:     "amplitude",
	Sf:            "sf",
	TradingStatus: "trading_status",
	Period:        "period",
	IsAdjusted:    "is_adjusted",
	DataQuality:   "data_quality",
	IsVerified:    "is_verified",
	DataSource:    "data_source",
	Version:       "version",
	CreatedAt:     "created_at",
	UpdatedAt:     "updated_at",
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
