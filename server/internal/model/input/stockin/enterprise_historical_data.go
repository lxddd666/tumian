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

// EnterpriseHistoricalDataUpdateFields 修改企业级历史行情数据表 (K线数据)字段过滤
type EnterpriseHistoricalDataUpdateFields struct {
	Symbol        string      `json:"symbol"        dc:"证券代码 (如: 000001.SZ, AAPL)"`
	T             *gtime.Time `json:"t"             dc:"交易时间 (精确到分钟或日)"`
	Date          *gtime.Time `json:"date"          dc:"交易日期 (衍生字段)"`
	Year          int         `json:"year"          dc:"交易年份"`
	Month         int         `json:"month"         dc:"交易月份"`
	Weekday       int         `json:"weekday"       dc:"星期几 (1=周日,7=周六)"`
	O             float64     `json:"o"             dc:"开盘价"`
	H             float64     `json:"h"             dc:"最高价"`
	L             float64     `json:"l"             dc:"最低价"`
	C             float64     `json:"c"             dc:"收盘价"`
	Pc            float64     `json:"pc"            dc:"前收盘价"`
	V             int64       `json:"v"             dc:"成交量 (股/手)"`
	A             float64     `json:"a"             dc:"成交额 (元)"`
	Change        float64     `json:"change"        dc:"涨跌额"`
	ChangePct     float64     `json:"changePct"     dc:"涨跌幅 (%)"`
	Amplitude     float64     `json:"amplitude"     dc:"振幅 (%)"`
	Sf            int         `json:"sf"            dc:"停牌标志: 0-正常, 1-停牌"`
	TradingStatus string      `json:"tradingStatus" dc:"交易状态描述"`
	Period        string      `json:"period"        dc:"数据周期: 1min, 5min, 15min, 30min, 60min, day, week, month"`
	IsAdjusted    int         `json:"isAdjusted"    dc:"是否复权: 0-不复权, 1-前复权, 2-后复权"`
	DataQuality   int         `json:"dataQuality"   dc:"数据质量: 0-异常, 1-正常, 2-补全"`
	IsVerified    int         `json:"isVerified"    dc:"是否已验证: 0-未验证, 1-已验证"`
	DataSource    string      `json:"dataSource"    dc:"数据来源"`
	Version       int         `json:"version"       dc:"数据版本"`
}

// EnterpriseHistoricalDataInsertFields 新增企业级历史行情数据表 (K线数据)字段过滤
type EnterpriseHistoricalDataInsertFields struct {
	Symbol        string      `json:"symbol"        dc:"证券代码 (如: 000001.SZ, AAPL)"`
	T             *gtime.Time `json:"t"             dc:"交易时间 (精确到分钟或日)"`
	Date          *gtime.Time `json:"date"          dc:"交易日期 (衍生字段)"`
	Year          int         `json:"year"          dc:"交易年份"`
	Month         int         `json:"month"         dc:"交易月份"`
	Weekday       int         `json:"weekday"       dc:"星期几 (1=周日,7=周六)"`
	O             float64     `json:"o"             dc:"开盘价"`
	H             float64     `json:"h"             dc:"最高价"`
	L             float64     `json:"l"             dc:"最低价"`
	C             float64     `json:"c"             dc:"收盘价"`
	Pc            float64     `json:"pc"            dc:"前收盘价"`
	V             int64       `json:"v"             dc:"成交量 (股/手)"`
	A             float64     `json:"a"             dc:"成交额 (元)"`
	Change        float64     `json:"change"        dc:"涨跌额"`
	ChangePct     float64     `json:"changePct"     dc:"涨跌幅 (%)"`
	Amplitude     float64     `json:"amplitude"     dc:"振幅 (%)"`
	Sf            int         `json:"sf"            dc:"停牌标志: 0-正常, 1-停牌"`
	TradingStatus string      `json:"tradingStatus" dc:"交易状态描述"`
	Period        string      `json:"period"        dc:"数据周期: 1min, 5min, 15min, 30min, 60min, day, week, month"`
	IsAdjusted    int         `json:"isAdjusted"    dc:"是否复权: 0-不复权, 1-前复权, 2-后复权"`
	DataQuality   int         `json:"dataQuality"   dc:"数据质量: 0-异常, 1-正常, 2-补全"`
	IsVerified    int         `json:"isVerified"    dc:"是否已验证: 0-未验证, 1-已验证"`
	DataSource    string      `json:"dataSource"    dc:"数据来源"`
	Version       int         `json:"version"       dc:"数据版本"`
}

