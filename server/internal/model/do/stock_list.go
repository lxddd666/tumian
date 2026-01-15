// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// StockList is the golang structure of table hg_stock_list for DAO operations like Where/Data.
type StockList struct {
	g.Meta       `orm:"table:hg_stock_list, do:true"`
	Id           any         // 自增主键
	Dm           any         // 股票代码 (唯一业务标识，如: 000001)
	Mc           any         // 股票名称 (如: 平安银行)
	Jys          any         // 交易所代码 (如: sh, sz, bj)
	ExchangeName any         // 交易所全称
	Symbol       any         // 标准股票代码 (如: 000001.SZ)
	Status       any         // 状态: 1-正常, 0-退市
	ListDate     *gtime.Time // 上市日期
	DataSource   any         // 数据来源
	CreatedAt    *gtime.Time // 创建时间
	UpdatedAt    *gtime.Time // 更新时间
}
