// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// StockAllCode is the golang structure of table hg_stock_all_code for DAO operations like Where/Data.
type StockAllCode struct {
	g.Meta    `orm:"table:hg_stock_all_code, do:true"`
	Dm        interface{} // 股票代码
	Mc        interface{} // 股票名称
	Jys       interface{} // 交易所
	CreatedAt *gtime.Time // 创建时间
}
