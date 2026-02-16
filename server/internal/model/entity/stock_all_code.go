// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// StockAllCode is the golang structure for table stock_all_code.
type StockAllCode struct {
	Dm        string      `json:"dm"        orm:"dm"         description:""`
	Mc        string      `json:"mc"        orm:"mc"         description:"股票名称"`
	Jys       string      `json:"jys"       orm:"jys"        description:"交易所"`
	CreatedAt *gtime.Time `json:"createdAt" orm:"created_at" description:"创建时间"`
	Industry  string      `json:"industry"  orm:"industry"   description:"所属行业"`
}
