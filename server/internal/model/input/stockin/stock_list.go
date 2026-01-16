// Package stockin
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2026 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
// @AutoGenerate Version 2.18.6
package stockin

import (
	"context"
	"hotgo/internal/consts"
	"hotgo/internal/model/entity"
	"hotgo/internal/model/input/form"
	"hotgo/utility/validate"

	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// StockListUpdateFields 修改股票列表核心表字段过滤
type StockListUpdateFields struct {
	Dm           string      `json:"dm"           dc:"股票代码 (唯一业务标识，如: 000001)"`
	Mc           string      `json:"mc"           dc:"股票名称 (如: 平安银行)"`
	Jys          string      `json:"jys"          dc:"交易所代码 (如: sh, sz, bj)"`
	ExchangeName string      `json:"exchangeName" dc:"交易所全称"`
	Symbol       string      `json:"symbol"       dc:"标准股票代码 (如: 000001.SZ)"`
	Status       int         `json:"status"       dc:"状态: 1-正常, 0-退市"`
	ListDate     *gtime.Time `json:"listDate"     dc:"上市日期"`
	DataSource   string      `json:"dataSource"   dc:"数据来源"`
}

// StockListInsertFields 新增股票列表核心表字段过滤
type StockListInsertFields struct {
	Dm           string      `json:"dm"           dc:"股票代码 (唯一业务标识，如: 000001)"`
	Mc           string      `json:"mc"           dc:"股票名称 (如: 平安银行)"`
	Jys          string      `json:"jys"          dc:"交易所代码 (如: sh, sz, bj)"`
	ExchangeName string      `json:"exchangeName" dc:"交易所全称"`
	Symbol       string      `json:"symbol"       dc:"标准股票代码 (如: 000001.SZ)"`
	Status       int         `json:"status"       dc:"状态: 1-正常, 0-退市"`
	ListDate     *gtime.Time `json:"listDate"     dc:"上市日期"`
	DataSource   string      `json:"dataSource"   dc:"数据来源"`
}

// StockListEditInp 修改/新增股票列表核心表
type StockListEditInp struct {
	entity.StockList
}

func (in *StockListEditInp) Filter(ctx context.Context) (err error) {
	// 验证股票代码 (唯一业务标识，如: 000001)
	if err := g.Validator().Rules("required").Data(in.Dm).Messages("股票代码 (唯一业务标识，如: 000001)不能为空").Run(ctx); err != nil {
		return err.Current()
	}
	// 验证股票名称 (如: 平安银行)
	if err := g.Validator().Rules("required").Data(in.Mc).Messages("股票名称 (如: 平安银行)不能为空").Run(ctx); err != nil {
		return err.Current()
	}
	// 验证交易所代码 (如: sh, sz, bj)
	if err := g.Validator().Rules("required").Data(in.Jys).Messages("交易所代码 (如: sh, sz, bj)不能为空").Run(ctx); err != nil {
		return err.Current()
	}
	// 验证状态: 1-正常, 0-退市
	if err := g.Validator().Rules("required").Data(in.Status).Messages("状态: 1-正常, 0-退市不能为空").Run(ctx); err != nil {
		return err.Current()
	}
	if err := g.Validator().Rules("in:1,2").Data(in.Status).Messages("状态: 1-正常, 0-退市值不正确").Run(ctx); err != nil {
		return err.Current()
	}
	return
}

type StockListEditModel struct{}

// StockListDeleteInp 删除股票列表核心表
type StockListDeleteInp struct {
	Id interface{} `json:"id" v:"required#自增主键不能为空" dc:"自增主键"`
}

func (in *StockListDeleteInp) Filter(ctx context.Context) (err error) {
	return
}

type StockListDeleteModel struct{}

// StockListViewInp 获取指定股票列表核心表信息
type StockListViewInp struct {
	Id int64 `json:"id" v:"required#自增主键不能为空" dc:"自增主键"`
}

func (in *StockListViewInp) Filter(ctx context.Context) (err error) {
	return
}

type StockListViewModel struct {
	entity.StockList
}

// StockListListInp 获取股票列表核心表列表
type StockListListInp struct {
	form.PageReq
	Id        int64         `json:"id"        dc:"自增主键"`
	Status    int           `json:"status"    dc:"状态: 1-正常, 0-退市"`
	CreatedAt []*gtime.Time `json:"createdAt" dc:"创建时间"`
}

func (in *StockListListInp) Filter(ctx context.Context) (err error) {
	return
}

type StockListListModel struct {
	Id           int64       `json:"id"           dc:"自增主键"`
	Dm           string      `json:"dm"           dc:"股票代码 (唯一业务标识，如: 000001)"`
	Mc           string      `json:"mc"           dc:"股票名称 (如: 平安银行)"`
	Jys          string      `json:"jys"          dc:"交易所代码 (如: sh, sz, bj)"`
	ExchangeName string      `json:"exchangeName" dc:"交易所全称"`
	Symbol       string      `json:"symbol"       dc:"标准股票代码 (如: 000001.SZ)"`
	Status       int         `json:"status"       dc:"状态: 1-正常, 0-退市"`
	ListDate     *gtime.Time `json:"listDate"     dc:"上市日期"`
	DataSource   string      `json:"dataSource"   dc:"数据来源"`
	CreatedAt    *gtime.Time `json:"createdAt"    dc:"创建时间"`
	UpdatedAt    *gtime.Time `json:"updatedAt"    dc:"更新时间"`
}

// StockListExportModel 导出股票列表核心表
type StockListExportModel struct {
	Id           int64       `json:"id"           dc:"自增主键"`
	Dm           string      `json:"dm"           dc:"股票代码 (唯一业务标识，如: 000001)"`
	Mc           string      `json:"mc"           dc:"股票名称 (如: 平安银行)"`
	Jys          string      `json:"jys"          dc:"交易所代码 (如: sh, sz, bj)"`
	ExchangeName string      `json:"exchangeName" dc:"交易所全称"`
	Symbol       string      `json:"symbol"       dc:"标准股票代码 (如: 000001.SZ)"`
	Status       int         `json:"status"       dc:"状态: 1-正常, 0-退市"`
	ListDate     *gtime.Time `json:"listDate"     dc:"上市日期"`
	DataSource   string      `json:"dataSource"   dc:"数据来源"`
	CreatedAt    *gtime.Time `json:"createdAt"    dc:"创建时间"`
	UpdatedAt    *gtime.Time `json:"updatedAt"    dc:"更新时间"`
}

// StockListStatusInp 更新股票列表核心表状态
type StockListStatusInp struct {
	Id     int64 `json:"id" v:"required#自增主键不能为空" dc:"自增主键"`
	Status int   `json:"status" dc:"状态"`
}

func (in *StockListStatusInp) Filter(ctx context.Context) (err error) {
	if in.Id <= 0 {
		err = gerror.New("自增主键不能为空")
		return
	}

	if in.Status <= 0 {
		err = gerror.New("状态不能为空")
		return
	}

	if !validate.InSlice(consts.StatusSlice, in.Status) {
		err = gerror.New("状态不正确")
		return
	}
	return
}

type StockListStatusModel struct{}

// StockListGetStockListInp 获取股票列表核心表数据
type StockListGetStockListInp struct {
	Symbol     string `json:"symbol" v:"required#股票代码不能为空" dc:"股票代码 (例如: 000001.SZ)"`
	Interval   string `json:"interval"  dc:"分时级别 (例如: d)"`
	AdjustType string `json:"adjustType"  dc:"除权类型 (例如: n)"`
	Token      string `json:"token"  dc:"token证书"`
	StartTime  string `json:"st" dc:"开始时间"`
	EndTime    string `json:"et" dc:"结束时间"`
	Limit      int    `json:"lt" dc:"最新条数"`
}

func (in *StockListGetStockListInp) Filter(ctx context.Context) (err error) {
	return
}
