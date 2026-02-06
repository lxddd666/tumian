// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// AtrData is the golang structure of table hg_atr_data for DAO operations like Where/Data.
type AtrData struct {
	g.Meta    `orm:"table:hg_atr_data, do:true"`
	Id        any         // 自增主键
	Symbol    any         // 股票或标的代码 (例如: AAPL, 000001.SZ)
	T         *gtime.Time // 交易时间 (统一为datetime类型，日线数据时间部分设为00:00:00)
	Tr        any         // 真实波幅 (True Range)
	Atr       any         // 平均真实波幅 (Average True Range)
	CreatedAt *gtime.Time // 数据创建时间
	UpdatedAt *gtime.Time // 数据更新时间
}
