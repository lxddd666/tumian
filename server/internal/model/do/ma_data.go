// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// MaData is the golang structure of table hg_ma_data for DAO operations like Where/Data.
type MaData struct {
	g.Meta       `orm:"table:hg_ma_data, do:true"`
	Id           any         // 自增主键
	Symbol       any         // 股票或标的代码 (例如: AAPL, 000001.SZ)
	T            *gtime.Time // 交易时间 (统一为datetime类型，日线数据时间部分设为00:00:00)
	IntervalType any         // 数据间隔: minute-短分时, day-日线
	Ma3          any         // MA3值
	Ma5          any         // MA5值
	Ma10         any         // MA10值
	Ma15         any         // MA15值
	Ma20         any         // MA20值
	Ma30         any         // MA30值
	Ma60         any         // MA60值
	Ma120        any         // MA120值
	Ma200        any         // MA200值
	Ma250        any         // MA250值
	CreatedAt    *gtime.Time // 数据创建时间
	UpdatedAt    *gtime.Time // 数据更新时间
}
