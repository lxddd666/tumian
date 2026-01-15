// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// BalanceSheet is the golang structure of table hg_balance_sheet for DAO operations like Where/Data.
type BalanceSheet struct {
	g.Meta        `orm:"table:hg_balance_sheet, do:true"`
	Id            any         // 自增主键
	Symbol        any         // 公司代码/股票代码 (例如: 000001.SZ, AAPL)
	Jzrq          *gtime.Time // 截止日期 (会计期间结束日)
	Plrq          *gtime.Time // 披露日期
	ReportYear    any         // 报告年度
	ReportQuarter any         // 报告季度 (1-4, 年报为NULL)
	ReportType    any         // 报告类型: annual-年报, quarter-季报, interim-中报
	Hbzj          any         // 货币资金
	Jyxjrzc       any         // 交易性金融资产
	Yspj          any         // 应收票据
	Yszk          any         // 应收账款
	Yfkx          any         // 预付款项
	Yslx          any         // 应收利息
	Ysgl          any         // 应收股利
	Qtysk         any         // 其他应收款
	Ch            any         // 存货
	Dfy           any         // 待摊费用
	Ynndqdfldzc   any         // 一年内到期的非流动资产
	Qtldzc        any         // 其他流动资产
	Ldzchj        any         // 流动资产合计
	Cqgqtz        any         // 长期股权投资
	Cqysk         any         // 长期应收款
	Gdzc          any         // 固定资产
	Zjgc          any         // 在建工程
	Wxzc          any         // 无形资产
	Sy            any         // 商誉
	Cqdtfy        any         // 长期待摊费用
	Dysdszc       any         // 递延所得税资产
	Qtfldzc       any         // 其他非流动资产
	Fldzchj       any         // 非流动资产合计
	Zczj          any         // 资产总计
	Dqjk          any         // 短期借款
	Jyxjrfz       any         // 交易性金融负债
	Yfpj          any         // 应付票据
	Yfzk          any         // 应付账款
	Ysk           any         // 预收账款
	Yfgzxc        any         // 应付职工薪酬
	Yjsf          any         // 应交税费
	Yflx          any         // 应付利息
	Yfgl          any         // 应付股利
	Qtfzk         any         // 其他应付款
	Ynndqdfldfz   any         // 一年内到期的非流动负债
	Qtldfz        any         // 其他流动负债
	Ldfzhj        any         // 流动负债合计
	Cqjk          any         // 长期借款
	Yfzq          any         // 应付债券
	Cqyfk         any         // 长期应付款
	Dysdsfz       any         // 递延所得税负债
	Qtfldfz       any         // 其他非流动负债
	Fldfzhj       any         // 非流动负债合计
	Fzhj          any         // 负债合计
	Sszb          any         // 实收资本(或股本)
	Zbgj          any         // 资本公积
	Ylgj          any         // 盈余公积
	Wfplr         any         // 未分配利润
	Gsmgdqsyhj    any         // 归属于母公司股东权益合计
	Ssgdqy        any         // 少数股东权益
	Syzqyhj       any         // 所有者权益合计
	Fzhgdqyzj     any         // 负债和股东权益总计
	DataSource    any         // 数据来源
	Currency      any         // 货币单位 (CNY, USD等)
	Unit          any         // 单位: yuan-元, wan-万元, qianwan-千万元
	IsAudited     any         // 是否审计: 0-未审计, 1-已审计
	Version       any         // 数据版本
	CreatedAt     *gtime.Time // 创建时间
	UpdatedAt     *gtime.Time // 更新时间
	CreatedBy     any         // 创建人
	UpdatedBy     any         // 更新人
}
