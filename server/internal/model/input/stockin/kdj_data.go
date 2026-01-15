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

// KdjDataUpdateFields 修改KDJ随机指标数据表字段过滤
type KdjDataUpdateFields struct {
	Symbol       string      `json:"symbol"       dc:"股票或标的代码 (例如: AAPL, 000001.SZ)"`
	T            *gtime.Time `json:"t"            dc:"交易时间 (统一为datetime类型，日线数据时间部分设为00:00:00)"`
	IntervalType string      `json:"intervalType" dc:"数据间隔: minute-短分时, day-日线"`
	K            float64     `json:"k"            dc:"K值"`
	D            float64     `json:"d"            dc:"D值"`
	J            float64     `json:"j"            dc:"J值"`
}

// KdjDataInsertFields 新增KDJ随机指标数据表字段过滤
type KdjDataInsertFields struct {
	Symbol       string      `json:"symbol"       dc:"股票或标的代码 (例如: AAPL, 000001.SZ)"`
	T            *gtime.Time `json:"t"            dc:"交易时间 (统一为datetime类型，日线数据时间部分设为00:00:00)"`
	IntervalType string      `json:"intervalType" dc:"数据间隔: minute-短分时, day-日线"`
	K            float64     `json:"k"            dc:"K值"`
	D            float64     `json:"d"            dc:"D值"`
	J            float64     `json:"j"            dc:"J值"`
}

// KdjDataEditInp 修改/新增KDJ随机指标数据表
type KdjDataEditInp struct {
	entity.KdjData
}

func (in *KdjDataEditInp) Filter(ctx context.Context) (err error) {
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

type KdjDataEditModel struct{}

// KdjDataDeleteInp 删除KDJ随机指标数据表
type KdjDataDeleteInp struct {
	Id interface{} `json:"id" v:"required#自增主键不能为空" dc:"自增主键"`
}

func (in *KdjDataDeleteInp) Filter(ctx context.Context) (err error) {
	return
}

type KdjDataDeleteModel struct{}

// KdjDataViewInp 获取指定KDJ随机指标数据表信息
type KdjDataViewInp struct {
	Id int64 `json:"id" v:"required#自增主键不能为空" dc:"自增主键"`
}

func (in *KdjDataViewInp) Filter(ctx context.Context) (err error) {
	return
}

type KdjDataViewModel struct {
	entity.KdjData
}

// KdjDataListInp 获取KDJ随机指标数据表列表
type KdjDataListInp struct {
	form.PageReq
	Id        int64         `json:"id"        dc:"自增主键"`
	CreatedAt []*gtime.Time `json:"createdAt" dc:"数据创建时间"`
}

func (in *KdjDataListInp) Filter(ctx context.Context) (err error) {
	return
}

type KdjDataListModel struct {
	Id           int64       `json:"id"           dc:"自增主键"`
	Symbol       string      `json:"symbol"       dc:"股票或标的代码 (例如: AAPL, 000001.SZ)"`
	T            *gtime.Time `json:"t"            dc:"交易时间 (统一为datetime类型，日线数据时间部分设为00:00:00)"`
	IntervalType string      `json:"intervalType" dc:"数据间隔: minute-短分时, day-日线"`
	K            float64     `json:"k"            dc:"K值"`
	D            float64     `json:"d"            dc:"D值"`
	J            float64     `json:"j"            dc:"J值"`
	CreatedAt    *gtime.Time `json:"createdAt"    dc:"数据创建时间"`
	UpdatedAt    *gtime.Time `json:"updatedAt"    dc:"数据更新时间"`
}

// KdjDataExportModel 导出KDJ随机指标数据表
type KdjDataExportModel struct {
	Id           int64       `json:"id"           dc:"自增主键"`
	Symbol       string      `json:"symbol"       dc:"股票或标的代码 (例如: AAPL, 000001.SZ)"`
	T            *gtime.Time `json:"t"            dc:"交易时间 (统一为datetime类型，日线数据时间部分设为00:00:00)"`
	IntervalType string      `json:"intervalType" dc:"数据间隔: minute-短分时, day-日线"`
	K            float64     `json:"k"            dc:"K值"`
	D            float64     `json:"d"            dc:"D值"`
	J            float64     `json:"j"            dc:"J值"`
	CreatedAt    *gtime.Time `json:"createdAt"    dc:"数据创建时间"`
	UpdatedAt    *gtime.Time `json:"updatedAt"    dc:"数据更新时间"`
}

type GetKdjInp struct {
	Symbol string `json:"symbol"       dc:"股票或标的代码 (例如: AAPL, 000001.SZ)"`
	T      int    `json:"t"            dc:"交易时间 (统一为datetime类型，日线数据时间部分设为00:00:00)"`
}
