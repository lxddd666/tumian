// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// BalanceSheet is the golang structure for table balance_sheet.
type BalanceSheet struct {
	Id            uint64      `json:"id"            orm:"id"             description:"自增主键"`
	Symbol        string      `json:"symbol"        orm:"symbol"         description:"公司代码/股票代码 (例如: 000001.SZ, AAPL)"`
	Jzrq          *gtime.Time `json:"jzrq"          orm:"jzrq"           description:"截止日期 (会计期间结束日)"`
	Plrq          *gtime.Time `json:"plrq"          orm:"plrq"           description:"披露日期"`
	ReportYear    int         `json:"reportYear"    orm:"report_year"    description:"报告年度"`
	ReportQuarter int         `json:"reportQuarter" orm:"report_quarter" description:"报告季度 (1-4, 年报为NULL)"`
	ReportType    string      `json:"reportType"    orm:"report_type"    description:"报告类型: annual-年报, quarter-季报, interim-中报"`
	Hbzj          float64     `json:"hbzj"          orm:"hbzj"           description:"货币资金"`
	Jyxjrzc       float64     `json:"jyxjrzc"       orm:"jyxjrzc"        description:"交易性金融资产"`
	Yspj          float64     `json:"yspj"          orm:"yspj"           description:"应收票据"`
	Yszk          float64     `json:"yszk"          orm:"yszk"           description:"应收账款"`
	Yfkx          float64     `json:"yfkx"          orm:"yfkx"           description:"预付款项"`
	Yslx          float64     `json:"yslx"          orm:"yslx"           description:"应收利息"`
	Ysgl          float64     `json:"ysgl"          orm:"ysgl"           description:"应收股利"`
	Qtysk         float64     `json:"qtysk"         orm:"qtysk"          description:"其他应收款"`
	Ch            float64     `json:"ch"            orm:"ch"             description:"存货"`
	Dfy           float64     `json:"dfy"           orm:"dfy"            description:"待摊费用"`
	Ynndqdfldzc   float64     `json:"ynndqdfldzc"   orm:"ynndqdfldzc"    description:"一年内到期的非流动资产"`
	Qtldzc        float64     `json:"qtldzc"        orm:"qtldzc"         description:"其他流动资产"`
	Ldzchj        float64     `json:"ldzchj"        orm:"ldzchj"         description:"流动资产合计"`
	Cqgqtz        float64     `json:"cqgqtz"        orm:"cqgqtz"         description:"长期股权投资"`
	Cqysk         float64     `json:"cqysk"         orm:"cqysk"          description:"长期应收款"`
	Gdzc          float64     `json:"gdzc"          orm:"gdzc"           description:"固定资产"`
	Zjgc          float64     `json:"zjgc"          orm:"zjgc"           description:"在建工程"`
	Wxzc          float64     `json:"wxzc"          orm:"wxzc"           description:"无形资产"`
	Sy            float64     `json:"sy"            orm:"sy"             description:"商誉"`
	Cqdtfy        float64     `json:"cqdtfy"        orm:"cqdtfy"         description:"长期待摊费用"`
	Dysdszc       float64     `json:"dysdszc"       orm:"dysdszc"        description:"递延所得税资产"`
	Qtfldzc       float64     `json:"qtfldzc"       orm:"qtfldzc"        description:"其他非流动资产"`
	Fldzchj       float64     `json:"fldzchj"       orm:"fldzchj"        description:"非流动资产合计"`
	Zczj          float64     `json:"zczj"          orm:"zczj"           description:"资产总计"`
	Dqjk          float64     `json:"dqjk"          orm:"dqjk"           description:"短期借款"`
	Jyxjrfz       float64     `json:"jyxjrfz"       orm:"jyxjrfz"        description:"交易性金融负债"`
	Yfpj          float64     `json:"yfpj"          orm:"yfpj"           description:"应付票据"`
	Yfzk          float64     `json:"yfzk"          orm:"yfzk"           description:"应付账款"`
	Ysk           float64     `json:"ysk"           orm:"ysk"            description:"预收账款"`
	Yfgzxc        float64     `json:"yfgzxc"        orm:"yfgzxc"         description:"应付职工薪酬"`
	Yjsf          float64     `json:"yjsf"          orm:"yjsf"           description:"应交税费"`
	Yflx          float64     `json:"yflx"          orm:"yflx"           description:"应付利息"`
	Yfgl          float64     `json:"yfgl"          orm:"yfgl"           description:"应付股利"`
	Qtfzk         float64     `json:"qtfzk"         orm:"qtfzk"          description:"其他应付款"`
	Ynndqdfldfz   float64     `json:"ynndqdfldfz"   orm:"ynndqdfldfz"    description:"一年内到期的非流动负债"`
	Qtldfz        float64     `json:"qtldfz"        orm:"qtldfz"         description:"其他流动负债"`
	Ldfzhj        float64     `json:"ldfzhj"        orm:"ldfzhj"         description:"流动负债合计"`
	Cqjk          float64     `json:"cqjk"          orm:"cqjk"           description:"长期借款"`
	Yfzq          float64     `json:"yfzq"          orm:"yfzq"           description:"应付债券"`
	Cqyfk         float64     `json:"cqyfk"         orm:"cqyfk"          description:"长期应付款"`
	Dysdsfz       float64     `json:"dysdsfz"       orm:"dysdsfz"        description:"递延所得税负债"`
	Qtfldfz       float64     `json:"qtfldfz"       orm:"qtfldfz"        description:"其他非流动负债"`
	Fldfzhj       float64     `json:"fldfzhj"       orm:"fldfzhj"        description:"非流动负债合计"`
	Fzhj          float64     `json:"fzhj"          orm:"fzhj"           description:"负债合计"`
	Sszb          float64     `json:"sszb"          orm:"sszb"           description:"实收资本(或股本)"`
	Zbgj          float64     `json:"zbgj"          orm:"zbgj"           description:"资本公积"`
	Ylgj          float64     `json:"ylgj"          orm:"ylgj"           description:"盈余公积"`
	Wfplr         float64     `json:"wfplr"         orm:"wfplr"          description:"未分配利润"`
	Gsmgdqsyhj    float64     `json:"gsmgdqsyhj"    orm:"gsmgdqsyhj"     description:"归属于母公司股东权益合计"`
	Ssgdqy        float64     `json:"ssgdqy"        orm:"ssgdqy"         description:"少数股东权益"`
	Syzqyhj       float64     `json:"syzqyhj"       orm:"syzqyhj"        description:"所有者权益合计"`
	Fzhgdqyzj     float64     `json:"fzhgdqyzj"     orm:"fzhgdqyzj"      description:"负债和股东权益总计"`
	DataSource    string      `json:"dataSource"    orm:"data_source"    description:"数据来源"`
	Currency      string      `json:"currency"      orm:"currency"       description:"货币单位 (CNY, USD等)"`
	Unit          string      `json:"unit"          orm:"unit"           description:"单位: yuan-元, wan-万元, qianwan-千万元"`
	IsAudited     int         `json:"isAudited"     orm:"is_audited"     description:"是否审计: 0-未审计, 1-已审计"`
	Version       int         `json:"version"       orm:"version"        description:"数据版本"`
	CreatedAt     *gtime.Time `json:"createdAt"     orm:"created_at"     description:"创建时间"`
	UpdatedAt     *gtime.Time `json:"updatedAt"     orm:"updated_at"     description:"更新时间"`
	CreatedBy     string      `json:"createdBy"     orm:"created_by"     description:"创建人"`
	UpdatedBy     string      `json:"updatedBy"     orm:"updated_by"     description:"更新人"`
}
