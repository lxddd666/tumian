// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// StockBasicInfo is the golang structure for table stock_basic_info.
type StockBasicInfo struct {
	Id             uint64      `json:"id"             orm:"id"               description:"自增主键"`
	Symbol         string      `json:"symbol"         orm:"symbol"           description:"股票代码"`
	Ii             string      `json:"ii"             orm:"ii"               description:""`
	Ei             string      `json:"ei"             orm:"ei"               description:""`
	Exchange       string      `json:"exchange"       orm:"exchange"         description:"交易所名称"`
	Name           string      `json:"name"           orm:"name"             description:""`
	ShortName      string      `json:"shortName"      orm:"short_name"       description:""`
	EnName         string      `json:"enName"         orm:"en_name"          description:""`
	Od             *gtime.Time `json:"od"             orm:"od"               description:"上市日期"`
	DataUpdateDate *gtime.Time `json:"dataUpdateDate" orm:"data_update_date" description:""`
	Pc             float64     `json:"pc"             orm:"pc"               description:""`
	Up             float64     `json:"up"             orm:"up"               description:""`
	Dp             float64     `json:"dp"             orm:"dp"               description:""`
	Pk             float64     `json:"pk"             orm:"pk"               description:""`
	Fv             int64       `json:"fv"             orm:"fv"               description:""`
	Tv             int64       `json:"tv"             orm:"tv"               description:""`
	FloatRatio     float64     `json:"floatRatio"     orm:"float_ratio"      description:"流通股比例 (%)"`
	Is             int         `json:"is"             orm:"is"               description:""`
	TradingStatus  string      `json:"tradingStatus"  orm:"trading_status"   description:"交易状态描述"`
	Industry       string      `json:"industry"       orm:"industry"         description:""`
	Sector         string      `json:"sector"         orm:"sector"           description:""`
	MarketType     string      `json:"marketType"     orm:"market_type"      description:""`
	DataSource     string      `json:"dataSource"     orm:"data_source"      description:""`
	IsActive       int         `json:"isActive"       orm:"is_active"        description:""`
	Version        int         `json:"version"        orm:"version"          description:""`
	CreatedAt      *gtime.Time `json:"createdAt"      orm:"created_at"       description:""`
	UpdatedAt      *gtime.Time `json:"updatedAt"      orm:"updated_at"       description:""`
}
