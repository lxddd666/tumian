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

// MaDataUpdateFields 修改移动平均线(MA)指标数据表字段过滤
type MaDataUpdateFields struct {
	Symbol       string      `json:"symbol"       dc:"股票或标的代码 (例如: AAPL, 000001.SZ)"`
	T            *gtime.Time `json:"t"            dc:"交易时间 (统一为datetime类型，日线数据时间部分设为00:00:00)"`
	IntervalType string      `json:"intervalType" dc:"数据间隔: minute-短分时, day-日线"`
	Ma3          float64     `json:"ma3"          dc:"MA3值"`
	Ma5          float64     `json:"ma5"          dc:"MA5值"`
	Ma10         float64     `json:"ma10"         dc:"MA10值"`
	Ma15         float64     `json:"ma15"         dc:"MA15值"`
	Ma20         float64     `json:"ma20"         dc:"MA20值"`
	Ma30         float64     `json:"ma30"         dc:"MA30值"`
	Ma60         float64     `json:"ma60"         dc:"MA60值"`
	Ma120        float64     `json:"ma120"        dc:"MA120值"`
	Ma200        float64     `json:"ma200"        dc:"MA200值"`
	Ma250        float64     `json:"ma250"        dc:"MA250值"`
}

// MaDataInsertFields 新增移动平均线(MA)指标数据表字段过滤
type MaDataInsertFields struct {
	Symbol       string      `json:"symbol"       dc:"股票或标的代码 (例如: AAPL, 000001.SZ)"`
	T            *gtime.Time `json:"t"            dc:"交易时间 (统一为datetime类型，日线数据时间部分设为00:00:00)"`
	IntervalType string      `json:"intervalType" dc:"数据间隔: minute-短分时, day-日线"`
	Ma3          float64     `json:"ma3"          dc:"MA3值"`
	Ma5          float64     `json:"ma5"          dc:"MA5值"`
	Ma10         float64     `json:"ma10"         dc:"MA10值"`
	Ma15         float64     `json:"ma15"         dc:"MA15值"`
	Ma20         float64     `json:"ma20"         dc:"MA20值"`
	Ma30         float64     `json:"ma30"         dc:"MA30值"`
	Ma60         float64     `json:"ma60"         dc:"MA60值"`
	Ma120        float64     `json:"ma120"        dc:"MA120值"`
	Ma200        float64     `json:"ma200"        dc:"MA200值"`
	Ma250        float64     `json:"ma250"        dc:"MA250值"`
}

// MaDataEditInp 修改/新增移动平均线(MA)指标数据表
type MaDataEditInp struct {
	entity.MaData
}

func (in *MaDataEditInp) Filter(ctx context.Context) (err error) {
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

type MaDataEditModel struct{}

// MaDataDeleteInp 删除移动平均线(MA)指标数据表
type MaDataDeleteInp struct {
	Id interface{} `json:"id" v:"required#自增主键不能为空" dc:"自增主键"`
}

func (in *MaDataDeleteInp) Filter(ctx context.Context) (err error) {
	return
}

type MaDataDeleteModel struct{}

// MaDataViewInp 获取指定移动平均线(MA)指标数据表信息
type MaDataViewInp struct {
	Id int64 `json:"id" v:"required#自增主键不能为空" dc:"自增主键"`
}

func (in *MaDataViewInp) Filter(ctx context.Context) (err error) {
	return
}

type MaDataViewModel struct {
	entity.MaData
}

// MaDataListInp 获取移动平均线(MA)指标数据表列表
type MaDataListInp struct {
	form.PageReq
	Id        int64         `json:"id"        dc:"自增主键"`
	CreatedAt []*gtime.Time `json:"createdAt" dc:"数据创建时间"`
}

func (in *MaDataListInp) Filter(ctx context.Context) (err error) {
	return
}

type MaDataListModel struct {
	Id           int64       `json:"id"           dc:"自增主键"`
	Symbol       string      `json:"symbol"       dc:"股票或标的代码 (例如: AAPL, 000001.SZ)"`
	T            *gtime.Time `json:"t"            dc:"交易时间 (统一为datetime类型，日线数据时间部分设为00:00:00)"`
	IntervalType string      `json:"intervalType" dc:"数据间隔: minute-短分时, day-日线"`
	Ma3          float64     `json:"ma3"          dc:"MA3值"`
	Ma5          float64     `json:"ma5"          dc:"MA5值"`
	Ma10         float64     `json:"ma10"         dc:"MA10值"`
	Ma15         float64     `json:"ma15"         dc:"MA15值"`
	Ma20         float64     `json:"ma20"         dc:"MA20值"`
	Ma30         float64     `json:"ma30"         dc:"MA30值"`
	Ma60         float64     `json:"ma60"         dc:"MA60值"`
	Ma120        float64     `json:"ma120"        dc:"MA120值"`
	Ma200        float64     `json:"ma200"        dc:"MA200值"`
	Ma250        float64     `json:"ma250"        dc:"MA250值"`
	CreatedAt    *gtime.Time `json:"createdAt"    dc:"数据创建时间"`
	UpdatedAt    *gtime.Time `json:"updatedAt"    dc:"数据更新时间"`
}

// MaDataExportModel 导出移动平均线(MA)指标数据表
type MaDataExportModel struct {
	Id           int64       `json:"id"           dc:"自增主键"`
	Symbol       string      `json:"symbol"       dc:"股票或标的代码 (例如: AAPL, 000001.SZ)"`
	T            *gtime.Time `json:"t"            dc:"交易时间 (统一为datetime类型，日线数据时间部分设为00:00:00)"`
	IntervalType string      `json:"intervalType" dc:"数据间隔: minute-短分时, day-日线"`
	Ma3          float64     `json:"ma3"          dc:"MA3值"`
	Ma5          float64     `json:"ma5"          dc:"MA5值"`
	Ma10         float64     `json:"ma10"         dc:"MA10值"`
	Ma15         float64     `json:"ma15"         dc:"MA15值"`
	Ma20         float64     `json:"ma20"         dc:"MA20值"`
	Ma30         float64     `json:"ma30"         dc:"MA30值"`
	Ma60         float64     `json:"ma60"         dc:"MA60值"`
	Ma120        float64     `json:"ma120"        dc:"MA120值"`
	Ma200        float64     `json:"ma200"        dc:"MA200值"`
	Ma250        float64     `json:"ma250"        dc:"MA250值"`
	CreatedAt    *gtime.Time `json:"createdAt"    dc:"数据创建时间"`
	UpdatedAt    *gtime.Time `json:"updatedAt"    dc:"数据更新时间"`
}
