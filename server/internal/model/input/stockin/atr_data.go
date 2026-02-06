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

// AtrDataUpdateFields 修改atr指标数据表字段过滤
type AtrDataUpdateFields struct {
	Symbol string      `json:"symbol" dc:"股票或标的代码 (例如: AAPL, 000001.SZ)"`
	T      *gtime.Time `json:"t"      dc:"交易时间 (统一为datetime类型，日线数据时间部分设为00:00:00)"`
	Tr     float64     `json:"tr"     dc:"真实波幅 (True Range)"`
	Atr    float64     `json:"atr"    dc:"平均真实波幅 (Average True Range)"`
}

// AtrDataInsertFields 新增atr指标数据表字段过滤
type AtrDataInsertFields struct {
	Symbol string      `json:"symbol" dc:"股票或标的代码 (例如: AAPL, 000001.SZ)"`
	T      *gtime.Time `json:"t"      dc:"交易时间 (统一为datetime类型，日线数据时间部分设为00:00:00)"`
	Tr     float64     `json:"tr"     dc:"真实波幅 (True Range)"`
	Atr    float64     `json:"atr"    dc:"平均真实波幅 (Average True Range)"`
}

// AtrDataEditInp 修改/新增atr指标数据表
type AtrDataEditInp struct {
	entity.AtrData
}

func (in *AtrDataEditInp) Filter(ctx context.Context) (err error) {
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

type AtrDataEditModel struct{}

// AtrDataDeleteInp 删除atr指标数据表
type AtrDataDeleteInp struct {
	Id interface{} `json:"id" v:"required#自增主键不能为空" dc:"自增主键"`
}

func (in *AtrDataDeleteInp) Filter(ctx context.Context) (err error) {
	return
}

type AtrDataDeleteModel struct{}

// AtrDataViewInp 获取指定atr指标数据表信息
type AtrDataViewInp struct {
	Id int64 `json:"id" v:"required#自增主键不能为空" dc:"自增主键"`
}

func (in *AtrDataViewInp) Filter(ctx context.Context) (err error) {
	return
}

type AtrDataViewModel struct {
	entity.AtrData
}

// AtrDataListInp 获取atr指标数据表列表
type AtrDataListInp struct {
	form.PageReq
	Id        int64         `json:"id"        dc:"自增主键"`
	CreatedAt []*gtime.Time `json:"createdAt" dc:"数据创建时间"`
}

func (in *AtrDataListInp) Filter(ctx context.Context) (err error) {
	return
}

type AtrDataListModel struct {
	Id        int64       `json:"id"        dc:"自增主键"`
	Symbol    string      `json:"symbol"    dc:"股票或标的代码 (例如: AAPL, 000001.SZ)"`
	T         *gtime.Time `json:"t"         dc:"交易时间 (统一为datetime类型，日线数据时间部分设为00:00:00)"`
	Tr        float64     `json:"tr"        dc:"真实波幅 (True Range)"`
	Atr       float64     `json:"atr"       dc:"平均真实波幅 (Average True Range)"`
	CreatedAt *gtime.Time `json:"createdAt" dc:"数据创建时间"`
	UpdatedAt *gtime.Time `json:"updatedAt" dc:"数据更新时间"`
}

// AtrDataExportModel 导出atr指标数据表
type AtrDataExportModel struct {
	Id        int64       `json:"id"        dc:"自增主键"`
	Symbol    string      `json:"symbol"    dc:"股票或标的代码 (例如: AAPL, 000001.SZ)"`
	T         *gtime.Time `json:"t"         dc:"交易时间 (统一为datetime类型，日线数据时间部分设为00:00:00)"`
	Tr        float64     `json:"tr"        dc:"真实波幅 (True Range)"`
	Atr       float64     `json:"atr"       dc:"平均真实波幅 (Average True Range)"`
	CreatedAt *gtime.Time `json:"createdAt" dc:"数据创建时间"`
	UpdatedAt *gtime.Time `json:"updatedAt" dc:"数据更新时间"`
}

// GetAtrDataInp 获取指定atr指标数据表信息
type GetAtrDataInp struct {
	Symbol string `json:"symbol" dc:"股票代码"`
}
