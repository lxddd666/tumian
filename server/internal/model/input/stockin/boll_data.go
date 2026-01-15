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

// BollDataUpdateFields 修改布林带(BOLL)指标数据表字段过滤
type BollDataUpdateFields struct {
	Symbol       string      `json:"symbol"       dc:"股票或标的代码 (例如: AAPL, 000001.SZ)"`
	T            *gtime.Time `json:"t"            dc:"交易时间 (统一为datetime类型，日线数据时间部分设为00:00:00)"`
	IntervalType string      `json:"intervalType" dc:"数据间隔: minute-短分时, day-日线"`
	U            float64     `json:"u"            dc:"上轨(Upper Band)"`
	M            float64     `json:"m"            dc:"中轨(Middle Band)"`
	D            float64     `json:"d"            dc:"下轨(Lower Band)"`
}

// BollDataInsertFields 新增布林带(BOLL)指标数据表字段过滤
type BollDataInsertFields struct {
	Symbol       string      `json:"symbol"       dc:"股票或标的代码 (例如: AAPL, 000001.SZ)"`
	T            *gtime.Time `json:"t"            dc:"交易时间 (统一为datetime类型，日线数据时间部分设为00:00:00)"`
	IntervalType string      `json:"intervalType" dc:"数据间隔: minute-短分时, day-日线"`
	U            float64     `json:"u"            dc:"上轨(Upper Band)"`
	M            float64     `json:"m"            dc:"中轨(Middle Band)"`
	D            float64     `json:"d"            dc:"下轨(Lower Band)"`
}

// BollDataEditInp 修改/新增布林带(BOLL)指标数据表
type BollDataEditInp struct {
	entity.BollData
}

func (in *BollDataEditInp) Filter(ctx context.Context) (err error) {
	// 验证股票或标的代码 (例如: AAPL, 000001.SZ)
	if err := g.Validator().Rules("required").Data(in.Symbol).Messages("股票或标的代码 (例如: AAPL, 000001.SZ)不能为空").Run(ctx); err != nil {
		return err.Current()
	}
	// 验证交易时间 (统一为datetime类型，日线数据时间部分设为00:00:00)
	if err := g.Validator().Rules("required").Data(in.T).Messages("交易时间 (统一为datetime类型，日线数据时间部分设为00:00:00)不能为空").Run(ctx); err != nil {
		return err.Current()
	}
	// 验证数据间隔: minute-短分时, day-日线
	if err := g.Validator().Rules("required").Data(in.IntervalType).Messages("数据间隔: minute-短分时, day-日线不能为空").Run(ctx); err != nil {
		return err.Current()
	}
	return
}

type BollDataEditModel struct{}

// BollDataDeleteInp 删除布林带(BOLL)指标数据表
type BollDataDeleteInp struct {
	Id interface{} `json:"id" v:"required#自增主键不能为空" dc:"自增主键"`
}

func (in *BollDataDeleteInp) Filter(ctx context.Context) (err error) {
	return
}

type BollDataDeleteModel struct{}

// BollDataViewInp 获取指定布林带(BOLL)指标数据表信息
type BollDataViewInp struct {
	Id int64 `json:"id" v:"required#自增主键不能为空" dc:"自增主键"`
}

func (in *BollDataViewInp) Filter(ctx context.Context) (err error) {
	return
}

type BollDataViewModel struct {
	entity.BollData
}

// BollDataListInp 获取布林带(BOLL)指标数据表列表
type BollDataListInp struct {
	form.PageReq
	Id        int64         `json:"id"        dc:"自增主键"`
	CreatedAt []*gtime.Time `json:"createdAt" dc:"数据创建时间"`
}

func (in *BollDataListInp) Filter(ctx context.Context) (err error) {
	return
}

type BollDataListModel struct {
	Id           int64       `json:"id"           dc:"自增主键"`
	Symbol       string      `json:"symbol"       dc:"股票或标的代码 (例如: AAPL, 000001.SZ)"`
	T            *gtime.Time `json:"t"            dc:"交易时间 (统一为datetime类型，日线数据时间部分设为00:00:00)"`
	IntervalType string      `json:"intervalType" dc:"数据间隔: minute-短分时, day-日线"`
	U            float64     `json:"u"            dc:"上轨(Upper Band)"`
	M            float64     `json:"m"            dc:"中轨(Middle Band)"`
	D            float64     `json:"d"            dc:"下轨(Lower Band)"`
	CreatedAt    *gtime.Time `json:"createdAt"    dc:"数据创建时间"`
	UpdatedAt    *gtime.Time `json:"updatedAt"    dc:"数据更新时间"`
}

// BollDataExportModel 导出布林带(BOLL)指标数据表
type BollDataExportModel struct {
	Id           int64       `json:"id"           dc:"自增主键"`
	Symbol       string      `json:"symbol"       dc:"股票或标的代码 (例如: AAPL, 000001.SZ)"`
	T            *gtime.Time `json:"t"            dc:"交易时间 (统一为datetime类型，日线数据时间部分设为00:00:00)"`
	IntervalType string      `json:"intervalType" dc:"数据间隔: minute-短分时, day-日线"`
	U            float64     `json:"u"            dc:"上轨(Upper Band)"`
	M            float64     `json:"m"            dc:"中轨(Middle Band)"`
	D            float64     `json:"d"            dc:"下轨(Lower Band)"`
	CreatedAt    *gtime.Time `json:"createdAt"    dc:"数据创建时间"`
	UpdatedAt    *gtime.Time `json:"updatedAt"    dc:"数据更新时间"`
}
