// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// StockScoreMain is the golang structure for table stock_score_main.
type StockScoreMain struct {
	Id                 uint64      `json:"id"                 orm:"id"                  description:"主键ID"`
	Symbol             string      `json:"symbol"             orm:"symbol"              description:"股票代码"`
	T                  *gtime.Time `json:"t"                  orm:"t"                   description:"评分日期"`
	ComprehensiveScore float64     `json:"comprehensiveScore" orm:"comprehensive_score" description:"综合评分"`
	PriceScore         float64     `json:"priceScore"         orm:"price_score"         description:"价格动量评分"`
	IncomeScore        float64     `json:"incomeScore"        orm:"income_score"        description:"收益预测评分"`
	FinancialScore     float64     `json:"financialScore"     orm:"financial_score"     description:"财务分析评分"`
	ValuationScore     float64     `json:"valuationScore"     orm:"valuation_score"     description:"相对估值评分"`
	RiskScore          float64     `json:"riskScore"          orm:"risk_score"          description:"风险评分"`
	DataSource         string      `json:"dataSource"         orm:"data_source"         description:"数据来源"`
	UpdatedAt          *gtime.Time `json:"updatedAt"          orm:"updated_at"          description:"更新时间"`
}
