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
	Id                   interface{} // 自增主键
	Symbol               interface{} // 股票代码 (如: 000001.SZ)
	Date                 *gtime.Time // 截止日期 (报告期截止日，如2025-03-31)
	ReportYear           interface{} // 报告年度
	ReportQuarter        interface{} // 报告季度 (1-4)
	FiscalPeriod         interface{} // 会计期间 (如2025Q1)
	Income               interface{} // 营业收入（万元）
	Expend               interface{} // 营业支出（万元）
	Profit               interface{} // 营业利润（万元）
	Totalp               interface{} // 利润总额（万元）
	Reprofit             interface{} // 净利润（万元）
	Basege               interface{} // 基本每股收益(元/股)
	Ettege               interface{} // 稀释每股收益(元/股)
	Otherp               interface{} // 其他综合收益（万元）
	Totalcp              interface{} // 综合收益总额（万元）
	GrossProfitMargin    interface{} // 毛利率(%)
	NetProfitMargin      interface{} // 净利率(%)
	OperatingProfitRatio interface{} // 营业利润率(%)
	ReportType           interface{} // 报告类型: 一季报, 中报, 三季报, 年报
	DataSource           interface{} // 数据来源 (如: 交易所财报)
	Currency             interface{} // 货币单位
	IsAudited            interface{} // 是否审计: 0-未审计, 1-已审计
	CreatedAt            *gtime.Time // 创建时间
	UpdatedAt            *gtime.Time // 更新时间
}
