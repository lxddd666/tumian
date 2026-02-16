// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// RsiData is the golang structure for table rsi_data.
type RsiData struct {
	Id        uint64      `json:"id"        orm:"id"         description:"自增主键"`
	Symbol    string      `json:"symbol"    orm:"symbol"     description:"股票或标的代码 (例如: AAPL, 000001.SZ)"`
	T         *gtime.Time `json:"t"         orm:"t"          description:"交易时间 (统一为datetime类型，日线数据时间部分设为00:00:00)"`
	Rsi       float64     `json:"rsi"       orm:"rsi"        description:"rsi"`
	CreatedAt *gtime.Time `json:"createdAt" orm:"created_at" description:"数据创建时间"`
	UpdatedAt *gtime.Time `json:"updatedAt" orm:"updated_at" description:"数据更新时间"`
	Rsi6      float64     `json:"rsi6"      orm:"rsi_6"      description:""`
	Rsi12     float64     `json:"rsi12"     orm:"rsi_12"     description:""`
	Rsi24     float64     `json:"rsi24"     orm:"rsi_24"     description:""`
}
