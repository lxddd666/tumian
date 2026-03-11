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

// StockAiJudgmentUpdateFields 修改ai 选股判断字段过滤
type StockAiJudgmentUpdateFields struct {
	Symbol                string      `json:"symbol"                dc:"股票或标的代码 (例如: AAPL, 000001.SZ)"`
	T                     *gtime.Time `json:"t"                     dc:"时间"`
	AiId                  int64       `json:"aiId"                  dc:"ai id"`
	IndicatorsJudgment    string      `json:"indicatorsJudgment"    dc:"指标判断"`
	IndicatorsFlag        int         `json:"indicatorsFlag"        dc:"指标判断 1是2否"`
	FinancialJudgment     string      `json:"financialJudgment"     dc:"财务判断"`
	FinancialFlag         int         `json:"financialFlag"         dc:"财务判断 1是2否"`
	ComprehensiveJudgment string      `json:"comprehensiveJudgment" dc:"综合判断"`
	ComprehensiveFlag     int         `json:"comprehensiveFlag"     dc:"综合判断 1是2否"`
}

// StockAiJudgmentInsertFields 新增ai 选股判断字段过滤
type StockAiJudgmentInsertFields struct {
	Symbol                string      `json:"symbol"                dc:"股票或标的代码 (例如: AAPL, 000001.SZ)"`
	T                     *gtime.Time `json:"t"                     dc:"时间"`
	AiId                  int64       `json:"aiId"                  dc:"ai id"`
	IndicatorsJudgment    string      `json:"indicatorsJudgment"    dc:"指标判断"`
	IndicatorsFlag        int         `json:"indicatorsFlag"        dc:"指标判断 1是2否"`
	FinancialJudgment     string      `json:"financialJudgment"     dc:"财务判断"`
	FinancialFlag         int         `json:"financialFlag"         dc:"财务判断 1是2否"`
	ComprehensiveJudgment string      `json:"comprehensiveJudgment" dc:"综合判断"`
	ComprehensiveFlag     int         `json:"comprehensiveFlag"     dc:"综合判断 1是2否"`
}

// StockAiJudgmentEditInp 修改/新增ai 选股判断
type StockAiJudgmentEditInp struct {
	entity.StockAiJudgment
}

func (in *StockAiJudgmentEditInp) Filter(ctx context.Context) (err error) {
	// 验证股票或标的代码 (例如: AAPL, 000001.SZ)
	if err := g.Validator().Rules("required").Data(in.Symbol).Messages("股票或标的代码 (例如: AAPL, 000001.SZ)不能为空").Run(ctx); err != nil {
		return err.Current()
	}
	// 验证时间
	if err := g.Validator().Rules("required").Data(in.T).Messages("时间不能为空").Run(ctx); err != nil {
		return err.Current()
	}
	// 验证ai id
	if err := g.Validator().Rules("required").Data(in.AiId).Messages("ai id不能为空").Run(ctx); err != nil {
		return err.Current()
	}
	return
}

type StockAiJudgmentEditModel struct{}

// StockAiJudgmentDeleteInp 删除ai 选股判断
type StockAiJudgmentDeleteInp struct {
	Id interface{} `json:"id" v:"required#自增主键不能为空" dc:"自增主键"`
}

type StockAiJudgmentAiJudgmentInp struct {
	Symbol string `json:"symbol"                dc:"股票或标的代码 (例如: AAPL, 000001.SZ)"`
}

func (in *StockAiJudgmentDeleteInp) Filter(ctx context.Context) (err error) {
	return
}

type StockAiJudgmentDeleteModel struct{}

// StockAiJudgmentViewInp 获取指定ai 选股判断信息
type StockAiJudgmentViewInp struct {
	Id int64 `json:"id" v:"required#自增主键不能为空" dc:"自增主键"`
}

func (in *StockAiJudgmentViewInp) Filter(ctx context.Context) (err error) {
	return
}

type StockAiJudgmentViewModel struct {
	entity.StockAiJudgment
}

// StockAiJudgmentListInp 获取ai 选股判断列表
type StockAiJudgmentListInp struct {
	form.PageReq
	Id        int64         `json:"id"        dc:"自增主键"`
	CreatedAt []*gtime.Time `json:"createdAt" dc:"创建时间"`
}

