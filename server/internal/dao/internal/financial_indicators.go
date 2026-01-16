// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// FinancialIndicatorsDao is the data access object for the table hg_financial_indicators.
type FinancialIndicatorsDao struct {
	table    string                     // table is the underlying table name of the DAO.
	group    string                     // group is the database configuration group name of the current DAO.
	columns  FinancialIndicatorsColumns // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler         // handlers for customized model modification.
}

// FinancialIndicatorsColumns defines and stores column names for the table hg_financial_indicators.
type FinancialIndicatorsColumns struct {
	Id        string // 主键ID
	Date      string // 报告日期 yyyy-MM-dd
	Symbol    string // 股票或标的代码 (例如: AAPL, 000001.SZ)
	Tbmg      string // 摊薄每股收益(元)
	Jqmg      string // 加权每股收益(元)
	Mgsy      string // 每股收益_调整后(元)
	Kfmg      string // 扣除非经常性损益后的每股收益(元)
	Mgjz      string // 每股净资产_调整前(元)
	Mgjzad    string // 每股净资产_调整后(元)
	Mgjy      string // 每股经营性现金流(元)
	Mggjj     string // 每股资本公积金(元)
	Mgwly     string // 每股未分配利润(元)
	Zclr      string // 总资产利润率(%)
	Zylr      string // 主营业务利润率(%)
	Zzlr      string // 总资产净利润率(%)
	Cblr      string // 成本费用利润率(%)
	Yylr      string // 营业利润率(%)
	Zycb      string // 主营业务成本率(%)
	Xsjl      string // 销售净利率(%)
	Gbbc      string // 股本报酬率(%)
	Jzbc      string // 净资产报酬率(%)
	Zcbc      string // 资产报酬率(%)
	Xsml      string // 销售毛利率(%)
	Xxbz      string // 三项费用比重
	Fzy       string // 非主营比重
	Zybz      string // 主营利润比重
	Gxff      string // 股息发放率(%)
	Tzsy      string // 投资收益率(%)
	Zyyw      string // 主营业务利润(元)
	Jzsy      string // 净资产收益率(%)
	Jqjz      string // 加权净资产收益率(%)
	Kflr      string // 扣除非经常性损益后的净利润(元)
	Zysr      string // 主营业务收入增长率(%)
	Jlzz      string // 净利润增长率(%)
	Jzzz      string // 净资产增长率(%)
	Zzzz      string // 总资产增长率(%)
	Yszz      string // 应收账款周转率(次)
	Yszzt     string // 应收账款周转天数(天)
	Chzz      string // 存货周转天数(天)
	Chzzl     string // 存货周转率(次)
	Gzzz      string // 固定资产周转率(次)
	Zzzzl     string // 总资产周转率(次)
	Zzzzt     string // 总资产周转天数(天)
	Ldzz      string // 流动资产周转率(次)
	Ldzzt     string // 流动资产周转天数(天)
	Gdzz      string // 股东权益周转率(次)
	Ldbl      string // 流动比率
	Sdbl      string // 速动比率
	Xjbl      string // 现金比率(%)
	Lxzf      string // 利息支付倍数
	Zjbl      string // 长期债务与营运资金比率(%)
	Gdqy      string // 股东权益比率(%)
	Cqfz      string // 长期负债比率(%)
	Gdgd      string // 股东权益与固定资产比率(%)
	Fzqy      string // 负债与所有者权益比率(%)
	Zczjbl    string // 长期资产与长期资金比率(%)
	Zblv      string // 资本化比率(%)
	Gdzcjz    string // 固定资产净值率(%)
	Zbgdh     string // 资本固定化比率(%)
	Cqbl      string // 产权比率(%)
	Qxjzb     string // 清算价值比率(%)
	Gdzcbz    string // 固定资产比重(%)
	Zcfzl     string // 资产负债率(%)
	Zzc       string // 总资产(元)
	Jyxj      string // 经营现金净流量对销售收入比率(%)
	Zcjyxj    string // 资产的经营现金流量回报率(%)
	Jylrb     string // 经营现金净流量与净利润的比率(%)
	Jyfzl     string // 经营现金净流量对负债比率(%)
	Xjlbl     string // 现金流量比率(%)
	Dqgptz    string // 短期股票投资(元)
	Dqzctz    string // 短期债券投资(元)
	Dqjytz    string // 短期其它经营性投资(元)
	Qcgptz    string // 长期股票投资(元)
	Cqzqtz    string // 长期债券投资(元)
	Cqjyxtz   string // 长期其它经营性投资(元)
	Yszk1     string // 1年以内应收帐款(元)
	Yszk12    string // 1-2年以内应收帐款(元)
	Yszk23    string // 2-3年以内应收帐款(元)
	Yszk3     string // 3年以内应收帐款(元)
	Yfhk1     string // 1年以内预付货款(元)
	Yfhk12    string // 1-2年以内预付货款(元)
	Yfhk23    string // 2-3年以内预付货款(元)
	Yfhk3     string // 3年以内预付货款(元)
	Ysk1      string // 1年以内其它应收款(元)
	Ysk12     string // 1-2年以内其它应收款(元)
	Ysk23     string // 2-3年以内其它应收款(元)
	Ysk3      string // 3年以内其它应收款(元)
	CreatedAt string // 数据创建时间
	UpdatedAt string // 数据更新时间
}

