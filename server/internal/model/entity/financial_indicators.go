// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// FinancialIndicators is the golang structure for table financial_indicators.
type FinancialIndicators struct {
	Id             uint64      `json:"id"             orm:"id"             description:"自增主键"`
	Symbol         string      `json:"symbol"         orm:"symbol"         description:"公司代码/股票代码 (例如: 000001.SZ, AAPL)"`
	Jzrq           *gtime.Time `json:"jzrq"           orm:"jzrq"           description:"截止日期 (报告期结束日)"`
	Plrq           *gtime.Time `json:"plrq"           orm:"plrq"           description:"披露日期"`
	ReportYear     int         `json:"reportYear"     orm:"report_year"    description:"报告年度"`
	ReportQuarter  int         `json:"reportQuarter"  orm:"report_quarter" description:"报告季度 (1-4)"`
	ReportType     string      `json:"reportType"     orm:"report_type"    description:"报告类型: annual-年报, quarter-季报"`
	FiscalPeriod   string      `json:"fiscalPeriod"   orm:"fiscal_period"  description:"会计期间 (衍生字段，如2023Q1)"`
	Mgzbgjj        float64     `json:"mgzbgjj"        orm:"mgzbgjj"        description:"每股资本公积金"`
	Mgjyhdxjl      float64     `json:"mgjyhdxjl"      orm:"mgjyhdxjl"      description:"每股经营活动现金流量"`
	Mgjzc          float64     `json:"mgjzc"          orm:"mgjzc"          description:"每股净资产"`
	Jbmgsy         float64     `json:"jbmgsy"         orm:"jbmgsy"         description:"基本每股收益"`
	Xsmgsy         float64     `json:"xsmgsy"         orm:"xsmgsy"         description:"稀释每股收益"`
	Mgwfplr        float64     `json:"mgwfplr"        orm:"mgwfplr"        description:"每股未分配利润"`
	Kfmgsy         float64     `json:"kfmgsy"         orm:"kfmgsy"         description:"扣非每股收益"`
	Jzcsyl         float64     `json:"jzcsyl"         orm:"jzcsyl"         description:"净资产收益率(%)"`
	Jqjzcsyl       float64     `json:"jqjzcsyl"       orm:"jqjzcsyl"       description:"加权净资产收益率(%)"`
	Tbjzcsyl       float64     `json:"tbjzcsyl"       orm:"tbjzcsyl"       description:"摊薄净资产收益率(%)"`
	Tbzzcsyl       float64     `json:"tbzzcsyl"       orm:"tbzzcsyl"       description:"摊薄总资产收益率(%)"`
	Xsmlv          float64     `json:"xsmlv"          orm:"xsmlv"          description:"销售毛利率(%)"`
	Mlv            float64     `json:"mlv"            orm:"mlv"            description:"毛利率(%)"`
	Jlv            float64     `json:"jlv"            orm:"jlv"            description:"净利率(%)"`
	Sjslv          float64     `json:"sjslv"          orm:"sjslv"          description:"实际税率(%)"`
	Zyyrsrzz       float64     `json:"zyyrsrzz"       orm:"zyyrsrzz"       description:"主营收入同比增长(%)"`
	Jlrzz          float64     `json:"jlrzz"          orm:"jlrzz"          description:"净利润同比增长(%)"`
	Gsmgsyzzdjlrzz float64     `json:"gsmgsyzzdjlrzz" orm:"gsmgsyzzdjlrzz" description:"归属于母公司所有者的净利润同比增长(%)"`
	Kfjlrzz        float64     `json:"kfjlrzz"        orm:"kfjlrzz"        description:"扣非净利润同比增长(%)"`
	Yyzsrgdhbzz    float64     `json:"yyzsrgdhbzz"    orm:"yyzsrgdhbzz"    description:"营业总收入滚动环比增长(%)"`
	Sljlrjqhbzz    float64     `json:"sljlrjqhbzz"    orm:"sljlrjqhbzz"    description:"归属净利润滚动环比增长(%)"`
	Kfjlrgdhbzz    float64     `json:"kfjlrgdhbzz"    orm:"kfjlrgdhbzz"    description:"扣非净利润滚动环比增长(%)"`
	Yskyysr        float64     `json:"yskyysr"        orm:"yskyysr"        description:"预收款/营业收入"`
	Xsxjlyysr      float64     `json:"xsxjlyysr"      orm:"xsxjlyysr"      description:"销售现金流/营业收入"`
	Zcfzl          float64     `json:"zcfzl"          orm:"zcfzl"          description:"资产负债比率(%)"`
	Chzzl          float64     `json:"chzzl"          orm:"chzzl"          description:"存货周转率(次)"`
	DataSource     string      `json:"dataSource"     orm:"data_source"    description:"数据来源"`
	Currency       string      `json:"currency"       orm:"currency"       description:"货币单位"`
	Unit           string      `json:"unit"           orm:"unit"           description:"单位: yuan-元"`
	IsCalculated   int         `json:"isCalculated"   orm:"is_calculated"  description:"是否为计算指标: 0-原始数据, 1-计算得出"`
	CalcVersion    string      `json:"calcVersion"    orm:"calc_version"   description:"计算版本"`
	CreatedAt      *gtime.Time `json:"createdAt"      orm:"created_at"     description:"创建时间"`
	UpdatedAt      *gtime.Time `json:"updatedAt"      orm:"updated_at"     description:"更新时间"`
}
