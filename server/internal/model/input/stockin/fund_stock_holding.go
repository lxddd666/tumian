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

// FundStockHoldingUpdateFields 修改基金持股明细表 (来源于基金定期报告)字段过滤
type FundStockHoldingUpdateFields struct {
	Jzrq       *gtime.Time `json:"jzrq"       dc:"截止日期 (报告期，如2025-12-31)"`
	T          *gtime.Time `json:"t"          dc:"交易时间 (衍生自jzrq，兼容时间序列查询)"`
	Jjmc       string      `json:"jjmc"       dc:"基金名称"`
	Jjdm       string      `json:"jjdm"       dc:"基金代码"`
	Symbol     string      `json:"symbol"     dc:"股票代码 (如: 000001.SZ)"`
	Ccsl       int64       `json:"ccsl"       dc:"持仓数量(股)"`
	Ltbl       float64     `json:"ltbl"       dc:"占流通股比例(%)"`
	Cgsz       float64     `json:"cgsz"       dc:"持股市值（元）"`
	Jzbl       float64     `json:"jzbl"       dc:"占净值比例（%）"`
	AvgCost    float64     `json:"avgCost"    dc:"估算持仓成本 (元/股)"`
	DataSource string      `json:"dataSource" dc:"数据来源 (如: 基金季报)"`
	ReportType string      `json:"reportType" dc:"报告类型: 季报, 中报, 年报"`
	IsLatest   int         `json:"isLatest"   dc:"是否为该基金对该股票的最新持仓: 0-否, 1-是"`
}

// FundStockHoldingInsertFields 新增基金持股明细表 (来源于基金定期报告)字段过滤
type FundStockHoldingInsertFields struct {
	Jzrq       *gtime.Time `json:"jzrq"       dc:"截止日期 (报告期，如2025-12-31)"`
	T          *gtime.Time `json:"t"          dc:"交易时间 (衍生自jzrq，兼容时间序列查询)"`
	Jjmc       string      `json:"jjmc"       dc:"基金名称"`
	Jjdm       string      `json:"jjdm"       dc:"基金代码"`
	Symbol     string      `json:"symbol"     dc:"股票代码 (如: 000001.SZ)"`
	Ccsl       int64       `json:"ccsl"       dc:"持仓数量(股)"`
	Ltbl       float64     `json:"ltbl"       dc:"占流通股比例(%)"`
	Cgsz       float64     `json:"cgsz"       dc:"持股市值（元）"`
	Jzbl       float64     `json:"jzbl"       dc:"占净值比例（%）"`
	AvgCost    float64     `json:"avgCost"    dc:"估算持仓成本 (元/股)"`
	DataSource string      `json:"dataSource" dc:"数据来源 (如: 基金季报)"`
	ReportType string      `json:"reportType" dc:"报告类型: 季报, 中报, 年报"`
	IsLatest   int         `json:"isLatest"   dc:"是否为该基金对该股票的最新持仓: 0-否, 1-是"`
}

// FundStockHoldingEditInp 修改/新增基金持股明细表 (来源于基金定期报告)
type FundStockHoldingEditInp struct {
	entity.FundStockHolding
}

func (in *FundStockHoldingEditInp) Filter(ctx context.Context) (err error) {
	// 验证截止日期 (报告期，如2025-12-31)
	if err := g.Validator().Rules("required").Data(in.Jzrq).Messages("截止日期 (报告期，如2025-12-31)不能为空").Run(ctx); err != nil {
		return err.Current()
	}
	// 验证基金名称
	if err := g.Validator().Rules("required").Data(in.Jjmc).Messages("基金名称不能为空").Run(ctx); err != nil {
		return err.Current()
	}
	// 验证基金代码
	if err := g.Validator().Rules("required").Data(in.Jjdm).Messages("基金代码不能为空").Run(ctx); err != nil {
		return err.Current()
	}
	// 验证股票代码 (如: 000001.SZ)
	if err := g.Validator().Rules("required").Data(in.Symbol).Messages("股票代码 (如: 000001.SZ)不能为空").Run(ctx); err != nil {
		return err.Current()
	}
	return
}

type FundStockHoldingEditModel struct{}

// FundStockHoldingDeleteInp 删除基金持股明细表 (来源于基金定期报告)
type FundStockHoldingDeleteInp struct {
	Id interface{} `json:"id" v:"required#自增主键不能为空" dc:"自增主键"`
}

func (in *FundStockHoldingDeleteInp) Filter(ctx context.Context) (err error) {
	return
}

type FundStockHoldingDeleteModel struct{}

