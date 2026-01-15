// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// EnterpriseHistoricalData is the golang structure of table hg_enterprise_historical_data for DAO operations like Where/Data.
type EnterpriseHistoricalData struct {
	g.Meta        `orm:"table:hg_enterprise_historical_data, do:true"`
	Id            any         // 自增主键
	Symbol        any         // 证券代码 (如: 000001.SZ, AAPL)
	T             *gtime.Time // 交易时间 (精确到分钟或日)
	Date          *gtime.Time // 交易日期 (衍生字段)
	Year          any         // 交易年份
	Month         any         // 交易月份
	Weekday       any         // 星期几 (1=周日,7=周六)
	O             any         // 开盘价
	H             any         // 最高价
	L             any         // 最低价
	C             any         // 收盘价
	Pc            any         // 前收盘价
	V             any         // 成交量 (股/手)
	A             any         // 成交额 (元)
	Change        any         // 涨跌额
	ChangePct     any         // 涨跌幅 (%)
	Amplitude     any         // 振幅 (%)
	Sf            any         // 停牌标志: 0-正常, 1-停牌
	TradingStatus any         // 交易状态描述
	Period        any         // 数据周期: 1min, 5min, 15min, 30min, 60min, day, week, month
	IsAdjusted    any         // 是否复权: 0-不复权, 1-前复权, 2-后复权
	DataQuality   any         // 数据质量: 0-异常, 1-正常, 2-补全
	IsVerified    any         // 是否已验证: 0-未验证, 1-已验证
	DataSource    any         // 数据来源
	Version       any         // 数据版本
	CreatedAt     *gtime.Time // 创建时间
	UpdatedAt     *gtime.Time // 更新时间
}
