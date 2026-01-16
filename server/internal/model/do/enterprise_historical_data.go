// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// EnterpriseHistoricalData is the golang structure of table hg_enterprise_historical_data for DAO operations like Where/Data.
type EnterpriseHistoricalData struct {
	g.Meta    `orm:"table:hg_enterprise_historical_data, do:true"`
	Id        any         // 主键ID
	Symbol    any         // 公司代码/股票代码 (例如: 000001.SZ, AAPL)
	T         *gtime.Time // 交易时间
	O         any         // 开盘价
	H         any         // 最高价
	L         any         // 最低价
	C         any         // 收盘价
	V         any         // 成交量
	A         any         // 成交额
	Pc        any         // 前收盘价
	Sf        any         // 停牌状态: 1停牌, 0不停牌
	CreatedAt *gtime.Time // 创建时间
	UpdatedAt *gtime.Time // 更新时间
}
