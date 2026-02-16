// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// SlowStochasticData is the golang structure for table slow_stochastic_data.
type SlowStochasticData struct {
	Id        uint64      `json:"id"        orm:"id"         description:"自增主键"`
	Symbol    string      `json:"symbol"    orm:"symbol"     description:"股票或标的代码 (例如: AAPL, 000001.SZ)"`
	T         *gtime.Time `json:"t"         orm:"t"          description:"交易时间 (统一为datetime类型，日线数据时间部分设为00:00:00)"`
	K         float64     `json:"k"         orm:"k"          description:"k"`
	D         float64     `json:"d"         orm:"d"          description:"d"`
	CreatedAt *gtime.Time `json:"createdAt" orm:"created_at" description:"数据创建时间"`
	UpdatedAt *gtime.Time `json:"updatedAt" orm:"updated_at" description:"数据更新时间"`
}
