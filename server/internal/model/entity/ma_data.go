// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// MaData is the golang structure for table ma_data.
type MaData struct {
	Id           uint64      `json:"id"           orm:"id"            description:"自增主键"`
	Symbol       string      `json:"symbol"       orm:"symbol"        description:"股票或标的代码 (例如: AAPL, 000001.SZ)"`
	T            *gtime.Time `json:"t"            orm:"t"             description:"交易时间 (统一为datetime类型，日线数据时间部分设为00:00:00)"`
	IntervalType string      `json:"intervalType" orm:"interval_type" description:"数据间隔: minute-短分时, day-日线"`
	Ma3          float64     `json:"ma3"          orm:"ma3"           description:"MA3值"`
	Ma5          float64     `json:"ma5"          orm:"ma5"           description:"MA5值"`
	Ma10         float64     `json:"ma10"         orm:"ma10"          description:"MA10值"`
	Ma15         float64     `json:"ma15"         orm:"ma15"          description:"MA15值"`
	Ma20         float64     `json:"ma20"         orm:"ma20"          description:"MA20值"`
	Ma30         float64     `json:"ma30"         orm:"ma30"          description:"MA30值"`
	Ma60         float64     `json:"ma60"         orm:"ma60"          description:"MA60值"`
	Ma120        float64     `json:"ma120"        orm:"ma120"         description:"MA120值"`
	Ma200        float64     `json:"ma200"        orm:"ma200"         description:"MA200值"`
	Ma250        float64     `json:"ma250"        orm:"ma250"         description:"MA250值"`
	CreatedAt    *gtime.Time `json:"createdAt"    orm:"created_at"    description:"数据创建时间"`
	UpdatedAt    *gtime.Time `json:"updatedAt"    orm:"updated_at"    description:"数据更新时间"`
}
