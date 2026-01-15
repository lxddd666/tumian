// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// QuarterlyProfit is the golang structure of table hg_quarterly_profit for DAO operations like Where/Data.
type QuarterlyProfit struct {
	g.Meta               `orm:"table:hg_quarterly_profit, do:true"`
	Id                   any         // 自增主键
	Symbol               any         // 股票代码 (如: 000001.SZ)
	Date                 *gtime.Time // 截止日期 (报告期截止日，如2025-03-31)
	ReportYear           any         // 报告年度
	ReportQuarter        any         // 报告季度 (1-4)
	FiscalPeriod         any         // 会计期间 (如2025Q1)
	Income               any         // 营业收入（万元）
	Expend               any         // 营业支出（万元）
	Profit               any         // 营业利润（万元）
	Totalp               any         // 利润总额（万元）
	Reprofit             any         // 净利润（万元）
	Basege               any         // 基本每股收益(元/股)
	Ettege               any         // 稀释每股收益(元/股)
	Otherp               any         // 其他综合收益（万元）
	Totalcp              any         // 综合收益总额（万元）
	GrossProfitMargin    any         // 毛利率(%)
	NetProfitMargin      any         // 净利率(%)
	OperatingProfitRatio any         // 营业利润率(%)
	ReportType           any         // 报告类型: 一季报, 中报, 三季报, 年报
	DataSource           any         // 数据来源 (如: 交易所财报)
	Currency             any         // 货币单位
	IsAudited            any         // 是否审计: 0-未审计, 1-已审计
	CreatedAt            *gtime.Time // 创建时间
	UpdatedAt            *gtime.Time // 更新时间
}
