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

// KstDataUpdateFields 修改kst指标数据表字段过滤
type KstDataUpdateFields struct {
	Symbol string      `json:"symbol" dc:"股票或标的代码 (例如: AAPL, 000001.SZ)"`
	T      *gtime.Time `json:"t"      dc:"交易时间 (统一为datetime类型，日线数据时间部分设为00:00:00)"`
	Kst    float64     `json:"kst"    dc:"kst"`
	Signal float64     `json:"signal" dc:"signal"`
}

// KstDataInsertFields 新增kst指标数据表字段过滤
type KstDataInsertFields struct {
	Symbol string      `json:"symbol" dc:"股票或标的代码 (例如: AAPL, 000001.SZ)"`
	T      *gtime.Time `json:"t"      dc:"交易时间 (统一为datetime类型，日线数据时间部分设为00:00:00)"`
	Kst    float64     `json:"kst"    dc:"kst"`
	Signal float64     `json:"signal" dc:"signal"`
}

// KstDataEditInp 修改/新增kst指标数据表
type KstDataEditInp struct {
	entity.KstData
}

func (in *KstDataEditInp) Filter(ctx context.Context) (err error) {
	// 验证股票或标的代码 (例如: AAPL, 000001.SZ)
	if err := g.Validator().Rules("required").Data(in.Symbol).Messages("股票或标的代码 (例如: AAPL, 000001.SZ)不能为空").Run(ctx); err != nil {
		return err.Current()
	}
	// 验证交易时间 (统一为datetime类型，日线数据时间部分设为00:00:00)
	if err := g.Validator().Rules("required").Data(in.T).Messages("交易时间 (统一为datetime类型，日线数据时间部分设为00:00:00)不能为空").Run(ctx); err != nil {
		return err.Current()
	}
	return
}

type KstDataEditModel struct{}

// KstDataDeleteInp 删除kst指标数据表
type KstDataDeleteInp struct {
	Id interface{} `json:"id" v:"required#自增主键不能为空" dc:"自增主键"`
}

func (in *KstDataDeleteInp) Filter(ctx context.Context) (err error) {
	return
}

type KstDataDeleteModel struct{}

// KstDataViewInp 获取指定kst指标数据表信息
type KstDataViewInp struct {
	Id int64 `json:"id" v:"required#自增主键不能为空" dc:"自增主键"`
}

func (in *KstDataViewInp) Filter(ctx context.Context) (err error) {
	return
}

type KstDataViewModel struct {
	entity.KstData
}

// KstDataListInp 获取kst指标数据表列表
type KstDataListInp struct {
	form.PageReq
	Id        int64         `json:"id"        dc:"自增主键"`
	CreatedAt []*gtime.Time `json:"createdAt" dc:"数据创建时间"`
}

func (in *KstDataListInp) Filter(ctx context.Context) (err error) {
	return
}

type KstDataListModel struct {
	Id        int64       `json:"id"        dc:"自增主键"`
	Symbol    string      `json:"symbol"    dc:"股票或标的代码 (例如: AAPL, 000001.SZ)"`
	T         *gtime.Time `json:"t"         dc:"交易时间 (统一为datetime类型，日线数据时间部分设为00:00:00)"`
	Kst       float64     `json:"kst"       dc:"kst"`
	Signal    float64     `json:"signal"    dc:"signal"`
	CreatedAt *gtime.Time `json:"createdAt" dc:"数据创建时间"`
	UpdatedAt *gtime.Time `json:"updatedAt" dc:"数据更新时间"`
}

// KstDataExportModel 导出kst指标数据表
type KstDataExportModel struct {
	Id        int64       `json:"id"        dc:"自增主键"`
	Symbol    string      `json:"symbol"    dc:"股票或标的代码 (例如: AAPL, 000001.SZ)"`
	T         *gtime.Time `json:"t"         dc:"交易时间 (统一为datetime类型，日线数据时间部分设为00:00:00)"`
	Kst       float64     `json:"kst"       dc:"kst"`
	Signal    float64     `json:"signal"    dc:"signal"`
	CreatedAt *gtime.Time `json:"createdAt" dc:"数据创建时间"`
	UpdatedAt *gtime.Time `json:"updatedAt" dc:"数据更新时间"`
}
