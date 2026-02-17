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

// WilliamsDataUpdateFields 修改wmsr指标数据表字段过滤
type WilliamsDataUpdateFields struct {
	Symbol string      `json:"symbol" dc:"股票或标的代码 (例如: AAPL, 000001.SZ)"`
	T      *gtime.Time `json:"t"      dc:"交易时间 (统一为datetime类型，日线数据时间部分设为00:00:00)"`
	R      float64     `json:"r"      dc:"r"`
}

// WilliamsDataInsertFields 新增wmsr指标数据表字段过滤
type WilliamsDataInsertFields struct {
	Symbol string      `json:"symbol" dc:"股票或标的代码 (例如: AAPL, 000001.SZ)"`
	T      *gtime.Time `json:"t"      dc:"交易时间 (统一为datetime类型，日线数据时间部分设为00:00:00)"`
	R      float64     `json:"r"      dc:"r"`
}

// WilliamsDataEditInp 修改/新增wmsr指标数据表
type WilliamsDataEditInp struct {
	entity.WilliamsData
}

func (in *WilliamsDataEditInp) Filter(ctx context.Context) (err error) {
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

type WilliamsDataEditModel struct{}

// WilliamsDataDeleteInp 删除wmsr指标数据表
type WilliamsDataDeleteInp struct {
	Id interface{} `json:"id" v:"required#自增主键不能为空" dc:"自增主键"`
}

func (in *WilliamsDataDeleteInp) Filter(ctx context.Context) (err error) {
	return
}

type WilliamsDataDeleteModel struct{}

// WilliamsDataViewInp 获取指定wmsr指标数据表信息
type WilliamsDataViewInp struct {
	Id int64 `json:"id" v:"required#自增主键不能为空" dc:"自增主键"`
}

type GetWilliamsDataInp struct {
	Symbol string `json:"symbol"  dc:"代码"`
}

func (in *WilliamsDataViewInp) Filter(ctx context.Context) (err error) {
	return
}

type WilliamsDataViewModel struct {
	entity.WilliamsData
}

// WilliamsDataListInp 获取wmsr指标数据表列表
type WilliamsDataListInp struct {
	form.PageReq
	Id        int64         `json:"id"        dc:"自增主键"`
	CreatedAt []*gtime.Time `json:"createdAt" dc:"数据创建时间"`
}

func (in *WilliamsDataListInp) Filter(ctx context.Context) (err error) {
	return
}

type WilliamsDataListModel struct {
	Id        int64       `json:"id"        dc:"自增主键"`
	Symbol    string      `json:"symbol"    dc:"股票或标的代码 (例如: AAPL, 000001.SZ)"`
	T         *gtime.Time `json:"t"         dc:"交易时间 (统一为datetime类型，日线数据时间部分设为00:00:00)"`
	R         float64     `json:"r"         dc:"r"`
	CreatedAt *gtime.Time `json:"createdAt" dc:"数据创建时间"`
	UpdatedAt *gtime.Time `json:"updatedAt" dc:"数据更新时间"`
}

// WilliamsDataExportModel 导出wmsr指标数据表
type WilliamsDataExportModel struct {
	Id        int64       `json:"id"        dc:"自增主键"`
	Symbol    string      `json:"symbol"    dc:"股票或标的代码 (例如: AAPL, 000001.SZ)"`
	T         *gtime.Time `json:"t"         dc:"交易时间 (统一为datetime类型，日线数据时间部分设为00:00:00)"`
	R         float64     `json:"r"         dc:"r"`
	CreatedAt *gtime.Time `json:"createdAt" dc:"数据创建时间"`
	UpdatedAt *gtime.Time `json:"updatedAt" dc:"数据更新时间"`
}
