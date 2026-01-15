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
	Id             interface{} // 自增主键
	Symbol         interface{} // 股票代码
	Ii             interface{} //
	Ei             interface{} //
	Exchange       interface{} // 交易所名称
	Name           interface{} //
	ShortName      interface{} //
	EnName         interface{} //
	Od             *gtime.Time // 上市日期
	DataUpdateDate *gtime.Time //
	Pc             interface{} //
	Up             interface{} //
	Dp             interface{} //
	Pk             interface{} //
	Fv             interface{} //
	Tv             interface{} //
	FloatRatio     interface{} // 流通股比例 (%)
	Is             interface{} //
	TradingStatus  interface{} // 交易状态描述
	Industry       interface{} //
	Sector         interface{} //
	MarketType     interface{} //
	DataSource     interface{} //
	IsActive       interface{} //
	Version        interface{} //
	CreatedAt      *gtime.Time //
	UpdatedAt      *gtime.Time //
}
