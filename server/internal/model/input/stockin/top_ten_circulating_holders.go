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

// TopTenCirculatingHoldersUpdateFields 修改公司十大流通股东表 (数据来源于定期报告)[citation:4]字段过滤
type TopTenCirculatingHoldersUpdateFields struct {
	Symbol        string      `json:"symbol"        dc:"公司代码/股票代码 (例如: 000001.SZ)"`
	Jzrq          *gtime.Time `json:"jzrq"          dc:"截止日期 (报告期结束日, 如2023-09-30)[citation:4]"`
	Ggrq          *gtime.Time `json:"ggrq"          dc:"公告日期 (信息发布日期)"`
	ReportYear    int         `json:"reportYear"    dc:"报告年度"`
	ReportQuarter int         `json:"reportQuarter" dc:"报告季度 (1-4)"`
	ReportType    string      `json:"reportType"    dc:"报告类型: annual-年报, half_year-中报, quarter-季报[citation:4]"`
	Gdmc          string      `json:"gdmc"          dc:"股东名称"`
	Gdlx          string      `json:"gdlx"          dc:"股东类型 (如: 基金、社保、个人等)"`
	Gfxz          string      `json:"gfxz"          dc:"股份性质 (如: 流通A股、限售A股等)"`
	Cgsl          int64       `json:"cgsl"          dc:"持股数量 (股)"`
	Cgbl          float64     `json:"cgbl"          dc:"持股比例 (%)"`
	Cgpm          int         `json:"cgpm"          dc:"持股排名 (1-10)"`
	Bdyy          string      `json:"bdyy"          dc:"变动原因"`
	BdType        string      `json:"bdType"        dc:"变动类型: increase-增持, decrease-减持, new-新进, unchanged-不变, other-其他"`
	DataSource    string      `json:"dataSource"    dc:"数据来源 (如: 交易所公告)[citation:4]"`
	IsLatest      int         `json:"isLatest"      dc:"是否为该报告期最新数据: 0-历史快照, 1-最新"`
	Version       int         `json:"version"       dc:"数据版本"`
}

// TopTenCirculatingHoldersInsertFields 新增公司十大流通股东表 (数据来源于定期报告)[citation:4]字段过滤
type TopTenCirculatingHoldersInsertFields struct {
	Symbol        string      `json:"symbol"        dc:"公司代码/股票代码 (例如: 000001.SZ)"`
	Jzrq          *gtime.Time `json:"jzrq"          dc:"截止日期 (报告期结束日, 如2023-09-30)[citation:4]"`
	Ggrq          *gtime.Time `json:"ggrq"          dc:"公告日期 (信息发布日期)"`
	ReportYear    int         `json:"reportYear"    dc:"报告年度"`
	ReportQuarter int         `json:"reportQuarter" dc:"报告季度 (1-4)"`
	ReportType    string      `json:"reportType"    dc:"报告类型: annual-年报, half_year-中报, quarter-季报[citation:4]"`
	Gdmc          string      `json:"gdmc"          dc:"股东名称"`
	Gdlx          string      `json:"gdlx"          dc:"股东类型 (如: 基金、社保、个人等)"`
	Gfxz          string      `json:"gfxz"          dc:"股份性质 (如: 流通A股、限售A股等)"`
	Cgsl          int64       `json:"cgsl"          dc:"持股数量 (股)"`
	Cgbl          float64     `json:"cgbl"          dc:"持股比例 (%)"`
	Cgpm          int         `json:"cgpm"          dc:"持股排名 (1-10)"`
	Bdyy          string      `json:"bdyy"          dc:"变动原因"`
	BdType        string      `json:"bdType"        dc:"变动类型: increase-增持, decrease-减持, new-新进, unchanged-不变, other-其他"`
	DataSource    string      `json:"dataSource"    dc:"数据来源 (如: 交易所公告)[citation:4]"`
	IsLatest      int         `json:"isLatest"      dc:"是否为该报告期最新数据: 0-历史快照, 1-最新"`
	Version       int         `json:"version"       dc:"数据版本"`
}

