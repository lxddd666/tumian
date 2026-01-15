// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// BollData is the golang structure for table boll_data.
type BollData struct {
	Id           uint64      `json:"id"           orm:"id"            description:"自增主键"`
	Symbol       string      `json:"symbol"       orm:"symbol"        description:"股票或标的代码 (例如: AAPL, 000001.SZ)"`
	T            *gtime.Time `json:"t"            orm:"t"             description:"交易时间 (统一为datetime类型，日线数据时间部分设为00:00:00)"`
	IntervalType string      `json:"intervalType" orm:"interval_type" description:"数据间隔: minute-短分时, day-日线"`
	U            float64     `json:"u"            orm:"u"             description:"上轨(Upper Band)"`
	M            float64     `json:"m"            orm:"m"             description:"中轨(Middle Band)"`
	D            float64     `json:"d"            orm:"d"             description:"下轨(Lower Band)"`
	CreatedAt    *gtime.Time `json:"createdAt"    orm:"created_at"    description:"数据创建时间"`
	UpdatedAt    *gtime.Time `json:"updatedAt"    orm:"updated_at"    description:"数据更新时间"`
}
