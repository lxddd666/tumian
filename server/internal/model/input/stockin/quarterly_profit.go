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

// QuarterlyProfitUpdateFields 修改季度利润数据表 (近一年各季度)字段过滤
type QuarterlyProfitUpdateFields struct {
	Symbol               string      `json:"symbol"               dc:"股票代码 (如: 000001.SZ)"`
	Date                 *gtime.Time `json:"date"                 dc:"截止日期 (报告期截止日，如2025-03-31)"`
	ReportYear           int         `json:"reportYear"           dc:"报告年度"`
	ReportQuarter        int         `json:"reportQuarter"        dc:"报告季度 (1-4)"`
	FiscalPeriod         string      `json:"fiscalPeriod"         dc:"会计期间 (如2025Q1)"`
	Income               float64     `json:"income"               dc:"营业收入（万元）"`
	Expend               float64     `json:"expend"               dc:"营业支出（万元）"`
	Profit               float64     `json:"profit"               dc:"营业利润（万元）"`
	Totalp               float64     `json:"totalp"               dc:"利润总额（万元）"`
	Reprofit             float64     `json:"reprofit"             dc:"净利润（万元）"`
	Basege               float64     `json:"basege"               dc:"基本每股收益(元/股)"`
	Ettege               float64     `json:"ettege"               dc:"稀释每股收益(元/股)"`
	Otherp               float64     `json:"otherp"               dc:"其他综合收益（万元）"`
	Totalcp              float64     `json:"totalcp"              dc:"综合收益总额（万元）"`
	GrossProfitMargin    float64     `json:"grossProfitMargin"    dc:"毛利率(%)"`
	NetProfitMargin      float64     `json:"netProfitMargin"      dc:"净利率(%)"`
	OperatingProfitRatio float64     `json:"operatingProfitRatio" dc:"营业利润率(%)"`
	ReportType           string      `json:"reportType"           dc:"报告类型: 一季报, 中报, 三季报, 年报"`
	DataSource           string      `json:"dataSource"           dc:"数据来源 (如: 交易所财报)"`
	Currency             string      `json:"currency"             dc:"货币单位"`
	IsAudited            int         `json:"isAudited"            dc:"是否审计: 0-未审计, 1-已审计"`
}

// QuarterlyProfitInsertFields 新增季度利润数据表 (近一年各季度)字段过滤
type QuarterlyProfitInsertFields struct {
	Symbol               string      `json:"symbol"               dc:"股票代码 (如: 000001.SZ)"`
	Date                 *gtime.Time `json:"date"                 dc:"截止日期 (报告期截止日，如2025-03-31)"`
	ReportYear           int         `json:"reportYear"           dc:"报告年度"`
	ReportQuarter        int         `json:"reportQuarter"        dc:"报告季度 (1-4)"`
	FiscalPeriod         string      `json:"fiscalPeriod"         dc:"会计期间 (如2025Q1)"`
	Income               float64     `json:"income"               dc:"营业收入（万元）"`
	Expend               float64     `json:"expend"               dc:"营业支出（万元）"`
	Profit               float64     `json:"profit"               dc:"营业利润（万元）"`
	Totalp               float64     `json:"totalp"               dc:"利润总额（万元）"`
	Reprofit             float64     `json:"reprofit"             dc:"净利润（万元）"`
	Basege               float64     `json:"basege"               dc:"基本每股收益(元/股)"`
	Ettege               float64     `json:"ettege"               dc:"稀释每股收益(元/股)"`
	Otherp               float64     `json:"otherp"               dc:"其他综合收益（万元）"`
	Totalcp              float64     `json:"totalcp"              dc:"综合收益总额（万元）"`
	GrossProfitMargin    float64     `json:"grossProfitMargin"    dc:"毛利率(%)"`
	NetProfitMargin      float64     `json:"netProfitMargin"      dc:"净利率(%)"`
	OperatingProfitRatio float64     `json:"operatingProfitRatio" dc:"营业利润率(%)"`
	ReportType           string      `json:"reportType"           dc:"报告类型: 一季报, 中报, 三季报, 年报"`
	DataSource           string      `json:"dataSource"           dc:"数据来源 (如: 交易所财报)"`
	Currency             string      `json:"currency"             dc:"货币单位"`
	IsAudited            int         `json:"isAudited"            dc:"是否审计: 0-未审计, 1-已审计"`
}

// QuarterlyProfitEditInp 修改/新增季度利润数据表 (近一年各季度)
type QuarterlyProfitEditInp struct {
	entity.QuarterlyProfit
}

func (in *QuarterlyProfitEditInp) Filter(ctx context.Context) (err error) {
	// 验证股票代码 (如: 000001.SZ)
	if err := g.Validator().Rules("required").Data(in.Symbol).Messages("股票代码 (如: 000001.SZ)不能为空").Run(ctx); err != nil {
		return err.Current()
	}
	// 验证截止日期 (报告期截止日，如2025-03-31)
	if err := g.Validator().Rules("required").Data(in.Date).Messages("截止日期 (报告期截止日，如2025-03-31)不能为空").Run(ctx); err != nil {
		return err.Current()
	}
	return
}

type QuarterlyProfitEditModel struct{}

// QuarterlyProfitDeleteInp 删除季度利润数据表 (近一年各季度)
type QuarterlyProfitDeleteInp struct {
	Id interface{} `json:"id" v:"required#自增主键不能为空" dc:"自增主键"`
}

func (in *QuarterlyProfitDeleteInp) Filter(ctx context.Context) (err error) {
	return
}