// financialIndicatorsColumns holds the columns for the table hg_financial_indicators.
var financialIndicatorsColumns = FinancialIndicatorsColumns{
	Id:        "id",
	Date:      "date",
	Symbol:    "symbol",
	Tbmg:      "tbmg",
	Jqmg:      "jqmg",
	Mgsy:      "mgsy",
	Kfmg:      "kfmg",
	Mgjz:      "mgjz",
	Mgjzad:    "mgjzad",
	Mgjy:      "mgjy",
	Mggjj:     "mggjj",
	Mgwly:     "mgwly",
	Zclr:      "zclr",
	Zylr:      "zylr",
	Zzlr:      "zzlr",
	Cblr:      "cblr",
	Yylr:      "yylr",
	Zycb:      "zycb",
	Xsjl:      "xsjl",
	Gbbc:      "gbbc",
	Jzbc:      "jzbc",
	Zcbc:      "zcbc",
	Xsml:      "xsml",
	Xxbz:      "xxbz",
	Fzy:       "fzy",
	Zybz:      "zybz",
	Gxff:      "gxff",
	Tzsy:      "tzsy",
	Zyyw:      "zyyw",
	Jzsy:      "jzsy",
	Jqjz:      "jqjz",
	Kflr:      "kflr",
	Zysr:      "zysr",
	Jlzz:      "jlzz",
	Jzzz:      "jzzz",
	Zzzz:      "zzzz",
	Yszz:      "yszz",
	Yszzt:     "yszzt",
	Chzz:      "chzz",
	Chzzl:     "chzzl",
	Gzzz:      "gzzz",
	Zzzzl:     "zzzzl",
	Zzzzt:     "zzzzt",
	Ldzz:      "ldzz",
	Ldzzt:     "ldzzt",
	Gdzz:      "gdzz",
	Ldbl:      "ldbl",
	Sdbl:      "sdbl",
	Xjbl:      "xjbl",
	Lxzf:      "lxzf",
	Zjbl:      "zjbl",
	Gdqy:      "gdqy",
	Cqfz:      "cqfz",
	Gdgd:      "gdgd",
	Fzqy:      "fzqy",
	Zczjbl:    "zczjbl",
	Zblv:      "zblv",
	Gdzcjz:    "gdzcjz",
	Zbgdh:     "zbgdh",
	Cqbl:      "cqbl",
	Qxjzb:     "qxjzb",
	Gdzcbz:    "gdzcbz",
	Zcfzl:     "zcfzl",
	Zzc:       "zzc",
	Jyxj:      "jyxj",
	Zcjyxj:    "zcjyxj",
	Jylrb:     "jylrb",
	Jyfzl:     "jyfzl",
	Xjlbl:     "xjlbl",
	Dqgptz:    "dqgptz",
	Dqzctz:    "dqzctz",
	Dqjytz:    "dqjytz",
	Qcgptz:    "qcgptz",
	Cqzqtz:    "cqzqtz",
	Cqjyxtz:   "cqjyxtz",
	Yszk1:     "yszk1",
	Yszk12:    "yszk12",
	Yszk23:    "yszk23",
	Yszk3:     "yszk3",
	Yfhk1:     "yfhk1",
	Yfhk12:    "yfhk12",
	Yfhk23:    "yfhk23",
	Yfhk3:     "yfhk3",
	Ysk1:      "ysk1",
	Ysk12:     "ysk12",
	Ysk23:     "ysk23",
	Ysk3:      "ysk3",
	CreatedAt: "created_at",
	UpdatedAt: "updated_at",
}

// NewFinancialIndicatorsDao creates and returns a new DAO object for table data access.
func NewFinancialIndicatorsDao(handlers ...gdb.ModelHandler) *FinancialIndicatorsDao {
	return &FinancialIndicatorsDao{
		group:    "default",
		table:    "hg_financial_indicators",
		columns:  financialIndicatorsColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *FinancialIndicatorsDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *FinancialIndicatorsDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *FinancialIndicatorsDao) Columns() FinancialIndicatorsColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *FinancialIndicatorsDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *FinancialIndicatorsDao) Ctx(ctx context.Context) *gdb.Model {
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
func (dao *FinancialIndicatorsDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
