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

// FinancialIndicatorsUpdateFields 修改财务指标分析表字段过滤
type FinancialIndicatorsUpdateFields struct {
	Symbol         string      `json:"symbol"         dc:"公司代码/股票代码 (例如: 000001.SZ, AAPL)"`
	Jzrq           *gtime.Time `json:"jzrq"           dc:"截止日期 (报告期结束日)"`
	Plrq           *gtime.Time `json:"plrq"           dc:"披露日期"`
	ReportYear     int         `json:"reportYear"     dc:"报告年度"`
	ReportQuarter  int         `json:"reportQuarter"  dc:"报告季度 (1-4)"`
	ReportType     string      `json:"reportType"     dc:"报告类型: annual-年报, quarter-季报"`
	FiscalPeriod   string      `json:"fiscalPeriod"   dc:"会计期间 (衍生字段，如2023Q1)"`
	Mgzbgjj        float64     `json:"mgzbgjj"        dc:"每股资本公积金"`
	Mgjyhdxjl      float64     `json:"mgjyhdxjl"      dc:"每股经营活动现金流量"`
	Mgjzc          float64     `json:"mgjzc"          dc:"每股净资产"`
	Jbmgsy         float64     `json:"jbmgsy"         dc:"基本每股收益"`
	Xsmgsy         float64     `json:"xsmgsy"         dc:"稀释每股收益"`
	Mgwfplr        float64     `json:"mgwfplr"        dc:"每股未分配利润"`
	Kfmgsy         float64     `json:"kfmgsy"         dc:"扣非每股收益"`
	Jzcsyl         float64     `json:"jzcsyl"         dc:"净资产收益率(%)"`
	Jqjzcsyl       float64     `json:"jqjzcsyl"       dc:"加权净资产收益率(%)"`
	Tbjzcsyl       float64     `json:"tbjzcsyl"       dc:"摊薄净资产收益率(%)"`
	Tbzzcsyl       float64     `json:"tbzzcsyl"       dc:"摊薄总资产收益率(%)"`
	Xsmlv          float64     `json:"xsmlv"          dc:"销售毛利率(%)"`
	Mlv            float64     `json:"mlv"            dc:"毛利率(%)"`
	Jlv            float64     `json:"jlv"            dc:"净利率(%)"`
	Sjslv          float64     `json:"sjslv"          dc:"实际税率(%)"`
	Zyyrsrzz       float64     `json:"zyyrsrzz"       dc:"主营收入同比增长(%)"`
	Jlrzz          float64     `json:"jlrzz"          dc:"净利润同比增长(%)"`
	Gsmgsyzzdjlrzz float64     `json:"gsmgsyzzdjlrzz" dc:"归属于母公司所有者的净利润同比增长(%)"`
	Kfjlrzz        float64     `json:"kfjlrzz"        dc:"扣非净利润同比增长(%)"`
	Yyzsrgdhbzz    float64     `json:"yyzsrgdhbzz"    dc:"营业总收入滚动环比增长(%)"`
	Sljlrjqhbzz    float64     `json:"sljlrjqhbzz"    dc:"归属净利润滚动环比增长(%)"`
	Kfjlrgdhbzz    float64     `json:"kfjlrgdhbzz"    dc:"扣非净利润滚动环比增长(%)"`
	Yskyysr        float64     `json:"yskyysr"        dc:"预收款/营业收入"`
	Xsxjlyysr      float64     `json:"xsxjlyysr"      dc:"销售现金流/营业收入"`
	Zcfzl          float64     `json:"zcfzl"          dc:"资产负债比率(%)"`
	Chzzl          float64     `json:"chzzl"          dc:"存货周转率(次)"`
	DataSource     string      `json:"dataSource"     dc:"数据来源"`
	Currency       string      `json:"currency"       dc:"货币单位"`
	Unit           string      `json:"unit"           dc:"单位: yuan-元"`
	IsCalculated   int         `json:"isCalculated"   dc:"是否为计算指标: 0-原始数据, 1-计算得出"`
	CalcVersion    string      `json:"calcVersion"    dc:"计算版本"`
}

