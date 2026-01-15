// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// MacdData is the golang structure for table macd_data.
type MacdData struct {
	Id           uint64      `json:"id"           orm:"id"            description:"自增主键"`
	Symbol       string      `json:"symbol"       orm:"symbol"        description:"股票或标的代码 (例如: AAPL, 000001.SH)"`
	T            *gtime.Time `json:"t"            orm:"t"             description:"交易时间 (统一为datetime类型，日线数据时间部分设为00:00:00)"`
	IntervalType string      `json:"intervalType" orm:"interval_type" description:"数据间隔: minute-短分时, day-日线"`
	Diff         float64     `json:"diff"         orm:"diff"          description:"DIFF值"`
	Dea          float64     `json:"dea"          orm:"dea"           description:"DEA值"`
	Macd         float64     `json:"macd"         orm:"macd"          description:"MACD值"`
	Ema12        float64     `json:"ema12"        orm:"ema12"         description:"EMA(12)值"`
	Ema26        float64     `json:"ema26"        orm:"ema26"         description:"EMA(26)值"`
	CreatedAt    *gtime.Time `json:"createdAt"    orm:"created_at"    description:"数据创建时间"`
	UpdatedAt    *gtime.Time `json:"updatedAt"    orm:"updated_at"    description:"数据更新时间"`
}
