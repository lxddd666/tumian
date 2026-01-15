// Package stockin
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2026 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
// @AutoGenerate Version 2.18.6
package stockin

import (
	"context"
	"hotgo/internal/model/entity"
	"hotgo/internal/model/input/form"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// IncomeStatementUpdateFields 修改利润表 (Income Statement)字段过滤
type IncomeStatementUpdateFields struct {
	Symbol             string      `json:"symbol"             dc:"股票代码 (如: 000001.SZ)"`
	Jzrq               *gtime.Time `json:"jzrq"               dc:"截止日期 (报告期截止日，如2025-12-31)"`
	Plrq               *gtime.Time `json:"plrq"               dc:"披露日期 (财报实际发布日期)"`
	ReportYear         int         `json:"reportYear"         dc:"报告年度"`
	ReportQuarter      int         `json:"reportQuarter"      dc:"报告季度 (1-4)"`
	ReportType         string      `json:"reportType"         dc:"报告类型: annual-年报, quarter-季报"`
	FiscalPeriod       string      `json:"fiscalPeriod"       dc:"会计期间"`
	Yysr               float64     `json:"yysr"               dc:"营业收入"`
	Yzbf               float64     `json:"yzbf"               dc:"已赚保费"`
	Fdczssr            float64     `json:"fdczssr"            dc:"房地产销售收入"`
	Qtywsr             float64     `json:"qtywsr"             dc:"其他业务收入"`
	Yyzsr              float64     `json:"yyzsr"              dc:"营业总收入"`
	Lxsr               float64     `json:"lxsr"               dc:"利息收入"`
	Sxfjyjsr           float64     `json:"sxfjyjsr"           dc:"手续费及佣金收入"`
	Btsr               float64     `json:"btsr"               dc:"补贴收入"`
	Ywsr               float64     `json:"ywsr"               dc:"营业外收入"`
	Qtsy               float64     `json:"qtsy"               dc:"其他收益"`
	Yycb               float64     `json:"yycb"               dc:"营业成本"`
	Fdczscb            float64     `json:"fdczscb"            dc:"房地产销售成本"`
	Qtywcb             float64     `json:"qtywcb"             dc:"其他业务成本"`
	Yyzcb              float64     `json:"yyzcb"              dc:"营业总成本"`
	Yysjjfj            float64     `json:"yysjjfj"            dc:"营业税金及附加"`
	Xsfy               float64     `json:"xsfy"               dc:"销售费用"`
	Glfy               float64     `json:"glfy"               dc:"管理费用"`
	Yffy               float64     `json:"yffy"               dc:"研发费用"`
	Cwfy               float64     `json:"cwfy"               dc:"财务费用"`
	Sxfjyjzc           float64     `json:"sxfjyjzc"           dc:"手续费及佣金支出"`
	Lxzc               float64     `json:"lxzc"               dc:"利息支出"`
	Tbj                float64     `json:"tbj"                dc:"退保金"`
	Pczjje             float64     `json:"pczjje"             dc:"赔付支出净额"`
	Tqbxhtzbjje        float64     `json:"tqbxhtzbjje"        dc:"提取保险合同准备金净额"`
	Bdhlzc             float64     `json:"bdhlzc"             dc:"保单红利支出"`
	Fbfy               float64     `json:"fbfy"               dc:"分保费用"`
	Zcjzss             float64     `json:"zcjzss"             dc:"资产减值损失"`
	Ywzc               float64     `json:"ywzc"               dc:"营业外支出"`
	Qtywlr             float64     `json:"qtywlr"             dc:"其他业务利润"`
	Yylr               float64     `json:"yylr"               dc:"营业利润"`
	Lrze               float64     `json:"lrze"               dc:"利润总额"`
	Jlr                float64     `json:"jlr"                dc:"净利润"`
	Jlrhfcjcx          float64     `json:"jlrhfcjcx"          dc:"净利润(扣除非经常性损益后)"`
	Gsmgsyzzdjlr       float64     `json:"gsmgsyzzdjlr"       dc:"归属于母公司所有者的净利润"`
	Bhbfzhbqsljlr      float64     `json:"bhbfzhbqsljlr"      dc:"被合并方在合并前实现净利润"`
	Tzsy               float64     `json:"tzsy"               dc:"投资收益"`
	Lyqyhhhqydtzsy     float64     `json:"lyqyhhhqydtzsy"     dc:"联营企业和合营企业的投资收益"`
	Gyjzbdsy           float64     `json:"gyjzbdsy"           dc:"公允价值变动收益"`
	Qhsy               float64     `json:"qhsy"               dc:"期货损益"`
	Tgsy               float64     `json:"tgsy"               dc:"托管收益"`
	Hdsy               float64     `json:"hdsy"               dc:"汇兑收益"`
	Fldzcczsy          float64     `json:"fldzcczsy"          dc:"非流动资产处置收益"`
	Sdsfy              float64     `json:"sdsfy"              dc:"所得税费用"`
	Ssgdsy             float64     `json:"ssgdsy"             dc:"少数股东损益"`
	Wqrtzss            float64     `json:"wqrtzss"            dc:"未确认投资损失"`
	Jbmgsy             float64     `json:"jbmgsy"             dc:"基本每股收益"`
	Xsmgsy             float64     `json:"xsmgsy"             dc:"稀释每股收益"`
	Zhsyz              float64     `json:"zhsyz"              dc:"综合收益总额"`
	Gsssgdzhsyz        float64     `json:"gsssgdzhsyz"        dc:"归属于少数股东的综合收益总额"`
	GrossMargin        float64     `json:"grossMargin"        dc:"毛利率(%)"`
	OperatingMargin    float64     `json:"operatingMargin"    dc:"营业利润率(%)"`
	NetMargin          float64     `json:"netMargin"          dc:"净利率(%)"`
	EffectiveTaxRate   float64     `json:"effectiveTaxRate"   dc:"实际税率(%)"`
	DataSource         string      `json:"dataSource"         dc:"数据来源"`
	Currency           string      `json:"currency"           dc:"货币单位"`
	Unit               string      `json:"unit"               dc:"单位: yuan-元, wan-万元"`
	AccountingStandard string      `json:"accountingStandard" dc:"会计准则 (如: CAS, IFRS)"`
	IsAudited          int         `json:"isAudited"          dc:"是否审计: 0-未审计, 1-已审计"`
	IsConsolidated     int         `json:"isConsolidated"     dc:"是否合并报表: 1-合并, 0-母公司"`
}

// IncomeStatementInsertFields 新增利润表 (Income Statement)字段过滤
type IncomeStatementInsertFields struct {
	Symbol             string      `json:"symbol"             dc:"股票代码 (如: 000001.SZ)"`
	Jzrq               *gtime.Time `json:"jzrq"               dc:"截止日期 (报告期截止日，如2025-12-31)"`
	Plrq               *gtime.Time `json:"plrq"               dc:"披露日期 (财报实际发布日期)"`
	ReportYear         int         `json:"reportYear"         dc:"报告年度"`
	ReportQuarter      int         `json:"reportQuarter"      dc:"报告季度 (1-4)"`
	ReportType         string      `json:"reportType"         dc:"报告类型: annual-年报, quarter-季报"`
	FiscalPeriod       string      `json:"fiscalPeriod"       dc:"会计期间"`
	Yysr               float64     `json:"yysr"               dc:"营业收入"`
	Yzbf               float64     `json:"yzbf"               dc:"已赚保费"`
	Fdczssr            float64     `json:"fdczssr"            dc:"房地产销售收入"`
	Qtywsr             float64     `json:"qtywsr"             dc:"其他业务收入"`
	Yyzsr              float64     `json:"yyzsr"              dc:"营业总收入"`
	Lxsr               float64     `json:"lxsr"               dc:"利息收入"`
	Sxfjyjsr           float64     `json:"sxfjyjsr"           dc:"手续费及佣金收入"`
	Btsr               float64     `json:"btsr"               dc:"补贴收入"`
	Ywsr               float64     `json:"ywsr"               dc:"营业外收入"`
	Qtsy               float64     `json:"qtsy"               dc:"其他收益"`
	Yycb               float64     `json:"yycb"               dc:"营业成本"`
	Fdczscb            float64     `json:"fdczscb"            dc:"房地产销售成本"`
	Qtywcb             float64     `json:"qtywcb"             dc:"其他业务成本"`
	Yyzcb              float64     `json:"yyzcb"              dc:"营业总成本"`
	Yysjjfj            float64     `json:"yysjjfj"            dc:"营业税金及附加"`
	Xsfy               float64     `json:"xsfy"               dc:"销售费用"`
	Glfy               float64     `json:"glfy"               dc:"管理费用"`
	Yffy               float64     `json:"yffy"               dc:"研发费用"`
	Cwfy               float64     `json:"cwfy"               dc:"财务费用"`
	Sxfjyjzc           float64     `json:"sxfjyjzc"           dc:"手续费及佣金支出"`
	Lxzc               float64     `json:"lxzc"               dc:"利息支出"`
	Tbj                float64     `json:"tbj"                dc:"退保金"`
	Pczjje             float64     `json:"pczjje"             dc:"赔付支出净额"`
	Tqbxhtzbjje        float64     `json:"tqbxhtzbjje"        dc:"提取保险合同准备金净额"`
	Bdhlzc             float64     `json:"bdhlzc"             dc:"保单红利支出"`
	Fbfy               float64     `json:"fbfy"               dc:"分保费用"`
	Zcjzss             float64     `json:"zcjzss"             dc:"资产减值损失"`
	Ywzc               float64     `json:"ywzc"               dc:"营业外支出"`
	Qtywlr             float64     `json:"qtywlr"             dc:"其他业务利润"`
	Yylr               float64     `json:"yylr"               dc:"营业利润"`
	Lrze               float64     `json:"lrze"               dc:"利润总额"`
	Jlr                float64     `json:"jlr"                dc:"净利润"`
	Jlrhfcjcx          float64     `json:"jlrhfcjcx"          dc:"净利润(扣除非经常性损益后)"`
	Gsmgsyzzdjlr       float64     `json:"gsmgsyzzdjlr"       dc:"归属于母公司所有者的净利润"`
	Bhbfzhbqsljlr      float64     `json:"bhbfzhbqsljlr"      dc:"被合并方在合并前实现净利润"`
	Tzsy               float64     `json:"tzsy"               dc:"投资收益"`
	Lyqyhhhqydtzsy     float64     `json:"lyqyhhhqydtzsy"     dc:"联营企业和合营企业的投资收益"`
	Gyjzbdsy           float64     `json:"gyjzbdsy"           dc:"公允价值变动收益"`
	Qhsy               float64     `json:"qhsy"               dc:"期货损益"`
	Tgsy               float64     `json:"tgsy"               dc:"托管收益"`
	Hdsy               float64     `json:"hdsy"               dc:"汇兑收益"`
	Fldzcczsy          float64     `json:"fldzcczsy"          dc:"非流动资产处置收益"`
	Sdsfy              float64     `json:"sdsfy"              dc:"所得税费用"`
	Ssgdsy             float64     `json:"ssgdsy"             dc:"少数股东损益"`
	Wqrtzss            float64     `json:"wqrtzss"            dc:"未确认投资损失"`
	Jbmgsy             float64     `json:"jbmgsy"             dc:"基本每股收益"`
	Xsmgsy             float64     `json:"xsmgsy"             dc:"稀释每股收益"`
	Zhsyz              float64     `json:"zhsyz"              dc:"综合收益总额"`
	Gsssgdzhsyz        float64     `json:"gsssgdzhsyz"        dc:"归属于少数股东的综合收益总额"`
	GrossMargin        float64     `json:"grossMargin"        dc:"毛利率(%)"`
	OperatingMargin    float64     `json:"operatingMargin"    dc:"营业利润率(%)"`
	NetMargin          float64     `json:"netMargin"          dc:"净利率(%)"`
	EffectiveTaxRate   float64     `json:"effectiveTaxRate"   dc:"实际税率(%)"`
	DataSource         string      `json:"dataSource"         dc:"数据来源"`
	Currency           string      `json:"currency"           dc:"货币单位"`
	Unit               string      `json:"unit"               dc:"单位: yuan-元, wan-万元"`
	AccountingStandard string      `json:"accountingStandard" dc:"会计准则 (如: CAS, IFRS)"`
	IsAudited          int         `json:"isAudited"          dc:"是否审计: 0-未审计, 1-已审计"`
	IsConsolidated     int         `json:"isConsolidated"     dc:"是否合并报表: 1-合并, 0-母公司"`
}

// IncomeStatementEditInp 修改/新增利润表 (Income Statement)
type IncomeStatementEditInp struct {
	entity.IncomeStatement
}

func (in *IncomeStatementEditInp) Filter(ctx context.Context) (err error) {
	// 验证股票代码 (如: 000001.SZ)
	if err := g.Validator().Rules("required").Data(in.Symbol).Messages("股票代码 (如: 000001.SZ)不能为空").Run(ctx); err != nil {
		return err.Current()
	}
	// 验证截止日期 (报告期截止日，如2025-12-31)
	if err := g.Validator().Rules("required").Data(in.Jzrq).Messages("截止日期 (报告期截止日，如2025-12-31)不能为空").Run(ctx); err != nil {
		return err.Current()
	}
	// 验证披露日期 (财报实际发布日期)
	if err := g.Validator().Rules("required").Data(in.Plrq).Messages("披露日期 (财报实际发布日期)不能为空").Run(ctx); err != nil {
		return err.Current()
	}
	// 验证报告类型: annual-年报, quarter-季报
	if err := g.Validator().Rules("required").Data(in.ReportType).Messages("报告类型: annual-年报, quarter-季报不能为空").Run(ctx); err != nil {
		return err.Current()
	}
	return
}

type IncomeStatementEditModel struct{}

// IncomeStatementDeleteInp 删除利润表 (Income Statement)
type IncomeStatementDeleteInp struct {
	Id interface{} `json:"id" v:"required#自增主键不能为空" dc:"自增主键"`
}

func (in *IncomeStatementDeleteInp) Filter(ctx context.Context) (err error) {
	return
}

type IncomeStatementDeleteModel struct{}

// IncomeStatementViewInp 获取指定利润表 (Income Statement)信息
type IncomeStatementViewInp struct {
	Id int64 `json:"id" v:"required#自增主键不能为空" dc:"自增主键"`
}

func (in *IncomeStatementViewInp) Filter(ctx context.Context) (err error) {
	return
}

type IncomeStatementViewModel struct {
	entity.IncomeStatement
}

// IncomeStatementListInp 获取利润表 (Income Statement)列表
type IncomeStatementListInp struct {
	form.PageReq
	Id        int64         `json:"id"        dc:"自增主键"`
	CreatedAt []*gtime.Time `json:"createdAt" dc:"创建时间"`
}

func (in *IncomeStatementListInp) Filter(ctx context.Context) (err error) {
	return
}

type IncomeStatementListModel struct {
	Id                 int64       `json:"id"                 dc:"自增主键"`
	Symbol             string      `json:"symbol"             dc:"股票代码 (如: 000001.SZ)"`
	Jzrq               *gtime.Time `json:"jzrq"               dc:"截止日期 (报告期截止日，如2025-12-31)"`
	Plrq               *gtime.Time `json:"plrq"               dc:"披露日期 (财报实际发布日期)"`
	ReportYear         int         `json:"reportYear"         dc:"报告年度"`
	ReportQuarter      int         `json:"reportQuarter"      dc:"报告季度 (1-4)"`
	ReportType         string      `json:"reportType"         dc:"报告类型: annual-年报, quarter-季报"`
	FiscalPeriod       string      `json:"fiscalPeriod"       dc:"会计期间"`
	Yysr               float64     `json:"yysr"               dc:"营业收入"`
	Yzbf               float64     `json:"yzbf"               dc:"已赚保费"`
	Fdczssr            float64     `json:"fdczssr"            dc:"房地产销售收入"`
	Qtywsr             float64     `json:"qtywsr"             dc:"其他业务收入"`
	Yyzsr              float64     `json:"yyzsr"              dc:"营业总收入"`
	Lxsr               float64     `json:"lxsr"               dc:"利息收入"`
	Sxfjyjsr           float64     `json:"sxfjyjsr"           dc:"手续费及佣金收入"`
	Btsr               float64     `json:"btsr"               dc:"补贴收入"`
	Ywsr               float64     `json:"ywsr"               dc:"营业外收入"`
	Qtsy               float64     `json:"qtsy"               dc:"其他收益"`
	Yycb               float64     `json:"yycb"               dc:"营业成本"`
	Fdczscb            float64     `json:"fdczscb"            dc:"房地产销售成本"`
	Qtywcb             float64     `json:"qtywcb"             dc:"其他业务成本"`
	Yyzcb              float64     `json:"yyzcb"              dc:"营业总成本"`
	Yysjjfj            float64     `json:"yysjjfj"            dc:"营业税金及附加"`
	Xsfy               float64     `json:"xsfy"               dc:"销售费用"`
	Glfy               float64     `json:"glfy"               dc:"管理费用"`
	Yffy               float64     `json:"yffy"               dc:"研发费用"`
	Cwfy               float64     `json:"cwfy"               dc:"财务费用"`
	Sxfjyjzc           float64     `json:"sxfjyjzc"           dc:"手续费及佣金支出"`
	Lxzc               float64     `json:"lxzc"               dc:"利息支出"`
	Tbj                float64     `json:"tbj"                dc:"退保金"`
	Pczjje             float64     `json:"pczjje"             dc:"赔付支出净额"`
	Tqbxhtzbjje        float64     `json:"tqbxhtzbjje"        dc:"提取保险合同准备金净额"`
	Bdhlzc             float64     `json:"bdhlzc"             dc:"保单红利支出"`
	Fbfy               float64     `json:"fbfy"               dc:"分保费用"`
	Zcjzss             float64     `json:"zcjzss"             dc:"资产减值损失"`
	Ywzc               float64     `json:"ywzc"               dc:"营业外支出"`
	Qtywlr             float64     `json:"qtywlr"             dc:"其他业务利润"`
	Yylr               float64     `json:"yylr"               dc:"营业利润"`
	Lrze               float64     `json:"lrze"               dc:"利润总额"`
	Jlr                float64     `json:"jlr"                dc:"净利润"`
	Jlrhfcjcx          float64     `json:"jlrhfcjcx"          dc:"净利润(扣除非经常性损益后)"`
	Gsmgsyzzdjlr       float64     `json:"gsmgsyzzdjlr"       dc:"归属于母公司所有者的净利润"`
	Bhbfzhbqsljlr      float64     `json:"bhbfzhbqsljlr"      dc:"被合并方在合并前实现净利润"`
	Tzsy               float64     `json:"tzsy"               dc:"投资收益"`
	Lyqyhhhqydtzsy     float64     `json:"lyqyhhhqydtzsy"     dc:"联营企业和合营企业的投资收益"`
	Gyjzbdsy           float64     `json:"gyjzbdsy"           dc:"公允价值变动收益"`
	Qhsy               float64     `json:"qhsy"               dc:"期货损益"`
	Tgsy               float64     `json:"tgsy"               dc:"托管收益"`
	Hdsy               float64     `json:"hdsy"               dc:"汇兑收益"`
	Fldzcczsy          float64     `json:"fldzcczsy"          dc:"非流动资产处置收益"`
	Sdsfy              float64     `json:"sdsfy"              dc:"所得税费用"`
	Ssgdsy             float64     `json:"ssgdsy"             dc:"少数股东损益"`
	Wqrtzss            float64     `json:"wqrtzss"            dc:"未确认投资损失"`
	Jbmgsy             float64     `json:"jbmgsy"             dc:"基本每股收益"`
	Xsmgsy             float64     `json:"xsmgsy"             dc:"稀释每股收益"`
	Zhsyz              float64     `json:"zhsyz"              dc:"综合收益总额"`
	Gsssgdzhsyz        float64     `json:"gsssgdzhsyz"        dc:"归属于少数股东的综合收益总额"`
	GrossMargin        float64     `json:"grossMargin"        dc:"毛利率(%)"`
	OperatingMargin    float64     `json:"operatingMargin"    dc:"营业利润率(%)"`
	NetMargin          float64     `json:"netMargin"          dc:"净利率(%)"`
	EffectiveTaxRate   float64     `json:"effectiveTaxRate"   dc:"实际税率(%)"`
	DataSource         string      `json:"dataSource"         dc:"数据来源"`
	Currency           string      `json:"currency"           dc:"货币单位"`
	Unit               string      `json:"unit"               dc:"单位: yuan-元, wan-万元"`
	AccountingStandard string      `json:"accountingStandard" dc:"会计准则 (如: CAS, IFRS)"`
	IsAudited          int         `json:"isAudited"          dc:"是否审计: 0-未审计, 1-已审计"`
	IsConsolidated     int         `json:"isConsolidated"     dc:"是否合并报表: 1-合并, 0-母公司"`
	CreatedAt          *gtime.Time `json:"createdAt"          dc:"创建时间"`
	UpdatedAt          *gtime.Time `json:"updatedAt"          dc:"更新时间"`
}

// IncomeStatementExportModel 导出利润表 (Income Statement)
type IncomeStatementExportModel struct {
	Id                 int64       `json:"id"                 dc:"自增主键"`
	Symbol             string      `json:"symbol"             dc:"股票代码 (如: 000001.SZ)"`
	Jzrq               *gtime.Time `json:"jzrq"               dc:"截止日期 (报告期截止日，如2025-12-31)"`
	Plrq               *gtime.Time `json:"plrq"               dc:"披露日期 (财报实际发布日期)"`
	ReportYear         int         `json:"reportYear"         dc:"报告年度"`
	ReportQuarter      int         `json:"reportQuarter"      dc:"报告季度 (1-4)"`
	ReportType         string      `json:"reportType"         dc:"报告类型: annual-年报, quarter-季报"`
	FiscalPeriod       string      `json:"fiscalPeriod"       dc:"会计期间"`
	Yysr               float64     `json:"yysr"               dc:"营业收入"`
	Yzbf               float64     `json:"yzbf"               dc:"已赚保费"`
	Fdczssr            float64     `json:"fdczssr"            dc:"房地产销售收入"`
	Qtywsr             float64     `json:"qtywsr"             dc:"其他业务收入"`
	Yyzsr              float64     `json:"yyzsr"              dc:"营业总收入"`
	Lxsr               float64     `json:"lxsr"               dc:"利息收入"`
	Sxfjyjsr           float64     `json:"sxfjyjsr"           dc:"手续费及佣金收入"`
	Btsr               float64     `json:"btsr"               dc:"补贴收入"`
	Ywsr               float64     `json:"ywsr"               dc:"营业外收入"`
	Qtsy               float64     `json:"qtsy"               dc:"其他收益"`
	Yycb               float64     `json:"yycb"               dc:"营业成本"`
	Fdczscb            float64     `json:"fdczscb"            dc:"房地产销售成本"`
	Qtywcb             float64     `json:"qtywcb"             dc:"其他业务成本"`
	Yyzcb              float64     `json:"yyzcb"              dc:"营业总成本"`
	Yysjjfj            float64     `json:"yysjjfj"            dc:"营业税金及附加"`
	Xsfy               float64     `json:"xsfy"               dc:"销售费用"`
	Glfy               float64     `json:"glfy"               dc:"管理费用"`
	Yffy               float64     `json:"yffy"               dc:"研发费用"`
	Cwfy               float64     `json:"cwfy"               dc:"财务费用"`
	Sxfjyjzc           float64     `json:"sxfjyjzc"           dc:"手续费及佣金支出"`
	Lxzc               float64     `json:"lxzc"               dc:"利息支出"`
	Tbj                float64     `json:"tbj"                dc:"退保金"`
	Pczjje             float64     `json:"pczjje"             dc:"赔付支出净额"`
	Tqbxhtzbjje        float64     `json:"tqbxhtzbjje"        dc:"提取保险合同准备金净额"`
	Bdhlzc             float64     `json:"bdhlzc"             dc:"保单红利支出"`
	Fbfy               float64     `json:"fbfy"               dc:"分保费用"`
	Zcjzss             float64     `json:"zcjzss"             dc:"资产减值损失"`
	Ywzc               float64     `json:"ywzc"               dc:"营业外支出"`
	Qtywlr             float64     `json:"qtywlr"             dc:"其他业务利润"`
	Yylr               float64     `json:"yylr"               dc:"营业利润"`
	Lrze               float64     `json:"lrze"               dc:"利润总额"`
	Jlr                float64     `json:"jlr"                dc:"净利润"`
	Jlrhfcjcx          float64     `json:"jlrhfcjcx"          dc:"净利润(扣除非经常性损益后)"`
	Gsmgsyzzdjlr       float64     `json:"gsmgsyzzdjlr"       dc:"归属于母公司所有者的净利润"`
	Bhbfzhbqsljlr      float64     `json:"bhbfzhbqsljlr"      dc:"被合并方在合并前实现净利润"`
	Tzsy               float64     `json:"tzsy"               dc:"投资收益"`
	Lyqyhhhqydtzsy     float64     `json:"lyqyhhhqydtzsy"     dc:"联营企业和合营企业的投资收益"`
	Gyjzbdsy           float64     `json:"gyjzbdsy"           dc:"公允价值变动收益"`
	Qhsy               float64     `json:"qhsy"               dc:"期货损益"`
	Tgsy               float64     `json:"tgsy"               dc:"托管收益"`
	Hdsy               float64     `json:"hdsy"               dc:"汇兑收益"`
	Fldzcczsy          float64     `json:"fldzcczsy"          dc:"非流动资产处置收益"`
	Sdsfy              float64     `json:"sdsfy"              dc:"所得税费用"`
	Ssgdsy             float64     `json:"ssgdsy"             dc:"少数股东损益"`
	Wqrtzss            float64     `json:"wqrtzss"            dc:"未确认投资损失"`
	Jbmgsy             float64     `json:"jbmgsy"             dc:"基本每股收益"`
	Xsmgsy             float64     `json:"xsmgsy"             dc:"稀释每股收益"`
	Zhsyz              float64     `json:"zhsyz"              dc:"综合收益总额"`
	Gsssgdzhsyz        float64     `json:"gsssgdzhsyz"        dc:"归属于少数股东的综合收益总额"`
	GrossMargin        float64     `json:"grossMargin"        dc:"毛利率(%)"`
	OperatingMargin    float64     `json:"operatingMargin"    dc:"营业利润率(%)"`
	NetMargin          float64     `json:"netMargin"          dc:"净利率(%)"`
	EffectiveTaxRate   float64     `json:"effectiveTaxRate"   dc:"实际税率(%)"`
	DataSource         string      `json:"dataSource"         dc:"数据来源"`
	Currency           string      `json:"currency"           dc:"货币单位"`
	Unit               string      `json:"unit"               dc:"单位: yuan-元, wan-万元"`
	AccountingStandard string      `json:"accountingStandard" dc:"会计准则 (如: CAS, IFRS)"`
	IsAudited          int         `json:"isAudited"          dc:"是否审计: 0-未审计, 1-已审计"`
	IsConsolidated     int         `json:"isConsolidated"     dc:"是否合并报表: 1-合并, 0-母公司"`
	CreatedAt          *gtime.Time `json:"createdAt"          dc:"创建时间"`
	UpdatedAt          *gtime.Time `json:"updatedAt"          dc:"更新时间"`
}

// IncomeStatementGetIncomeStatementInp 获取利润表数据
type IncomeStatementGetIncomeStatementInp struct {
	Symbol    string `json:"symbol" v:"required#股票代码不能为空" dc:"股票代码 (例如: 000001.SZ)"`
	Token     string `json:"token" v:"required#token证书不能为空" dc:"token证书"`
	StartTime string `json:"st" dc:"开始时间"`
	EndTime   string `json:"et" dc:"结束时间"`
}

func (in *IncomeStatementGetIncomeStatementInp) Filter(ctx context.Context) (err error) {
	return
}