func (in *StockAiJudgmentListInp) Filter(ctx context.Context) (err error) {
	return
}

type StockAiJudgmentListModel struct {
	Id                int64       `json:"id"                dc:"自增主键"`
	Symbol            string      `json:"symbol"            dc:"股票或标的代码 (例如: AAPL, 000001.SZ)"`
	T                 *gtime.Time `json:"t"                 dc:"时间"`
	AiId              int64       `json:"aiId"              dc:"ai id"`
	IndicatorsFlag    int         `json:"indicatorsFlag"    dc:"指标判断 1是2否"`
	FinancialFlag     int         `json:"financialFlag"     dc:"财务判断 1是2否"`
	ComprehensiveFlag int         `json:"comprehensiveFlag" dc:"综合判断 1是2否"`
	CreatedAt         *gtime.Time `json:"createdAt"         dc:"创建时间"`
}

// StockAiJudgmentExportModel 导出ai 选股判断
type StockAiJudgmentExportModel struct {
	Id                       uint64      `json:"id"                       orm:"id"                         description:"自增主键"`
	Symbol                   string      `json:"symbol"                   orm:"symbol"                     description:"股票或标的代码 (例如: AAPL, 000001.SZ)"`
	TotalScore               int         `json:"totalScore" dc:"总分"`
	Mc                       string      `json:"mc"                       orm:"mc"                         description:"股票名称"`
	IndicatorsJudgment       string      `json:"indicatorsJudgment"       orm:"indicators_judgment"        description:"指标判断"`
	IndicatorsFlag           string      `json:"indicatorsFlag"           orm:"indicators_flag"            description:"指标判断 1是2否"`
	FinancialJudgment        string      `json:"financialJudgment"        orm:"financial_judgment"         description:"财务判断"`
	FinancialFlag            string      `json:"financialFlag"            orm:"financial_flag"             description:"财务判断 1是2否"`
	FinancialRatingJudgment  string      `json:"financialRatingJudgment"  orm:"financial_rating_judgment"  description:"财报评分"`
	FinancialRatingFlag      string      `json:"financialRatingFlag"      orm:"financial_rating_flag"      description:"财报评分 1是2否"`
	MarketAnalysisJudgment   string      `json:"marketAnalysisJudgment"   orm:"market_analysis_judgment"   description:"市场资金分析"`
	MarketAnalysisFlag       string      `json:"marketAnalysisFlag"       orm:"market_analysis_flag"       description:"市场资金分析 1是2否"`
	ValueAssessmentJudgment  string      `json:"valueAssessmentJudgment"  orm:"value_assessment_judgment"  description:"价值估算"`
	ValueAssessmentFlag      string      `json:"valueAssessmentFlag"      orm:"value_assessment_flag"      description:"价值估算 1是2否"`
	ComprehensiveJudgment    string      `json:"comprehensiveJudgment"    orm:"comprehensive_judgment"     description:"综合判断"`
	ComprehensiveFlag        string      `json:"comprehensiveFlag"        orm:"comprehensive_flag"         description:"综合判断 1是2否"`
	Target                   float64     `json:"target"                   orm:"target"                     description:"开仓止盈价格"`
	Stop                     float64     `json:"stop"                     orm:"stop"                       description:"开仓止损价格"`
	JudgmentIndicatorsScript string      `json:"judgmentIndicatorsScript" orm:"judgment_indicators_script" description:"指标判断话术"`
	JudgmentFinancialScript  string      `json:"judgmentFinancialScript"  orm:"judgment_financial_script"  description:"财务判断话术"`
	FinancialMttScript       string      `json:"financialMttScript"       orm:"financial_mtt_script"       description:"财报mtt评分"`
	SupportScript            string      `json:"supportScript"            orm:"support_script"             description:"支撑位压力位"`
	ValuationMttScript       string      `json:"valuationMttScript"       orm:"valuation_mtt_script"       description:"估值"`
	CreatedAt                *gtime.Time `json:"createdAt"                orm:"created_at"                 description:"创建时间"`
	AiId                     int64       `json:"aiId"                     orm:"ai_id"                      description:"ai id"`
	AiName                   string      `json:"aiName"                   orm:"ai_name"                    description:"ai名称"`
	T                        *gtime.Time `json:"t"                        orm:"t"                          description:"时间"`
}