// FinancialIndicatorsInsertFields 新增财务指标分析表字段过滤
type FinancialIndicatorsInsertFields struct {
	Symbol         string      `json:"symbol"         dc:"公司代码/股票代码 (例如: 000001.SZ, AAPL)"`
	Jzrq           *gtime.Time `json:"jzrq"           dc:"截止日期 (报告期结束日)"`
	Plrq           *gtime.Time `json:"plrq"           dc:"披露日期"`
	ReportYear     int         `json:"reportYear"     dc:"报告年度"`
	ReportQuarter  int         `json:"reportQuarter"  dc:"报告季度 (1-4)"`
	ReportType     string      `json:"reportType"     dc:"报告类型: annual-年报, quarter-季报"`
	FiscalPeriod   string      `json:"fiscalPeriod"   dc:"会计期间 (衍生字段，如2023Q1)"`
	Mgzbgjj        float64     `json:"mgzbgjj"        dc:"每股资本公积金"`
	Mgjyhdxjl      float64     `json:"mgjyhdxjl"      dc:"每股经营活动现金流量"`
	Mgjzc          float64     `json:"mgjzc"          dc:"每股净资产"`
	Jbmgsy         float64     `json:"jbmgsy"         dc:"基本每股收益"`
	Xsmgsy         float64     `json:"xsmgsy"         dc:"稀释每股收益"`
	Mgwfplr        float64     `json:"mgwfplr"        dc:"每股未分配利润"`
	Kfmgsy         float64     `json:"kfmgsy"         dc:"扣非每股收益"`
	Jzcsyl         float64     `json:"jzcsyl"         dc:"净资产收益率(%)"`
	Jqjzcsyl       float64     `json:"jqjzcsyl"       dc:"加权净资产收益率(%)"`
	Tbjzcsyl       float64     `json:"tbjzcsyl"       dc:"摊薄净资产收益率(%)"`
	Tbzzcsyl       float64     `json:"tbzzcsyl"       dc:"摊薄总资产收益率(%)"`
	Xsmlv          float64     `json:"xsmlv"          dc:"销售毛利率(%)"`
	Mlv            float64     `json:"mlv"            dc:"毛利率(%)"`
	Jlv            float64     `json:"jlv"            dc:"净利率(%)"`
	Sjslv          float64     `json:"sjslv"          dc:"实际税率(%)"`
	Zyyrsrzz       float64     `json:"zyyrsrzz"       dc:"主营收入同比增长(%)"`
	Jlrzz          float64     `json:"jlrzz"          dc:"净利润同比增长(%)"`
	Gsmgsyzzdjlrzz float64     `json:"gsmgsyzzdjlrzz" dc:"归属于母公司所有者的净利润同比增长(%)"`
	Kfjlrzz        float64     `json:"kfjlrzz"        dc:"扣非净利润同比增长(%)"`
	Yyzsrgdhbzz    float64     `json:"yyzsrgdhbzz"    dc:"营业总收入滚动环比增长(%)"`
	Sljlrjqhbzz    float64     `json:"sljlrjqhbzz"    dc:"归属净利润滚动环比增长(%)"`
	Kfjlrgdhbzz    float64     `json:"kfjlrgdhbzz"    dc:"扣非净利润滚动环比增长(%)"`
	Yskyysr        float64     `json:"yskyysr"        dc:"预收款/营业收入"`
	Xsxjlyysr      float64     `json:"xsxjlyysr"      dc:"销售现金流/营业收入"`
	Zcfzl          float64     `json:"zcfzl"          dc:"资产负债比率(%)"`
	Chzzl          float64     `json:"chzzl"          dc:"存货周转率(次)"`
	DataSource     string      `json:"dataSource"     dc:"数据来源"`
	Currency       string      `json:"currency"       dc:"货币单位"`
	Unit           string      `json:"unit"           dc:"单位: yuan-元"`
	IsCalculated   int         `json:"isCalculated"   dc:"是否为计算指标: 0-原始数据, 1-计算得出"`
	CalcVersion    string      `json:"calcVersion"    dc:"计算版本"`
}

