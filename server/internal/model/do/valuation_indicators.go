// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// ValuationIndicators is the golang structure of table hg_valuation_indicators for DAO operations like Where/Data.
type ValuationIndicators struct {
	g.Meta              `orm:"table:hg_valuation_indicators, do:true"`
	Id                  any         // 主键ID
	Symbol              any         // 股票代码
	Mc                  any         // 股票名称
	T                   *gtime.Time // 数据日期
	CalcWindow          any         // 分位值计算窗口: 1Y, 3Y, 5Y, 10Y
	PeTtm               any         // 市盈率(TTM) - 股价/每股收益(TTM)
	PePercentile30      any         // 市盈率30分位值
	PePercentile70      any         // 市盈率70分位值
	PePercentileCurrent any         // 当前市盈率历史百分位
	Pb                  any         // 市净率 - 股价/每股净资产
	PbPercentile30      any         // 市净率30分位值
	PbPercentile70      any         // 市净率70分位值
	PbPercentileCurrent any         // 当前市净率历史百分位
}
