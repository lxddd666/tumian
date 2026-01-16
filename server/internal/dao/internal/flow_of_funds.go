// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// FlowOfFundsDao is the data access object for the table hg_flow_of_funds.
type FlowOfFundsDao struct {
	table   string             // table is the underlying table name of the DAO.
	group   string             // group is the database configuration group name of the current DAO.
	columns FlowOfFundsColumns // columns contains all the column names of Table for convenient usage.
}

// FlowOfFundsColumns defines and stores column names for the table hg_flow_of_funds.
type FlowOfFundsColumns struct {
	Id         string // 主键ID
	Symbol     string // 股票代码 (如: 000001.SZ)
	T          string // 交易时间 (通常为HHMMSS格式的整数)
	Zmbzds     string // 主买单总单数
	Zmszds     string // 主卖单总单数
	Dddx       string // 大单动向
	Zddy       string // 涨跌动因
	Ddcf       string // 大单差分
	Zmbzdszl   string // 主买单总单数增量
	Zmszdszl   string // 主卖单总单数增量
	Cjbszl     string // 成交笔数增量
	Zmbtdcje   string // 主买特大单成交额
	Zmbddcje   string // 主买大单成交额
	Zmbzdcje   string // 主买中单成交额
	Zmbxdcje   string // 主买小单成交额
	Zmstdcje   string // 主卖特大单成交额
	Zmsddcje   string // 主卖大单成交额
	Zmszdcje   string // 主卖中单成交额
	Zmsxdcje   string // 主卖小单成交额
	Bdmbtdcje  string // 被动买特大单成交额
	Bdmbddcje  string // 被动买大单成交额
	Bdmbzdcje  string // 被动买中单成交额
	Bdmbxdcje  string // 被动买小单成交额
	Bdmstdcje  string // 被动卖特大单成交额
	Bdmsddcje  string // 被动卖大单成交额
	Bdmszdcje  string // 被动卖中单成交额
	Bdmsxdcje  string // 被动卖小单成交额
	Zmbtdcjl   string // 主买特大单成交量
	Zmbddcjl   string // 主买大单成交量
	Zmbzdcjl   string // 主买中单成交量
	Zmbxdcjl   string // 主买小单成交量
	Zmstdcjl   string // 主卖特大单成交量
	Zmsddcjl   string // 主卖大单成交量
	Zmszdcjl   string // 主卖中单成交量
	Zmsxdcjl   string // 主卖小单成交量
	Bdmbtdcjl  string // 被动买特大单成交量
	Bdmbddcjl  string // 被动买大单成交量
	Bdmbzdcjl  string // 被动买中单成交量
	Bdmbxdcjl  string // 被动买小单成交量
	Bdmstdcjl  string // 被动卖特大单成交量
	Bdmsddcjl  string // 被动卖大单成交量
	Bdmszdcjl  string // 被动卖中单成交量
	Bdmsxdcjl  string // 被动卖小单成交量
	Zmbtdcjzl  string // 主买特大单成交额增量
	Zmbddcjzl  string // 主买大单成交额增量
	Zmbtdcjzlv string // 主买特大单成交量增量
	Zmbddcjzlv string // 主买大单成交量增量
	CreatedAt  string // 创建时间
	UpdatedAt  string // 更新时间
}

// flowOfFundsColumns holds the columns for the table hg_flow_of_funds.
var flowOfFundsColumns = FlowOfFundsColumns{
	Id:         "id",
	Symbol:     "symbol",
	T:          "t",
	Zmbzds:     "zmbzds",
	Zmszds:     "zmszds",
	Dddx:       "dddx",
	Zddy:       "zddy",
	Ddcf:       "ddcf",
	Zmbzdszl:   "zmbzdszl",
	Zmszdszl:   "zmszdszl",
	Cjbszl:     "cjbszl",
	Zmbtdcje:   "zmbtdcje",
	Zmbddcje:   "zmbddcje",
	Zmbzdcje:   "zmbzdcje",
	Zmbxdcje:   "zmbxdcje",
	Zmstdcje:   "zmstdcje",
	Zmsddcje:   "zmsddcje",
	Zmszdcje:   "zmszdcje",
	Zmsxdcje:   "zmsxdcje",
	Bdmbtdcje:  "bdmbtdcje",
	Bdmbddcje:  "bdmbddcje",
	Bdmbzdcje:  "bdmbzdcje",
	Bdmbxdcje:  "bdmbxdcje",
	Bdmstdcje:  "bdmstdcje",
	Bdmsddcje:  "bdmsddcje",
	Bdmszdcje:  "bdmszdcje",
	Bdmsxdcje:  "bdmsxdcje",
	Zmbtdcjl:   "zmbtdcjl",
	Zmbddcjl:   "zmbddcjl",
	Zmbzdcjl:   "zmbzdcjl",
	Zmbxdcjl:   "zmbxdcjl",
	Zmstdcjl:   "zmstdcjl",
	Zmsddcjl:   "zmsddcjl",
	Zmszdcjl:   "zmszdcjl",
	Zmsxdcjl:   "zmsxdcjl",
	Bdmbtdcjl:  "bdmbtdcjl",
	Bdmbddcjl:  "bdmbddcjl",
	Bdmbzdcjl:  "bdmbzdcjl",
	Bdmbxdcjl:  "bdmbxdcjl",
	Bdmstdcjl:  "bdmstdcjl",
	Bdmsddcjl:  "bdmsddcjl",
	Bdmszdcjl:  "bdmszdcjl",
	Bdmsxdcjl:  "bdmsxdcjl",
	Zmbtdcjzl:  "zmbtdcjzl",
	Zmbddcjzl:  "zmbddcjzl",
	Zmbtdcjzlv: "zmbtdcjzlv",
	Zmbddcjzlv: "zmbddcjzlv",
	CreatedAt:  "created_at",
	UpdatedAt:  "updated_at",
}

// NewFlowOfFundsDao creates and returns a new DAO object for table data access.
func NewFlowOfFundsDao() *FlowOfFundsDao {
	return &FlowOfFundsDao{
		group:   "default",
		table:   "hg_flow_of_funds",
		columns: flowOfFundsColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *FlowOfFundsDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *FlowOfFundsDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *FlowOfFundsDao) Columns() FlowOfFundsColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *FlowOfFundsDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *FlowOfFundsDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rolls back the transaction and returns the error if function f returns a non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note: Do not commit or roll back the transaction in function f,
// as it is automatically handled by this function.
func (dao *FlowOfFundsDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
