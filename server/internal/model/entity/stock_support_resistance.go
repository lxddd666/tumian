// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// StockSupportResistance is the golang structure for table stock_support_resistance.
type StockSupportResistance struct {
	Id        uint64      `json:"id"        orm:"id"         description:"主键ID"`
	T         *gtime.Time `json:"t"         orm:"t"          description:"时间"`
	Symbol    string      `json:"symbol"    orm:"symbol"     description:"股票或标的代码 (例如: AAPL, 000001.SZ)"`
	Price     float64     `json:"price"     orm:"price"      description:"当前股票价格"`
	Yl        float64     `json:"yl"        orm:"yl"         description:"压力位"`
	Zc        float64     `json:"zc"        orm:"zc"         description:"支撑位"`
	CreatedAt *gtime.Time `json:"createdAt" orm:"created_at" description:"创建时间"`
}