// EnterpriseHistoricalDataEditInp 修改/新增企业级历史行情数据表 (K线数据)
type EnterpriseHistoricalDataEditInp struct {
	entity.EnterpriseHistoricalData
}

func (in *EnterpriseHistoricalDataEditInp) Filter(ctx context.Context) (err error) {
	// 验证证券代码 (如: 000001.SZ, AAPL)
	if err := g.Validator().Rules("required").Data(in.Symbol).Messages("证券代码 (如: 000001.SZ, AAPL)不能为空").Run(ctx); err != nil {
		return err.Current()
	}
	// 验证交易时间 (精确到分钟或日)
	if err := g.Validator().Rules("required").Data(in.T).Messages("交易时间 (精确到分钟或日)不能为空").Run(ctx); err != nil {
		return err.Current()
	}
	// 验证停牌标志: 0-正常, 1-停牌
	if err := g.Validator().Rules("required").Data(in.Sf).Messages("停牌标志: 0-正常, 1-停牌不能为空").Run(ctx); err != nil {
		return err.Current()
	}
	// 验证数据周期: 1min, 5min, 15min, 30min, 60min, day, week, month
	if err := g.Validator().Rules("required").Data(in.Period).Messages("数据周期: 1min, 5min, 15min, 30min, 60min, day, week, month不能为空").Run(ctx); err != nil {
		return err.Current()
	}
	return
}

type EnterpriseHistoricalDataEditModel struct{}

// EnterpriseHistoricalDataDeleteInp 删除企业级历史行情数据表 (K线数据)
type EnterpriseHistoricalDataDeleteInp struct {
	Id interface{} `json:"id" v:"required#自增主键不能为空" dc:"自增主键"`
}

func (in *EnterpriseHistoricalDataDeleteInp) Filter(ctx context.Context) (err error) {
	return
}

type EnterpriseHistoricalDataDeleteModel struct{}

// EnterpriseHistoricalDataViewInp 获取指定企业级历史行情数据表 (K线数据)信息
type EnterpriseHistoricalDataViewInp struct {
	Id int64 `json:"id" v:"required#自增主键不能为空" dc:"自增主键"`
}

func (in *EnterpriseHistoricalDataViewInp) Filter(ctx context.Context) (err error) {
	return
}

type EnterpriseHistoricalDataViewModel struct {
	entity.EnterpriseHistoricalData
}

// EnterpriseHistoricalDataListInp 获取企业级历史行情数据表 (K线数据)列表
type EnterpriseHistoricalDataListInp struct {
	form.PageReq
	Id        int64         `json:"id"        dc:"自增主键"`
	CreatedAt []*gtime.Time `json:"createdAt" dc:"创建时间"`
}

func (in *EnterpriseHistoricalDataListInp) Filter(ctx context.Context) (err error) {
	return
}