// FinancialIndicatorsEditInp 修改/新增财务指标分析表
type FinancialIndicatorsEditInp struct {
	entity.FinancialIndicators
}

func (in *FinancialIndicatorsEditInp) Filter(ctx context.Context) (err error) {
	// 验证公司代码/股票代码 (例如: 000001.SZ, AAPL)
	if err := g.Validator().Rules("required").Data(in.Symbol).Messages("公司代码/股票代码 (例如: 000001.SZ, AAPL)不能为空").Run(ctx); err != nil {
		return err.Current()
	}
	// 验证截止日期 (报告期结束日)
	if err := g.Validator().Rules("required").Data(in.Jzrq).Messages("截止日期 (报告期结束日)不能为空").Run(ctx); err != nil {
		return err.Current()
	}
	// 验证披露日期
	if err := g.Validator().Rules("required").Data(in.Plrq).Messages("披露日期不能为空").Run(ctx); err != nil {
		return err.Current()
	}
	// 验证报告年度
	if err := g.Validator().Rules("required").Data(in.ReportYear).Messages("报告年度不能为空").Run(ctx); err != nil {
		return err.Current()
	}
	// 验证报告类型: annual-年报, quarter-季报
	if err := g.Validator().Rules("required").Data(in.ReportType).Messages("报告类型: annual-年报, quarter-季报不能为空").Run(ctx); err != nil {
		return err.Current()
	}
	return
}

type FinancialIndicatorsEditModel struct{}

// FinancialIndicatorsDeleteInp 删除财务指标分析表
type FinancialIndicatorsDeleteInp struct {
	Id interface{} `json:"id" v:"required#自增主键不能为空" dc:"自增主键"`
}

func (in *FinancialIndicatorsDeleteInp) Filter(ctx context.Context) (err error) {
	return
}

type FinancialIndicatorsDeleteModel struct{}

// FinancialIndicatorsViewInp 获取指定财务指标分析表信息
type FinancialIndicatorsViewInp struct {
	Id int64 `json:"id" v:"required#自增主键不能为空" dc:"自增主键"`
}

func (in *FinancialIndicatorsViewInp) Filter(ctx context.Context) (err error) {
	return
}

type FinancialIndicatorsViewModel struct {
	entity.FinancialIndicators
}

// FinancialIndicatorsListInp 获取财务指标分析表列表
type FinancialIndicatorsListInp struct {
	form.PageReq
	Id        int64         `json:"id"        dc:"自增主键"`
	CreatedAt []*gtime.Time `json:"createdAt" dc:"创建时间"`
}

func (in *FinancialIndicatorsListInp) Filter(ctx context.Context) (err error) {
	return
}

