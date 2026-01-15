// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// IncomeStatement is the golang structure for table income_statement.
type IncomeStatement struct {
	Id                 uint64      `json:"id"                 orm:"id"                  description:"自增主键"`
	Symbol             string      `json:"symbol"             orm:"symbol"              description:"股票代码 (如: 000001.SZ)"`
	Jzrq               *gtime.Time `json:"jzrq"               orm:"jzrq"                description:"截止日期 (报告期截止日，如2025-12-31)"`
	Plrq               *gtime.Time `json:"plrq"               orm:"plrq"                description:"披露日期 (财报实际发布日期)"`
	ReportYear         int         `json:"reportYear"         orm:"report_year"         description:"报告年度"`
	ReportQuarter      int         `json:"reportQuarter"      orm:"report_quarter"      description:"报告季度 (1-4)"`
	ReportType         string      `json:"reportType"         orm:"report_type"         description:"报告类型: annual-年报, quarter-季报"`
	FiscalPeriod       string      `json:"fiscalPeriod"       orm:"fiscal_period"       description:"会计期间"`
	Yysr               float64     `json:"yysr"               orm:"yysr"                description:"营业收入"`
	Yzbf               float64     `json:"yzbf"               orm:"yzbf"                description:"已赚保费"`
	Fdczssr            float64     `json:"fdczssr"            orm:"fdczssr"             description:"房地产销售收入"`
	Qtywsr             float64     `json:"qtywsr"             orm:"qtywsr"              description:"其他业务收入"`
	Yyzsr              float64     `json:"yyzsr"              orm:"yyzsr"               description:"营业总收入"`
	Lxsr               float64     `json:"lxsr"               orm:"lxsr"                description:"利息收入"`
	Sxfjyjsr           float64     `json:"sxfjyjsr"           orm:"sxfjyjsr"            description:"手续费及佣金收入"`
	Btsr               float64     `json:"btsr"               orm:"btsr"                description:"补贴收入"`
	Ywsr               float64     `json:"ywsr"               orm:"ywsr"                description:"营业外收入"`
	Qtsy               float64     `json:"qtsy"               orm:"qtsy"                description:"其他收益"`
	Yycb               float64     `json:"yycb"               orm:"yycb"                description:"营业成本"`
	Fdczscb            float64     `json:"fdczscb"            orm:"fdczscb"             description:"房地产销售成本"`
	Qtywcb             float64     `json:"qtywcb"             orm:"qtywcb"              description:"其他业务成本"`
	Yyzcb              float64     `json:"yyzcb"              orm:"yyzcb"               description:"营业总成本"`
	Yysjjfj            float64     `json:"yysjjfj"            orm:"yysjjfj"             description:"营业税金及附加"`
	Xsfy               float64     `json:"xsfy"               orm:"xsfy"                description:"销售费用"`
	Glfy               float64     `json:"glfy"               orm:"glfy"                description:"管理费用"`
	Yffy               float64     `json:"yffy"               orm:"yffy"                description:"研发费用"`
	Cwfy               float64     `json:"cwfy"               orm:"cwfy"                description:"财务费用"`
	Sxfjyjzc           float64     `json:"sxfjyjzc"           orm:"sxfjyjzc"            description:"手续费及佣金支出"`
	Lxzc               float64     `json:"lxzc"               orm:"lxzc"                description:"利息支出"`
	Tbj                float64     `json:"tbj"                orm:"tbj"                 description:"退保金"`
	Pczjje             float64     `json:"pczjje"             orm:"pczjje"              description:"赔付支出净额"`
	Tqbxhtzbjje        float64     `json:"tqbxhtzbjje"        orm:"tqbxhtzbjje"         description:"提取保险合同准备金净额"`
	Bdhlzc             float64     `json:"bdhlzc"             orm:"bdhlzc"              description:"保单红利支出"`
	Fbfy               float64     `json:"fbfy"               orm:"fbfy"                description:"分保费用"`
	Zcjzss             float64     `json:"zcjzss"             orm:"zcjzss"              description:"资产减值损失"`
	Ywzc               float64     `json:"ywzc"               orm:"ywzc"                description:"营业外支出"`
	Qtywlr             float64     `json:"qtywlr"             orm:"qtywlr"              description:"其他业务利润"`
	Yylr               float64     `json:"yylr"               orm:"yylr"                description:"营业利润"`
	Lrze               float64     `json:"lrze"               orm:"lrze"                description:"利润总额"`
	Jlr                float64     `json:"jlr"                orm:"jlr"                 description:"净利润"`
	Jlrhfcjcx          float64     `json:"jlrhfcjcx"          orm:"jlrhfcjcx"           description:"净利润(扣除非经常性损益后)"`
	Gsmgsyzzdjlr       float64     `json:"gsmgsyzzdjlr"       orm:"gsmgsyzzdjlr"        description:"归属于母公司所有者的净利润"`
	Bhbfzhbqsljlr      float64     `json:"bhbfzhbqsljlr"      orm:"bhbfzhbqsljlr"       description:"被合并方在合并前实现净利润"`
	Tzsy               float64     `json:"tzsy"               orm:"tzsy"                description:"投资收益"`
	Lyqyhhhqydtzsy     float64     `json:"lyqyhhhqydtzsy"     orm:"lyqyhhhqydtzsy"      description:"联营企业和合营企业的投资收益"`
	Gyjzbdsy           float64     `json:"gyjzbdsy"           orm:"gyjzbdsy"            description:"公允价值变动收益"`
	Qhsy               float64     `json:"qhsy"               orm:"qhsy"                description:"期货损益"`
	Tgsy               float64     `json:"tgsy"               orm:"tgsy"                description:"托管收益"`
	Hdsy               float64     `json:"hdsy"               orm:"hdsy"                description:"汇兑收益"`
	Fldzcczsy          float64     `json:"fldzcczsy"          orm:"fldzcczsy"           description:"非流动资产处置收益"`
	Sdsfy              float64     `json:"sdsfy"              orm:"sdsfy"               description:"所得税费用"`
	Ssgdsy             float64     `json:"ssgdsy"             orm:"ssgdsy"              description:"少数股东损益"`
	Wqrtzss            float64     `json:"wqrtzss"            orm:"wqrtzss"             description:"未确认投资损失"`
	Jbmgsy             float64     `json:"jbmgsy"             orm:"jbmgsy"              description:"基本每股收益"`
	Xsmgsy             float64     `json:"xsmgsy"             orm:"xsmgsy"              description:"稀释每股收益"`
	Zhsyz              float64     `json:"zhsyz"              orm:"zhsyz"               description:"综合收益总额"`
	Gsssgdzhsyz        float64     `json:"gsssgdzhsyz"        orm:"gsssgdzhsyz"         description:"归属于少数股东的综合收益总额"`
	GrossMargin        float64     `json:"grossMargin"        orm:"gross_margin"        description:"毛利率(%)"`
	OperatingMargin    float64     `json:"operatingMargin"    orm:"operating_margin"    description:"营业利润率(%)"`
	NetMargin          float64     `json:"netMargin"          orm:"net_margin"          description:"净利率(%)"`
	EffectiveTaxRate   float64     `json:"effectiveTaxRate"   orm:"effective_tax_rate"  description:"实际税率(%)"`
	DataSource         string      `json:"dataSource"         orm:"data_source"         description:"数据来源"`
	Currency           string      `json:"currency"           orm:"currency"            description:"货币单位"`
	Unit               string      `json:"unit"               orm:"unit"                description:"单位: yuan-元, wan-万元"`
	AccountingStandard string      `json:"accountingStandard" orm:"accounting_standard" description:"会计准则 (如: CAS, IFRS)"`
	IsAudited          int         `json:"isAudited"          orm:"is_audited"          description:"是否审计: 0-未审计, 1-已审计"`
	IsConsolidated     int         `json:"isConsolidated"     orm:"is_consolidated"     description:"是否合并报表: 1-合并, 0-母公司"`
	CreatedAt          *gtime.Time `json:"createdAt"          orm:"created_at"          description:"创建时间"`
	UpdatedAt          *gtime.Time `json:"updatedAt"          orm:"updated_at"          description:"更新时间"`
}
