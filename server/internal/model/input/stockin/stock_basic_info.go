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

// StockBasicInfoUpdateFields 修改股票基础信息表字段过滤
type StockBasicInfoUpdateFields struct {
	Symbol         string      `json:"symbol"         dc:"股票代码"`
	Ii             string      `json:"ii"             dc:"ii"`
	Ei             string      `json:"ei"             dc:"ei"`
	Exchange       string      `json:"exchange"       dc:"交易所名称"`
	Name           string      `json:"name"           dc:"name"`
	ShortName      string      `json:"shortName"      dc:"short_name"`
	EnName         string      `json:"enName"         dc:"en_name"`
	Od             *gtime.Time `json:"od"             dc:"上市日期"`
	DataUpdateDate *gtime.Time `json:"dataUpdateDate" dc:"data_update_date"`
	Pc             float64     `json:"pc"             dc:"pc"`
	Up             float64     `json:"up"             dc:"up"`
	Dp             float64     `json:"dp"             dc:"dp"`
	Pk             float64     `json:"pk"             dc:"pk"`
	Fv             int64       `json:"fv"             dc:"fv"`
	Tv             int64       `json:"tv"             dc:"tv"`
	FloatRatio     float64     `json:"floatRatio"     dc:"流通股比例 (%)"`
	Is             int         `json:"is"             dc:"is"`
	TradingStatus  string      `json:"tradingStatus"  dc:"交易状态描述"`
	Industry       string      `json:"industry"       dc:"industry"`
	Sector         string      `json:"sector"         dc:"sector"`
	MarketType     string      `json:"marketType"     dc:"market_type"`
	DataSource     string      `json:"dataSource"     dc:"data_source"`
	IsActive       int         `json:"isActive"       dc:"is_active"`
	Version        int         `json:"version"        dc:"version"`
}

// StockBasicInfoInsertFields 新增股票基础信息表字段过滤
type StockBasicInfoInsertFields struct {
	Symbol         string      `json:"symbol"         dc:"股票代码"`
	Ii             string      `json:"ii"             dc:"ii"`
	Ei             string      `json:"ei"             dc:"ei"`
	Exchange       string      `json:"exchange"       dc:"交易所名称"`
	Name           string      `json:"name"           dc:"name"`
	ShortName      string      `json:"shortName"      dc:"short_name"`
	EnName         string      `json:"enName"         dc:"en_name"`
	Od             *gtime.Time `json:"od"             dc:"上市日期"`
	DataUpdateDate *gtime.Time `json:"dataUpdateDate" dc:"data_update_date"`
	Pc             float64     `json:"pc"             dc:"pc"`
	Up             float64     `json:"up"             dc:"up"`
	Dp             float64     `json:"dp"             dc:"dp"`
	Pk             float64     `json:"pk"             dc:"pk"`
	Fv             int64       `json:"fv"             dc:"fv"`
	Tv             int64       `json:"tv"             dc:"tv"`
	FloatRatio     float64     `json:"floatRatio"     dc:"流通股比例 (%)"`
	Is             int         `json:"is"             dc:"is"`
	TradingStatus  string      `json:"tradingStatus"  dc:"交易状态描述"`
	Industry       string      `json:"industry"       dc:"industry"`
	Sector         string      `json:"sector"         dc:"sector"`
	MarketType     string      `json:"marketType"     dc:"market_type"`
	DataSource     string      `json:"dataSource"     dc:"data_source"`
	IsActive       int         `json:"isActive"       dc:"is_active"`
	Version        int         `json:"version"        dc:"version"`
}

// StockBasicInfoEditInp 修改/新增股票基础信息表
type StockBasicInfoEditInp struct {
	entity.StockBasicInfo
}

func (in *StockBasicInfoEditInp) Filter(ctx context.Context) (err error) {
	// 验证股票代码
	if err := g.Validator().Rules("required").Data(in.Symbol).Messages("股票代码不能为空").Run(ctx); err != nil {
		return err.Current()
	}
	// 验证ii
	if err := g.Validator().Rules("required").Data(in.Ii).Messages("ii不能为空").Run(ctx); err != nil {
		return err.Current()
	}
	// 验证ei
	if err := g.Validator().Rules("required").Data(in.Ei).Messages("ei不能为空").Run(ctx); err != nil {
		return err.Current()
	}
	// 验证name
	if err := g.Validator().Rules("required").Data(in.Name).Messages("name不能为空").Run(ctx); err != nil {
		return err.Current()
	}
	// 验证上市日期
	if err := g.Validator().Rules("required").Data(in.Od).Messages("上市日期不能为空").Run(ctx); err != nil {
		return err.Current()
	}
	// 验证data_update_date
	if err := g.Validator().Rules("required").Data(in.DataUpdateDate).Messages("data_update_date不能为空").Run(ctx); err != nil {
		return err.Current()
	}
	// 验证is
	if err := g.Validator().Rules("required").Data(in.Is).Messages("is不能为空").Run(ctx); err != nil {
		return err.Current()
	}
	return
}

type StockBasicInfoEditModel struct{}

// StockBasicInfoDeleteInp 删除股票基础信息表
type StockBasicInfoDeleteInp struct {
	Id interface{} `json:"id" v:"required#自增主键不能为空" dc:"自增主键"`
}

func (in *StockBasicInfoDeleteInp) Filter(ctx context.Context) (err error) {
	return
}

type StockBasicInfoDeleteModel struct{}

