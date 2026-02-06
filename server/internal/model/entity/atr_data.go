// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// AtrData is the golang structure for table atr_data.
type AtrData struct {
	Id        uint64      `json:"id"        orm:"id"         description:"自增主键"`
	Symbol    string      `json:"symbol"    orm:"symbol"     description:"股票或标的代码 (例如: AAPL, 000001.SZ)"`
	T         *gtime.Time `json:"t"         orm:"t"          description:"交易时间 (统一为datetime类型，日线数据时间部分设为00:00:00)"`
	Tr        float64     `json:"tr"        orm:"tr"         description:"真实波幅 (True Range)"`
	Atr       float64     `json:"atr"       orm:"atr"        description:"平均真实波幅 (Average True Range)"`
	CreatedAt *gtime.Time `json:"createdAt" orm:"created_at" description:"数据创建时间"`
	UpdatedAt *gtime.Time `json:"updatedAt" orm:"updated_at" description:"数据更新时间"`
}
