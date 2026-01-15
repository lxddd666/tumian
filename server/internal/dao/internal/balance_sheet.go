// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// BalanceSheetDao is the data access object for the table hg_balance_sheet.
type BalanceSheetDao struct {
	table    string              // table is the underlying table name of the DAO.
	group    string              // group is the database configuration group name of the current DAO.
	columns  BalanceSheetColumns // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler  // handlers for customized model modification.
}

// BalanceSheetColumns defines and stores column names for the table hg_balance_sheet.
type BalanceSheetColumns struct {
	Id            string // 自增主键
	Symbol        string // 公司代码/股票代码 (例如: 000001.SZ, AAPL)
	Jzrq          string // 截止日期 (会计期间结束日)
	Plrq          string // 披露日期
	ReportYear    string // 报告年度
	ReportQuarter string // 报告季度 (1-4, 年报为NULL)
	ReportType    string // 报告类型: annual-年报, quarter-季报, interim-中报
	Hbzj          string // 货币资金
	Jyxjrzc       string // 交易性金融资产
	Yspj          string // 应收票据
	Yszk          string // 应收账款
	Yfkx          string // 预付款项
	Yslx          string // 应收利息
	Ysgl          string // 应收股利
	Qtysk         string // 其他应收款
	Ch            string // 存货
	Dfy           string // 待摊费用
	Ynndqdfldzc   string // 一年内到期的非流动资产
	Qtldzc        string // 其他流动资产
	Ldzchj        string // 流动资产合计
	Cqgqtz        string // 长期股权投资
	Cqysk         string // 长期应收款
	Gdzc          string // 固定资产
	Zjgc          string // 在建工程
	Wxzc          string // 无形资产
	Sy            string // 商誉
	Cqdtfy        string // 长期待摊费用
	Dysdszc       string // 递延所得税资产
	Qtfldzc       string // 其他非流动资产
	Fldzchj       string // 非流动资产合计
	Zczj          string // 资产总计
	Dqjk          string // 短期借款
	Jyxjrfz       string // 交易性金融负债
	Yfpj          string // 应付票据
	Yfzk          string // 应付账款
	Ysk           string // 预收账款
	Yfgzxc        string // 应付职工薪酬
	Yjsf          string // 应交税费
	Yflx          string // 应付利息
	Yfgl          string // 应付股利
	Qtfzk         string // 其他应付款
	Ynndqdfldfz   string // 一年内到期的非流动负债
	Qtldfz        string // 其他流动负债
	Ldfzhj        string // 流动负债合计
	Cqjk          string // 长期借款
	Yfzq          string // 应付债券
	Cqyfk         string // 长期应付款
	Dysdsfz       string // 递延所得税负债
	Qtfldfz       string // 其他非流动负债
	Fldfzhj       string // 非流动负债合计
	Fzhj          string // 负债合计
	Sszb          string // 实收资本(或股本)
	Zbgj          string // 资本公积
	Ylgj          string // 盈余公积
	Wfplr         string // 未分配利润
	Gsmgdqsyhj    string // 归属于母公司股东权益合计
	Ssgdqy        string // 少数股东权益
	Syzqyhj       string // 所有者权益合计
	Fzhgdqyzj     string // 负债和股东权益总计
	DataSource    string // 数据来源
	Currency      string // 货币单位 (CNY, USD等)
	Unit          string // 单位: yuan-元, wan-万元, qianwan-千万元
	IsAudited     string // 是否审计: 0-未审计, 1-已审计
	Version       string // 数据版本
	CreatedAt     string // 创建时间
	UpdatedAt     string // 更新时间
	CreatedBy     string // 创建人
	UpdatedBy     string // 更新人
}

// balanceSheetColumns holds the columns for the table hg_balance_sheet.
var balanceSheetColumns = BalanceSheetColumns{
	Id:            "id",
	Symbol:        "symbol",
	Jzrq:          "jzrq",
	Plrq:          "plrq",
	ReportYear:    "report_year",
	ReportQuarter: "report_quarter",
	ReportType:    "report_type",
	Hbzj:          "hbzj",
	Jyxjrzc:       "jyxjrzc",
	Yspj:          "yspj",
	Yszk:          "yszk",
	Yfkx:          "yfkx",
	Yslx:          "yslx",
	Ysgl:          "ysgl",
	Qtysk:         "qtysk",
	Ch:            "ch",
	Dfy:           "dfy",
	Ynndqdfldzc:   "ynndqdfldzc",
	Qtldzc:        "qtldzc",
	Ldzchj:        "ldzchj",
	Cqgqtz:        "cqgqtz",
	Cqysk:         "cqysk",
	Gdzc:          "gdzc",
	Zjgc:          "zjgc",
	Wxzc:          "wxzc",
	Sy:            "sy",
	Cqdtfy:        "cqdtfy",
	Dysdszc:       "dysdszc",
	Qtfldzc:       "qtfldzc",
	Fldzchj:       "fldzchj",
	Zczj:          "zczj",
	Dqjk:          "dqjk",
	Jyxjrfz:       "jyxjrfz",
	Yfpj:          "yfpj",
	Yfzk:          "yfzk",
	Ysk:           "ysk",
	Yfgzxc:        "yfgzxc",
	Yjsf:          "yjsf",
	Yflx:          "yflx",
	Yfgl:          "yfgl",
	Qtfzk:         "qtfzk",
	Ynndqdfldfz:   "ynndqdfldfz",
	Qtldfz:        "qtldfz",
	Ldfzhj:        "ldfzhj",
	Cqjk:          "cqjk",
	Yfzq:          "yfzq",
	Cqyfk:         "cqyfk",
	Dysdsfz:       "dysdsfz",
	Qtfldfz:       "qtfldfz",
	Fldfzhj:       "fldfzhj",
	Fzhj:          "fzhj",
	Sszb:          "sszb",
	Zbgj:          "zbgj",
	Ylgj:          "ylgj",
	Wfplr:         "wfplr",
	Gsmgdqsyhj:    "gsmgdqsyhj",
	Ssgdqy:        "ssgdqy",
	Syzqyhj:       "syzqyhj",
	Fzhgdqyzj:     "fzhgdqyzj",
	DataSource:    "data_source",
	Currency:      "currency",
	Unit:          "unit",
	IsAudited:     "is_audited",
	Version:       "version",
	CreatedAt:     "created_at",
	UpdatedAt:     "updated_at",
	CreatedBy:     "created_by",
	UpdatedBy:     "updated_by",
}

// NewBalanceSheetDao creates and returns a new DAO object for table data access.
func NewBalanceSheetDao(handlers ...gdb.ModelHandler) *BalanceSheetDao {
	return &BalanceSheetDao{
		group:    "default",
		table:    "hg_balance_sheet",
		columns:  balanceSheetColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *BalanceSheetDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *BalanceSheetDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *BalanceSheetDao) Columns() BalanceSheetColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *BalanceSheetDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *BalanceSheetDao) Ctx(ctx context.Context) *gdb.Model {
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
func (dao *BalanceSheetDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
