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

// FastkDataUpdateFields 修改fastk指标数据表字段过滤
type FastkDataUpdateFields struct {
	Symbol string      `json:"symbol" dc:"股票或标的代码 (例如: AAPL, 000001.SZ)"`
	T      *gtime.Time `json:"t"      dc:"交易时间 (统一为datetime类型，日线数据时间部分设为00:00:00)"`
	K      float64     `json:"k"      dc:"kst"`
	D      float64     `json:"d"      dc:"signal"`
}

// FastkDataInsertFields 新增fastk指标数据表字段过滤
type FastkDataInsertFields struct {
	Symbol string      `json:"symbol" dc:"股票或标的代码 (例如: AAPL, 000001.SZ)"`
	T      *gtime.Time `json:"t"      dc:"交易时间 (统一为datetime类型，日线数据时间部分设为00:00:00)"`
	K      float64     `json:"k"      dc:"kst"`
	D      float64     `json:"d"      dc:"signal"`
}

// FastkDataEditInp 修改/新增fastk指标数据表
type FastkDataEditInp struct {
	entity.FastkData
}

func (in *FastkDataEditInp) Filter(ctx context.Context) (err error) {
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

type FastkDataEditModel struct{}

// FastkDataDeleteInp 删除fastk指标数据表
type FastkDataDeleteInp struct {
	Id interface{} `json:"id" v:"required#自增主键不能为空" dc:"自增主键"`
}

func (in *FastkDataDeleteInp) Filter(ctx context.Context) (err error) {
	return
}

type FastkDataDeleteModel struct{}

// FastkDataViewInp 获取指定fastk指标数据表信息
type FastkDataViewInp struct {
	Id int64 `json:"id" v:"required#自增主键不能为空" dc:"自增主键"`
}

func (in *FastkDataViewInp) Filter(ctx context.Context) (err error) {
	return
}

type FastkDataViewModel struct {
	entity.FastkData
}

// FastkDataListInp 获取fastk指标数据表列表
type FastkDataListInp struct {
	form.PageReq
	Id        int64         `json:"id"        dc:"自增主键"`
	CreatedAt []*gtime.Time `json:"createdAt" dc:"数据创建时间"`
}

func (in *FastkDataListInp) Filter(ctx context.Context) (err error) {
	return
}

type FastkDataListModel struct {
	Id        int64       `json:"id"        dc:"自增主键"`
	Symbol    string      `json:"symbol"    dc:"股票或标的代码 (例如: AAPL, 000001.SZ)"`
	T         *gtime.Time `json:"t"         dc:"交易时间 (统一为datetime类型，日线数据时间部分设为00:00:00)"`
	K         float64     `json:"k"         dc:"kst"`
	D         float64     `json:"d"         dc:"signal"`
	CreatedAt *gtime.Time `json:"createdAt" dc:"数据创建时间"`
	UpdatedAt *gtime.Time `json:"updatedAt" dc:"数据更新时间"`
}

// FastkDataExportModel 导出fastk指标数据表
type FastkDataExportModel struct {
	Id        int64       `json:"id"        dc:"自增主键"`
	Symbol    string      `json:"symbol"    dc:"股票或标的代码 (例如: AAPL, 000001.SZ)"`
	T         *gtime.Time `json:"t"         dc:"交易时间 (统一为datetime类型，日线数据时间部分设为00:00:00)"`
	K         float64     `json:"k"         dc:"kst"`
	D         float64     `json:"d"         dc:"signal"`
	CreatedAt *gtime.Time `json:"createdAt" dc:"数据创建时间"`
	UpdatedAt *gtime.Time `json:"updatedAt" dc:"数据更新时间"`
}
