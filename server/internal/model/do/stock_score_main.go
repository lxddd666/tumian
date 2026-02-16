// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// StockScoreMain is the golang structure of table hg_stock_score_main for DAO operations like Where/Data.
type StockScoreMain struct {
	g.Meta             `orm:"table:hg_stock_score_main, do:true"`
	Id                 any         // 主键ID
	Symbol             any         //
	Mc                 any         // mc
	T                  *gtime.Time // 评分日期
	ComprehensiveScore any         // 综合评分
	PriceScore         any         // 价格动量评分
	IncomeScore        any         // 收益预测评分
	FinancialScore     any         // 财务分析评分
	ValuationScore     any         // 相对估值评分
	RiskScore          any         // 风险评分
	DataSource         any         // 数据来源
	UpdatedAt          *gtime.Time // 更新时间
}
