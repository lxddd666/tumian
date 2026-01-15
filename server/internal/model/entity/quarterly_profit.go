// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// QuarterlyProfit is the golang structure for table quarterly_profit.
type QuarterlyProfit struct {
	Id                   uint64      `json:"id"                   orm:"id"                     description:"自增主键"`
	Symbol               string      `json:"symbol"               orm:"symbol"                 description:"股票代码 (如: 000001.SZ)"`
	Date                 *gtime.Time `json:"date"                 orm:"date"                   description:"截止日期 (报告期截止日，如2025-03-31)"`
	ReportYear           int         `json:"reportYear"           orm:"report_year"            description:"报告年度"`
	ReportQuarter        int         `json:"reportQuarter"        orm:"report_quarter"         description:"报告季度 (1-4)"`
	FiscalPeriod         string      `json:"fiscalPeriod"         orm:"fiscal_period"          description:"会计期间 (如2025Q1)"`
	Income               float64     `json:"income"               orm:"income"                 description:"营业收入（万元）"`
	Expend               float64     `json:"expend"               orm:"expend"                 description:"营业支出（万元）"`
	Profit               float64     `json:"profit"               orm:"profit"                 description:"营业利润（万元）"`
	Totalp               float64     `json:"totalp"               orm:"totalp"                 description:"利润总额（万元）"`
	Reprofit             float64     `json:"reprofit"             orm:"reprofit"               description:"净利润（万元）"`
	Basege               float64     `json:"basege"               orm:"basege"                 description:"基本每股收益(元/股)"`
	Ettege               float64     `json:"ettege"               orm:"ettege"                 description:"稀释每股收益(元/股)"`
	Otherp               float64     `json:"otherp"               orm:"otherp"                 description:"其他综合收益（万元）"`
	Totalcp              float64     `json:"totalcp"              orm:"totalcp"                description:"综合收益总额（万元）"`
	GrossProfitMargin    float64     `json:"grossProfitMargin"    orm:"gross_profit_margin"    description:"毛利率(%)"`
	NetProfitMargin      float64     `json:"netProfitMargin"      orm:"net_profit_margin"      description:"净利率(%)"`
	OperatingProfitRatio float64     `json:"operatingProfitRatio" orm:"operating_profit_ratio" description:"营业利润率(%)"`
	ReportType           string      `json:"reportType"           orm:"report_type"            description:"报告类型: 一季报, 中报, 三季报, 年报"`
	DataSource           string      `json:"dataSource"           orm:"data_source"            description:"数据来源 (如: 交易所财报)"`
	Currency             string      `json:"currency"             orm:"currency"               description:"货币单位"`
	IsAudited            int         `json:"isAudited"            orm:"is_audited"             description:"是否审计: 0-未审计, 1-已审计"`
	CreatedAt            *gtime.Time `json:"createdAt"            orm:"created_at"             description:"创建时间"`
	UpdatedAt            *gtime.Time `json:"updatedAt"            orm:"updated_at"             description:"更新时间"`
}
