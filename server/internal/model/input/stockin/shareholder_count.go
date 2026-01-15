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

// ShareholderCountUpdateFields 修改公司股东户数统计表 (按报告期统计)字段过滤
type ShareholderCountUpdateFields struct {
	Symbol        string      `json:"symbol"        dc:"公司代码/股票代码 (例如: 000001.SZ, 600000.SS)"`
	Jzrq          *gtime.Time `json:"jzrq"          dc:"截止日期 (报告期结束日, 如2023-09-30)"`
	ReportYear    int         `json:"reportYear"    dc:"报告年度"`
	ReportQuarter int         `json:"reportQuarter" dc:"报告季度 (1-4)"`
	ReportType    string      `json:"reportType"    dc:"报告类型: annual-年报, half_year-中报, quarter-季报"`
	Gdzs          int         `json:"gdzs"          dc:"股东总数 (户)"`
	Agdhs         int         `json:"agdhs"         dc:"A股东户数 (户)"`
	Bgdhs         int         `json:"bgdhs"         dc:"B股东户数 (户)"`
	Hgdhs         int         `json:"hgdhs"         dc:"H股东户数 (户)"`
	Yltgdhs       int         `json:"yltgdhs"       dc:"已流通股东户数 (户)"`
	Wltgdhs       int         `json:"wltgdhs"       dc:"未流通股东户数 (户)"`
	AgRatio       float64     `json:"agRatio"       dc:"A股股东占比(%)"`
	YltRatio      float64     `json:"yltRatio"      dc:"已流通股东占比(%)"`
	DataSource    string      `json:"dataSource"    dc:"数据来源"`
	IsLatest      int         `json:"isLatest"      dc:"是否为该报告期最新数据"`
}

// ShareholderCountInsertFields 新增公司股东户数统计表 (按报告期统计)字段过滤
type ShareholderCountInsertFields struct {
	Symbol        string      `json:"symbol"        dc:"公司代码/股票代码 (例如: 000001.SZ, 600000.SS)"`
	Jzrq          *gtime.Time `json:"jzrq"          dc:"截止日期 (报告期结束日, 如2023-09-30)"`
	ReportYear    int         `json:"reportYear"    dc:"报告年度"`
	ReportQuarter int         `json:"reportQuarter" dc:"报告季度 (1-4)"`
	ReportType    string      `json:"reportType"    dc:"报告类型: annual-年报, half_year-中报, quarter-季报"`
	Gdzs          int         `json:"gdzs"          dc:"股东总数 (户)"`
	Agdhs         int         `json:"agdhs"         dc:"A股东户数 (户)"`
	Bgdhs         int         `json:"bgdhs"         dc:"B股东户数 (户)"`
	Hgdhs         int         `json:"hgdhs"         dc:"H股东户数 (户)"`
	Yltgdhs       int         `json:"yltgdhs"       dc:"已流通股东户数 (户)"`
	Wltgdhs       int         `json:"wltgdhs"       dc:"未流通股东户数 (户)"`
	AgRatio       float64     `json:"agRatio"       dc:"A股股东占比(%)"`
	YltRatio      float64     `json:"yltRatio"      dc:"已流通股东占比(%)"`
	DataSource    string      `json:"dataSource"    dc:"数据来源"`
	IsLatest      int         `json:"isLatest"      dc:"是否为该报告期最新数据"`
}

// ShareholderCountEditInp 修改/新增公司股东户数统计表 (按报告期统计)
type ShareholderCountEditInp struct {
	entity.ShareholderCount
}

func (in *ShareholderCountEditInp) Filter(ctx context.Context) (err error) {
	// 验证公司代码/股票代码 (例如: 000001.SZ, 600000.SS)
	if err := g.Validator().Rules("required").Data(in.Symbol).Messages("公司代码/股票代码 (例如: 000001.SZ, 600000.SS)不能为空").Run(ctx); err != nil {
		return err.Current()
	}
	// 验证截止日期 (报告期结束日, 如2023-09-30)
	if err := g.Validator().Rules("required").Data(in.Jzrq).Messages("截止日期 (报告期结束日, 如2023-09-30)不能为空").Run(ctx); err != nil {
		return err.Current()
	}
	// 验证报告年度
	if err := g.Validator().Rules("required").Data(in.ReportYear).Messages("报告年度不能为空").Run(ctx); err != nil {
		return err.Current()
	}
	// 验证报告类型: annual-年报, half_year-中报, quarter-季报
	if err := g.Validator().Rules("required").Data(in.ReportType).Messages("报告类型: annual-年报, half_year-中报, quarter-季报不能为空").Run(ctx); err != nil {
		return err.Current()
	}
	// 验证股东总数 (户)
	if err := g.Validator().Rules("required").Data(in.Gdzs).Messages("股东总数 (户)不能为空").Run(ctx); err != nil {
		return err.Current()
	}
	return
}