// TopTenCirculatingHoldersEditInp 修改/新增公司十大流通股东表 (数据来源于定期报告)[citation:4]
type TopTenCirculatingHoldersEditInp struct {
	entity.TopTenCirculatingHolders
}

func (in *TopTenCirculatingHoldersEditInp) Filter(ctx context.Context) (err error) {
	// 验证公司代码/股票代码 (例如: 000001.SZ)
	if err := g.Validator().Rules("required").Data(in.Symbol).Messages("公司代码/股票代码 (例如: 000001.SZ)不能为空").Run(ctx); err != nil {
		return err.Current()
	}
	// 验证截止日期 (报告期结束日, 如2023-09-30)[citation:4]
	if err := g.Validator().Rules("required").Data(in.Jzrq).Messages("截止日期 (报告期结束日, 如2023-09-30)[citation:4]不能为空").Run(ctx); err != nil {
		return err.Current()
	}
	// 验证公告日期 (信息发布日期)
	if err := g.Validator().Rules("required").Data(in.Ggrq).Messages("公告日期 (信息发布日期)不能为空").Run(ctx); err != nil {
		return err.Current()
	}
	// 验证报告年度
	if err := g.Validator().Rules("required").Data(in.ReportYear).Messages("报告年度不能为空").Run(ctx); err != nil {
		return err.Current()
	}
	// 验证报告类型: annual-年报, half_year-中报, quarter-季报[citation:4]
	if err := g.Validator().Rules("required").Data(in.ReportType).Messages("报告类型: annual-年报, half_year-中报, quarter-季报[citation:4]不能为空").Run(ctx); err != nil {
		return err.Current()
	}
	// 验证股东名称
	if err := g.Validator().Rules("required").Data(in.Gdmc).Messages("股东名称不能为空").Run(ctx); err != nil {
		return err.Current()
	}
	// 验证持股排名 (1-10)
	if err := g.Validator().Rules("required").Data(in.Cgpm).Messages("持股排名 (1-10)不能为空").Run(ctx); err != nil {
		return err.Current()
	}
	return
}

type TopTenCirculatingHoldersEditModel struct{}

// TopTenCirculatingHoldersDeleteInp 删除公司十大流通股东表 (数据来源于定期报告)[citation:4]
type TopTenCirculatingHoldersDeleteInp struct {
	Id interface{} `json:"id" v:"required#自增主键不能为空" dc:"自增主键"`
}

func (in *TopTenCirculatingHoldersDeleteInp) Filter(ctx context.Context) (err error) {
	return
}

type TopTenCirculatingHoldersDeleteModel struct{}

// TopTenCirculatingHoldersViewInp 获取指定公司十大流通股东表 (数据来源于定期报告)[citation:4]信息
type TopTenCirculatingHoldersViewInp struct {
	Id int64 `json:"id" v:"required#自增主键不能为空" dc:"自增主键"`
}

func (in *TopTenCirculatingHoldersViewInp) Filter(ctx context.Context) (err error) {
	return
}

type TopTenCirculatingHoldersViewModel struct {
	entity.TopTenCirculatingHolders
}

// TopTenCirculatingHoldersListInp 获取公司十大流通股东表 (数据来源于定期报告)[citation:4]列表
type TopTenCirculatingHoldersListInp struct {
	form.PageReq
	Id        int64         `json:"id"        dc:"自增主键"`
	CreatedAt []*gtime.Time `json:"createdAt" dc:"创建时间"`
}

func (in *TopTenCirculatingHoldersListInp) Filter(ctx context.Context) (err error) {
	return
}

