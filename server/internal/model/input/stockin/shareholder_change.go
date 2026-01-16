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

// ShareholderChangeUpdateFields 修改股东户数变化记录表 (记录相邻报告期的户数变化)字段过滤
type ShareholderChangeUpdateFields struct {
	Symbol          string      `json:"symbol"          dc:"股票代码 (如: 000001.SZ)"`
	Jzrq            *gtime.Time `json:"jzrq"            dc:"截止日期 (统计截止日，如2025-12-31)[citation:9]"`
	Gdhs            int         `json:"gdhs"            dc:"股东户数 (统计截止日的总户数)[citation:3][citation:6]"`
	Bh              float64     `json:"bh"              dc:"比上期变化情况 (通常为百分比，如 -6.23 表示减少6.23%)[citation:9]"`
	ChangeDirection string      `json:"changeDirection" dc:"变化方向 (衍生字段)"`
	DataSource      string      `json:"dataSource"      dc:"数据来源 (如: 交易所公告[citation:6]、互动平台[citation:3])"`
	AnnDate         *gtime.Time `json:"annDate"         dc:"公告日期 (信息发布日期)[citation:1]"`
}

// ShareholderChangeInsertFields 新增股东户数变化记录表 (记录相邻报告期的户数变化)字段过滤
type ShareholderChangeInsertFields struct {
	Symbol          string      `json:"symbol"          dc:"股票代码 (如: 000001.SZ)"`
	Jzrq            *gtime.Time `json:"jzrq"            dc:"截止日期 (统计截止日，如2025-12-31)[citation:9]"`
	Gdhs            int         `json:"gdhs"            dc:"股东户数 (统计截止日的总户数)[citation:3][citation:6]"`
	Bh              float64     `json:"bh"              dc:"比上期变化情况 (通常为百分比，如 -6.23 表示减少6.23%)[citation:9]"`
	ChangeDirection string      `json:"changeDirection" dc:"变化方向 (衍生字段)"`
	DataSource      string      `json:"dataSource"      dc:"数据来源 (如: 交易所公告[citation:6]、互动平台[citation:3])"`
	AnnDate         *gtime.Time `json:"annDate"         dc:"公告日期 (信息发布日期)[citation:1]"`
}

// ShareholderChangeEditInp 修改/新增股东户数变化记录表 (记录相邻报告期的户数变化)
type ShareholderChangeEditInp struct {
	entity.ShareholderChange
}

func (in *ShareholderChangeEditInp) Filter(ctx context.Context) (err error) {
	// 验证股票代码 (如: 000001.SZ)
	if err := g.Validator().Rules("required").Data(in.Symbol).Messages("股票代码 (如: 000001.SZ)不能为空").Run(ctx); err != nil {
		return err.Current()
	}
	// 验证截止日期 (统计截止日，如2025-12-31)[citation:9]
	if err := g.Validator().Rules("required").Data(in.Jzrq).Messages("截止日期 (统计截止日，如2025-12-31)[citation:9]不能为空").Run(ctx); err != nil {
		return err.Current()
	}
	return
}

type ShareholderChangeEditModel struct{}

// ShareholderChangeDeleteInp 删除股东户数变化记录表 (记录相邻报告期的户数变化)
type ShareholderChangeDeleteInp struct {
	Id interface{} `json:"id" v:"required#自增主键不能为空" dc:"自增主键"`
}

func (in *ShareholderChangeDeleteInp) Filter(ctx context.Context) (err error) {
	return
}

type ShareholderChangeDeleteModel struct{}

// ShareholderChangeViewInp 获取指定股东户数变化记录表 (记录相邻报告期的户数变化)信息
type ShareholderChangeViewInp struct {
	Id int64 `json:"id" v:"required#自增主键不能为空" dc:"自增主键"`
}

func (in *ShareholderChangeViewInp) Filter(ctx context.Context) (err error) {
	return
}

type ShareholderChangeViewModel struct {
	entity.ShareholderChange
}

// ShareholderChangeListInp 获取股东户数变化记录表 (记录相邻报告期的户数变化)列表
type ShareholderChangeListInp struct {
	form.PageReq
	Id        int64         `json:"id"        dc:"自增主键"`
	CreatedAt []*gtime.Time `json:"createdAt" dc:"创建时间"`
}

func (in *ShareholderChangeListInp) Filter(ctx context.Context) (err error) {
	return
}

type ShareholderChangeListModel struct {
	Id              int64       `json:"id"              dc:"自增主键"`
	Symbol          string      `json:"symbol"          dc:"股票代码 (如: 000001.SZ)"`
	Jzrq            *gtime.Time `json:"jzrq"            dc:"截止日期 (统计截止日，如2025-12-31)[citation:9]"`
	Gdhs            int         `json:"gdhs"            dc:"股东户数 (统计截止日的总户数)[citation:3][citation:6]"`
	Bh              float64     `json:"bh"              dc:"比上期变化情况 (通常为百分比，如 -6.23 表示减少6.23%)[citation:9]"`
	ChangeDirection string      `json:"changeDirection" dc:"变化方向 (衍生字段)"`
	DataSource      string      `json:"dataSource"      dc:"数据来源 (如: 交易所公告[citation:6]、互动平台[citation:3])"`
	AnnDate         *gtime.Time `json:"annDate"         dc:"公告日期 (信息发布日期)[citation:1]"`
	CreatedAt       *gtime.Time `json:"createdAt"       dc:"创建时间"`
	UpdatedAt       *gtime.Time `json:"updatedAt"       dc:"更新时间"`
}

// ShareholderChangeExportModel 导出股东户数变化记录表 (记录相邻报告期的户数变化)
type ShareholderChangeExportModel struct {
	Id              int64       `json:"id"              dc:"自增主键"`
	Symbol          string      `json:"symbol"          dc:"股票代码 (如: 000001.SZ)"`
	Jzrq            *gtime.Time `json:"jzrq"            dc:"截止日期 (统计截止日，如2025-12-31)[citation:9]"`
	Gdhs            int         `json:"gdhs"            dc:"股东户数 (统计截止日的总户数)[citation:3][citation:6]"`
	Bh              float64     `json:"bh"              dc:"比上期变化情况 (通常为百分比，如 -6.23 表示减少6.23%)[citation:9]"`
	ChangeDirection string      `json:"changeDirection" dc:"变化方向 (衍生字段)"`
	DataSource      string      `json:"dataSource"      dc:"数据来源 (如: 交易所公告[citation:6]、互动平台[citation:3])"`
	AnnDate         *gtime.Time `json:"annDate"         dc:"公告日期 (信息发布日期)[citation:1]"`
	CreatedAt       *gtime.Time `json:"createdAt"       dc:"创建时间"`
	UpdatedAt       *gtime.Time `json:"updatedAt"       dc:"更新时间"`
}

// ShareholderChangeGetShareholderChangeInp 获取股东户数变化记录表数据
type ShareholderChangeGetShareholderChangeInp struct {
	Symbol     string `json:"symbol" v:"required#股票代码不能为空" dc:"股票代码 (例如: 000001.SZ)"`
	Interval   string `json:"interval"  dc:"分时级别 (例如: d)"`
	AdjustType string `json:"adjustType"  dc:"除权类型 (例如: n)"`
	Token      string `json:"token"  dc:"token证书"`
	StartTime  string `json:"st" dc:"开始时间"`
	EndTime    string `json:"et" dc:"结束时间"`
	Limit      int    `json:"lt" dc:"最新条数"`
}

func (in *ShareholderChangeGetShareholderChangeInp) Filter(ctx context.Context) (err error) {
	return
}
