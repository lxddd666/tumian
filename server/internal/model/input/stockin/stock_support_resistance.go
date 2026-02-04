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

// StockSupportResistanceUpdateFields 修改股票支撑阻力表字段过滤
type StockSupportResistanceUpdateFields struct {
	T      *gtime.Time `json:"t"      dc:"时间"`
	Symbol string      `json:"symbol" dc:"股票或标的代码 (例如: AAPL, 000001.SZ)"`
	Price  float64     `json:"price"  dc:"当前股票价格"`
	Yl     float64     `json:"yl"     dc:"压力位"`
	Zc     float64     `json:"zc"     dc:"支撑位"`
}

// StockSupportResistanceInsertFields 新增股票支撑阻力表字段过滤
type StockSupportResistanceInsertFields struct {
	T      *gtime.Time `json:"t"      dc:"时间"`
	Symbol string      `json:"symbol" dc:"股票或标的代码 (例如: AAPL, 000001.SZ)"`
	Price  float64     `json:"price"  dc:"当前股票价格"`
	Yl     float64     `json:"yl"     dc:"压力位"`
	Zc     float64     `json:"zc"     dc:"支撑位"`
}

// StockSupportResistanceEditInp 修改/新增股票支撑阻力表
type StockSupportResistanceEditInp struct {
	entity.StockSupportResistance
}

func (in *StockSupportResistanceEditInp) Filter(ctx context.Context) (err error) {
	// 验证时间
	if err := g.Validator().Rules("required").Data(in.T).Messages("时间不能为空").Run(ctx); err != nil {
		return err.Current()
	}
	// 验证股票或标的代码 (例如: AAPL, 000001.SZ)
	if err := g.Validator().Rules("required").Data(in.Symbol).Messages("股票或标的代码 (例如: AAPL, 000001.SZ)不能为空").Run(ctx); err != nil {
		return err.Current()
	}
	// 验证当前股票价格
	if err := g.Validator().Rules("regex:(^[0-9]{1,10}$)|(^[0-9]{1,10}[\\.]{1}[0-9]{1,2}$)").Data(in.Price).Messages("当前股票价格最多允许输入10位整数及2位小数").Run(ctx); err != nil {
		return err.Current()
	}
	return
}

type StockSupportResistanceEditModel struct{}

// StockSupportResistanceDeleteInp 删除股票支撑阻力表
type StockSupportResistanceDeleteInp struct {
	Id interface{} `json:"id" v:"required#主键ID不能为空" dc:"主键ID"`
}

func (in *StockSupportResistanceDeleteInp) Filter(ctx context.Context) (err error) {
	return
}

type StockSupportResistanceDeleteModel struct{}

// StockSupportResistanceViewInp 获取指定股票支撑阻力表信息
type StockSupportResistanceViewInp struct {
	Id int64 `json:"id" v:"required#主键ID不能为空" dc:"主键ID"`
}

func (in *StockSupportResistanceViewInp) Filter(ctx context.Context) (err error) {
	return
}

type StockSupportResistanceViewModel struct {
	entity.StockSupportResistance
}

// StockSupportResistanceListInp 获取股票支撑阻力表列表
type StockSupportResistanceListInp struct {
	form.PageReq
	Id        int64         `json:"id"        dc:"主键ID"`
	CreatedAt []*gtime.Time `json:"createdAt" dc:"创建时间"`
}

func (in *StockSupportResistanceListInp) Filter(ctx context.Context) (err error) {
	return
}

type StockSupportResistanceListModel struct {
	Id        int64       `json:"id"        dc:"主键ID"`
	T         *gtime.Time `json:"t"         dc:"时间"`
	Symbol    string      `json:"symbol"    dc:"股票或标的代码 (例如: AAPL, 000001.SZ)"`
	Price     float64     `json:"price"     dc:"当前股票价格"`
	Yl        float64     `json:"yl"        dc:"压力位"`
	Zc        float64     `json:"zc"        dc:"支撑位"`
	CreatedAt *gtime.Time `json:"createdAt" dc:"创建时间"`
}

// StockSupportResistanceExportModel 导出股票支撑阻力表
type StockSupportResistanceExportModel struct {
	Id        int64       `json:"id"        dc:"主键ID"`
	T         *gtime.Time `json:"t"         dc:"时间"`
	Symbol    string      `json:"symbol"    dc:"股票或标的代码 (例如: AAPL, 000001.SZ)"`
	Price     float64     `json:"price"     dc:"当前股票价格"`
	Yl        float64     `json:"yl"        dc:"压力位"`
	Zc        float64     `json:"zc"        dc:"支撑位"`
	CreatedAt *gtime.Time `json:"createdAt" dc:"创建时间"`
}
