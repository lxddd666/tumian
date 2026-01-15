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
	g.Meta         `orm:"table:hg_financial_indicators, do:true"`
	Id             any         // 自增主键
	Symbol         any         // 公司代码/股票代码 (例如: 000001.SZ, AAPL)
	Jzrq           *gtime.Time // 截止日期 (报告期结束日)
	Plrq           *gtime.Time // 披露日期
	ReportYear     any         // 报告年度
	ReportQuarter  any         // 报告季度 (1-4)
	ReportType     any         // 报告类型: annual-年报, quarter-季报
	FiscalPeriod   any         // 会计期间 (衍生字段，如2023Q1)
	Mgzbgjj        any         // 每股资本公积金
	Mgjyhdxjl      any         // 每股经营活动现金流量
	Mgjzc          any         // 每股净资产
	Jbmgsy         any         // 基本每股收益
	Xsmgsy         any         // 稀释每股收益
	Mgwfplr        any         // 每股未分配利润
	Kfmgsy         any         // 扣非每股收益
	Jzcsyl         any         // 净资产收益率(%)
	Jqjzcsyl       any         // 加权净资产收益率(%)
	Tbjzcsyl       any         // 摊薄净资产收益率(%)
	Tbzzcsyl       any         // 摊薄总资产收益率(%)
	Xsmlv          any         // 销售毛利率(%)
	Mlv            any         // 毛利率(%)
	Jlv            any         // 净利率(%)
	Sjslv          any         // 实际税率(%)
	Zyyrsrzz       any         // 主营收入同比增长(%)
	Jlrzz          any         // 净利润同比增长(%)
	Gsmgsyzzdjlrzz any         // 归属于母公司所有者的净利润同比增长(%)
	Kfjlrzz        any         // 扣非净利润同比增长(%)
	Yyzsrgdhbzz    any         // 营业总收入滚动环比增长(%)
	Sljlrjqhbzz    any         // 归属净利润滚动环比增长(%)
	Kfjlrgdhbzz    any         // 扣非净利润滚动环比增长(%)
	Yskyysr        any         // 预收款/营业收入
	Xsxjlyysr      any         // 销售现金流/营业收入
	Zcfzl          any         // 资产负债比率(%)
	Chzzl          any         // 存货周转率(次)
	DataSource     any         // 数据来源
	Currency       any         // 货币单位
	Unit           any         // 单位: yuan-元
	IsCalculated   any         // 是否为计算指标: 0-原始数据, 1-计算得出
	CalcVersion    any         // 计算版本
	CreatedAt      *gtime.Time // 创建时间
	UpdatedAt      *gtime.Time // 更新时间
}
