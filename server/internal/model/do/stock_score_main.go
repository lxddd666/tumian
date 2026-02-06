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
	Id                 interface{} // 主键ID
	Symbol             interface{} // 股票代码
	Mc                 interface{} // mc
	T                  *gtime.Time // 评分日期
	ComprehensiveScore interface{} // 综合评分
	PriceScore         interface{} // 价格动量评分
	IncomeScore        interface{} // 收益预测评分
	FinancialScore     interface{} // 财务分析评分
	ValuationScore     interface{} // 相对估值评分
	RiskScore          interface{} // 风险评分
	DataSource         interface{} // 数据来源
	UpdatedAt          *gtime.Time // 更新时间
}
