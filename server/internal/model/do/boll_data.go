// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// BollData is the golang structure of table hg_boll_data for DAO operations like Where/Data.
type BollData struct {
	g.Meta       `orm:"table:hg_boll_data, do:true"`
	Id           interface{} // 自增主键
	Symbol       interface{} // 股票或标的代码 (例如: AAPL, 000001.SZ)
	T            *gtime.Time // 交易时间 (统一为datetime类型，日线数据时间部分设为00:00:00)
	IntervalType interface{} // 数据间隔: minute-短分时, day-日线
	U            interface{} // 上轨(Upper Band)
	M            interface{} // 中轨(Middle Band)
	D            interface{} // 下轨(Lower Band)
	CreatedAt    *gtime.Time // 数据创建时间
	UpdatedAt    *gtime.Time // 数据更新时间
}