type FinancialIndicatorsListModel struct {
	Id             int64       `json:"id"             dc:"自增主键"`
	Symbol         string      `json:"symbol"         dc:"公司代码/股票代码 (例如: 000001.SZ, AAPL)"`
	Jzrq           *gtime.Time `json:"jzrq"           dc:"截止日期 (报告期结束日)"`
	Plrq           *gtime.Time `json:"plrq"           dc:"披露日期"`
	ReportYear     int         `json:"reportYear"     dc:"报告年度"`
	ReportQuarter  int         `json:"reportQuarter"  dc:"报告季度 (1-4)"`
	ReportType     string      `json:"reportType"     dc:"报告类型: annual-年报, quarter-季报"`
	FiscalPeriod   string      `json:"fiscalPeriod"   dc:"会计期间 (衍生字段，如2023Q1)"`
	Mgzbgjj        float64     `json:"mgzbgjj"        dc:"每股资本公积金"`
	Mgjyhdxjl      float64     `json:"mgjyhdxjl"      dc:"每股经营活动现金流量"`
	Mgjzc          float64     `json:"mgjzc"          dc:"每股净资产"`
	Jbmgsy         float64     `json:"jbmgsy"         dc:"基本每股收益"`
	Xsmgsy         float64     `json:"xsmgsy"         dc:"稀释每股收益"`
	Mgwfplr        float64     `json:"mgwfplr"        dc:"每股未分配利润"`
	Kfmgsy         float64     `json:"kfmgsy"         dc:"扣非每股收益"`
	Jzcsyl         float64     `json:"jzcsyl"         dc:"净资产收益率(%)"`
	Jqjzcsyl       float64     `json:"jqjzcsyl"       dc:"加权净资产收益率(%)"`
	Tbjzcsyl       float64     `json:"tbjzcsyl"       dc:"摊薄净资产收益率(%)"`
	Tbzzcsyl       float64     `json:"tbzzcsyl"       dc:"摊薄总资产收益率(%)"`
	Xsmlv          float64     `json:"xsmlv"          dc:"销售毛利率(%)"`
	Mlv            float64     `json:"mlv"            dc:"毛利率(%)"`
	Jlv            float64     `json:"jlv"            dc:"净利率(%)"`
	Sjslv          float64     `json:"sjslv"          dc:"实际税率(%)"`
	Zyyrsrzz       float64     `json:"zyyrsrzz"       dc:"主营收入同比增长(%)"`
	Jlrzz          float64     `json:"jlrzz"          dc:"净利润同比增长(%)"`
	Gsmgsyzzdjlrzz float64     `json:"gsmgsyzzdjlrzz" dc:"归属于母公司所有者的净利润同比增长(%)"`
	Kfjlrzz        float64     `json:"kfjlrzz"        dc:"扣非净利润同比增长(%)"`
	Yyzsrgdhbzz    float64     `json:"yyzsrgdhbzz"    dc:"营业总收入滚动环比增长(%)"`
	Sljlrjqhbzz    float64     `json:"sljlrjqhbzz"    dc:"归属净利润滚动环比增长(%)"`
	Kfjlrgdhbzz    float64     `json:"kfjlrgdhbzz"    dc:"扣非净利润滚动环比增长(%)"`
	Yskyysr        float64     `json:"yskyysr"        dc:"预收款/营业收入"`
	Xsxjlyysr      float64     `json:"xsxjlyysr"      dc:"销售现金流/营业收入"`
	Zcfzl          float64     `json:"zcfzl"          dc:"资产负债比率(%)"`
	Chzzl          float64     `json:"chzzl"          dc:"存货周转率(次)"`
	DataSource     string      `json:"dataSource"     dc:"数据来源"`
	Currency       string      `json:"currency"       dc:"货币单位"`
	Unit           string      `json:"unit"           dc:"单位: yuan-元"`
	IsCalculated   int         `json:"isCalculated"   dc:"是否为计算指标: 0-原始数据, 1-计算得出"`
	CalcVersion    string      `json:"calcVersion"    dc:"计算版本"`
	CreatedAt      *gtime.Time `json:"createdAt"      dc:"创建时间"`
	UpdatedAt      *gtime.Time `json:"updatedAt"      dc:"更新时间"`
}

