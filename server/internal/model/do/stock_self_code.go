// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// StockSelfCode is the golang structure of table hg_stock_self_code for DAO operations like Where/Data.
type StockSelfCode struct {
	g.Meta    `orm:"table:hg_stock_self_code, do:true"`
	Dm        any         // 股票代码
	Mc        any         // 股票名称
	Jys       any         // 交易所
	CreatedAt *gtime.Time // 创建时间
}
