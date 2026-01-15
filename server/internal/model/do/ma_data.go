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
	Id           interface{} // 自增主键
	Symbol       interface{} // 股票或标的代码 (例如: AAPL, 000001.SZ)
	T            *gtime.Time // 交易时间 (统一为datetime类型，日线数据时间部分设为00:00:00)
	IntervalType interface{} // 数据间隔: minute-短分时, day-日线
	Ma3          interface{} // MA3值
	Ma5          interface{} // MA5值
	Ma10         interface{} // MA10值
	Ma15         interface{} // MA15值
	Ma20         interface{} // MA20值
	Ma30         interface{} // MA30值
	Ma60         interface{} // MA60值
	Ma120        interface{} // MA120值
	Ma200        interface{} // MA200值
	Ma250        interface{} // MA250值
	CreatedAt    *gtime.Time // 数据创建时间
	UpdatedAt    *gtime.Time // 数据更新时间
}