// FinancialIndicatorsExportModel 导出财务指标分析表
type FinancialIndicatorsExportModel struct {
	Id             int64       `json:"id"             dc:"自增主键"`
	Symbol         string      `json:"symbol"         dc:"公司代码/股票代码 (例如: 000001.SZ, AAPL)"`
	Jzrq           *gtime.Time `json:"jzrq"           dc:"截止日期 (报告期结束日)"`
	Plrq           *gtime.Time `json:"plrq"           dc:"披露日期"`
	ReportYear     int         `json:"reportYear"     dc:"报告年度"`
	ReportQuarter  int         `json:"reportQuarter"  dc:"报告季度 (1-4)"`
	ReportType     string      `json:"reportType"     dc:"报告类型: annual-年报, quarter-季报"`
	FiscalPeriod   string      `json:"fiscalPeriod"   dc:"会计期间 (衍生字段，如2023Q1)"`
	Mgzbgjj        float64     `json:"mgzbgjj"        dc:"每股资本公积金"`
	Mgjyhdxjl      float64     `json:"mgjyhdxjl"      dc:"每股经营活动现金流量"`
	Mgjzc          float64     `json:"mgjzc"          dc:"每股净资产"`
	Jbmgsy         float64     `json:"jbmgsy"         dc:"基本每股收益"`
	Xsmgsy         float64     `json:"xsmgsy"         dc:"稀释每股收益"`
	Mgwfplr        float64     `json:"mgwfplr"        dc:"每股未分配利润"`
	Kfmgsy         float64     `json:"kfmgsy"         dc:"扣非每股收益"`
	Jzcsyl         float64     `json:"jzcsyl"         dc:"净资产收益率(%)"`
	Jqjzcsyl       float64     `json:"jqjzcsyl"       dc:"加权净资产收益率(%)"`
	Tbjzcsyl       float64     `json:"tbjzcsyl"       dc:"摊薄净资产收益率(%)"`
	Tbzzcsyl       float64     `json:"tbzzcsyl"       dc:"摊薄总资产收益率(%)"`
	Xsmlv          float64     `json:"xsmlv"          dc:"销售毛利率(%)"`
	Mlv            float64     `json:"mlv"            dc:"毛利率(%)"`
	Jlv            float64     `json:"jlv"            dc:"净利率(%)"`
	Sjslv          float64     `json:"sjslv"          dc:"实际税率(%)"`
	Zyyrsrzz       float64     `json:"zyyrsrzz"       dc:"主营收入同比增长(%)"`
	Jlrzz          float64     `json:"jlrzz"          dc:"净利润同比增长(%)"`
	Gsmgsyzzdjlrzz float64     `json:"gsmgsyzzdjlrzz" dc:"归属于母公司所有者的净利润同比增长(%)"`
	Kfjlrzz        float64     `json:"kfjlrzz"        dc:"扣非净利润同比增长(%)"`
	Yyzsrgdhbzz    float64     `json:"yyzsrgdhbzz"    dc:"营业总收入滚动环比增长(%)"`
	Sljlrjqhbzz    float64     `json:"sljlrjqhbzz"    dc:"归属净利润滚动环比增长(%)"`
	Kfjlrgdhbzz    float64     `json:"kfjlrgdhbzz"    dc:"扣非净利润滚动环比增长(%)"`
	Yskyysr        float64     `json:"yskyysr"        dc:"预收款/营业收入"`
	Xsxjlyysr      float64     `json:"xsxjlyysr"      dc:"销售现金流/营业收入"`
	Zcfzl          float64     `json:"zcfzl"          dc:"资产负债比率(%)"`
	Chzzl          float64     `json:"chzzl"          dc:"存货周转率(次)"`
	DataSource     string      `json:"dataSource"     dc:"数据来源"`
	Currency       string      `json:"currency"       dc:"货币单位"`
	Unit           string      `json:"unit"           dc:"单位: yuan-元"`
	IsCalculated   int         `json:"isCalculated"   dc:"是否为计算指标: 0-原始数据, 1-计算得出"`
	CalcVersion    string      `json:"calcVersion"    dc:"计算版本"`
	CreatedAt      *gtime.Time `json:"createdAt"      dc:"创建时间"`
	UpdatedAt      *gtime.Time `json:"updatedAt"      dc:"更新时间"`
}

// FinancialIndicatorsGetFinancialIndicatorsInp 获取财务指标分析表数据
type FinancialIndicatorsGetFinancialIndicatorsInp struct {
	Symbol     string `json:"symbol" v:"required#股票代码不能为空" dc:"股票代码 (例如: 000001.SZ)"`
	Interval   string `json:"interval"  dc:"分时级别 (例如: d)"`
	AdjustType string `json:"adjustType"  dc:"除权类型 (例如: n)"`
	Token      string `json:"token"  dc:"token证书"`
	StartTime  string `json:"st" dc:"开始时间"`
	EndTime    string `json:"et" dc:"结束时间"`
	Limit      int    `json:"lt" dc:"最新条数"`
}

func (in *FinancialIndicatorsGetFinancialIndicatorsInp) Filter(ctx context.Context) (err error) {
	return
}