// FundStockHoldingViewInp 获取指定基金持股明细表 (来源于基金定期报告)信息
type FundStockHoldingViewInp struct {
	Id int64 `json:"id" v:"required#自增主键不能为空" dc:"自增主键"`
}

func (in *FundStockHoldingViewInp) Filter(ctx context.Context) (err error) {
	return
}

type FundStockHoldingViewModel struct {
	entity.FundStockHolding
}

// FundStockHoldingListInp 获取基金持股明细表 (来源于基金定期报告)列表
type FundStockHoldingListInp struct {
	form.PageReq
	Id        int64         `json:"id"        dc:"自增主键"`
	CreatedAt []*gtime.Time `json:"createdAt" dc:"创建时间"`
}

func (in *FundStockHoldingListInp) Filter(ctx context.Context) (err error) {
	return
}

type FundStockHoldingListModel struct {
	Id         int64       `json:"id"         dc:"自增主键"`
	Jzrq       *gtime.Time `json:"jzrq"       dc:"截止日期 (报告期，如2025-12-31)"`
	T          *gtime.Time `json:"t"          dc:"交易时间 (衍生自jzrq，兼容时间序列查询)"`
	Jjmc       string      `json:"jjmc"       dc:"基金名称"`
	Jjdm       string      `json:"jjdm"       dc:"基金代码"`
	Symbol     string      `json:"symbol"     dc:"股票代码 (如: 000001.SZ)"`
	Ccsl       int64       `json:"ccsl"       dc:"持仓数量(股)"`
	Ltbl       float64     `json:"ltbl"       dc:"占流通股比例(%)"`
	Cgsz       float64     `json:"cgsz"       dc:"持股市值（元）"`
	Jzbl       float64     `json:"jzbl"       dc:"占净值比例（%）"`
	AvgCost    float64     `json:"avgCost"    dc:"估算持仓成本 (元/股)"`
	DataSource string      `json:"dataSource" dc:"数据来源 (如: 基金季报)"`
	ReportType string      `json:"reportType" dc:"报告类型: 季报, 中报, 年报"`
	IsLatest   int         `json:"isLatest"   dc:"是否为该基金对该股票的最新持仓: 0-否, 1-是"`
	CreatedAt  *gtime.Time `json:"createdAt"  dc:"创建时间"`
	UpdatedAt  *gtime.Time `json:"updatedAt"  dc:"更新时间"`
}

// FundStockHoldingExportModel 导出基金持股明细表 (来源于基金定期报告)
type FundStockHoldingExportModel struct {
	Id         int64       `json:"id"         dc:"自增主键"`
	Jzrq       *gtime.Time `json:"jzrq"       dc:"截止日期 (报告期，如2025-12-31)"`
	T          *gtime.Time `json:"t"          dc:"交易时间 (衍生自jzrq，兼容时间序列查询)"`
	Jjmc       string      `json:"jjmc"       dc:"基金名称"`
	Jjdm       string      `json:"jjdm"       dc:"基金代码"`
	Symbol     string      `json:"symbol"     dc:"股票代码 (如: 000001.SZ)"`
	Ccsl       int64       `json:"ccsl"       dc:"持仓数量(股)"`
	Ltbl       float64     `json:"ltbl"       dc:"占流通股比例(%)"`
	Cgsz       float64     `json:"cgsz"       dc:"持股市值（元）"`
	Jzbl       float64     `json:"jzbl"       dc:"占净值比例（%）"`
	AvgCost    float64     `json:"avgCost"    dc:"估算持仓成本 (元/股)"`
	DataSource string      `json:"dataSource" dc:"数据来源 (如: 基金季报)"`
	ReportType string      `json:"reportType" dc:"报告类型: 季报, 中报, 年报"`
	IsLatest   int         `json:"isLatest"   dc:"是否为该基金对该股票的最新持仓: 0-否, 1-是"`
	CreatedAt  *gtime.Time `json:"createdAt"  dc:"创建时间"`
	UpdatedAt  *gtime.Time `json:"updatedAt"  dc:"更新时间"`
}

// FundStockHoldingGetFundStockHoldingInp 获取基金持股明细表数据
type FundStockHoldingGetFundStockHoldingInp struct {
	Symbol string `json:"symbol" v:"required#股票代码不能为空" dc:"股票代码 (例如: 000001.SZ)"`
	Token  string `json:"token"  dc:"token证书"`
}

func (in *FundStockHoldingGetFundStockHoldingInp) Filter(ctx context.Context) (err error) {
	return
}
