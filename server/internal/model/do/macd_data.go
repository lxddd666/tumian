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
	Id           interface{} // 自增主键
	Symbol       interface{} // 股票或标的代码 (例如: AAPL, 000001.SH)
	T            *gtime.Time // 交易时间 (统一为datetime类型，日线数据时间部分设为00:00:00)
	IntervalType interface{} // 数据间隔: minute-短分时, day-日线
	Diff         interface{} // DIFF值
	Dea          interface{} // DEA值
	Macd         interface{} // MACD值
	Ema12        interface{} // EMA(12)值
	Ema26        interface{} // EMA(26)值
	CreatedAt    *gtime.Time // 数据创建时间
	UpdatedAt    *gtime.Time // 数据更新时间
}
