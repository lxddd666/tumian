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

	"github.com/gogf/gf/v2/os/gtime"
)

// StockScoreMainUpdateFields 修改股票评分主表字段过滤
type StockScoreMainUpdateFields struct {
	StockCode              string      `json:"stockCode"              dc:"股票代码"`
	ScoreDate              *gtime.Time `json:"scoreDate"              dc:"评分日期"`
	ComprehensiveScore     float64     `json:"comprehensiveScore"     dc:"综合评分"`
	PriceMomentumScore     float64     `json:"priceMomentumScore"     dc:"价格动量评分"`
	EarningsForecastScore  float64     `json:"earningsForecastScore"  dc:"收益预测评分"`
	FinancialAnalysisScore float64     `json:"financialAnalysisScore" dc:"财务分析评分"`
	RelativeValuationScore float64     `json:"relativeValuationScore" dc:"相对估值评分"`
	PeTtm                  float64     `json:"peTtm"                  dc:"市盈率(TTM)"`
	PePercentile30         float64     `json:"pePercentile30"         dc:"市盈率30分位值"`
	PePercentile70         float64     `json:"pePercentile70"         dc:"市盈率70分位值"`
	Pb                     float64     `json:"pb"                     dc:"市净率"`
	PbPercentile30         float64     `json:"pbPercentile30"         dc:"市净率30分位值"`
	PbPercentile70         float64     `json:"pbPercentile70"         dc:"市净率70分位值"`
	DataSource             string      `json:"dataSource"             dc:"数据来源"`
	UpdateTime             *gtime.Time `json:"updateTime"             dc:"更新时间"`
}

// StockScoreMainInsertFields 新增股票评分主表字段过滤
type StockScoreMainInsertFields struct {
	StockCode              string      `json:"stockCode"              dc:"股票代码"`
	ScoreDate              *gtime.Time `json:"scoreDate"              dc:"评分日期"`
	ComprehensiveScore     float64     `json:"comprehensiveScore"     dc:"综合评分"`
	PriceMomentumScore     float64     `json:"priceMomentumScore"     dc:"价格动量评分"`
	EarningsForecastScore  float64     `json:"earningsForecastScore"  dc:"收益预测评分"`
	FinancialAnalysisScore float64     `json:"financialAnalysisScore" dc:"财务分析评分"`
	RelativeValuationScore float64     `json:"relativeValuationScore" dc:"相对估值评分"`
	PeTtm                  float64     `json:"peTtm"                  dc:"市盈率(TTM)"`
	PePercentile30         float64     `json:"pePercentile30"         dc:"市盈率30分位值"`
	PePercentile70         float64     `json:"pePercentile70"         dc:"市盈率70分位值"`
	Pb                     float64     `json:"pb"                     dc:"市净率"`
	PbPercentile30         float64     `json:"pbPercentile30"         dc:"市净率30分位值"`
	PbPercentile70         float64     `json:"pbPercentile70"         dc:"市净率70分位值"`
	DataSource             string      `json:"dataSource"             dc:"数据来源"`
	UpdateTime             *gtime.Time `json:"updateTime"             dc:"更新时间"`
}

// StockScoreMainEditInp 修改/新增股票评分主表
type StockScoreMainEditInp struct {
	entity.StockScoreMain
}

func (in *StockScoreMainEditInp) Filter(ctx context.Context) (err error) {
	// 验证股票代码

	// 验证更新时间

	return
}

type StockScoreMainEditModel struct{}

// StockScoreMainDeleteInp 删除股票评分主表
type StockScoreMainDeleteInp struct {
	Id interface{} `json:"id" v:"required#主键ID不能为空" dc:"主键ID"`
}

func (in *StockScoreMainDeleteInp) Filter(ctx context.Context) (err error) {
	return
}

type StockScoreMainDeleteModel struct{}

// StockScoreMainViewInp 获取指定股票评分主表信息
type StockScoreMainViewInp struct {
	Id int64 `json:"id" v:"required#主键ID不能为空" dc:"主键ID"`
}

func (in *StockScoreMainViewInp) Filter(ctx context.Context) (err error) {
	return
}

type StockScoreMainViewModel struct {
	entity.StockScoreMain
}

// StockScoreMainListInp 获取股票评分主表列表
type StockScoreMainListInp struct {
	form.PageReq
	Id int64 `json:"id" dc:"主键ID"`
}

func (in *StockScoreMainListInp) Filter(ctx context.Context) (err error) {
	return
}

type StockScoreMainListModel struct {
	Id                     int64       `json:"id"                     dc:"主键ID"`
	StockCode              string      `json:"stockCode"              dc:"股票代码"`
	ScoreDate              *gtime.Time `json:"scoreDate"              dc:"评分日期"`
	ComprehensiveScore     float64     `json:"comprehensiveScore"     dc:"综合评分"`
	PriceMomentumScore     float64     `json:"priceMomentumScore"     dc:"价格动量评分"`
	EarningsForecastScore  float64     `json:"earningsForecastScore"  dc:"收益预测评分"`
	FinancialAnalysisScore float64     `json:"financialAnalysisScore" dc:"财务分析评分"`
	RelativeValuationScore float64     `json:"relativeValuationScore" dc:"相对估值评分"`
	PeTtm                  float64     `json:"peTtm"                  dc:"市盈率(TTM)"`
	PePercentile30         float64     `json:"pePercentile30"         dc:"市盈率30分位值"`
	PePercentile70         float64     `json:"pePercentile70"         dc:"市盈率70分位值"`
	Pb                     float64     `json:"pb"                     dc:"市净率"`
	PbPercentile30         float64     `json:"pbPercentile30"         dc:"市净率30分位值"`
	PbPercentile70         float64     `json:"pbPercentile70"         dc:"市净率70分位值"`
	DataSource             string      `json:"dataSource"             dc:"数据来源"`
	UpdateTime             *gtime.Time `json:"updateTime"             dc:"更新时间"`
}

// StockScoreMainExportModel 导出股票评分主表
type StockScoreMainExportModel struct {
	Id                     int64       `json:"id"                     dc:"主键ID"`
	StockCode              string      `json:"stockCode"              dc:"股票代码"`
	ScoreDate              *gtime.Time `json:"scoreDate"              dc:"评分日期"`
	ComprehensiveScore     float64     `json:"comprehensiveScore"     dc:"综合评分"`
	PriceMomentumScore     float64     `json:"priceMomentumScore"     dc:"价格动量评分"`
	EarningsForecastScore  float64     `json:"earningsForecastScore"  dc:"收益预测评分"`
	FinancialAnalysisScore float64     `json:"financialAnalysisScore" dc:"财务分析评分"`
	RelativeValuationScore float64     `json:"relativeValuationScore" dc:"相对估值评分"`
	PeTtm                  float64     `json:"peTtm"                  dc:"市盈率(TTM)"`
	PePercentile30         float64     `json:"pePercentile30"         dc:"市盈率30分位值"`
	PePercentile70         float64     `json:"pePercentile70"         dc:"市盈率70分位值"`
	Pb                     float64     `json:"pb"                     dc:"市净率"`
	PbPercentile30         float64     `json:"pbPercentile30"         dc:"市净率30分位值"`
	PbPercentile70         float64     `json:"pbPercentile70"         dc:"市净率70分位值"`
	DataSource             string      `json:"dataSource"             dc:"数据来源"`
	UpdateTime             *gtime.Time `json:"updateTime"             dc:"更新时间"`
}
