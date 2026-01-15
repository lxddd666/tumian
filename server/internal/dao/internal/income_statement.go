// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// IncomeStatementDao is the data access object for the table hg_income_statement.
type IncomeStatementDao struct {
	table    string                 // table is the underlying table name of the DAO.
	group    string                 // group is the database configuration group name of the current DAO.
	columns  IncomeStatementColumns // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler     // handlers for customized model modification.
}

// IncomeStatementColumns defines and stores column names for the table hg_income_statement.
type IncomeStatementColumns struct {
	Id                 string // 自增主键
	Symbol             string // 股票代码 (如: 000001.SZ)
	Jzrq               string // 截止日期 (报告期截止日，如2025-12-31)
	Plrq               string // 披露日期 (财报实际发布日期)
	ReportYear         string // 报告年度
	ReportQuarter      string // 报告季度 (1-4)
	ReportType         string // 报告类型: annual-年报, quarter-季报
	FiscalPeriod       string // 会计期间
	Yysr               string // 营业收入
	Yzbf               string // 已赚保费
	Fdczssr            string // 房地产销售收入
	Qtywsr             string // 其他业务收入
	Yyzsr              string // 营业总收入
	Lxsr               string // 利息收入
	Sxfjyjsr           string // 手续费及佣金收入
	Btsr               string // 补贴收入
	Ywsr               string // 营业外收入
	Qtsy               string // 其他收益
	Yycb               string // 营业成本
	Fdczscb            string // 房地产销售成本
	Qtywcb             string // 其他业务成本
	Yyzcb              string // 营业总成本
	Yysjjfj            string // 营业税金及附加
	Xsfy               string // 销售费用
	Glfy               string // 管理费用
	Yffy               string // 研发费用
	Cwfy               string // 财务费用
	Sxfjyjzc           string // 手续费及佣金支出
	Lxzc               string // 利息支出
	Tbj                string // 退保金
	Pczjje             string // 赔付支出净额
	Tqbxhtzbjje        string // 提取保险合同准备金净额
	Bdhlzc             string // 保单红利支出
	Fbfy               string // 分保费用
	Zcjzss             string // 资产减值损失
	Ywzc               string // 营业外支出
	Qtywlr             string // 其他业务利润
	Yylr               string // 营业利润
	Lrze               string // 利润总额
	Jlr                string // 净利润
	Jlrhfcjcx          string // 净利润(扣除非经常性损益后)
	Gsmgsyzzdjlr       string // 归属于母公司所有者的净利润
	Bhbfzhbqsljlr      string // 被合并方在合并前实现净利润
	Tzsy               string // 投资收益
	Lyqyhhhqydtzsy     string // 联营企业和合营企业的投资收益
	Gyjzbdsy           string // 公允价值变动收益
	Qhsy               string // 期货损益
	Tgsy               string // 托管收益
	Hdsy               string // 汇兑收益
	Fldzcczsy          string // 非流动资产处置收益
	Sdsfy              string // 所得税费用
	Ssgdsy             string // 少数股东损益
	Wqrtzss            string // 未确认投资损失
	Jbmgsy             string // 基本每股收益
	Xsmgsy             string // 稀释每股收益
	Zhsyz              string // 综合收益总额
	Gsssgdzhsyz        string // 归属于少数股东的综合收益总额
	GrossMargin        string // 毛利率(%)
	OperatingMargin    string // 营业利润率(%)
	NetMargin          string // 净利率(%)
	EffectiveTaxRate   string // 实际税率(%)
	DataSource         string // 数据来源
	Currency           string // 货币单位
	Unit               string // 单位: yuan-元, wan-万元
	AccountingStandard string // 会计准则 (如: CAS, IFRS)
	IsAudited          string // 是否审计: 0-未审计, 1-已审计
	IsConsolidated     string // 是否合并报表: 1-合并, 0-母公司
	CreatedAt          string // 创建时间
	UpdatedAt          string // 更新时间
}

// incomeStatementColumns holds the columns for the table hg_income_statement.
var incomeStatementColumns = IncomeStatementColumns{
	Id:                 "id",
	Symbol:             "symbol",
	Jzrq:               "jzrq",
	Plrq:               "plrq",
	ReportYear:         "report_year",
	ReportQuarter:      "report_quarter",
	ReportType:         "report_type",
	FiscalPeriod:       "fiscal_period",
	Yysr:               "yysr",
	Yzbf:               "yzbf",
	Fdczssr:            "fdczssr",
	Qtywsr:             "qtywsr",
	Yyzsr:              "yyzsr",
	Lxsr:               "lxsr",
	Sxfjyjsr:           "sxfjyjsr",
	Btsr:               "btsr",
	Ywsr:               "ywsr",
	Qtsy:               "qtsy",
	Yycb:               "yycb",
	Fdczscb:            "fdczscb",
	Qtywcb:             "qtywcb",
	Yyzcb:              "yyzcb",
	Yysjjfj:            "yysjjfj",
	Xsfy:               "xsfy",
	Glfy:               "glfy",
	Yffy:               "yffy",
	Cwfy:               "cwfy",
	Sxfjyjzc:           "sxfjyjzc",
	Lxzc:               "lxzc",
	Tbj:                "tbj",
	Pczjje:             "pczjje",
	Tqbxhtzbjje:        "tqbxhtzbjje",
	Bdhlzc:             "bdhlzc",
	Fbfy:               "fbfy",
	Zcjzss:             "zcjzss",
	Ywzc:               "ywzc",
	Qtywlr:             "qtywlr",
	Yylr:               "yylr",
	Lrze:               "lrze",
	Jlr:                "jlr",
	Jlrhfcjcx:          "jlrhfcjcx",
	Gsmgsyzzdjlr:       "gsmgsyzzdjlr",
	Bhbfzhbqsljlr:      "bhbfzhbqsljlr",
	Tzsy:               "tzsy",
	Lyqyhhhqydtzsy:     "lyqyhhhqydtzsy",
	Gyjzbdsy:           "gyjzbdsy",
	Qhsy:               "qhsy",
	Tgsy:               "tgsy",
	Hdsy:               "hdsy",
	Fldzcczsy:          "fldzcczsy",
	Sdsfy:              "sdsfy",
	Ssgdsy:             "ssgdsy",
	Wqrtzss:            "wqrtzss",
	Jbmgsy:             "jbmgsy",
	Xsmgsy:             "xsmgsy",
	Zhsyz:              "zhsyz",
	Gsssgdzhsyz:        "gsssgdzhsyz",
	GrossMargin:        "gross_margin",
	OperatingMargin:    "operating_margin",
	NetMargin:          "net_margin",
	EffectiveTaxRate:   "effective_tax_rate",
	DataSource:         "data_source",
	Currency:           "currency",
	Unit:               "unit",
	AccountingStandard: "accounting_standard",
	IsAudited:          "is_audited",
	IsConsolidated:     "is_consolidated",
	CreatedAt:          "created_at",
	UpdatedAt:          "updated_at",
}

// NewIncomeStatementDao creates and returns a new DAO object for table data access.
func NewIncomeStatementDao(handlers ...gdb.ModelHandler) *IncomeStatementDao {
	return &IncomeStatementDao{
		group:    "default",
		table:    "hg_income_statement",
		columns:  incomeStatementColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *IncomeStatementDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *IncomeStatementDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *IncomeStatementDao) Columns() IncomeStatementColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *IncomeStatementDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *IncomeStatementDao) Ctx(ctx context.Context) *gdb.Model {
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
func (dao *IncomeStatementDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