type QuarterlyProfitDeleteModel struct{}

// QuarterlyProfitViewInp 获取指定季度利润数据表 (近一年各季度)信息
type QuarterlyProfitViewInp struct {
	Id int64 `json:"id" v:"required#自增主键不能为空" dc:"自增主键"`
}

func (in *QuarterlyProfitViewInp) Filter(ctx context.Context) (err error) {
	return
}

type QuarterlyProfitViewModel struct {
	entity.QuarterlyProfit
}

// QuarterlyProfitListInp 获取季度利润数据表 (近一年各季度)列表
type QuarterlyProfitListInp struct {
	form.PageReq
	Id        int64         `json:"id"        dc:"自增主键"`
	CreatedAt []*gtime.Time `json:"createdAt" dc:"创建时间"`
}

func (in *QuarterlyProfitListInp) Filter(ctx context.Context) (err error) {
	return
}

type QuarterlyProfitListModel struct {
	Id                   int64       `json:"id"                   dc:"自增主键"`
	Symbol               string      `json:"symbol"               dc:"股票代码 (如: 000001.SZ)"`
	Date                 *gtime.Time `json:"date"                 dc:"截止日期 (报告期截止日，如2025-03-31)"`
	ReportYear           int         `json:"reportYear"           dc:"报告年度"`
	ReportQuarter        int         `json:"reportQuarter"        dc:"报告季度 (1-4)"`
	FiscalPeriod         string      `json:"fiscalPeriod"         dc:"会计期间 (如2025Q1)"`
	Income               float64     `json:"income"               dc:"营业收入（万元）"`
	Expend               float64     `json:"expend"               dc:"营业支出（万元）"`
	Profit               float64     `json:"profit"               dc:"营业利润（万元）"`
	Totalp               float64     `json:"totalp"               dc:"利润总额（万元）"`
	Reprofit             float64     `json:"reprofit"             dc:"净利润（万元）"`
	Basege               float64     `json:"basege"               dc:"基本每股收益(元/股)"`
	Ettege               float64     `json:"ettege"               dc:"稀释每股收益(元/股)"`
	Otherp               float64     `json:"otherp"               dc:"其他综合收益（万元）"`
	Totalcp              float64     `json:"totalcp"              dc:"综合收益总额（万元）"`
	GrossProfitMargin    float64     `json:"grossProfitMargin"    dc:"毛利率(%)"`
	NetProfitMargin      float64     `json:"netProfitMargin"      dc:"净利率(%)"`
	OperatingProfitRatio float64     `json:"operatingProfitRatio" dc:"营业利润率(%)"`
	ReportType           string      `json:"reportType"           dc:"报告类型: 一季报, 中报, 三季报, 年报"`
	DataSource           string      `json:"dataSource"           dc:"数据来源 (如: 交易所财报)"`
	Currency             string      `json:"currency"             dc:"货币单位"`
	IsAudited            int         `json:"isAudited"            dc:"是否审计: 0-未审计, 1-已审计"`
	CreatedAt            *gtime.Time `json:"createdAt"            dc:"创建时间"`
	UpdatedAt            *gtime.Time `json:"updatedAt"            dc:"更新时间"`
}

// QuarterlyProfitExportModel 导出季度利润数据表 (近一年各季度)
type QuarterlyProfitExportModel struct {
	Id                   int64       `json:"id"                   dc:"自增主键"`
	Symbol               string      `json:"symbol"               dc:"股票代码 (如: 000001.SZ)"`
	Date                 *gtime.Time `json:"date"                 dc:"截止日期 (报告期截止日，如2025-03-31)"`
	ReportYear           int         `json:"reportYear"           dc:"报告年度"`
	ReportQuarter        int         `json:"reportQuarter"        dc:"报告季度 (1-4)"`
	FiscalPeriod         string      `json:"fiscalPeriod"         dc:"会计期间 (如2025Q1)"`
	Income               float64     `json:"income"               dc:"营业收入（万元）"`
	Expend               float64     `json:"expend"               dc:"营业支出（万元）"`
	Profit               float64     `json:"profit"               dc:"营业利润（万元）"`
	Totalp               float64     `json:"totalp"               dc:"利润总额（万元）"`
	Reprofit             float64     `json:"reprofit"             dc:"净利润（万元）"`
	Basege               float64     `json:"basege"               dc:"基本每股收益(元/股)"`
	Ettege               float64     `json:"ettege"               dc:"稀释每股收益(元/股)"`
	Otherp               float64     `json:"otherp"               dc:"其他综合收益（万元）"`
	Totalcp              float64     `json:"totalcp"              dc:"综合收益总额（万元）"`
	GrossProfitMargin    float64     `json:"grossProfitMargin"    dc:"毛利率(%)"`
	NetProfitMargin      float64     `json:"netProfitMargin"      dc:"净利率(%)"`
	OperatingProfitRatio float64     `json:"operatingProfitRatio" dc:"营业利润率(%)"`
	ReportType           string      `json:"reportType"           dc:"报告类型: 一季报, 中报, 三季报, 年报"`
	DataSource           string      `json:"dataSource"           dc:"数据来源 (如: 交易所财报)"`
	Currency             string      `json:"currency"             dc:"货币单位"`
	IsAudited            int         `json:"isAudited"            dc:"是否审计: 0-未审计, 1-已审计"`
	CreatedAt            *gtime.Time `json:"createdAt"            dc:"创建时间"`
	UpdatedAt            *gtime.Time `json:"updatedAt"            dc:"更新时间"`
}