// StockBasicInfoViewInp 获取指定股票基础信息表信息
type StockBasicInfoViewInp struct {
	Id int64 `json:"id" v:"required#自增主键不能为空" dc:"自增主键"`
}

func (in *StockBasicInfoViewInp) Filter(ctx context.Context) (err error) {
	return
}

type StockBasicInfoViewModel struct {
	entity.StockBasicInfo
}

// StockBasicInfoListInp 获取股票基础信息表列表
type StockBasicInfoListInp struct {
	form.PageReq
	Id        int64         `json:"id"        dc:"自增主键"`
	CreatedAt []*gtime.Time `json:"createdAt" dc:"created_at"`
}

func (in *StockBasicInfoListInp) Filter(ctx context.Context) (err error) {
	return
}

type StockBasicInfoListModel struct {
	Id             int64       `json:"id"             dc:"自增主键"`
	Symbol         string      `json:"symbol"         dc:"股票代码"`
	Ii             string      `json:"ii"             dc:"ii"`
	Ei             string      `json:"ei"             dc:"ei"`
	Exchange       string      `json:"exchange"       dc:"交易所名称"`
	Name           string      `json:"name"           dc:"name"`
	ShortName      string      `json:"shortName"      dc:"short_name"`
	EnName         string      `json:"enName"         dc:"en_name"`
	Od             *gtime.Time `json:"od"             dc:"上市日期"`
	DataUpdateDate *gtime.Time `json:"dataUpdateDate" dc:"data_update_date"`
	Pc             float64     `json:"pc"             dc:"pc"`
	Up             float64     `json:"up"             dc:"up"`
	Dp             float64     `json:"dp"             dc:"dp"`
	Pk             float64     `json:"pk"             dc:"pk"`
	Fv             int64       `json:"fv"             dc:"fv"`
	Tv             int64       `json:"tv"             dc:"tv"`
	FloatRatio     float64     `json:"floatRatio"     dc:"流通股比例 (%)"`
	Is             int         `json:"is"             dc:"is"`
	TradingStatus  string      `json:"tradingStatus"  dc:"交易状态描述"`
	Industry       string      `json:"industry"       dc:"industry"`
	Sector         string      `json:"sector"         dc:"sector"`
	MarketType     string      `json:"marketType"     dc:"market_type"`
	DataSource     string      `json:"dataSource"     dc:"data_source"`
	IsActive       int         `json:"isActive"       dc:"is_active"`
	Version        int         `json:"version"        dc:"version"`
	CreatedAt      *gtime.Time `json:"createdAt"      dc:"created_at"`
	UpdatedAt      *gtime.Time `json:"updatedAt"      dc:"updated_at"`
}

// StockBasicInfoExportModel 导出股票基础信息表
type StockBasicInfoExportModel struct {
	Id             int64       `json:"id"             dc:"自增主键"`
	Symbol         string      `json:"symbol"         dc:"股票代码"`
	Ii             string      `json:"ii"             dc:"ii"`
	Ei             string      `json:"ei"             dc:"ei"`
	Exchange       string      `json:"exchange"       dc:"交易所名称"`
	Name           string      `json:"name"           dc:"name"`
	ShortName      string      `json:"shortName"      dc:"short_name"`
	EnName         string      `json:"enName"         dc:"en_name"`
	Od             *gtime.Time `json:"od"             dc:"上市日期"`
	DataUpdateDate *gtime.Time `json:"dataUpdateDate" dc:"data_update_date"`
	Pc             float64     `json:"pc"             dc:"pc"`
	Up             float64     `json:"up"             dc:"up"`
	Dp             float64     `json:"dp"             dc:"dp"`
	Pk             float64     `json:"pk"             dc:"pk"`
	Fv             int64       `json:"fv"             dc:"fv"`
	Tv             int64       `json:"tv"             dc:"tv"`
	FloatRatio     float64     `json:"floatRatio"     dc:"流通股比例 (%)"`
	Is             int         `json:"is"             dc:"is"`
	TradingStatus  string      `json:"tradingStatus"  dc:"交易状态描述"`
	Industry       string      `json:"industry"       dc:"industry"`
	Sector         string      `json:"sector"         dc:"sector"`
	MarketType     string      `json:"marketType"     dc:"market_type"`
	DataSource     string      `json:"dataSource"     dc:"data_source"`
	IsActive       int         `json:"isActive"       dc:"is_active"`
	Version        int         `json:"version"        dc:"version"`
	CreatedAt      *gtime.Time `json:"createdAt"      dc:"created_at"`
	UpdatedAt      *gtime.Time `json:"updatedAt"      dc:"updated_at"`
}

// StockBasicInfoGetStockBasicInfoInp 获取股票基础信息表数据
type StockBasicInfoGetStockBasicInfoInp struct {
	Symbol     string `json:"symbol" v:"required#股票代码不能为空" dc:"股票代码 (例如: 000001.SZ)"`
	Interval   string `json:"interval" v:"required#分时级别不能为空" dc:"分时级别 (例如: d)"`
	AdjustType string `json:"adjustType" v:"required#除权类型不能为空" dc:"除权类型 (例如: n)"`
	Token      string `json:"token" v:"required#token证书不能为空" dc:"token证书"`
	StartTime  string `json:"st" dc:"开始时间"`
	EndTime    string `json:"et" dc:"结束时间"`
	Limit      int    `json:"lt" dc:"最新条数"`
}

func (in *StockBasicInfoGetStockBasicInfoInp) Filter(ctx context.Context) (err error) {
	return
}
