// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// StockAiJudgment is the golang structure of table hg_stock_ai_judgment for DAO operations like Where/Data.
type StockAiJudgment struct {
	g.Meta                   `orm:"table:hg_stock_ai_judgment, do:true"`
	Id                       interface{} // 自增主键
	Symbol                   interface{} // 股票或标的代码 (例如: AAPL, 000001.SZ)
	Mc                       interface{} // 股票名称
	IndicatorsJudgment       interface{} // 指标判断
	IndicatorsFlag           interface{} // 指标判断 1是2否
	FinancialJudgment        interface{} // 财务判断
	FinancialFlag            interface{} // 财务判断 1是2否
	FinancialRatingJudgment  interface{} // 财报评分
	FinancialRatingFlag      interface{} // 财报评分 1是2否
	MarketAnalysisJudgment   interface{} // 市场资金分析
	MarketAnalysisFlag       interface{} // 市场资金分析 1是2否
	ValueAssessmentJudgment  interface{} // 价值估算
	ValueAssessmentFlag      interface{} // 价值估算 1是2否
	ComprehensiveJudgment    interface{} // 综合判断
	ComprehensiveFlag        interface{} // 综合判断 1是2否
	Target                   interface{} // 开仓止盈价格
	Stop                     interface{} // 开仓止损价格
	JudgmentIndicatorsScript interface{} // 指标判断话术
	JudgmentFinancialScript  interface{} // 财务判断话术
	FinancialMttScript       interface{} // 财报mtt评分
	SupportScript            interface{} // 支撑位压力位
	ValuationMttScript       interface{} // 估值
	CreatedAt                *gtime.Time // 创建时间
	AiId                     interface{} // ai id
	AiName                   interface{} // ai名称
	T                        *gtime.Time // 时间
}
