// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// MacdData is the golang structure of table hg_macd_data for DAO operations like Where/Data.
type MacdData struct {
	g.Meta       `orm:"table:hg_macd_data, do:true"`
	Id           any         // 自增主键
	Symbol       any         // 股票或标的代码 (例如: AAPL, 000001.SH)
	T            *gtime.Time // 交易时间 (统一为datetime类型，日线数据时间部分设为00:00:00)
	IntervalType any         // 数据间隔: minute-短分时, day-日线
	Diff         any         // DIFF值
	Dea          any         // DEA值
	Macd         any         // MACD值
	Ema12        any         // EMA(12)值
	Ema26        any         // EMA(26)值
	CreatedAt    *gtime.Time // 数据创建时间
	UpdatedAt    *gtime.Time // 数据更新时间
}
