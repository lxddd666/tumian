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

// StockSelfCodeUpdateFields 修改自选股票字段过滤
type StockSelfCodeUpdateFields struct {
	Mc string `json:"mc" dc:"股票名称"`
}

// StockSelfCodeInsertFields 新增自选股票字段过滤
type StockSelfCodeInsertFields struct {
	Mc string `json:"mc" dc:"股票名称"`
}

// StockSelfCodeEditInp 修改/新增自选股票
type StockSelfCodeEditInp struct {
	entity.StockSelfCode
}

func (in *StockSelfCodeEditInp) Filter(ctx context.Context) (err error) {
	// 验证股票名称
	if err := g.Validator().Rules("required").Data(in.Mc).Messages("股票名称不能为空").Run(ctx); err != nil {
		return err.Current()
	}
	return
}

type StockSelfCodeEditModel struct{}

// StockSelfCodeDeleteInp 删除自选股票
type StockSelfCodeDeleteInp struct {
	Dm interface{} `json:"dm" v:"required#股票代码不能为空" dc:"股票代码"`
}

// SelfCodeIndicatorsApiInp 技术指标获取--APi
type SelfCodeIndicatorsApiInp struct {
	Code string `json:"code"`
}

func (in *StockSelfCodeDeleteInp) Filter(ctx context.Context) (err error) {
	return
}

type StockSelfCodeDeleteModel struct{}

// StockSelfCodeViewInp 获取指定自选股票信息
type StockSelfCodeViewInp struct {
	Dm string `json:"dm" v:"required#股票代码不能为空" dc:"股票代码"`
}

func (in *StockSelfCodeViewInp) Filter(ctx context.Context) (err error) {
	return
}

type StockSelfCodeViewModel struct {
	entity.StockSelfCode
}

// StockSelfCodeListInp 获取自选股票列表
type StockSelfCodeListInp struct {
	form.PageReq
	Dm        string        `json:"dm"        dc:"股票代码"`
	Jys       string        `json:"jys"       dc:"交易所"`
	CreatedAt []*gtime.Time `json:"createdAt" dc:"创建时间"`
}

func (in *StockSelfCodeListInp) Filter(ctx context.Context) (err error) {
	return
}

type StockSelfCodeListModel struct {
	Dm        string      `json:"dm"        dc:"股票代码"`
	Mc        string      `json:"mc"        dc:"股票名称"`
	Jys       string      `json:"jys"       dc:"交易所"`
	CreatedAt *gtime.Time `json:"createdAt" dc:"创建时间"`
}

// StockSelfCodeExportModel 导出自选股票
type StockSelfCodeExportModel struct {
	Dm        string      `json:"dm"        dc:"股票代码"`
	Mc        string      `json:"mc"        dc:"股票名称"`
	Jys       string      `json:"jys"       dc:"交易所"`
	CreatedAt *gtime.Time `json:"createdAt" dc:"创建时间"`
}
