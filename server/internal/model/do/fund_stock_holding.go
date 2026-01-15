// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// FundStockHolding is the golang structure of table hg_fund_stock_holding for DAO operations like Where/Data.
type FundStockHolding struct {
	g.Meta     `orm:"table:hg_fund_stock_holding, do:true"`
	Id         any         // 自增主键
	Jzrq       *gtime.Time // 截止日期 (报告期，如2025-12-31)
	T          *gtime.Time // 交易时间 (衍生自jzrq，兼容时间序列查询)
	Jjmc       any         // 基金名称
	Jjdm       any         // 基金代码
	Symbol     any         // 股票代码 (如: 000001.SZ)
	Ccsl       any         // 持仓数量(股)
	Ltbl       any         // 占流通股比例(%)
	Cgsz       any         // 持股市值（元）
	Jzbl       any         // 占净值比例（%）
	AvgCost    any         // 估算持仓成本 (元/股)
	DataSource any         // 数据来源 (如: 基金季报)
	ReportType any         // 报告类型: 季报, 中报, 年报
	IsLatest   any         // 是否为该基金对该股票的最新持仓: 0-否, 1-是
	CreatedAt  *gtime.Time // 创建时间
	UpdatedAt  *gtime.Time // 更新时间
}
