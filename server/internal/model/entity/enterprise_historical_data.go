// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// EnterpriseHistoricalData is the golang structure for table enterprise_historical_data.
type EnterpriseHistoricalData struct {
	Id        int64       `json:"id"        orm:"id"         description:"主键ID"`
	Symbol    string      `json:"symbol"    orm:"symbol"     description:"公司代码/股票代码 (例如: 000001.SZ, AAPL)"`
	T         *gtime.Time `json:"t"         orm:"t"          description:"交易时间"`
	O         float64     `json:"o"         orm:"o"          description:"开盘价"`
	H         float64     `json:"h"         orm:"h"          description:"最高价"`
	L         float64     `json:"l"         orm:"l"          description:"最低价"`
	C         float64     `json:"c"         orm:"c"          description:"收盘价"`
	V         float64     `json:"v"         orm:"v"          description:"成交量"`
	A         float64     `json:"a"         orm:"a"          description:"成交额"`
	Pc        float64     `json:"pc"        orm:"pc"         description:"前收盘价"`
	Sf        int         `json:"sf"        orm:"sf"         description:"停牌状态: 1停牌, 0不停牌"`
	CreatedAt *gtime.Time `json:"createdAt" orm:"created_at" description:"创建时间"`
	UpdatedAt *gtime.Time `json:"updatedAt" orm:"updated_at" description:"更新时间"`
}
