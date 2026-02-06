// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// StockSupportResistance is the golang structure of table hg_stock_support_resistance for DAO operations like Where/Data.
type StockSupportResistance struct {
	g.Meta    `orm:"table:hg_stock_support_resistance, do:true"`
	Id        interface{} // 主键ID
	T         *gtime.Time // 时间
	Symbol    interface{} // 股票或标的代码 (例如: AAPL, 000001.SZ)
	Price     interface{} // 当前股票价格
	Yl        interface{} // 压力位
	Zc        interface{} // 支撑位
	CreatedAt *gtime.Time // 创建时间
	Mc        interface{} // mc
}
