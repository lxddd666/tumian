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
	Id           interface{} // 自增主键
	Dm           interface{} // 股票代码 (唯一业务标识，如: 000001)
	Mc           interface{} // 股票名称 (如: 平安银行)
	Jys          interface{} // 交易所代码 (如: sh, sz, bj)
	ExchangeName interface{} // 交易所全称
	Symbol       interface{} // 标准股票代码 (如: 000001.SZ)
	Status       interface{} // 状态: 1-正常, 0-退市
	ListDate     *gtime.Time // 上市日期
	DataSource   interface{} // 数据来源
	CreatedAt    *gtime.Time // 创建时间
	UpdatedAt    *gtime.Time // 更新时间
}