type ShareholderCountEditModel struct{}

// ShareholderCountDeleteInp 删除公司股东户数统计表 (按报告期统计)
type ShareholderCountDeleteInp struct {
	Id interface{} `json:"id" v:"required#自增主键不能为空" dc:"自增主键"`
}

func (in *ShareholderCountDeleteInp) Filter(ctx context.Context) (err error) {
	return
}

type ShareholderCountDeleteModel struct{}

// ShareholderCountViewInp 获取指定公司股东户数统计表 (按报告期统计)信息
type ShareholderCountViewInp struct {
	Id int64 `json:"id" v:"required#自增主键不能为空" dc:"自增主键"`
}

func (in *ShareholderCountViewInp) Filter(ctx context.Context) (err error) {
	return
}

type ShareholderCountViewModel struct {
	entity.ShareholderCount
}

// ShareholderCountListInp 获取公司股东户数统计表 (按报告期统计)列表
type ShareholderCountListInp struct {
	form.PageReq
	Id        int64         `json:"id"        dc:"自增主键"`
	CreatedAt []*gtime.Time `json:"createdAt" dc:"创建时间"`
}

func (in *ShareholderCountListInp) Filter(ctx context.Context) (err error) {
	return
}

type ShareholderCountListModel struct {
	Id            int64       `json:"id"            dc:"自增主键"`
	Symbol        string      `json:"symbol"        dc:"公司代码/股票代码 (例如: 000001.SZ, 600000.SS)"`
	Jzrq          *gtime.Time `json:"jzrq"          dc:"截止日期 (报告期结束日, 如2023-09-30)"`
	ReportYear    int         `json:"reportYear"    dc:"报告年度"`
	ReportQuarter int         `json:"reportQuarter" dc:"报告季度 (1-4)"`
	ReportType    string      `json:"reportType"    dc:"报告类型: annual-年报, half_year-中报, quarter-季报"`
	Gdzs          int         `json:"gdzs"          dc:"股东总数 (户)"`
	Agdhs         int         `json:"agdhs"         dc:"A股东户数 (户)"`
	Bgdhs         int         `json:"bgdhs"         dc:"B股东户数 (户)"`
	Hgdhs         int         `json:"hgdhs"         dc:"H股东户数 (户)"`
	Yltgdhs       int         `json:"yltgdhs"       dc:"已流通股东户数 (户)"`
	Wltgdhs       int         `json:"wltgdhs"       dc:"未流通股东户数 (户)"`
	AgRatio       float64     `json:"agRatio"       dc:"A股股东占比(%)"`
	YltRatio      float64     `json:"yltRatio"      dc:"已流通股东占比(%)"`
	DataSource    string      `json:"dataSource"    dc:"数据来源"`
	IsLatest      int         `json:"isLatest"      dc:"是否为该报告期最新数据"`
	CreatedAt     *gtime.Time `json:"createdAt"     dc:"创建时间"`
	UpdatedAt     *gtime.Time `json:"updatedAt"     dc:"更新时间"`
}

// ShareholderCountExportModel 导出公司股东户数统计表 (按报告期统计)
type ShareholderCountExportModel struct {
	Id            int64       `json:"id"            dc:"自增主键"`
	Symbol        string      `json:"symbol"        dc:"公司代码/股票代码 (例如: 000001.SZ, 600000.SS)"`
	Jzrq          *gtime.Time `json:"jzrq"          dc:"截止日期 (报告期结束日, 如2023-09-30)"`
	ReportYear    int         `json:"reportYear"    dc:"报告年度"`
	ReportQuarter int         `json:"reportQuarter" dc:"报告季度 (1-4)"`
	ReportType    string      `json:"reportType"    dc:"报告类型: annual-年报, half_year-中报, quarter-季报"`
	Gdzs          int         `json:"gdzs"          dc:"股东总数 (户)"`
	Agdhs         int         `json:"agdhs"         dc:"A股东户数 (户)"`
	Bgdhs         int         `json:"bgdhs"         dc:"B股东户数 (户)"`
	Hgdhs         int         `json:"hgdhs"         dc:"H股东户数 (户)"`
	Yltgdhs       int         `json:"yltgdhs"       dc:"已流通股东户数 (户)"`
	Wltgdhs       int         `json:"wltgdhs"       dc:"未流通股东户数 (户)"`
	AgRatio       float64     `json:"agRatio"       dc:"A股股东占比(%)"`
	YltRatio      float64     `json:"yltRatio"      dc:"已流通股东占比(%)"`
	DataSource    string      `json:"dataSource"    dc:"数据来源"`
	IsLatest      int         `json:"isLatest"      dc:"是否为该报告期最新数据"`
	CreatedAt     *gtime.Time `json:"createdAt"     dc:"创建时间"`
	UpdatedAt     *gtime.Time `json:"updatedAt"     dc:"更新时间"`
}
