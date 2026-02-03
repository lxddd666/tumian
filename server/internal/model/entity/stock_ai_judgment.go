// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// StockAiJudgment is the golang structure for table stock_ai_judgment.
type StockAiJudgment struct {
	Id                       uint64      `json:"id"                       orm:"id"                         description:"自增主键"`
	Symbol                   string      `json:"symbol"                   orm:"symbol"                     description:"股票或标的代码 (例如: AAPL, 000001.SZ)"`
	Mc                       string      `json:"mc"                       orm:"mc"                         description:"股票名称"`
	IndicatorsJudgment       string      `json:"indicatorsJudgment"       orm:"indicators_judgment"        description:"指标判断"`
	IndicatorsFlag           string      `json:"indicatorsFlag"           orm:"indicators_flag"            description:"指标判断 1是2否"`
	FinancialJudgment        string      `json:"financialJudgment"        orm:"financial_judgment"         description:"财务判断"`
	FinancialFlag            string      `json:"financialFlag"            orm:"financial_flag"             description:"财务判断 1是2否"`
	ComprehensiveJudgment    string      `json:"comprehensiveJudgment"    orm:"comprehensive_judgment"     description:"综合判断"`
	ComprehensiveFlag        string      `json:"comprehensiveFlag"        orm:"comprehensive_flag"         description:"综合判断 1是2否"`
	Target                   float64     `json:"target"                   orm:"target"                     description:"开仓止盈价格"`
	Stop                     float64     `json:"stop"                     orm:"stop"                       description:"开仓止损价格"`
	JudgmentIndicatorsScript string      `json:"judgmentIndicatorsScript" orm:"judgment_indicators_script" description:"指标判断话术"`
	JudgmentFinancialScript  string      `json:"judgmentFinancialScript"  orm:"judgment_financial_script"  description:"财务判断话术"`
	CreatedAt                *gtime.Time `json:"createdAt"                orm:"created_at"                 description:"创建时间"`
	AiId                     int64       `json:"aiId"                     orm:"ai_id"                      description:"ai id"`
	AiName                   string      `json:"aiName"                   orm:"ai_name"                    description:"ai名称"`
	T                        *gtime.Time `json:"t"                        orm:"t"                          description:"时间"`
}
