// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// KdjData is the golang structure of table hg_kdj_data for DAO operations like Where/Data.
type KdjData struct {
	g.Meta       `orm:"table:hg_kdj_data, do:true"`
	Id           interface{} // 自增主键
	Symbol       interface{} // 股票或标的代码 (例如: AAPL, 000001.SZ)
	T            *gtime.Time // 交易时间 (统一为datetime类型，日线数据时间部分设为00:00:00)
	IntervalType interface{} // 数据间隔: minute-短分时, day-日线
	K            interface{} // K值
	D            interface{} // D值
	J            interface{} // J值
	CreatedAt    *gtime.Time // 数据创建时间
	UpdatedAt    *gtime.Time // 数据更新时间
}