type EnterpriseHistoricalDataListModel struct {
	Id            int64       `json:"id"            dc:"自增主键"`
	Symbol        string      `json:"symbol"        dc:"证券代码 (如: 000001.SZ, AAPL)"`
	T             *gtime.Time `json:"t"             dc:"交易时间 (精确到分钟或日)"`
	Date          *gtime.Time `json:"date"          dc:"交易日期 (衍生字段)"`
	Year          int         `json:"year"          dc:"交易年份"`
	Month         int         `json:"month"         dc:"交易月份"`
	Weekday       int         `json:"weekday"       dc:"星期几 (1=周日,7=周六)"`
	O             float64     `json:"o"             dc:"开盘价"`
	H             float64     `json:"h"             dc:"最高价"`
	L             float64     `json:"l"             dc:"最低价"`
	C             float64     `json:"c"             dc:"收盘价"`
	Pc            float64     `json:"pc"            dc:"前收盘价"`
	V             int64       `json:"v"             dc:"成交量 (股/手)"`
	A             float64     `json:"a"             dc:"成交额 (元)"`
	Change        float64     `json:"change"        dc:"涨跌额"`
	ChangePct     float64     `json:"changePct"     dc:"涨跌幅 (%)"`
	Amplitude     float64     `json:"amplitude"     dc:"振幅 (%)"`
	Sf            int         `json:"sf"            dc:"停牌标志: 0-正常, 1-停牌"`
	TradingStatus string      `json:"tradingStatus" dc:"交易状态描述"`
	Period        string      `json:"period"        dc:"数据周期: 1min, 5min, 15min, 30min, 60min, day, week, month"`
	IsAdjusted    int         `json:"isAdjusted"    dc:"是否复权: 0-不复权, 1-前复权, 2-后复权"`
	DataQuality   int         `json:"dataQuality"   dc:"数据质量: 0-异常, 1-正常, 2-补全"`
	IsVerified    int         `json:"isVerified"    dc:"是否已验证: 0-未验证, 1-已验证"`
	DataSource    string      `json:"dataSource"    dc:"数据来源"`
	Version       int         `json:"version"       dc:"数据版本"`
	CreatedAt     *gtime.Time `json:"createdAt"     dc:"创建时间"`
	UpdatedAt     *gtime.Time `json:"updatedAt"     dc:"更新时间"`
}

// EnterpriseHistoricalDataExportModel 导出企业级历史行情数据表 (K线数据)
type EnterpriseHistoricalDataExportModel struct {
	Id            int64       `json:"id"            dc:"自增主键"`
	Symbol        string      `json:"symbol"        dc:"证券代码 (如: 000001.SZ, AAPL)"`
	T             *gtime.Time `json:"t"             dc:"交易时间 (精确到分钟或日)"`
	Date          *gtime.Time `json:"date"          dc:"交易日期 (衍生字段)"`
	Year          int         `json:"year"          dc:"交易年份"`
	Month         int         `json:"month"         dc:"交易月份"`
	Weekday       int         `json:"weekday"       dc:"星期几 (1=周日,7=周六)"`
	O             float64     `json:"o"             dc:"开盘价"`
	H             float64     `json:"h"             dc:"最高价"`
	L             float64     `json:"l"             dc:"最低价"`
	C             float64     `json:"c"             dc:"收盘价"`
	Pc            float64     `json:"pc"            dc:"前收盘价"`
	V             int64       `json:"v"             dc:"成交量 (股/手)"`
	A             float64     `json:"a"             dc:"成交额 (元)"`
	Change        float64     `json:"change"        dc:"涨跌额"`
	ChangePct     float64     `json:"changePct"     dc:"涨跌幅 (%)"`
	Amplitude     float64     `json:"amplitude"     dc:"振幅 (%)"`
	Sf            int         `json:"sf"            dc:"停牌标志: 0-正常, 1-停牌"`
	TradingStatus string      `json:"tradingStatus" dc:"交易状态描述"`
	Period        string      `json:"period"        dc:"数据周期: 1min, 5min, 15min, 30min, 60min, day, week, month"`
	IsAdjusted    int         `json:"isAdjusted"    dc:"是否复权: 0-不复权, 1-前复权, 2-后复权"`
	DataQuality   int         `json:"dataQuality"   dc:"数据质量: 0-异常, 1-正常, 2-补全"`
	IsVerified    int         `json:"isVerified"    dc:"是否已验证: 0-未验证, 1-已验证"`
	DataSource    string      `json:"dataSource"    dc:"数据来源"`
	Version       int         `json:"version"       dc:"数据版本"`
	CreatedAt     *gtime.Time `json:"createdAt"     dc:"创建时间"`
	UpdatedAt     *gtime.Time `json:"updatedAt"     dc:"更新时间"`
}