type TopTenCirculatingHoldersListModel struct {
	Id            int64       `json:"id"            dc:"自增主键"`
	Symbol        string      `json:"symbol"        dc:"公司代码/股票代码 (例如: 000001.SZ)"`
	Jzrq          *gtime.Time `json:"jzrq"          dc:"截止日期 (报告期结束日, 如2023-09-30)[citation:4]"`
	Ggrq          *gtime.Time `json:"ggrq"          dc:"公告日期 (信息发布日期)"`
	ReportYear    int         `json:"reportYear"    dc:"报告年度"`
	ReportQuarter int         `json:"reportQuarter" dc:"报告季度 (1-4)"`
	ReportType    string      `json:"reportType"    dc:"报告类型: annual-年报, half_year-中报, quarter-季报[citation:4]"`
	Gdmc          string      `json:"gdmc"          dc:"股东名称"`
	Gdlx          string      `json:"gdlx"          dc:"股东类型 (如: 基金、社保、个人等)"`
	Gfxz          string      `json:"gfxz"          dc:"股份性质 (如: 流通A股、限售A股等)"`
	Cgsl          int64       `json:"cgsl"          dc:"持股数量 (股)"`
	Cgbl          float64     `json:"cgbl"          dc:"持股比例 (%)"`
	Cgpm          int         `json:"cgpm"          dc:"持股排名 (1-10)"`
	BdType        string      `json:"bdType"        dc:"变动类型: increase-增持, decrease-减持, new-新进, unchanged-不变, other-其他"`
	DataSource    string      `json:"dataSource"    dc:"数据来源 (如: 交易所公告)[citation:4]"`
	IsLatest      int         `json:"isLatest"      dc:"是否为该报告期最新数据: 0-历史快照, 1-最新"`
	Version       int         `json:"version"       dc:"数据版本"`
	CreatedAt     *gtime.Time `json:"createdAt"     dc:"创建时间"`
	UpdatedAt     *gtime.Time `json:"updatedAt"     dc:"更新时间"`
}

// TopTenCirculatingHoldersExportModel 导出公司十大流通股东表 (数据来源于定期报告)[citation:4]
type TopTenCirculatingHoldersExportModel struct {
	Id            int64       `json:"id"            dc:"自增主键"`
	Symbol        string      `json:"symbol"        dc:"公司代码/股票代码 (例如: 000001.SZ)"`
	Jzrq          *gtime.Time `json:"jzrq"          dc:"截止日期 (报告期结束日, 如2023-09-30)[citation:4]"`
	Ggrq          *gtime.Time `json:"ggrq"          dc:"公告日期 (信息发布日期)"`
	ReportYear    int         `json:"reportYear"    dc:"报告年度"`
	ReportQuarter int         `json:"reportQuarter" dc:"报告季度 (1-4)"`
	ReportType    string      `json:"reportType"    dc:"报告类型: annual-年报, half_year-中报, quarter-季报[citation:4]"`
	Gdmc          string      `json:"gdmc"          dc:"股东名称"`
	Gdlx          string      `json:"gdlx"          dc:"股东类型 (如: 基金、社保、个人等)"`
	Gfxz          string      `json:"gfxz"          dc:"股份性质 (如: 流通A股、限售A股等)"`
	Cgsl          int64       `json:"cgsl"          dc:"持股数量 (股)"`
	Cgbl          float64     `json:"cgbl"          dc:"持股比例 (%)"`
	Cgpm          int         `json:"cgpm"          dc:"持股排名 (1-10)"`
	BdType        string      `json:"bdType"        dc:"变动类型: increase-增持, decrease-减持, new-新进, unchanged-不变, other-其他"`
	DataSource    string      `json:"dataSource"    dc:"数据来源 (如: 交易所公告)[citation:4]"`
	IsLatest      int         `json:"isLatest"      dc:"是否为该报告期最新数据: 0-历史快照, 1-最新"`
	Version       int         `json:"version"       dc:"数据版本"`
	CreatedAt     *gtime.Time `json:"createdAt"     dc:"创建时间"`
	UpdatedAt     *gtime.Time `json:"updatedAt"     dc:"更新时间"`
}
