// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// FundStockHolding is the golang structure for table fund_stock_holding.
type FundStockHolding struct {
	Id         uint64      `json:"id"         orm:"id"          description:"自增主键"`
	Jzrq       *gtime.Time `json:"jzrq"       orm:"jzrq"        description:"截止日期 (报告期，如2025-12-31)"`
	T          *gtime.Time `json:"t"          orm:"t"           description:"交易时间 (衍生自jzrq，兼容时间序列查询)"`
	Jjmc       string      `json:"jjmc"       orm:"jjmc"        description:"基金名称"`
	Jjdm       string      `json:"jjdm"       orm:"jjdm"        description:"基金代码"`
	Symbol     string      `json:"symbol"     orm:"symbol"      description:"股票代码 (如: 000001.SZ)"`
	Ccsl       int64       `json:"ccsl"       orm:"ccsl"        description:"持仓数量(股)"`
	Ltbl       float64     `json:"ltbl"       orm:"ltbl"        description:"占流通股比例(%)"`
	Cgsz       float64     `json:"cgsz"       orm:"cgsz"        description:"持股市值（元）"`
	Jzbl       float64     `json:"jzbl"       orm:"jzbl"        description:"占净值比例（%）"`
	AvgCost    float64     `json:"avgCost"    orm:"avg_cost"    description:"估算持仓成本 (元/股)"`
	DataSource string      `json:"dataSource" orm:"data_source" description:"数据来源 (如: 基金季报)"`
	ReportType string      `json:"reportType" orm:"report_type" description:"报告类型: 季报, 中报, 年报"`
	IsLatest   int         `json:"isLatest"   orm:"is_latest"   description:"是否为该基金对该股票的最新持仓: 0-否, 1-是"`
	CreatedAt  *gtime.Time `json:"createdAt"  orm:"created_at"  description:"创建时间"`
	UpdatedAt  *gtime.Time `json:"updatedAt"  orm:"updated_at"  description:"更新时间"`
}
