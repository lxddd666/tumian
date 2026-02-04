// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// ValuationIndicators is the golang structure for table valuation_indicators.
type ValuationIndicators struct {
	Id                  uint64      `json:"id"                  orm:"id"                    description:"主键ID"`
	Symbol              string      `json:"symbol"              orm:"symbol"                description:"股票代码"`
	Mc                  string      `json:"mc"                  orm:"mc"                    description:"股票名称"`
	T                   *gtime.Time `json:"t"                   orm:"t"                     description:"数据日期"`
	CalcWindow          string      `json:"calcWindow"          orm:"calc_window"           description:"分位值计算窗口: 1Y, 3Y, 5Y, 10Y"`
	PeTtm               float64     `json:"peTtm"               orm:"pe_ttm"                description:"市盈率(TTM) - 股价/每股收益(TTM)"`
	PePercentile30      float64     `json:"pePercentile30"      orm:"pe_percentile_30"      description:"市盈率30分位值"`
	PePercentile70      float64     `json:"pePercentile70"      orm:"pe_percentile_70"      description:"市盈率70分位值"`
	PePercentileCurrent float64     `json:"pePercentileCurrent" orm:"pe_percentile_current" description:"当前市盈率历史百分位"`
	Pb                  float64     `json:"pb"                  orm:"pb"                    description:"市净率 - 股价/每股净资产"`
	PbPercentile30      float64     `json:"pbPercentile30"      orm:"pb_percentile_30"      description:"市净率30分位值"`
	PbPercentile70      float64     `json:"pbPercentile70"      orm:"pb_percentile_70"      description:"市净率70分位值"`
	PbPercentileCurrent float64     `json:"pbPercentileCurrent" orm:"pb_percentile_current" description:"当前市净率历史百分位"`
}
