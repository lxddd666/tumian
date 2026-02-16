// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// StockSelfCode is the golang structure for table stock_self_code.
type StockSelfCode struct {
	Dm        string      `json:"dm"        orm:"dm"         description:"股票代码"`
	Mc        string      `json:"mc"        orm:"mc"         description:"股票名称"`
	Jys       string      `json:"jys"       orm:"jys"        description:"交易所"`
	CreatedAt *gtime.Time `json:"createdAt" orm:"created_at" description:"创建时间"`
	Industry  string      `json:"industry"  orm:"industry"   description:"industry"`
}
