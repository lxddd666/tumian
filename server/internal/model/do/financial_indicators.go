// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// FinancialIndicators is the golang structure of table hg_financial_indicators for DAO operations like Where/Data.
type FinancialIndicators struct {
	g.Meta    `orm:"table:hg_financial_indicators, do:true"`
	Id        any         // 主键ID
	Date      *gtime.Time // 报告日期 yyyy-MM-dd
	Symbol    any         // 股票或标的代码 (例如: AAPL, 000001.SZ)
	Tbmg      any         // 摊薄每股收益(元)
	Jqmg      any         // 加权每股收益(元)
	Mgsy      any         // 每股收益_调整后(元)
	Kfmg      any         // 扣除非经常性损益后的每股收益(元)
	Mgjz      any         // 每股净资产_调整前(元)
	Mgjzad    any         // 每股净资产_调整后(元)
	Mgjy      any         // 每股经营性现金流(元)
	Mggjj     any         // 每股资本公积金(元)
	Mgwly     any         // 每股未分配利润(元)
	Zclr      any         // 总资产利润率(%)
	Zylr      any         // 主营业务利润率(%)
	Zzlr      any         // 总资产净利润率(%)
	Cblr      any         // 成本费用利润率(%)
	Yylr      any         // 营业利润率(%)
	Zycb      any         // 主营业务成本率(%)
	Xsjl      any         // 销售净利率(%)
	Gbbc      any         // 股本报酬率(%)
	Jzbc      any         // 净资产报酬率(%)
	Zcbc      any         // 资产报酬率(%)
	Xsml      any         // 销售毛利率(%)
	Xxbz      any         // 三项费用比重
	Fzy       any         // 非主营比重
	Zybz      any         // 主营利润比重
	Gxff      any         // 股息发放率(%)
	Tzsy      any         // 投资收益率(%)
	Zyyw      any         // 主营业务利润(元)
	Jzsy      any         // 净资产收益率(%)
	Jqjz      any         // 加权净资产收益率(%)
	Kflr      any         // 扣除非经常性损益后的净利润(元)
	Zysr      any         // 主营业务收入增长率(%)
	Jlzz      any         // 净利润增长率(%)
	Jzzz      any         // 净资产增长率(%)
	Zzzz      any         // 总资产增长率(%)
	Yszz      any         // 应收账款周转率(次)
	Yszzt     any         // 应收账款周转天数(天)
	Chzz      any         // 存货周转天数(天)
	Chzzl     any         // 存货周转率(次)
	Gzzz      any         // 固定资产周转率(次)
	Zzzzl     any         // 总资产周转率(次)
	Zzzzt     any         // 总资产周转天数(天)
	Ldzz      any         // 流动资产周转率(次)
	Ldzzt     any         // 流动资产周转天数(天)
	Gdzz      any         // 股东权益周转率(次)
	Ldbl      any         // 流动比率
	Sdbl      any         // 速动比率
	Xjbl      any         // 现金比率(%)
	Lxzf      any         // 利息支付倍数
	Zjbl      any         // 长期债务与营运资金比率(%)
	Gdqy      any         // 股东权益比率(%)
	Cqfz      any         // 长期负债比率(%)
	Gdgd      any         // 股东权益与固定资产比率(%)
	Fzqy      any         // 负债与所有者权益比率(%)
	Zczjbl    any         // 长期资产与长期资金比率(%)
	Zblv      any         // 资本化比率(%)
	Gdzcjz    any         // 固定资产净值率(%)
	Zbgdh     any         // 资本固定化比率(%)
	Cqbl      any         // 产权比率(%)
	Qxjzb     any         // 清算价值比率(%)
	Gdzcbz    any         // 固定资产比重(%)
	Zcfzl     any         // 资产负债率(%)
	Zzc       any         // 总资产(元)
	Jyxj      any         // 经营现金净流量对销售收入比率(%)
	Zcjyxj    any         // 资产的经营现金流量回报率(%)
	Jylrb     any         // 经营现金净流量与净利润的比率(%)
	Jyfzl     any         // 经营现金净流量对负债比率(%)
	Xjlbl     any         // 现金流量比率(%)
	Dqgptz    any         // 短期股票投资(元)
	Dqzctz    any         // 短期债券投资(元)
	Dqjytz    any         // 短期其它经营性投资(元)
	Qcgptz    any         // 长期股票投资(元)
	Cqzqtz    any         // 长期债券投资(元)
	Cqjyxtz   any         // 长期其它经营性投资(元)
	Yszk1     any         // 1年以内应收帐款(元)
	Yszk12    any         // 1-2年以内应收帐款(元)
	Yszk23    any         // 2-3年以内应收帐款(元)
	Yszk3     any         // 3年以内应收帐款(元)
	Yfhk1     any         // 1年以内预付货款(元)
	Yfhk12    any         // 1-2年以内预付货款(元)
	Yfhk23    any         // 2-3年以内预付货款(元)
	Yfhk3     any         // 3年以内预付货款(元)
	Ysk1      any         // 1年以内其它应收款(元)
	Ysk12     any         // 1-2年以内其它应收款(元)
	Ysk23     any         // 2-3年以内其它应收款(元)
	Ysk3      any         // 3年以内其它应收款(元)
	CreatedAt *gtime.Time // 数据创建时间
	UpdatedAt *gtime.Time // 数据更新时间
}
