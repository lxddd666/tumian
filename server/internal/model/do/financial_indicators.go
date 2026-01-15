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
	Id             interface{} // 自增主键
	Symbol         interface{} // 公司代码/股票代码 (例如: 000001.SZ, AAPL)
	Jzrq           *gtime.Time // 截止日期 (报告期结束日)
	Plrq           *gtime.Time // 披露日期
	ReportYear     interface{} // 报告年度
	ReportQuarter  interface{} // 报告季度 (1-4)
	ReportType     interface{} // 报告类型: annual-年报, quarter-季报
	FiscalPeriod   interface{} // 会计期间 (衍生字段，如2023Q1)
	Mgzbgjj        interface{} // 每股资本公积金
	Mgjyhdxjl      interface{} // 每股经营活动现金流量
	Mgjzc          interface{} // 每股净资产
	Jbmgsy         interface{} // 基本每股收益
	Xsmgsy         interface{} // 稀释每股收益
	Mgwfplr        interface{} // 每股未分配利润
	Kfmgsy         interface{} // 扣非每股收益
	Jzcsyl         interface{} // 净资产收益率(%)
	Jqjzcsyl       interface{} // 加权净资产收益率(%)
	Tbjzcsyl       interface{} // 摊薄净资产收益率(%)
	Tbzzcsyl       interface{} // 摊薄总资产收益率(%)
	Xsmlv          interface{} // 销售毛利率(%)
	Mlv            interface{} // 毛利率(%)
	Jlv            interface{} // 净利率(%)
	Sjslv          interface{} // 实际税率(%)
	Zyyrsrzz       interface{} // 主营收入同比增长(%)
	Jlrzz          interface{} // 净利润同比增长(%)
	Gsmgsyzzdjlrzz interface{} // 归属于母公司所有者的净利润同比增长(%)
	Kfjlrzz        interface{} // 扣非净利润同比增长(%)
	Yyzsrgdhbzz    interface{} // 营业总收入滚动环比增长(%)
	Sljlrjqhbzz    interface{} // 归属净利润滚动环比增长(%)
	Kfjlrgdhbzz    interface{} // 扣非净利润滚动环比增长(%)
	Yskyysr        interface{} // 预收款/营业收入
	Xsxjlyysr      interface{} // 销售现金流/营业收入
	Zcfzl          interface{} // 资产负债比率(%)
	Chzzl          interface{} // 存货周转率(次)
	DataSource     interface{} // 数据来源
	Currency       interface{} // 货币单位
	Unit           interface{} // 单位: yuan-元
	IsCalculated   interface{} // 是否为计算指标: 0-原始数据, 1-计算得出
	CalcVersion    interface{} // 计算版本
	CreatedAt      *gtime.Time // 创建时间
	UpdatedAt      *gtime.Time // 更新时间
}
