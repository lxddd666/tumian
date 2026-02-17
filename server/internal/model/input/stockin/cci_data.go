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

// CciDataUpdateFields 修改cci指标数据表字段过滤
type CciDataUpdateFields struct {
	Symbol string      `json:"symbol" dc:"股票或标的代码 (例如: AAPL, 000001.SZ)"`
	T      *gtime.Time `json:"t"      dc:"交易时间 (统一为datetime类型，日线数据时间部分设为00:00:00)"`
	Cci    float64     `json:"cci"    dc:"cci"`
}

// CciDataInsertFields 新增cci指标数据表字段过滤
type CciDataInsertFields struct {
	Symbol string      `json:"symbol" dc:"股票或标的代码 (例如: AAPL, 000001.SZ)"`
	T      *gtime.Time `json:"t"      dc:"交易时间 (统一为datetime类型，日线数据时间部分设为00:00:00)"`
	Cci    float64     `json:"cci"    dc:"cci"`
}

// CciDataEditInp 修改/新增cci指标数据表
type CciDataEditInp struct {
	entity.CciData
}

func (in *CciDataEditInp) Filter(ctx context.Context) (err error) {
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

type CciDataEditModel struct{}

// CciDataDeleteInp 删除cci指标数据表
type CciDataDeleteInp struct {
	Id interface{} `json:"id" v:"required#自增主键不能为空" dc:"自增主键"`
}

func (in *CciDataDeleteInp) Filter(ctx context.Context) (err error) {
	return
}

type CciDataDeleteModel struct{}

// CciDataViewInp 获取指定cci指标数据表信息
type CciDataViewInp struct {
	Id int64 `json:"id" v:"required#自增主键不能为空" dc:"自增主键"`
}

// GetCciDataInp 获取指定cci指标数据表信息
type GetCciDataInp struct {
	Symbol string `json:"symbol"  dc:"代码"`
}

func (in *CciDataViewInp) Filter(ctx context.Context) (err error) {
	return
}

type CciDataViewModel struct {
	entity.CciData
}

// CciDataListInp 获取cci指标数据表列表
type CciDataListInp struct {
	form.PageReq
	Id        int64         `json:"id"        dc:"自增主键"`
	CreatedAt []*gtime.Time `json:"createdAt" dc:"数据创建时间"`
}

func (in *CciDataListInp) Filter(ctx context.Context) (err error) {
	return
}

type CciDataListModel struct {
	Id        int64       `json:"id"        dc:"自增主键"`
	Symbol    string      `json:"symbol"    dc:"股票或标的代码 (例如: AAPL, 000001.SZ)"`
	T         *gtime.Time `json:"t"         dc:"交易时间 (统一为datetime类型，日线数据时间部分设为00:00:00)"`
	Cci       float64     `json:"cci"       dc:"cci"`
	CreatedAt *gtime.Time `json:"createdAt" dc:"数据创建时间"`
	UpdatedAt *gtime.Time `json:"updatedAt" dc:"数据更新时间"`
}

// CciDataExportModel 导出cci指标数据表
type CciDataExportModel struct {
	Id        int64       `json:"id"        dc:"自增主键"`
	Symbol    string      `json:"symbol"    dc:"股票或标的代码 (例如: AAPL, 000001.SZ)"`
	T         *gtime.Time `json:"t"         dc:"交易时间 (统一为datetime类型，日线数据时间部分设为00:00:00)"`
	Cci       float64     `json:"cci"       dc:"cci"`
	CreatedAt *gtime.Time `json:"createdAt" dc:"数据创建时间"`
	UpdatedAt *gtime.Time `json:"updatedAt" dc:"数据更新时间"`
}
