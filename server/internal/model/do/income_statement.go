// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// IncomeStatement is the golang structure of table hg_income_statement for DAO operations like Where/Data.
type IncomeStatement struct {
	g.Meta             `orm:"table:hg_income_statement, do:true"`
	Id                 any         // 自增主键
	Symbol             any         // 股票代码 (如: 000001.SZ)
	Jzrq               *gtime.Time // 截止日期 (报告期截止日，如2025-12-31)
	Plrq               *gtime.Time // 披露日期 (财报实际发布日期)
	ReportYear         any         // 报告年度
	ReportQuarter      any         // 报告季度 (1-4)
	ReportType         any         // 报告类型: annual-年报, quarter-季报
	FiscalPeriod       any         // 会计期间
	Yysr               any         // 营业收入
	Yzbf               any         // 已赚保费
	Fdczssr            any         // 房地产销售收入
	Qtywsr             any         // 其他业务收入
	Yyzsr              any         // 营业总收入
	Lxsr               any         // 利息收入
	Sxfjyjsr           any         // 手续费及佣金收入
	Btsr               any         // 补贴收入
	Ywsr               any         // 营业外收入
	Qtsy               any         // 其他收益
	Yycb               any         // 营业成本
	Fdczscb            any         // 房地产销售成本
	Qtywcb             any         // 其他业务成本
	Yyzcb              any         // 营业总成本
	Yysjjfj            any         // 营业税金及附加
	Xsfy               any         // 销售费用
	Glfy               any         // 管理费用
	Yffy               any         // 研发费用
	Cwfy               any         // 财务费用
	Sxfjyjzc           any         // 手续费及佣金支出
	Lxzc               any         // 利息支出
	Tbj                any         // 退保金
	Pczjje             any         // 赔付支出净额
	Tqbxhtzbjje        any         // 提取保险合同准备金净额
	Bdhlzc             any         // 保单红利支出
	Fbfy               any         // 分保费用
	Zcjzss             any         // 资产减值损失
	Ywzc               any         // 营业外支出
	Qtywlr             any         // 其他业务利润
	Yylr               any         // 营业利润
	Lrze               any         // 利润总额
	Jlr                any         // 净利润
	Jlrhfcjcx          any         // 净利润(扣除非经常性损益后)
	Gsmgsyzzdjlr       any         // 归属于母公司所有者的净利润
	Bhbfzhbqsljlr      any         // 被合并方在合并前实现净利润
	Tzsy               any         // 投资收益
	Lyqyhhhqydtzsy     any         // 联营企业和合营企业的投资收益
	Gyjzbdsy           any         // 公允价值变动收益
	Qhsy               any         // 期货损益
	Tgsy               any         // 托管收益
	Hdsy               any         // 汇兑收益
	Fldzcczsy          any         // 非流动资产处置收益
	Sdsfy              any         // 所得税费用
	Ssgdsy             any         // 少数股东损益
	Wqrtzss            any         // 未确认投资损失
	Jbmgsy             any         // 基本每股收益
	Xsmgsy             any         // 稀释每股收益
	Zhsyz              any         // 综合收益总额
	Gsssgdzhsyz        any         // 归属于少数股东的综合收益总额
	GrossMargin        any         // 毛利率(%)
	OperatingMargin    any         // 营业利润率(%)
	NetMargin          any         // 净利率(%)
	EffectiveTaxRate   any         // 实际税率(%)
	DataSource         any         // 数据来源
	Currency           any         // 货币单位
	Unit               any         // 单位: yuan-元, wan-万元
	AccountingStandard any         // 会计准则 (如: CAS, IFRS)
	IsAudited          any         // 是否审计: 0-未审计, 1-已审计
	IsConsolidated     any         // 是否合并报表: 1-合并, 0-母公司
	CreatedAt          *gtime.Time // 创建时间
	UpdatedAt          *gtime.Time // 更新时间
}
