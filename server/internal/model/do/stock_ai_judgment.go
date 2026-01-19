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
	g.Meta                `orm:"table:hg_stock_ai_judgment, do:true"`
	Id                    any         // 自增主键
	Symbol                any         // 股票或标的代码 (例如: AAPL, 000001.SZ)
	T                     *gtime.Time // 时间
	AiId                  any         // ai id
	IndicatorsJudgment    any         // 指标判断
	IndicatorsFlag        any         // 指标判断 1是2否
	FinancialJudgment     any         // 财务判断
	FinancialFlag         any         // 财务判断 1是2否
	ComprehensiveJudgment any         // 综合判断
	ComprehensiveFlag     any         // 综合判断 1是2否
	CreatedAt             *gtime.Time // 创建时间
}
