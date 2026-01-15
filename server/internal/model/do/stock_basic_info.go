// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// StockBasicInfo is the golang structure of table hg_stock_basic_info for DAO operations like Where/Data.
type StockBasicInfo struct {
	g.Meta         `orm:"table:hg_stock_basic_info, do:true"`
	Id             any         // 自增主键
	Symbol         any         // 股票代码
	Ii             any         //
	Ei             any         //
	Exchange       any         // 交易所名称
	Name           any         //
	ShortName      any         //
	EnName         any         //
	Od             *gtime.Time // 上市日期
	DataUpdateDate *gtime.Time //
	Pc             any         //
	Up             any         //
	Dp             any         //
	Pk             any         //
	Fv             any         //
	Tv             any         //
	FloatRatio     any         // 流通股比例 (%)
	Is             any         //
	TradingStatus  any         // 交易状态描述
	Industry       any         //
	Sector         any         //
	MarketType     any         //
	DataSource     any         //
	IsActive       any         //
	Version        any         //
	CreatedAt      *gtime.Time //
	UpdatedAt      *gtime.Time //
}
