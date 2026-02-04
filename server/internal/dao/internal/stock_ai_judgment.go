// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// StockAiJudgmentDao is the data access object for the table hg_stock_ai_judgment.
type StockAiJudgmentDao struct {
	table    string                 // table is the underlying table name of the DAO.
	group    string                 // group is the database configuration group name of the current DAO.
	columns  StockAiJudgmentColumns // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler     // handlers for customized model modification.
}

// StockAiJudgmentColumns defines and stores column names for the table hg_stock_ai_judgment.
type StockAiJudgmentColumns struct {
	Id                       string // 自增主键
	Symbol                   string // 股票或标的代码 (例如: AAPL, 000001.SZ)
	Mc                       string // 股票名称
	IndicatorsJudgment       string // 指标判断
	IndicatorsFlag           string // 指标判断 1是2否
	FinancialJudgment        string // 财务判断
	FinancialFlag            string // 财务判断 1是2否
	ComprehensiveJudgment    string // 综合判断
	ComprehensiveFlag        string // 综合判断 1是2否
	Target                   string // 开仓止盈价格
	Stop                     string // 开仓止损价格
	JudgmentIndicatorsScript string // 指标判断话术
	JudgmentFinancialScript  string // 财务判断话术
	FinancialMttScript       string // 财报mtt评分
	SupportScript            string // 支撑位压力位
	ValuationMttScript       string // 估值
	CreatedAt                string // 创建时间
	AiId                     string // ai id
	AiName                   string // ai名称
	T                        string // 时间
}

// stockAiJudgmentColumns holds the columns for the table hg_stock_ai_judgment.
var stockAiJudgmentColumns = StockAiJudgmentColumns{
	Id:                       "id",
	Symbol:                   "symbol",
	Mc:                       "mc",
	IndicatorsJudgment:       "indicators_judgment",
	IndicatorsFlag:           "indicators_flag",
	FinancialJudgment:        "financial_judgment",
	FinancialFlag:            "financial_flag",
	ComprehensiveJudgment:    "comprehensive_judgment",
	ComprehensiveFlag:        "comprehensive_flag",
	Target:                   "target",
	Stop:                     "stop",
	JudgmentIndicatorsScript: "judgment_indicators_script",
	JudgmentFinancialScript:  "judgment_financial_script",
	FinancialMttScript:       "financial_mtt_script",
	SupportScript:            "support_script",
	ValuationMttScript:       "valuation_mtt_script",
	CreatedAt:                "created_at",
	AiId:                     "ai_id",
	AiName:                   "ai_name",
	T:                        "t",
}

// NewStockAiJudgmentDao creates and returns a new DAO object for table data access.
func NewStockAiJudgmentDao(handlers ...gdb.ModelHandler) *StockAiJudgmentDao {
	return &StockAiJudgmentDao{
		group:    "default",
		table:    "hg_stock_ai_judgment",
		columns:  stockAiJudgmentColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *StockAiJudgmentDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *StockAiJudgmentDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *StockAiJudgmentDao) Columns() StockAiJudgmentColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *StockAiJudgmentDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *StockAiJudgmentDao) Ctx(ctx context.Context) *gdb.Model {
	model := dao.DB().Model(dao.table)
	for _, handler := range dao.handlers {
		model = handler(model)
	}
	return model.Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rolls back the transaction and returns the error if function f returns a non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note: Do not commit or roll back the transaction in function f,
// as it is automatically handled by this function.
func (dao *StockAiJudgmentDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
