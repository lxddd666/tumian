// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// StockScoreMainDao is the data access object for the table hg_stock_score_main.
type StockScoreMainDao struct {
	table    string                // table is the underlying table name of the DAO.
	group    string                // group is the database configuration group name of the current DAO.
	columns  StockScoreMainColumns // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler    // handlers for customized model modification.
}

// StockScoreMainColumns defines and stores column names for the table hg_stock_score_main.
type StockScoreMainColumns struct {
	Id                 string // 主键ID
	Symbol             string //
	Mc                 string // mc
	T                  string // 评分日期
	ComprehensiveScore string // 综合评分
	PriceScore         string // 价格动量评分
	IncomeScore        string // 收益预测评分
	FinancialScore     string // 财务分析评分
	ValuationScore     string // 相对估值评分
	RiskScore          string // 风险评分
	DataSource         string // 数据来源
	UpdatedAt          string // 更新时间
}

// stockScoreMainColumns holds the columns for the table hg_stock_score_main.
var stockScoreMainColumns = StockScoreMainColumns{
	Id:                 "id",
	Symbol:             "symbol",
	Mc:                 "mc",
	T:                  "t",
	ComprehensiveScore: "comprehensive_score",
	PriceScore:         "price_score",
	IncomeScore:        "income_score",
	FinancialScore:     "financial_score",
	ValuationScore:     "valuation_score",
	RiskScore:          "risk_score",
	DataSource:         "data_source",
	UpdatedAt:          "updated_at",
}

// NewStockScoreMainDao creates and returns a new DAO object for table data access.
func NewStockScoreMainDao(handlers ...gdb.ModelHandler) *StockScoreMainDao {
	return &StockScoreMainDao{
		group:    "default",
		table:    "hg_stock_score_main",
		columns:  stockScoreMainColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *StockScoreMainDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *StockScoreMainDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *StockScoreMainDao) Columns() StockScoreMainColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *StockScoreMainDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *StockScoreMainDao) Ctx(ctx context.Context) *gdb.Model {
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
func (dao *StockScoreMainDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
